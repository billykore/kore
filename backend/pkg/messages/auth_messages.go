package messages

const (
	InvalidUsernameOrPassword    = "Username or password was incorrect."
	LoginFailed                  = "Login failed."
	LogoutSucceed                = "Logout Succeed."
	LogoutFailed                 = "Logout failed."
	UserAlreadyLoggedOut         = "User already logged out."
	FailedLoginAttemptNotExpired = "Login failed. You can login again after 24 hours of last failed login attempt."
	MaxLoginAttemptReached       = "Login was failed 5 times. Please try again later."
)
