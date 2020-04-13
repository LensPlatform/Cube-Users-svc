package design

import (
	. "goa.design/goa/v3/dsl"
)

// Credentials defines the credentials to use for authenticating to service methods.
var Credentials = Type("Creds", func() {
	Field(1, "jwt", String, "JWT token", func() {
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	})
	Field(2, "api_key", String, "API Key", func() {
		Example("abcdef12345")
	})
	Field(3, "oauth_token", String, "OAuth2 token", func() {
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	})
	Required("jwt", "api_key", "oauth_token")
})

var User = Type("User", func() {
	Field(1, "body", Any)
})

var Team = Type("Team", func() {
	Field(1, "body", Any)
})

var Group = Type("Group", func() {
	Field(1, "body", Any)
})

var Subscription = Type("Subscription", func() {
	Field(1, "body", Any)
})

var Profile = Type("Profile", func() {
	Field(1, "body", Any)
})
