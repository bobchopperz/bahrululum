package constants

type Role string

const (
	RoleUser   Role = "user"
	RoleMentor Role = "mentor"
	RoleAdmin  Role = "admin"
)

var AllRoles = []Role{
	RoleUser,
	RoleMentor,
	RoleAdmin,
}

func (r Role) String() string {
	return string(r)
}

func IsValidRole(role string) bool {
	for _, validRole := range AllRoles {
		if validRole.String() == role {
			return true
		}
	}
	return false
}

func ParseRole(role string) (Role, bool) {
	for _, validRole := range AllRoles {
		if validRole.String() == role {
			return validRole, true
		}
	}
	return "", false
}