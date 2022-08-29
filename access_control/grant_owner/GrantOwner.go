package grant_owner

type GrantOwner struct {
	grantOwnerError error
}

func New(gateway GrantOwnerGateway) GrantOwner {
	return GrantOwner{}
}

func (uc GrantOwner) Grant(userId int64, resourceId int64) {

}

func (uc GrantOwner) HasError() bool {
	return uc.grantOwnerError != nil
}

func (uc GrantOwner) Error() error {
	return uc.grantOwnerError
}
