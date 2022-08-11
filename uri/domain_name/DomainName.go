package domain_name

import (
	"github.com/attestify/go-kernel/error/validation_error"
	"regexp"
)

type DomainName struct {
	value string
}

func New(value string) (DomainName, error) {
	length := len([]rune(value))
	if length < 1 || length > 255 {
		return DomainName{}, validation_error.New("The domain name value must be at least one (1) character, " +
			"and no greater than two-hundred fifty-five (255) characters.")
	}

	if !isValidDomainNameValue(value) {
		return DomainName{}, validation_error.New("The domain name can only be ASCII characters and hyphens.  " +
			"The domain name cannot start with a hyphen.")
	}

	return DomainName{
		value: value,
	}, nil
}

func (dn *DomainName) Value() string {
	return dn.value
}

func (dn *DomainName) Equals(compare DomainName) bool {
	return dn.Value() == compare.Value()
}

func isValidDomainNameValue(value string) bool {
	result := false
	result, _ = regexp.MatchString(`^[a-z0-9]+(-[a-z0-9]+)*`, value)
	return result
}
