package modify_access

import "github.com/attestify/go-kernel/access_control"

// ModifyAccessGateway provides the behavior for persisting a role record
// Expected Alternative Behaviors
//  - If the access permission_list does not exist, ModifyAccessGateway will modify and respond with success
//	- If the combination already exists, ModifyAccessGateway gateway will respond as if it was a successful record attempt.
// Returns one (1) type of error:
//  - InternalError - Returned if there is an error invoking the implementation of this interface
type ModifyAccessGateway interface {
	Modify(accessControl access_control.AccessControl)
	Error() error
	HasError() bool
}
