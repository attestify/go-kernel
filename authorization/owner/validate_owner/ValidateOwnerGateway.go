package validate_owner

import (
	"github.com/attestify/go-kernel/authorization/owner_control"
	"github.com/attestify/go-kernel/identity/id"
)

type ValidateOwnerGateway interface {
	Validate(userId id.Id, resourceId id.Id) owner_control.OwnerControl
	HasError() bool
	Error() error
}
