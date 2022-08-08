package host

import "github.com/attestify/go-kernel/uri/registered_name"

// Host is a uri host as defined in RFC3986 Section 3.2.2.
// Source: https://datatracker.ietf.org/doc/html/rfc3986#section-3.2.2
type Host struct {
	hostType string
	value string
}

func NewFromRegisteredName(regName registered_name.RegisteredName) (Host, error) {
	return Host {
		hostType: "reg-name",
		value: regName.Value(),
	}, nil
}

func (host *Host) Value() string {
	return host.value
}

func (host *Host) HostType() string {
	return host.hostType
}
