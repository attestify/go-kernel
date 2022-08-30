package grant_owner

import "github.com/attestify/go-kernel/access_control"

type GrantOwnerGateway interface {
	Grant(accessControl access_control.AccessControl)
	HasError() bool
	Error() error
}
