package access_control

import "github.com/attestify/go-kernel/identity/id"

// GrantAccessGateway provides the behavior for persisting a role record
// Expected Alternative Behaviors
//  - If the access control does not exist, GrantAccessGateway will grant and respond with success
//	- If the combination already exists, GrantAccessGateway gateway will respond as if it was a successful record attempt.
// Returns one (1) type of error:
//  - InternalError - Returned if there is an error invoking the implementation of this interface
type GrantAccessGateway interface {
	Grant(userId id.Id, entityId id.Id, entity string) error
}