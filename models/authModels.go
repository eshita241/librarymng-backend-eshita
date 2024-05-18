package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Auth struct {
	ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string     `gorm:"type:varchar(100);not null"`
	Email     string     `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string     `gorm:"type:varchar(100);not null"`
	Role      *string    `gorm:"type:varchar(50);default:'user';not null"`
	Provider  *string    `gorm:"type:varchar(50);default:'local';not null"`
	Photo     *string    `gorm:"not null;default:'default.png'"`
	Verified  *bool      `gorm:"not null;default:false"`
	CreatedAt *time.Time `gorm:"not null;default:now()"`
	UpdatedAt *time.Time `gorm:"not null;default:now()"`
}

type AuthResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FilterUserRecord(auth *Auth) AuthResponse {
	return AuthResponse{
		ID:        *auth.ID,
		Name:      auth.Name,
		Email:     auth.Email,
		Role:      *auth.Role,
		Photo:     *auth.Photo,
		Provider:  *auth.Provider,
		CreatedAt: *auth.CreatedAt,
		UpdatedAt: *auth.UpdatedAt,
	}
}

// Initialize a new validator instance.
var validate = validator.New()

// ErrorResponse represents a structured validation error message.
type ErrorResponse struct {
	Field string `json:"field"`           // The field that caused the validation error.
	Tag   string `json:"tag"`             // The validation rule that failed.
	Value string `json:"value,omitempty"` // The value that caused the validation error (optional).
}

// ValidateStruct takes any payload struct and returns a slice of validation error messages.
func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse // Slice to hold the validation errors.

	// Perform the validation on the payload struct.
	err := validate.Struct(payload)
	if err != nil {
		// Iterate over the validation errors.
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace() // The field name causing the error.
			element.Tag = err.Tag()               // The validation tag that failed.
			element.Value = err.Param()           // The parameter associated with the validation tag (if any).
			errors = append(errors, &element)     // Append the error details to the errors slice.
		}
	}
	return errors // Return the slice of validation errors.
}

type SignUpInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8"`
	Photo           string `json:"photo"`
}

type SignInInput struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}
