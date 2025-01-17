package auth

import "fmt"

// This file is generated by goenums; DO NOT EDIT

// Role wraps an enum value of type role to enforce type safety.
type Role struct {
	role
}

// role is the underlying type of Role and should not be used directly.
type role int

const (
	unknownRole role = iota
	adminRole
	userRole
	guestRole
)

var (
	strRoleArray = [...]string{
		adminRole: "ADMIN",
		userRole:  "USER",
		guestRole: "GUEST",
	}

	typeRoleMap = map[string]role{
		"ADMIN": adminRole,
		"USER":  userRole,
		"GUEST": guestRole,
	}
)

// String returns the string representation of the enum value
func (r role) String() string {
	return strRoleArray[r]
}

// ParseRole attempts to parse the given value into Role.
// It supports string, fmt.Stringer, int, int64, and int32.
// If the value is not a valid Role or the value is not a supported type, it will return the enums unknown value (unknownRole).
func ParseRole(a any) Role {
	switch v := a.(type) {
	case Role:
		return v
	case string:
		return Role{stringToRole(v)}
	case fmt.Stringer:
		return Role{stringToRole(v.String())}
	case int:
		return Role{role(v)}
	case int64:
		return Role{role(int(v))}
	case int32:
		return Role{role(int(v))}
	}
	return Role{unknownRole}
}

// stringToRole attempts to parse the given string into Role.
// If the value is not a valid Role, it will return the enums unknown value (unknownRole).
func stringToRole(s string) role {
	if v, ok := typeRoleMap[s]; ok {
		return v
	}
	return unknownRole
}

// IsValid returns true if the enum value is valid.
// The unknown value (unknownRole) is not considered valid.
func (r role) IsValid() bool {
	return r >= role(1) && r <= role(len(strRoleArray))
}

// RoleContainer contains all possible values of type Role.
type RoleContainer struct {
	UNKNOWN Role
	ADMIN   Role
	USER    Role
	GUEST   Role
}

// Roles is a global instance of RoleContainer that contains all possible values of type Role.
var Roles = RoleContainer{
	UNKNOWN: Role{unknownRole},
	ADMIN:   Role{adminRole},
	USER:    Role{userRole},
	GUEST:   Role{guestRole},
}

// All returns a slice containing all possible values of type Role.
func (c RoleContainer) All() []Role {
	return []Role{
		c.ADMIN,
		c.USER,
		c.GUEST,
	}
}

// MarshalJSON returns the JSON representation of the enum value.
func (r *Role) MarshalJSON() ([]byte, error) {
	return []byte(`"` + r.String() + `"`), nil
}

// UnmarshalJSON parses the JSON representation of the enum value.
func (r *Role) UnmarshalJSON(b []byte) error {
	*r = ParseRole(string(b))
	return nil
}
