package token

import (
	"fmt"
	"github.com/credstack/credstack-lib/proto/response"
	"github.com/credstack/credstack-lib/secret"
	"github.com/golang-jwt/jwt/v5"
)

/*
generateHS256 - Generates arbitrary HS256 tokens with the claims that are passed as an argument to the function. It is
expected that a base64 encoded secret string (like the ones generated from secret.RandString) is used as the secret here.
When used with ClientCredentials flow, the client secret is expected here. As a result, the KID field is not added to the
header with this function either as both the issuing and validating party must both know the client secret
*/
func generateHS256(clientSecret string, claims jwt.RegisteredClaims, expiresIn uint32) (*response.TokenResponse, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	/*
		Unlike RS256 tokens, the client secret is simply used to sign the token with SigningMethodHS256. This provides
		s shared secret that both the issuer and the validator a shared secret that both parties can agree on. Client
		secrets are issued pretty simply, as its really just a base64 encoded version of the byte result of rand.Read
	*/
	decodedBytes := []byte(clientSecret)
	decoded, err := secret.DecodeBase64(decodedBytes, uint32(len(decodedBytes)))
	if err != nil {
		return nil, err
	}

	/*
		Then we just sign the token with the decoded secret.
	*/
	signedString, err := token.SignedString(decoded)
	if err != nil {
		return nil, fmt.Errorf("%w (%v)", ErrFailedToSignToken, err)
	}

	resp, err := MarshalTokenResponse(signedString, expiresIn)
	if err != nil {
		return nil, fmt.Errorf("%w (%v)", ErrMarshalTokenResponse, err)
	}

	return resp, nil
}
