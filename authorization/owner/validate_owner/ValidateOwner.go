package validate_owner

import (
	"github.com/attestify/go-kernel/authorization/owner_control"
	"github.com/attestify/go-kernel/error/internal_error"
	"github.com/attestify/go-kernel/identity/id"
)

type ValidateOwner struct {
	gateway ValidateOwnerGateway
	usecaseError error
}

func New(gateway ValidateOwnerGateway) ValidateOwner {
	var gatewayError error
	if gateway == nil {
		gatewayError = internal_error.New("The provided ValidateOwnerGateway is nil. Please provide an instance of a ValidateOwnerGateway.")
	}
	return ValidateOwner{
		gateway: gateway,
		usecaseError: gatewayError,
	}
}

func (usecase ValidateOwner) Validate(userId int64, resourceId int64) owner_control.OwnerControl {
	return usecase.gateway.Validate(id.New(userId), id.New(resourceId))
}

func (usecase ValidateOwner) HasError() bool {
	return usecase.usecaseError != nil
}

func (usecase ValidateOwner) Error() error {
	return usecase.usecaseError
}
