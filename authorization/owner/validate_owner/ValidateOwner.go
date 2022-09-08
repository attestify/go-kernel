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

func (usecase *ValidateOwner) Validate(userId int64, resourceId int64) owner_control.OwnerControl {
	if usecase.HasError() {
		return owner_control.MarkAsNotOwner(userId, resourceId)
	}
	ownerControl := usecase.gateway.Validate(id.New(userId), id.New(resourceId))
	if usecase.gateway.HasError() {
		usecase.usecaseError = usecase.gateway.Error()
	}
	return ownerControl
}

func (usecase ValidateOwner) HasError() bool {
	return usecase.usecaseError != nil
}

func (usecase ValidateOwner) Error() error {
	return usecase.usecaseError
}
