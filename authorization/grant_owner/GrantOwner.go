package grant_owner

import (
	"github.com/attestify/go-kernel/authorization"
	"github.com/attestify/go-kernel/error/internal_error"
)

type GrantOwner struct {
	accessControl     authorization.AccessControl
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
	usecase.setAccessControl(userId, resourceId)
	usecase.grantOwner()
}

func (usecase *GrantOwner) setAccessControl(userId int64, resourceId int64) {
	if usecase.HasError() {
		return
	}
	usecase.accessControl = authorization.New(userId, resourceId, []string{})
}

func (usecase *GrantOwner) grantOwner() {
	if usecase.HasError() {
		return
	}
	usecase.grantOwnerGateway.Grant(usecase.accessControl)
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
