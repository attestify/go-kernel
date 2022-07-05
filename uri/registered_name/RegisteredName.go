package registered_name

import (
	"errors"
	"github.com/attestify/go-kernel/uri/domain_name"
	"github.com/attestify/go-kernel/uri/top_level_domain"
)

type RegisteredName struct {
	tld        top_level_domain.TopLevelDomain
	domainName domain_name.DomainName
}

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
		tld:        *tldInstance,
		domainName: *domainNameInstance,
	}, nil
}

func generateError(err error) error {
	errorMessage := "Error creating a Registered Name: " + err.Error()
	return errors.New(errorMessage)
}

func (rn *RegisteredName) Value() string {
	return rn.domainName.Value() + "." + rn.tld.Value()
}

func (rn *RegisteredName) Equals(compare *RegisteredName) bool {
	return rn.Value() == compare.Value()
}
