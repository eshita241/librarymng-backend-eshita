package helpers

import (
	"librarymng-backend/models"

	"github.com/asaskevich/govalidator"
)

func ValidateBook(Book models.Book) error {
	_, err := govalidator.ValidateStruct(Book)
	return err
}
func ValidateUser(user models.User) error {
	_, err := govalidator.ValidateStruct(user)
	return err
}
