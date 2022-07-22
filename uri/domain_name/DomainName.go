package domain_name

import (
	"errors"
	"regexp"
)

type DomainName struct {
	value string
}

func New(value string) (*DomainName, error) {
	length := len([]rune(value))
	if length < 1 || length > 255 {
		return &DomainName{}, errors.New("The domain name value must be atleast one (1) character, and no greather than two-hundred fifty-five (255) characters.")
	}

	if !isValidDomainNameValue(value) {
		return &DomainName{}, errors.New("The domain name can only be ASCII characters and hyphens.  The domain name cannot start with a hyphen.")
	}

	return &DomainName{
		value: value,
	}, nil
}

func (dn *DomainName) Value() string {
	return dn.value
}

func (dn *DomainName) Equals(compare *DomainName) bool {
	return dn.Value() == compare.Value()
}

func isValidDomainNameValue(value string) bool {
	result := false
	result, _ = regexp.MatchString(`^[a-z0-9]+(-[a-z0-9]+)*`, value)
	return result
}
