package register

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type UserRegistrationData struct {
	FullName     string `json:"full_name"`
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

func encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func Register(w http.ResponseWriter, r *http.Request) {
	userData := UserRegistrationData{}

	json.NewDecoder(r.Body).Decode(&userData)
	defer r.Body.Close()
	fmt.Println(userData)

	// Encrypt password
	hash, err := encryptPassword(userData.Password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error encrypting password", http.StatusInternalServerError)
		return
	}

	userData.Password = hash

	// Store user in db, which will return user id

	// Start a session with user id

	fmt.Println(hash)

	w.Write([]byte("baka"))
}
