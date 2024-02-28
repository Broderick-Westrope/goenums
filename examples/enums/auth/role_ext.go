package auth

// CanAccessAdminPanel checks if the role has access to the admin panel.
func (r Role) CanAccessAdminPanel() bool {
	return r == Role{adminRole}
}

// Description provides a human-readable description of the role.
func (r Role) Description() string {
	switch r.role {
	case adminRole:
		return "Admins have full access to the system."
	case userRole:
		return "Users have access to standard features."
	case guestRole:
		return "Guests have limited access to features."
	default:
		return "Unknown role"
	}
}

// IsHigherThan checks if the role has higher privileges than the specified role.
func (r Role) IsHigherThan(other Role) bool {
	return r.role < other.role
}
