package auth

// IsActive checks if the status is active.
func (a *AccountStatus) IsActive() bool {
	return (*a).accountStatus == activeAccountStatus
}

// CanLogin determines if the status allows for user login.
func (a *AccountStatus) CanLogin() bool {
	switch a.accountStatus {
	case activeAccountStatus:
		return true
	case inactiveAccountStatus, suspendedAccountStatus:
		return false
	default:
		return false
	}
}

// StatusMessage provides a user-friendly message describing the status.
func (a *AccountStatus) StatusMessage() string {
	switch a.accountStatus {
	case activeAccountStatus:
		return "Your account is active."
	case inactiveAccountStatus:
		return "Your account is inactive. Please contact support."
	case suspendedAccountStatus:
		return "Your account has been suspended. Please contact support for more information."
	default:
		return "Unknown account status."
	}
}
