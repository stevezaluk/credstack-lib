package token

import (
	credstackError "github.com/credstack/credstack-lib/errors"
	"github.com/credstack/credstack-lib/proto/response"
)

// ErrMarshalTokenResponse - An error that gets returned
var ErrMarshalTokenResponse = credstackError.NewError(400, "ERR_MARSHAL_TOKEN_RESPONSE", "token: Failed to marshal token into token response")

/*
MarshalTokenResponse - Converts jwt.Token structures into response.TokenResponse structures so that they can be returned
effectively

TODO: Need support for id tokens and refresh tokens here
TODO: Expires in is not rendering properly, showing expiration instead of token lifetime
TODO: This function feels kind of clunky...
*/
func MarshalTokenResponse(accessToken string, expiration uint32) (*response.TokenResponse, error) {
	/*
		After we actually sign our token, we can quickly convert it back into a response.TokenResponse structure so that
		it can be returned from the API
	*/
	signedResponse := &response.TokenResponse{
		AccessToken:  accessToken,
		TokenType:    "Bearer",
		ExpiresIn:    expiration, // this is bad, not creating future support for custom expiration's
		RefreshToken: "",
		IdToken:      "",
		Scope:        "",
	}

	return signedResponse, nil
}
