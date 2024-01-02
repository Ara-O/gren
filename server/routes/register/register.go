package register

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ara-o/gren/database"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRegistrationData struct {
	FullName     string `json:"full_name" bson:"full_name"`
	EmailAddress string `json:"email_address" bson:"email_address"`
	Password     string `json:"password" bson:"password"`
}

type Session struct {
	SessionID string    `json:"session_id" bson:"session_id"`
	UserID    string    `json:"user_id" bson:"user_id"`
	IAT       time.Time `json:"iat" bson:"iat"`
}

func (userData *UserRegistrationData) sanitize() {
	userData.FullName = strings.TrimSpace(userData.FullName)
	userData.EmailAddress = strings.ToLower(strings.TrimSpace(userData.EmailAddress))
	userData.Password = strings.TrimSpace(userData.Password)
}

func encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func validateAndFormalizeInput(userData *UserRegistrationData) error {
	// Sanitize user input
	userData.sanitize()

	// Encrypt password
	hash, err := encryptPassword(userData.Password)
	if err != nil {
		fmt.Println(err)
		return err
	}

	userData.Password = hash

	return nil
}

func Register(w http.ResponseWriter, r *http.Request) {
	userData := UserRegistrationData{}

	err := json.NewDecoder(r.Body).Decode(&userData)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "An error occured. Please try again", http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	if err := validateAndFormalizeInput(&userData); err != nil {
		http.Error(w, "An error occured. Please try again", http.StatusInternalServerError)
		return
	}

	conn := database.GetDatabase()

	coll := conn.Collection("users")

	var existingUser UserRegistrationData

	err = coll.FindOne(context.TODO(), bson.D{primitive.E{Key: "email_address", Value: userData.EmailAddress}}).Decode(&existingUser)

	if err != nil {
		if err != mongo.ErrNoDocuments {
			fmt.Println(err)
			http.Error(w, "There was an error", http.StatusInternalServerError)
			return
		}
	}

	if existingUser != (UserRegistrationData{}) {
		http.Error(w, "A user already exists", http.StatusInternalServerError)
		return
	}

	res, err := coll.InsertOne(context.Background(), userData)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error inserting user into database", http.StatusInternalServerError)
		return
	}

	insertedId := res.InsertedID.(primitive.ObjectID).Hex()

	sessionColl := conn.Collection("sessions")

	session := Session{
		SessionID: uuid.New().String(),
		UserID:    insertedId,
		IAT:       time.Now(),
	}

	//Insert session into database
	res, err = sessionColl.InsertOne(context.TODO(), session)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "An error has occured", http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    res.InsertedID.(primitive.ObjectID).Hex(),
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(24 * time.Hour), // Set the expiration time accordingly
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	return
}
