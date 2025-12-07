package service

import (
	"errors"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Если строка содержит любой символ кроме .-
// вернуть true
func сontainsSymbol(s string) bool {
	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' {
			return true
		}
	}
	return false
}

func Convert(s string) (string, error) {
	var result string
	if сontainsSymbol(s) {
		result = morse.ToMorse(s)
	} else {
		result = morse.ToText(s)
	}
	if len(result) == 0 {
		return result, errors.New("cannot convert string")
	}
	return result, nil
}
