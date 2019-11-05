package main

import (
	"errors"
)

type Dictionary map[string]string

var ErrNotFound = errors.New("Could not find supplied word")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
