package grant_access

import (
	"github.com/attestify/go-kernel/access_control"
	"github.com/attestify/go-kernel/error/internal_error"
	"github.com/attestify/go-kernel/identity/id"
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

func (usecase GrantAccess) Grant(userId int64, entityId int64, entity string) error {

	err := usecase.gateway.GrantAccess(id.New(userId), id.New(entityId), entity)
	if err != nil {
		return err
	}
	return nil
}
