package access_control

import "github.com/attestify/go-kernel/identity/id"

// RevokeAccessGateway provides the behavior for remove an access permission
// Expected Alternative Behaviors
//  - If the access permission exists, GrantAccessGateway will remove and respond with success
//	- If the does not exists, GrantAccessGateway will respond with success.
// Returns one (1) type of error:
//   - InternalError - Returned if there is an error invoking the implementation of this interface
type RevokeAccessGateway interface {
	Revoke(userId id.Id, entityId id.Id, entity string) error
}
