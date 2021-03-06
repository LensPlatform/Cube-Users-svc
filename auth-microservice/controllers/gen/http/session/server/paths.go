// Code generated by goa v3.0.6, DO NOT EDIT.
//
// HTTP request path constructors for the session service.
//
// Command:
// $ goa gen github.com/anshap1719/authentication/design

package server

// RefreshSessionPath returns the URL path to the session service refresh HTTP endpoint.
func RefreshSessionPath() string {
	return "/auth/session"
}

// LogoutSessionPath returns the URL path to the session service logout HTTP endpoint.
func LogoutSessionPath() string {
	return "/auth/logout"
}

// LogoutOtherSessionPath returns the URL path to the session service logout-other HTTP endpoint.
func LogoutOtherSessionPath() string {
	return "/auth/logout/all"
}

// LogoutSpecificSessionPath returns the URL path to the session service logout-specific HTTP endpoint.
func LogoutSpecificSessionPath() string {
	return "/auth/logout/:session-id"
}

// GetSessionsSessionPath returns the URL path to the session service get-sessions HTTP endpoint.
func GetSessionsSessionPath() string {
	return "/auth/sessions"
}

// RedeemTokenSessionPath returns the URL path to the session service redeemToken HTTP endpoint.
func RedeemTokenSessionPath() string {
	return "/auth/token"
}

// CleanSessionsSessionPath returns the URL path to the session service clean-sessions HTTP endpoint.
func CleanSessionsSessionPath() string {
	return "/auth/clean/sessions"
}

// CleanLoginTokenSessionPath returns the URL path to the session service clean-login-token HTTP endpoint.
func CleanLoginTokenSessionPath() string {
	return "/auth/clean/token/login"
}

// CleanMergeTokenSessionPath returns the URL path to the session service clean-merge-token HTTP endpoint.
func CleanMergeTokenSessionPath() string {
	return "/auth/clean/token/merge"
}
