package grant_access

import (
	"github.com/attestify/go-kernel/access_control"
	"github.com/attestify/go-kernel/error/internal_error"
)

type GrantAccess struct {
	gateway access_control.GrantAccessGateway
}

func New(gateway access_control.GrantAccessGateway) (GrantAccess, error) {

	if gateway == nil {
		return GrantAccess{}, internal_error.New("the provided GrantAccessGateway is nil, please provide a valid instance of an GrantAccessGateway")
	}

	return GrantAccess{
		gateway: gateway,
	}, nil
}

func (usecase GrantAccess) Grant(userId int64, resourceId int64, resource string, permissions []string) error {

	accessControl, _ := access_control.New(userId, resourceId, resource, permissions)
	err := usecase.gateway.Grant(accessControl)
	if err != nil {
		return err
	}
	return nil
}
