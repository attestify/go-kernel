package grant_owner

import (
	"github.com/attestify/go-kernel/authorization/owner_control"
	"github.com/attestify/go-kernel/error/internal_error"
)

type GrantOwner struct {
	ownerControl      owner_control.OwnerControl
	grantOwnerGateway GrantOwnerGateway
	grantOwnerError   error
}

func New(gateway GrantOwnerGateway) GrantOwner {
	var gatewayError error
	if gateway == nil {
		gatewayError = internal_error.New("the provided ModifyAccessGateway is nil, please provide a valid instance of an ModifyAccessGateway")
	}
	return GrantOwner{
		grantOwnerError:   gatewayError,
		grantOwnerGateway: gateway,
	}
}

func (usecase *GrantOwner) Grant(userId int64, resourceId int64) {
	usecase.setOwnerControl(userId, resourceId)
	usecase.grantOwner()
}

func (usecase *GrantOwner) setOwnerControl(userId int64, resourceId int64) {
	if usecase.HasError() {
		return
	}
	usecase.ownerControl = owner_control.MarkAsOwner(userId, resourceId)
}

func (usecase *GrantOwner) grantOwner() {
	if usecase.HasError() {
		return
	}
	usecase.grantOwnerGateway.Grant(usecase.ownerControl)
	if usecase.grantOwnerGateway.HasError() {
		usecase.grantOwnerError = usecase.grantOwnerGateway.Error()
	}
}

func (usecase GrantOwner) HasError() bool {
	return usecase.grantOwnerError != nil
}

func (usecase GrantOwner) Error() error {
	return usecase.grantOwnerError
}
