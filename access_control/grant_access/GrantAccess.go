package grant_access

import (
	"github.com/attestify/go-kernel/access_control"
	"github.com/attestify/go-kernel/error/internal_error"
)

type GrantAccess struct {
	gateway       GrantAccessGateway
	accessControl access_control.AccessControl
	usecaseError  error
}

func New(gateway GrantAccessGateway) GrantAccess {
	var grantAccessError error
	if gateway == nil {
		grantAccessError = internal_error.New("the provided GrantAccessGateway is nil, please provide a valid instance of an GrantAccessGateway")
	}
	return GrantAccess{
		gateway:      gateway,
		usecaseError: grantAccessError,
	}
}

func (usecase *GrantAccess) Grant(userId int64, resourceId int64, resource string, permissions []string) {
	usecase.setAccessControl(userId, resourceId, resource, permissions)
	usecase.grantAccessControl()
}

func (usecase *GrantAccess) setAccessControl(userId int64, resourceId int64, resource string, permissions []string) {
	if usecase.HasError() {
		return
	}
	usecase.accessControl = access_control.New(userId, resourceId, resource, permissions)
}

func (usecase *GrantAccess) grantAccessControl() {
	if usecase.HasError() {
		return
	}
	usecase.usecaseError = usecase.gateway.Grant(usecase.accessControl)
}

// Error returns the current error.  This can be nil.
func (usecase GrantAccess) Error() error {
	return usecase.usecaseError
}

// HasError informs you if there is currently an error state
func (usecase GrantAccess) HasError() bool {
	return usecase.usecaseError != nil
}
