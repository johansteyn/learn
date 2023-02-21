package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

const key = "my secure jwt key"

type UserClaim struct {
	jwt.RegisteredClaims
	ID    int
	Email string
	Name  string
}

func main() {
	fmt.Println("JWT")
	fmt.Println()

	// ------------ Generating tokens ------------
	// First, generate a token with an empty payload
	// jwt.RegisteredClaims contains common/standard JWT claims that are usually present in a
	// JWT token payload such as iat (token issue time), exp (token expiry time) and many more.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{})
	// For some reason using RS256 gives error: key is invalid
	//token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.RegisteredClaims{})
	jwtToken, err := token.SignedString([]byte(key))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("JWT Token: %s\n", jwtToken)

	// Second, generate a token with payload
	// Instead of using jwt.RegisteredClaims directly, use my custom UserClaim
	// which will result in a payload that contains the specified fields
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{},
		ID:               1,
		Email:            "alice.cooper@example.com",
		Name:             "Alice Cooper",
	})
	jwtToken, err = token.SignedString([]byte(key))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("JWT Token: %s\n", jwtToken)

	// ------------ Parsing tokens ------------
	var userClaim UserClaim
	jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiRW1haWwiOiJhbGljZS5jb29wZXJAZXhhbXBsZS5jb20iLCJOYW1lIjoiQWxpY2UgQ29vcGVyIn0.DKsnZ25Eapy1ry3LjwTyvIipUJq1QoCtn3DTZyD2qTY"

	// First, parse a token unverified (though this is not recommended)
	parser := jwt.NewParser()
	token, _, err = parser.ParseUnverified(jwtToken, &userClaim)
	fmt.Printf("Parsed User Claim: %d %s %s\n", userClaim.ID, userClaim.Email, userClaim.Name)

	// Second, parse the token verified by providing as 3rd parameter a function that returns the secret key bytes
	token, err = jwt.ParseWithClaims(jwtToken, &userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	// Check token validity
	if !token.Valid {
		fmt.Println("Invalid token!")
		return
	}
	fmt.Printf("Parsed User Claim: %d %s %s\n", userClaim.ID, userClaim.Email, userClaim.Name)
}
