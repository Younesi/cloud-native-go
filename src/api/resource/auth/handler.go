package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"myapp/src/api"

	"github.com/golang-jwt/jwt/v5"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// Parse login credentials
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Verify credentials (replace with your actual authentication logic)
	if !verifyCredentials(creds.Username, creds.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": creds.Username,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(api.GetSecretKey())
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Return token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func verifyCredentials(username, password string) bool {
	// Implement your authentication logic here
	// This could involve checking against a database, etc.
	// For this example, we'll use a simple check
	return username == "testuser" && password == "testpass"
}
