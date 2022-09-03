package grant_owner

import "github.com/attestify/go-kernel/authorization"

type GrantOwnerGateway interface {
	Grant(accessControl authorization.AccessControl)
	HasError() bool
	Error() error
}
