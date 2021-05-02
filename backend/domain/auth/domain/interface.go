package auth

import godd "github.com/pagongamedev/go-dd"

// Service interface
// Helper Service Use Gobal (First Uppercase) , API Service Use Local (First Lowercase)
type Service interface {
	// checkUsernameRead(username string) (bool *godd.Error)
	// checkPasswordRead(username string) (bool *godd.Error)
	registerCreate(username string, password string, phone string) (interface{}, *godd.Error)
}

// Repository interface
type Repository interface {
	UserRegister(username string, hashsalt string, phone string) (interface{}, *godd.Error)
}
