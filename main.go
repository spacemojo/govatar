package main

import (
	"crypto/md5"
	"fmt"
)

const (
	IMAGE_URL string = "https://www.gravatar.com/avatar/%s"
)

func main() {

	fmt.Println("Testing MD5 hash")
	defaultImageURL := DefaultGravatarImageURL("fred.leclerc@gmail.com")
	fmt.Printf("DefaultImageURL : %s", defaultImageURL)

}

func DefaultGravatarImageURL(email string) string {
	emailBytes := []byte(email)
	hash := fmt.Sprintf("%x", md5.Sum(emailBytes))
	return fmt.Sprintf(IMAGE_URL, hash)
}
