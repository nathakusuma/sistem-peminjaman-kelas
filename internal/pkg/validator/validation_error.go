package validator

import "encoding/json"

type validationError struct {
	Tag         string `json:"tag"`
	Param       string `json:"param"`
	Translation string `json:"translation"`
}

type ValidationErrors []map[string]validationError

func (v ValidationErrors) Error() string {
	j, err := json.Marshal(v)
	if err != nil {
		return ""
	}

	return string(j)
}

func (v ValidationErrors) Serialize() any {
	return v
}
