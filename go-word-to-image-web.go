/*
Go-word-to-image-web is a web interface to go-word-to-image.
*/
package main

import (
	"fmt"
	"log"

	wi "github.com/maratart/go-word-to-image/wordtoimage"
)

func main() {
	word := "test"
	imageURL, err := wi.GetImageURL(word)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(imageURL)
}
