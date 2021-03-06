package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(wordkey string) (string, error) {
	found, ok := d[wordkey]

	if !ok {
		return "", ErrNotFound
	}

	return found, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
