package storegs

import (
	"errors"

	"golang.org/x/text"
)


var (
	ErrURLNotFound = errors.New(text: "url not fount") 
	ErrURLExists = errors.New(text: "url exists")
)

