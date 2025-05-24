package validator

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"sync"

	idlocales "github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	idtranslations "github.com/go-playground/validator/v10/translations/id"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/log"
)

type IValidator interface {
	ValidateStruct(data interface{}) error
	ValidateVariable(data interface{}, tag string) error
}

type validatorStruct struct {
	validator  *validator.Validate
	translator ut.Translator
}

var (
	validatorInstance IValidator
	once              sync.Once
)

func NewValidator() IValidator {
	once.Do(func() {
		id := idlocales.New()
		translator := ut.New(id, id)

		trans, found := translator.GetTranslator("id")
		if !found {
			log.Error(context.Background()).Msg("Translator not found")
		}

		val := validator.New()

		val.RegisterTagNameFunc(func(fld reflect.StructField) string {
			for _, tagName := range []string{"json", "form", "query", "param"} {
				if tag := fld.Tag.Get(tagName); tag != "" {
					name := strings.SplitN(tag, ",", 2)[0]
					if name == "-" {
						continue
					}
					return name
				}
			}
			// Fall back to field name if no tags are found
			return fld.Name
		})

		err := idtranslations.RegisterDefaultTranslations(val, trans)
		if err != nil {
			log.Error(context.Background()).Err(err).Msg("Failed to register translations")
		}

		validatorInstance = &validatorStruct{
			validator:  val,
			translator: trans,
		}
	})

	return validatorInstance
}

func (v *validatorStruct) ValidateStruct(data interface{}) error {
	if err := v.validator.Struct(data); err != nil {
		return v.handleValidationErrors(err)
	}
	return nil
}

func (v *validatorStruct) ValidateVariable(data interface{}, tag string) error {
	if err := v.validator.Var(data, tag); err != nil {
		return v.handleValidationErrors(err)
	}
	return nil
}

func (v *validatorStruct) handleValidationErrors(err error) error {
	var valErrs validator.ValidationErrors
	if errors.As(err, &valErrs) {
		length := len(valErrs)
		resp := make(ValidationErrors, length)
		for i, err := range valErrs {
			fieldPath := err.Namespace()

			parts := strings.Split(fieldPath, ".")
			if len(parts) > 1 {
				parts = parts[1:]
			}

			jsonTag := strings.ToLower(strings.Join(parts, "."))

			resp[i] = map[string]validationError{
				jsonTag: {
					Tag:         err.Tag(),
					Param:       err.Param(),
					Translation: err.Translate(v.translator),
				},
			}
		}
		return resp
	}

	log.Error(context.Background()).Err(err).Msg("Unexpected validation error")
	return err
}
