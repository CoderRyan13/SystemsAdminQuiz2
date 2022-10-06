//Filename: cmd/api/randomize.go

package main

import (
	"crypto/rand"
	"fmt"
	"net/http"

	"AWDquiz2.ryanarmstrong.net/internal/data"
	"AWDquiz2.ryanarmstrong.net/internal/validator"
)

//type Tools struct{}

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+_#$-!~"

func (app *application) generateRandomString(length int64) string {
	s := make([]rune, length)
	r := []rune(randomStringSource)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x := p.Uint64()
		y := uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)

}

// showInfoHandler for the "Post /v1/info" endpoint
func (app *application) showInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Our target decode destination
	var input struct {
		Info string `json:"info"`
	}
	// Initialize a new json.Decoder instance
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Copy the values from the input struct to a new Information struct
	info := &data.Information{
		Info: input.Info,
	}
	// Initialize a new Validator instance
	v := validator.New()

	// Check the map to determine if there were any validation errors
	if data.ValidateInfo(v, info); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	// Display the request
	fmt.Fprintf(w, "%+v\n", input)
}

// randomStringHandler for the "Get /v1/randomstring/:id" endpoint
func (app *application) randomStringHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Pass the random string from the function generateRandomString to a variable
	randomString := app.generateRandomString(id)
	// Display the random string
	fmt.Fprintf(w, "%+v\n", randomString)
}
