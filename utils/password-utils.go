package utils

import "golang.org/x/crypto/bcrypt"

func HashAndSaltPassword(password string) (string, error) {
	// Generate a salt with a cost factor of 14 (you can adjust the cost factor)
	salt, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(salt), nil
}

// CheckPassword compares a plaintext password with a hashed and salted password
func CheckPassword(plaintextPassword, hashedPassword string) error {
	// Compare the plaintext password with the hashed and salted password
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plaintextPassword))
}
