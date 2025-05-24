package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/enum"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/errorpkg"
)

type IJwt interface {
	Create(userID uuid.UUID, role enum.UserRole) (string, error)
	Decode(tokenString string, claims *Claims) error
	Validate(token string) (ValidateJWTResponse, error)
}

type Claims struct {
	jwt.RegisteredClaims
	Role enum.UserRole `json:"role"`
}

type ValidateJWTResponse struct {
	UserID uuid.UUID
	Role   enum.UserRole
}

type jwtStruct struct {
	exp    time.Duration
	secret []byte
}

func NewJwt(exp time.Duration, secret []byte) IJwt {
	return &jwtStruct{
		exp:    exp,
		secret: secret,
	}
}

func (j *jwtStruct) Create(userID uuid.UUID, role enum.UserRole) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.exp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Role: role,
	}

	unsignedJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedJWT, err := unsignedJWT.SignedString(j.secret)
	if err != nil {
		return "", err
	}

	return signedJWT, nil
}

func (j *jwtStruct) Decode(tokenString string, claims *Claims) error {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(_ *jwt.Token) (any, error) {
		return j.secret, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return jwt.ErrSignatureInvalid
	}

	return nil
}

func (j *jwtStruct) Validate(token string) (ValidateJWTResponse, error) {
	var claims Claims
	err := j.Decode(token, &claims)
	if err != nil {
		return ValidateJWTResponse{}, errorpkg.ErrInvalidBearerToken()
	}

	expirationTime, err := claims.GetExpirationTime()
	if err != nil {
		return ValidateJWTResponse{}, errorpkg.ErrInvalidBearerToken()
	}

	if expirationTime.Before(time.Now()) {
		return ValidateJWTResponse{}, errorpkg.ErrInvalidBearerToken()
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return ValidateJWTResponse{}, errorpkg.ErrInvalidBearerToken()
	}

	return ValidateJWTResponse{
		UserID: userID,
		Role:   claims.Role,
	}, nil
}
