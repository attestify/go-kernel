package revoke_access

import (
	"github.com/attestify/go-kernel/authorization/access_control"
)

// RevokeAccessGateway provides the behavior for remove an access permission_list
// Expected Alternative Behaviors
//  - If the access permission_list exists, ModifyAccessGateway will remove and respond with success
//	- If the does not exists, ModifyAccessGateway will respond with success.
// Returns one (1) type of error:
//   - InternalError - Returned if there is an error invoking the implementation of this interface
type RevokeAccessGateway interface {
	Revoke(accessControl access_control.AccessControl)
	Error() error
	HasError() bool
}
