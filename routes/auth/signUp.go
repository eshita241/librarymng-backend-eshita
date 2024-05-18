package auth

import (
	"librarymng-backend/database"
	"librarymng-backend/models"
	"strings"

	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

// SignUpUser is a handler function for the user sign-up process.
func SignUpUser(c *fiber.Ctx) error {
	// Declare a variable to hold the incoming request payload.
	var payload *models.SignUpInput

	// Parse the JSON request body into the payload variable.
	// Parsing, in the context of computing and programming, refers to the process of analyzing a string of symbols, either in natural language or in computer languages, according to the rules of a formal grammar.
	// It involves taking input in the form of a string of text and converting it into a more structured format that a program can process more easily.
	if err := c.BodyParser(&payload); err != nil {
		// Return a 400 Bad Request status if parsing fails.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// Validate the input fields using a validation function.
	errors := models.ValidateStruct(payload)
	if errors != nil {
		// Return a 400 Bad Request status if validation errors are found.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	// Check if the password and password confirmation match.
	if payload.Password != payload.PasswordConfirm {
		// Return a 400 Bad Request status if passwords do not match.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Passwords do not match"})
	}

	// Hash the password using bcrypt.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		// Return a 400 Bad Request status if hashing fails.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// Create a new user record with the hashed password.
	newUser := models.Auth{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email), // Convert email to lowercase for consistency.
		Password: string(hashedPassword),         // Store the hashed password.
		Photo:    &payload.Photo,                 // Store the photo URL if provided.
		//We used pointer here because we have access over updating photos, also to store small data instead of large data.
	}

	// Insert the new user record into the database.
	result := database.Database.Db.Create(&newUser)

	// Handle potential database errors.
	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		// Return a 409 Conflict status if the email already exists.
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "User with that email already exists"})
	} else if result.Error != nil {
		// Return a 502 Bad Gateway status for other database errors.
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	// Return a 201 Created status with the new user data, excluding sensitive fields.
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": models.FilterUserRecord(&newUser)}})
}

/*
When a POST request is made to the /api/auth/register endpoint, the Fiber framework will trigger this route function to parse and validate the incoming data according to the validation rules specified in the models.ValidateStruct() struct.
Next, the function will check if the Password and PasswordConfirm fields match. Upon successful validation, it will proceed to hash the password for security and utilize GORM’s DB.Create() method to store the new user’s information in the database.
In the event that a user with the same email already exists, a 409 Conflict error will be returned to the client. On the other hand, if the registration is successful, a filtered version of the newly registered user will be returned in the JSON response.
*/

/*
1. Body Parsing (`c.BodyParser(&payload)`):
   Purpose: To convert the raw JSON request body into a Go struct.
   Scope: It only checks whether the incoming JSON can be successfully parsed into the specified struct. This includes:
     - Correct JSON syntax.
     - Correct data types for fields (e.g., strings, numbers).

2. Input Validation (`models.ValidateStruct(payload)`):
   - Purpose: To check the business logic rules and constraints on the parsed data.
   - Scope: It validates the values of the fields to ensure they meet specific criteria. This can include:
     - Required fields (e.g., ensuring `email` and `password` are not empty).
     - Format constraints (e.g., email must be a valid email format).
     - Length constraints (e.g., password must be at least 8 characters long).
     - Cross-field validation (e.g., password and password confirmation match).

1. Parsing Step:
   - Error Example: If the JSON is malformed, such as `{"email": "user@example.com", "password": 123}`, `BodyParser` will fail because `password` should be a string, not a number. This ensures that the basic structure and types are correct.
   - Handling: If `BodyParser` fails, the request is immediately rejected with a 400 Bad Request response.

2. Validation Step:
   - Error Example: If the JSON is syntactically correct but does not meet business logic rules, such as `{"email": "invalid-email", "password": "short"}`, `ValidateStruct` will catch these issues.
   - Handling: If `ValidateStruct` finds validation errors, the request is rejected with detailed error messages specifying which fields are invalid and why.
*/

/*
Hashing Passwords: Hashing transforms the password into a fixed-size string of characters, which is typically a one-way process (it is computationally infeasible to reverse the hash to get the original password).
Security: Even if someone gains access to the hashed passwords, they cannot easily obtain the original passwords.
Salting: bcrypt automatically generates a salt and includes it in the hash. A salt is a random value added to the password before hashing, which ensures that even identical passwords will have different hashes.
*/
