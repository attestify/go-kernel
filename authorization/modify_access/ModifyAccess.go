package modify_access

import (
	"github.com/attestify/go-kernel/authorization"
	"github.com/attestify/go-kernel/error/internal_error"
)

type ModifyAccess struct {
	gateway       ModifyAccessGateway
	accessControl authorization.AccessControl
	usecaseError  error
}

func New(gateway ModifyAccessGateway) ModifyAccess {
	var modifyAccessError error
	if gateway == nil {
		modifyAccessError = internal_error.New("the provided ModifyAccessGateway is nil, please provide a valid instance of an ModifyAccessGateway")
	}
	return ModifyAccess{
		gateway:      gateway,
		usecaseError: modifyAccessError,
	}
}

func (usecase *ModifyAccess) Modify(userId int64, resourceId int64, permissions []string) {
	usecase.setAccessControl(userId, resourceId, permissions)
	usecase.modifyAccessControl()
}

func (usecase *ModifyAccess) setAccessControl(userId int64, resourceId int64, permissions []string) {
	if usecase.HasError() {
		return
	}
	usecase.accessControl = authorization.New(userId, resourceId, permissions)
}

func (usecase *ModifyAccess) modifyAccessControl() {
	if usecase.HasError() {
		return
	}
	usecase.gateway.Modify(usecase.accessControl)
	if usecase.gateway.HasError() {
		usecase.usecaseError = usecase.gateway.Error()
	}
}

// Error returns the current error.  This can be nil.
func (usecase ModifyAccess) Error() error {
	return usecase.usecaseError
}

// HasError informs you if there is currently an error state
func (usecase ModifyAccess) HasError() bool {
	return usecase.usecaseError != nil
}
