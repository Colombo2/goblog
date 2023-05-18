package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(passwd []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(passwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hashed)
}

func Compare(hpasswd string, ppasswd[]byte) bool {
    byteHash := []byte(hpasswd)
    err := bcrypt.CompareHashAndPassword(byteHash, ppasswd)
    if err != nil {
        log.Println(err)
        return false
    }
    
    return true
}