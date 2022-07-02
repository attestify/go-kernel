package top_level_domain

import (
	"errors"
	"regexp"
)

type TopLevelDomain struct {
	value string
}

func New(value string) (*TopLevelDomain, error) {
	if len(value) < 1 {
		return &TopLevelDomain{}, errors.New("The top level domain value must be atleast one (1) character.")
	}

	if !isOnlyLetters(value) {
		return &TopLevelDomain{}, errors.New("The top level domain value can only be letters.")
	}

	return &TopLevelDomain{
		value: value,
	}, nil
}

func (tld *TopLevelDomain) Value() string {
	return tld.value
}

func isOnlyLetters(value string) bool {
	result, err := regexp.MatchString(`^[A-Za-z]+$`, value)
	if err != nil {
		return false
	}
	return result
}
