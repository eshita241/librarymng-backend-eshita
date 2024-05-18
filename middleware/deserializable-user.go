package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/wpcodevo/golang-fiber-jwt/initializers"
	"github.com/wpcodevo/golang-fiber-jwt/models"
)

func DeserializeUser(c *fiber.Ctx) error {
	// Initialize variables to store the JWT token string.
	var tokenString string
	authorization := c.Get("Authorization")

	// Check if the Authorization header starts with "Bearer ".
	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		// If no token in Authorization header, check for token in cookies.
		tokenString = c.Cookies("token")
	}

	// If tokenString is empty, return unauthorized status.
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	// Load application configuration settings.
	config, _ := initializers.LoadConfig(".")

	// Parse and validate the JWT token.
	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		// Verify the signing method of the token.
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		// Return the secret key used for signing the token.
		return []byte(config.JwtSecret), nil
	})
	if err != nil {
		// Handle token parsing or validation errors.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalid token: %v", err)})
	}

	// Verify token claims and retrieve user information.
	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid token claim"})
	}

	// Retrieve user information from the database based on token claims.
	var user models.User
	initializers.DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))

	// Verify if the user associated with the token still exists.
	if user.ID.String() != fmt.Sprint(claims["sub"]) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "the user belonging to this token no longer exists"})
	}

	// Store user information in Fiber's local storage for subsequent middleware or handlers.
	c.Locals("user", models.FilterUserRecord(&user))

	// Proceed to the next middleware or handler in the chain.
	return c.Next()
}

//The DeserializeUser middleware function is designed to deserialize (parse and validate) a JWT token from either the Authorization header or a cookie in a Fiber application.
//Its primary purpose is to authenticate and authorize incoming requests by verifying the validity of the JWT token and retrieving user information associated with the token.

/*
Validating the password during the login process (`bcrypt.CompareHashAndPassword`) ensures that the user attempting to log in provides the correct credentials. This initial validation step confirms the user's identity and allows them to proceed to receive a JWT token for subsequent authenticated requests.
The middleware validation (`DeserializeUser` function or similar) serves a different purpose in the application's workflow:

1. Middleware Purpose:
   - The middleware (`DeserializeUser` in this case) is responsible for validating and decoding the JWT token from incoming requests.
   - It checks if the token is present and valid (not expired or tampered with).
   - Retrieves the user information associated with the token, if valid, and stores it in the request context (`c.Locals("user", ...)`) for use by subsequent handlers or middleware.
   - Ensures that subsequent requests that require authentication can access user information without repeating the JWT decoding and validation logic.

2. Why Validate in Middleware?
   - **Session Persistence:** The middleware validates the JWT token to maintain session persistence across multiple requests. Once a user logs in and receives a JWT token, subsequent requests can include this token. The middleware ensures these requests are authenticated without re-verifying credentials on each request.
   - **Authorization:** Besides authentication, middleware often checks for authorization roles or permissions associated with the user. This ensures that authenticated users have the necessary privileges to access specific resources or perform actions.
   - **Security Layer:** Middleware validation serves as a security layer, preventing unauthorized access attempts by rejecting requests with invalid or expired tokens.

3. **Overall Workflow:**
   - During login, `bcrypt.CompareHashAndPassword` verifies the user's password correctness against the stored hash, allowing the issuance of a JWT token upon successful validation.
   - The JWT token is then used in subsequent requests to validate the user's identity and authorize access to protected resources via middleware validation.
   - This separation of concerns (password validation during login vs. token validation in middleware) enhances security, scalability, and maintainability of the authentication system.

In summary, validating the password during login and validating the JWT token in middleware serve complementary purposes: one establishes initial identity verification, while the other maintains authenticated sessions and access control throughout the application's interaction with authenticated users.
*/
