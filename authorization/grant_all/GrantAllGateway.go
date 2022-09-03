package grant_all

import "github.com/attestify/go-kernel/authorization"

type GrantAllGateway interface {
	Grant(control authorization.AccessControl)
	Error() error
	HasError() bool
}
