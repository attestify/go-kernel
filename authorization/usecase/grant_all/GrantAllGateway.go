package grant_all

import (
	"github.com/attestify/go-kernel/authorization/access_control"
)

type GrantAllGateway interface {
	Grant(control access_control.AccessControl)
	Error() error
	HasError() bool
}
