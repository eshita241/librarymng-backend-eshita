package validator

import (
	"librarymng-backend/models" // Importing the models package

	"github.com/asaskevich/govalidator" // Importing the govalidator package for struct validation
)

// ValidateBook validates a Book instance using the govalidator package.
func ValidateBook(Book models.Book) error {
	// Validate the Book struct using govalidator.ValidateStruct
	_, err := govalidator.ValidateStruct(Book)
	return err // Return any validation error
}

// ValidateUser validates a User instance using the govalidator package.
func ValidateUser(user models.User) error {
	// Validate the User struct using govalidator.ValidateStruct
	_, err := govalidator.ValidateStruct(user)
	return err // Return any validation error
}

// ValidateIssue validates an Issue instance using the govalidator package.
func ValidateIssue(issue models.Issue) error {
	// Validate the Issue struct using govalidator.ValidateStruct
	_, err := govalidator.ValidateStruct(issue)
	return err // Return any validation error
}
