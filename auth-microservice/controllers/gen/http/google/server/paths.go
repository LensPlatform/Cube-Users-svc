// Code generated by goa v3.0.6, DO NOT EDIT.
//
// HTTP request path constructors for the google service.
//
// Command:
// $ goa gen github.com/anshap1719/authentication/design

package server

// RegisterURLGooglePath returns the URL path to the google service register-url HTTP endpoint.
func RegisterURLGooglePath() string {
	return "/google/register-start"
}

// AttachToAccountGooglePath returns the URL path to the google service attach-to-account HTTP endpoint.
func AttachToAccountGooglePath() string {
	return "/google/attach"
}

// DetachFromAccountGooglePath returns the URL path to the google service detach-from-account HTTP endpoint.
func DetachFromAccountGooglePath() string {
	return "/google/detach"
}

// ReceiveGooglePath returns the URL path to the google service receive HTTP endpoint.
func ReceiveGooglePath() string {
	return "/google/receive"
}
