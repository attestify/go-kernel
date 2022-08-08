package registered_name

import (
	"errors"
	"github.com/attestify/go-kernel/uri/domain_name"
	"github.com/attestify/go-kernel/uri/top_level_domain"
	"strings"
)

// RegisteredName is a registered name intended for lookup in the DNS uses
// the syntax defined in Section 3.5 of [RFC1034] and Section 2.1 of [RFC1123]
type RegisteredName struct {
	tld        top_level_domain.TopLevelDomain
	domainName domain_name.DomainName
}

// New Instantiates a RegisteredName class from separated top level domain
// and domain name string objects.
func New(tld string, domainName string) (*RegisteredName, error) {
	tldInstance, err := top_level_domain.New(tld)
	if err != nil {
		return &RegisteredName{}, generateError(err)
	}

	domainNameInstance, err := domain_name.New(domainName)
	if err != nil {
		return &RegisteredName{}, generateError(err)
	}

	return &RegisteredName{
		tld:        tldInstance,
		domainName: domainNameInstance,
	}, nil
}

// NewFromString Instantiates a RegisteredName class from a
// full registered name string.  Example: "attestify.io"
func NewFromString(registeredName string) (*RegisteredName, error) {

	tld := registeredName[strings.LastIndex(registeredName, ".")+1:]
	domainName := strings.TrimRight(registeredName,"."+tld)

	return New(tld, domainName)
}

// generateError constructs an error message for the RegisteredName struct
func generateError(err error) error {
	errorMessage := "Error creating a Registered Name: " + err.Error()
	return errors.New(errorMessage)
}

// Value returns the combined value of the RegisteredName
// top level domain and domain name as a single string.
func (rn *RegisteredName) Value() string {
	return rn.domainName.Value() + "." + rn.tld.Value()
}

// Equals compares the current RegisteredName to another instance
// of a RegisteredName object to asses equality
func (rn *RegisteredName) Equals(compare *RegisteredName) bool {
	return rn.Value() == compare.Value()
}
