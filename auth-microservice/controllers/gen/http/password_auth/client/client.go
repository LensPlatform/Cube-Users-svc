// Code generated by goa v3.0.6, DO NOT EDIT.
//
// password-auth client HTTP transport
//
// Command:
// $ goa gen github.com/anshap1719/authentication/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the password-auth service endpoint HTTP clients.
type Client struct {
	// Register Doer is the HTTP client used to make requests to the register
	// endpoint.
	RegisterDoer goahttp.Doer

	// Login Doer is the HTTP client used to make requests to the login endpoint.
	LoginDoer goahttp.Doer

	// Remove Doer is the HTTP client used to make requests to the remove endpoint.
	RemoveDoer goahttp.Doer

	// ChangePassword Doer is the HTTP client used to make requests to the
	// change-password endpoint.
	ChangePasswordDoer goahttp.Doer

	// Reset Doer is the HTTP client used to make requests to the reset endpoint.
	ResetDoer goahttp.Doer

	// ConfirmReset Doer is the HTTP client used to make requests to the
	// confirm-reset endpoint.
	ConfirmResetDoer goahttp.Doer

	// CheckEmailAvailable Doer is the HTTP client used to make requests to the
	// check-email-available endpoint.
	CheckEmailAvailableDoer goahttp.Doer

	// CheckPhoneAvailable Doer is the HTTP client used to make requests to the
	// check-phone-available endpoint.
	CheckPhoneAvailableDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the password-auth service
// servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		RegisterDoer:            doer,
		LoginDoer:               doer,
		RemoveDoer:              doer,
		ChangePasswordDoer:      doer,
		ResetDoer:               doer,
		ConfirmResetDoer:        doer,
		CheckEmailAvailableDoer: doer,
		CheckPhoneAvailableDoer: doer,
		CORSDoer:                doer,
		RestoreResponseBody:     restoreBody,
		scheme:                  scheme,
		host:                    host,
		decoder:                 dec,
		encoder:                 enc,
	}
}

// Register returns an endpoint that makes HTTP requests to the password-auth
// service register server.
func (c *Client) Register() goa.Endpoint {
	var (
		encodeRequest  = EncodeRegisterRequest(c.encoder)
		decodeResponse = DecodeRegisterResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRegisterRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RegisterDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("password-auth", "register", err)
		}
		return decodeResponse(resp)
	}
}

// Login returns an endpoint that makes HTTP requests to the password-auth
// service login server.
func (c *Client) Login() goa.Endpoint {
	var (
		encodeRequest  = EncodeLoginRequest(c.encoder)
		decodeResponse = DecodeLoginResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildLoginRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.LoginDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("password-auth", "login", err)
		}
		return decodeResponse(resp)
	}
}

// Remove returns an endpoint that makes HTTP requests to the password-auth
// service remove server.
func (c *Client) Remove() goa.Endpoint {
	var (
		encodeRequest  = EncodeRemoveRequest(c.encoder)
		decodeResponse = DecodeRemoveResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRemoveRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RemoveDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("password-auth", "remove", err)
		}
		return decodeResponse(resp)
	}
}

// ChangePassword returns an endpoint that makes HTTP requests to the
// password-auth service change-password server.
func (c *Client) ChangePassword() goa.Endpoint {
	var (
		encodeRequest  = EncodeChangePasswordRequest(c.encoder)
		decodeResponse = DecodeChangePasswordResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildChangePasswordRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ChangePasswordDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("password-auth", "change-password", err)
		}
		return decodeResponse(resp)
	}
}

// Reset returns an endpoint that makes HTTP requests to the password-auth
// service reset server.
func (c *Client) Reset() goa.Endpoint {
	var (
		encodeRequest  = EncodeResetRequest(c.encoder)
		decodeResponse = DecodeResetResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildResetRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ResetDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("password-auth", "reset", err)
		}
		return decodeResponse(resp)
	}
}

// ConfirmReset returns an endpoint that makes HTTP requests to the
// password-auth service confirm-reset server.
func (c *Client) ConfirmReset() goa.Endpoint {
	var (
		encodeRequest  = EncodeConfirmResetRequest(c.encoder)
		decodeResponse = DecodeConfirmResetResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildConfirmResetRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ConfirmResetDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("password-auth", "confirm-reset", err)
		}
		return decodeResponse(resp)
	}
}

// CheckEmailAvailable returns an endpoint that makes HTTP requests to the
// password-auth service check-email-available server.
func (c *Client) CheckEmailAvailable() goa.Endpoint {
	var (
		encodeRequest  = EncodeCheckEmailAvailableRequest(c.encoder)
		decodeResponse = DecodeCheckEmailAvailableResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCheckEmailAvailableRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CheckEmailAvailableDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("password-auth", "check-email-available", err)
		}
		return decodeResponse(resp)
	}
}

// CheckPhoneAvailable returns an endpoint that makes HTTP requests to the
// password-auth service check-phone-available server.
func (c *Client) CheckPhoneAvailable() goa.Endpoint {
	var (
		encodeRequest  = EncodeCheckPhoneAvailableRequest(c.encoder)
		decodeResponse = DecodeCheckPhoneAvailableResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCheckPhoneAvailableRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CheckPhoneAvailableDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("password-auth", "check-phone-available", err)
		}
		return decodeResponse(resp)
	}
}
