package grant_owner

import (
	"github.com/attestify/go-kernel/authorization/owner_control"
)

type GrantOwnerGateway interface {
	Grant(ownerControl owner_control.OwnerControl)
	HasError() bool
	Error() error
}
