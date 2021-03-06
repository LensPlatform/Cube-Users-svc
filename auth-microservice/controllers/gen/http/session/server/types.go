// Code generated by goa v3.0.6, DO NOT EDIT.
//
// session HTTP server types
//
// Command:
// $ goa gen github.com/anshap1719/authentication/design

package server

import (
	session "github.com/anshap1719/authentication/controllers/gen/session"
	sessionviews "github.com/anshap1719/authentication/controllers/gen/session/views"
	goa "goa.design/goa/v3/pkg"
)

// RedeemTokenRequestBody is the type of the "session" service "redeemToken"
// endpoint HTTP request body.
type RedeemTokenRequestBody struct {
	// A merge token for merging into an account
	Token     *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
	UserAgent *string `form:"User-Agent,omitempty" json:"User-Agent,omitempty" xml:"User-Agent,omitempty"`
}

// GetSessionsResponseBody is the type of the "session" service "get-sessions"
// endpoint HTTP response body.
type GetSessionsResponseBody struct {
	CurrentSession *SessionResponseBody          `form:"currentSession,omitempty" json:"currentSession,omitempty" xml:"currentSession,omitempty"`
	OtherSessions  SessionResponseBodyCollection `form:"otherSessions,omitempty" json:"otherSessions,omitempty" xml:"otherSessions,omitempty"`
}

// SessionResponseBody is used to define fields on response body types.
type SessionResponseBody struct {
	// Unique unchanging session ID
	ID string `form:"id" json:"id" xml:"id"`
	// ID of the user this session is for
	UserID string `form:"userId" json:"userId" xml:"userId"`
	// Time that this session was last used
	LastUsed string `form:"lastUsed" json:"lastUsed" xml:"lastUsed"`
	// The browser and browser version connected with this session
	Browser string `form:"browser" json:"browser" xml:"browser"`
	// The OS of the system where this session was used
	Os string `form:"os" json:"os" xml:"os"`
	// The last IP address where this session was used
	IP string `form:"ip" json:"ip" xml:"ip"`
	// A humanReadable string describing the last known location of the session
	Location string `form:"location" json:"location" xml:"location"`
	// The latitude of the last known location of the session
	Latitude string `form:"latitude" json:"latitude" xml:"latitude"`
	// The longitude of the last known location of the session
	Longitude string `form:"longitude" json:"longitude" xml:"longitude"`
	// Whether the session was from a mobile device
	IsMobile bool `form:"isMobile" json:"isMobile" xml:"isMobile"`
	// The URL of the Google map to show the location, suitable for using in an img
	// tag
	MapURL string `form:"mapUrl" json:"mapUrl" xml:"mapUrl"`
}

// SessionResponseBodyCollection is used to define fields on response body
// types.
type SessionResponseBodyCollection []*SessionResponseBody

// NewGetSessionsResponseBody builds the HTTP response body from the result of
// the "get-sessions" endpoint of the "session" service.
func NewGetSessionsResponseBody(res *sessionviews.AllSessionsView) *GetSessionsResponseBody {
	body := &GetSessionsResponseBody{}
	if res.CurrentSession != nil {
		body.CurrentSession = marshalSessionviewsSessionViewToSessionResponseBody(res.CurrentSession)
	}
	if res.OtherSessions != nil {
		body.OtherSessions = make([]*SessionResponseBody, len(res.OtherSessions))
		for i, val := range res.OtherSessions {
			body.OtherSessions[i] = marshalSessionviewsSessionViewToSessionResponseBody(val)
		}
	}
	return body
}

// NewRefreshPayload builds a session service refresh endpoint payload.
func NewRefreshPayload(xSession string, aPIKey *string) *session.RefreshPayload {
	return &session.RefreshPayload{
		XSession: &xSession,
		APIKey:   aPIKey,
	}
}

// NewLogoutPayload builds a session service logout endpoint payload.
func NewLogoutPayload(authorization *string, xSession *string, aPIKey *string) *session.LogoutPayload {
	return &session.LogoutPayload{
		Authorization: authorization,
		XSession:      xSession,
		APIKey:        aPIKey,
	}
}

// NewLogoutOtherPayload builds a session service logout-other endpoint payload.
func NewLogoutOtherPayload(authorization *string, xSession *string, aPIKey *string) *session.LogoutOtherPayload {
	return &session.LogoutOtherPayload{
		Authorization: authorization,
		XSession:      xSession,
		APIKey:        aPIKey,
	}
}

// NewLogoutSpecificPayload builds a session service logout-specific endpoint
// payload.
func NewLogoutSpecificPayload(sessionID string, authorization *string, xSession *string, aPIKey *string) *session.LogoutSpecificPayload {
	return &session.LogoutSpecificPayload{
		SessionID:     &sessionID,
		Authorization: authorization,
		XSession:      xSession,
		APIKey:        aPIKey,
	}
}

// NewGetSessionsPayload builds a session service get-sessions endpoint payload.
func NewGetSessionsPayload(authorization *string, xSession *string, aPIKey *string) *session.GetSessionsPayload {
	return &session.GetSessionsPayload{
		Authorization: authorization,
		XSession:      xSession,
		APIKey:        aPIKey,
	}
}

// NewRedeemTokenPayload builds a session service redeemToken endpoint payload.
func NewRedeemTokenPayload(body *RedeemTokenRequestBody, aPIKey *string) *session.RedeemTokenPayload {
	v := &session.RedeemTokenPayload{
		Token:     *body.Token,
		UserAgent: body.UserAgent,
	}
	v.APIKey = aPIKey
	return v
}

// NewCleanSessionsPayload builds a session service clean-sessions endpoint
// payload.
func NewCleanSessionsPayload(authorization *string, xSession *string, aPIKey *string) *session.CleanSessionsPayload {
	return &session.CleanSessionsPayload{
		Authorization: authorization,
		XSession:      xSession,
		APIKey:        aPIKey,
	}
}

// NewCleanLoginTokenPayload builds a session service clean-login-token
// endpoint payload.
func NewCleanLoginTokenPayload(authorization *string, xSession *string, aPIKey *string) *session.CleanLoginTokenPayload {
	return &session.CleanLoginTokenPayload{
		Authorization: authorization,
		XSession:      xSession,
		APIKey:        aPIKey,
	}
}

// NewCleanMergeTokenPayload builds a session service clean-merge-token
// endpoint payload.
func NewCleanMergeTokenPayload(authorization *string, xSession *string, aPIKey *string) *session.CleanMergeTokenPayload {
	return &session.CleanMergeTokenPayload{
		Authorization: authorization,
		XSession:      xSession,
		APIKey:        aPIKey,
	}
}

// ValidateRedeemTokenRequestBody runs the validations defined on
// RedeemTokenRequestBody
func ValidateRedeemTokenRequestBody(body *RedeemTokenRequestBody) (err error) {
	if body.Token == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("token", "body"))
	}
	if body.Token != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.token", *body.Token, goa.FormatUUID))
	}
	return
}

// ValidateSessionResponseBody runs the validations defined on
// SessionResponseBody
func ValidateSessionResponseBody(body *SessionResponseBody) (err error) {
	err = goa.MergeErrors(err, goa.ValidateFormat("body.lastUsed", body.LastUsed, goa.FormatDateTime))

	return
}

// ValidateSessionResponseBodyCollection runs the validations defined on
// SessionResponseBodyCollection
func ValidateSessionResponseBodyCollection(body SessionResponseBodyCollection) (err error) {
	for _, e := range body {
		if e != nil {
			if err2 := ValidateSessionResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}
