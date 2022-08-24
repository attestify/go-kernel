package permission_list

import (
	"github.com/attestify/go-kernel/error/validation_error"
	"github.com/attestify/go-kernel/strings"
)

// PermissionList
// Expected behaviour
// - Only characters allowed are alphabet characters
// - Will remove any numeric values
// - Replace any special characters with a dash
// - Must only start with an alpha character, and can only end with an alpha character
type PermissionList struct {
	permissions []string
	listError   error
}

func New() PermissionList {
	return PermissionList{}
}

func (list *PermissionList) AddPermission(permission string) {
	if list.HasError() {
		return
	}
	cleanedPermission := list.cleanPermission(permission)
	if list.HasError() { return }
	if list.ContainsPermission(cleanedPermission) {
		return
	}
	list.permissions = append(list.permissions, cleanedPermission)
}

func (list *PermissionList) AddManyPermissions(permissions []string) {
	for _, permissionInList := range permissions {
		list.AddPermission(permissionInList)
	}
}

func (list PermissionList) ContainsPermission(permission string) bool {
	for _, _permission := range list.permissions {
		if permission == _permission {
			return true
		}
	}
	return false
}

func (list PermissionList) GetAllPermissions() []string {
	return list.permissions
}

func (list PermissionList) Error() error {
	return list.listError
}

func (list PermissionList) HasError() bool {
	return list.listError != nil
}

func (list *PermissionList) cleanPermission(permission string) string {
	permission = strings.CleanAndLower(permission)
	permission = strings.RemoveAllNumbers(permission)
	permission = strings.ReplaceSpecialCharactersWithDash(permission)
	permission = strings.CleanLeadAndTrailSpecialCharacter(permission)
	if strings.LengthIsZero(permission) {
		list.listError = validation_error.New("The permissions must be at least one alphabetical character.")
	}
	return permission
}
