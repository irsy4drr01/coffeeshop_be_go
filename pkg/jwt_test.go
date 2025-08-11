package pkg

// import (
// 	"os"
// 	"testing"
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// )

// func TestGenerateToken(t *testing.T) {
// 	// Set up the JWT secret in environment variables for testing purposes.
// 	secret := "testsecret"
// 	os.Setenv("JWT_SECRET", secret)

// 	// Create a new JWT claim
// 	uuid := "1234"
// 	email := "test@example.com"
// 	role := "user"
// 	claims := NewJWT(uuid, email, role)

// 	// Test token generation
// 	token, err := claims.GenerateToken()
// 	if err != nil {
// 		t.Fatalf("Expected no error while generating token, got %v", err)
// 	}
// 	if token == "" {
// 		t.Fatalf("Expected a valid token, got empty string")
// 	}
// }

// func TestVerifyToken(t *testing.T) {
// 	// Set up the JWT secret in environment variables for testing purposes.
// 	secret := "testsecret"
// 	os.Setenv("JWT_SECRET", secret)

// 	// Create a new JWT claim
// 	uuid := "1234"
// 	email := "test@example.com"
// 	role := "user"
// 	claims := NewJWT(uuid, email, role)

// 	// Generate a token for verification testing
// 	token, err := claims.GenerateToken()
// 	if err != nil {
// 		t.Fatalf("Expected no error while generating token, got %v", err)
// 	}

// 	// Test token verification
// 	verifiedClaims, err := VerifyToken(token)
// 	if err != nil {
// 		t.Fatalf("Expected no error while verifying token, got %v", err)
// 	}
// 	if verifiedClaims.Uuid != uuid {
// 		t.Fatalf("Expected UUID to be %s, got %s", uuid, verifiedClaims.Uuid)
// 	}
// 	if verifiedClaims.Email != email {
// 		t.Fatalf("Expected email to be %s, got %s", email, verifiedClaims.Email)
// 	}
// 	if verifiedClaims.Role != role {
// 		t.Fatalf("Expected role to be %s, got %s", role, verifiedClaims.Role)
// 	}

// 	// Test token expiration
// 	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(-time.Minute))
// 	expiredToken, err := claims.GenerateToken()
// 	if err != nil {
// 		t.Fatalf("Expected no error while generating expired token, got %v", err)
// 	}

// 	_, err = VerifyToken(expiredToken)
// 	if err == nil {
// 		t.Fatalf("Expected error when verifying expired token, but got no error")
// 	}
// }

// func TestInvalidToken(t *testing.T) {
// 	// Test with an invalid token
// 	invalidToken := "invalidtoken"
// 	_, err := VerifyToken(invalidToken)
// 	if err == nil {
// 		t.Fatalf("Expected error when verifying invalid token, but got no error")
// 	}
// }
