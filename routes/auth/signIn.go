package auth

import (
	"fmt"
	"librarymng-backend/database"
	"librarymng-backend/initializers"
	"librarymng-backend/models"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

// SignInUser is a handler function for the user sign-in process.
func SignInUser(c *fiber.Ctx) error {
	var payload *models.SignInInput

	// Parse the JSON request body into the payload variable.
	if err := c.BodyParser(&payload); err != nil {
		// Return a 400 Bad Request status if parsing fails.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// Validate the input fields using a validation function.
	errors := models.ValidateStruct(payload)
	if errors != nil {
		// Return a 400 Bad Request status if validation errors are found.
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Find the user in the database by email.
	var user models.Auth
	result := database.Database.Db.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		// Return a 400 Bad Request status if the user is not found.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid email or Password"})
	}

	// Compare the provided password with the stored hashed password.
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		// Return a 400 Bad Request status if the password is incorrect.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid email or Password"})
	}

	// Load configuration settings.
	config, _ := initializers.LoadConfig(".")

	// Create a new JWT token.
	tokenByte := jwt.New(jwt.SigningMethodHS256) //HMAC-SHA256 is a secure algorithm used to create a digital signature for the token

	// Get the current time in UTC.
	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims) //Retrieves the claims (payload data) associated with the JWT token. jwt.MapClaims allows you to define arbitrary key-value pairs as token claims.

	// Set the token claims.
	claims["sub"] = user.ID
	claims["exp"] = now.Add(config.JwtExpiresIn).Unix() // Token expiration time.
	claims["iat"] = now.Unix()                          // Token issued at time.
	claims["nbf"] = now.Unix()                          // Token not before time.

	// Sign the token with the secret key.
	tokenString, err := tokenByte.SignedString([]byte(config.JwtSecret)) //okenByte is an instance of jwt.Token that has been configured with claims (like sub, exp, iat, nbf) and a signing method (jwt.SigningMethodHS256 in this case).
	//tokenByte.SignedString([]byte(config.JwtSecret)) is a method call on tokenByte.This method takes a secret key ([]byte(config.JwtSecret)) as an argument. This secret key is used to digitally sign the token.
	if err != nil {
		// Return a 502 Bad Gateway status if signing the token fails.
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("generating JWT Token failed: %v", err)})
	}

	// Set the JWT token in an HTTP-only cookie.
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(config.JwtMaxAge) * time.Minute), // Use Expires to set cookie max age.
		Secure:   false,                                                         //Secure: false`: If set to true, the cookie will only be sent over HTTPS connections. Setting it to false means the cookie can also be sent over HTTP connections, though this is less secure.
		HTTPOnly: true,                                                          //HTTPOnly: true`: This is the key aspect of an HTTP-only cookie. When set to true, it instucts the browser that the cookie should not be accessible via JavaScript. This helps mitigate certain types of cross-site scripting (XSS) attacks because JavaScript running on a page cannot access the cookie's content.
		Domain:   "localhost",                                                   //`Domain: "localhost"`: Specifies the domain for which the cookie is valid. In this case, it's set to "localhost", meaning the cookie is only sent to and from the server on the localhost domain.
	})

	// Return a 200 OK status with the JWT token.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "token": tokenString})
}

/*
This Fiber context handler will be responsible for processing user sign-in requests.
When this Fiber context handler is triggered, it will parse the incoming data, validate the fields according to the models.SignInInput struct, and check the database to see if a user with the provided email exists.
If the user exists, the handler will compare the provided password with the one stored in the database by utilizing the bcrypt.CompareHashAndPassword() method. This will confirm if the user has provided the correct credentials.
Upon successful authentication, a JWT token will be generated using and returned to the client in the form of an HTTP-only cookie.
*/

/*
JWT (JSON Web Token) tokens are typically created during the login process, not during signup.
Login vs. Signup
1. Signup Process:
   - During signup, a new user is registering and providing their credentials (like email and password).
   - The primary action during signup is to validate the provided information, hash the password, and store the user's data securely in the database.
   - JWT tokens are not typically created during signup because the user has not yet been authenticated and verified.

2. Login Process:
   - During login, a user who has already signed up is providing their credentials (email and password) to authenticate themselves.
   - Once the credentials are validated (email exists and password matches), a JWT token is generated.
   - This JWT token serves as a secure way to authenticate subsequent requests without requiring the user to resend their credentials with every request.
   It typically contains claims (like user ID, expiration time, etc.) that are signed using a secret key known only to the server. This ensures the token's integrity and prevents tampering.

Benefits of JWT Tokens
- **Stateless**: Since JWT tokens contain all necessary information within themselves, they are stateless and do not require server-side storage.
- **Security**: Tokens are signed using a secret key, ensuring that their content cannot be tampered with.
- **Efficiency**: Once issued, tokens allow clients to access authorized resources without repeatedly sending credentials, improving performance.
*/

/*
An HTTP-only cookie is a type of cookie that is designed to be inaccessible to JavaScript running in the browser. Here’s an explanation of each parameter in the context of setting an HTTP-only cookie for JWT token storage:
HTTP-only cookies enhance security by preventing client-side JavaScript from accessing sensitive cookie information. This helps protect against XSS attacks where malicious scripts attempt to steal session cookies.
setting `HTTPOnly: true` ensures that the JWT token stored in the "token" cookie can only be accessed and used by the server. This strengthens the security of your application by reducing the risk of client-side attacks that might attempt to steal or manipulate the JWT token.
*/
