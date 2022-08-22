package standard_permission

type StandardPermission string

const (
	Create      = "create"
	Read        = "read"
	Update      = "update"
	Delete      = "delete"
	CreateChild = "create-child"
	ReadChild   = "read-child"
	UpdateChild = "update-child"
	DeleteChild = "delete-child"
)
