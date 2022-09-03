package grant_all

import (
	"github.com/attestify/go-kernel/authorization/access_control"
	"github.com/attestify/go-kernel/error/internal_error"
)

type GrantAll struct {
	accessControl   access_control.AccessControl
	grantAllGateway GrantAllGateway
	usecaseError    error
}

func New(gateway GrantAllGateway) GrantAll {
	var gatewayError error
	if gateway == nil {
		gatewayError = internal_error.New("The provided GrantAllGateway is nil. Please provide a valid instance of an GrantAllGateway.")
	}
	return GrantAll{
		grantAllGateway: gateway,
		usecaseError:    gatewayError,
	}
}

func (uc *GrantAll) Grant(resourceId int64, permissions []string) {
	uc.setAccessControl(resourceId, permissions)
	uc.grantAll()
}

func (uc *GrantAll) setAccessControl(resourceId int64, permissions []string) {
	if uc.HasError() {
		return
	}
	uc.accessControl = access_control.New(0, resourceId, permissions)
}

func (uc *GrantAll) grantAll() {
	if uc.HasError() {
		return
	}
	uc.grantAllGateway.Grant(uc.accessControl)
	if uc.grantAllGateway.HasError() {
		uc.usecaseError = uc.grantAllGateway.Error()
	}
}

func (uc GrantAll) HasError() bool {
	return uc.usecaseError != nil
}

func (uc GrantAll) Error() error {
	return uc.usecaseError
}
