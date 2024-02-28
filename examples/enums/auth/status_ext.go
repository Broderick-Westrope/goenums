package auth

// IsActive checks if the status is active.
func (s Status) IsActive() bool {
	return s == Status{activeStatus}
}

// CanLogin determines if the status allows for user login.
func (s Status) CanLogin() bool {
	switch s.status {
	case activeStatus:
		return true
	case inactiveStatus, suspendedStatus:
		return false
	default:
		return false
	}
}

// StatusMessage provides a user-friendly message describing the status.
func (s Status) StatusMessage() string {
	switch s.status {
	case activeStatus:
		return "Your account is active."
	case inactiveStatus:
		return "Your account is inactive. Please contact support."
	case suspendedStatus:
		return "Your account has been suspended. Please contact support for more information."
	default:
		return "Unknown account status."
	}
}
