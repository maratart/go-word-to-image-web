/*
Go-word-to-image-web is a web interface to go-word-to-image.
*/
package main

import (
	"fmt"
	"log"
	"net/http"

	wi "github.com/maratart/go-word-to-image/wordtoimage"
)

const form = `
<div style="text-align: center; margin-top: 50px;">
<form action="post">
<input type="text" name="word" size="40" value="%s" autofocus="autofocus" >
<input type="submit" value="Translate to image">
</form>
</div>
`
const picture = `
<div style="text-align: center">
<img src='%s' style='max-width: 600px; max-height: 600px;' />
</div>
`

func handlerShowImage(w http.ResponseWriter, r *http.Request) {
	// set content-type for the answer page
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// get word
	word := r.FormValue("word")
	// show form
	fmt.Fprintf(w, form, word)
	if len(word) > 0 {
		// translate word to image
		imageURL, err := wi.GetImageURL(word)
		// show image
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Error: %s", err)
		} else {
			fmt.Fprintf(w, picture, imageURL)
		}
	}
}

func main() {
	http.HandleFunc("/", handlerShowImage)
	http.ListenAndServe(":8001", nil)
}
