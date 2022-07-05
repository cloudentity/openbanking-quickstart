package utils

import "github.com/golang-jwt/jwt/v4"

type JwtClaims map[string]interface{}

func (j *JwtClaims) Valid() error {
	return nil
}

func unpack(token string) (JwtClaims, error) {
	var (
		claims = JwtClaims{}
		parser jwt.Parser
		err    error
	)

	if _, _, err = parser.ParseUnverified(token, &claims); err != nil {
		return claims, err
	}

	return claims, nil
}

type ResponseClaims struct {
	State            string
	Code             string
	Error            string
	ErrorDescription string
}

func DecodeResponseToken(token string) (ResponseClaims, error) {
	var (
		claims         map[string]interface{}
		responseClaims ResponseClaims
		val            string
		ok             bool
		err            error
	)

	if claims, err = unpack(token); err != nil {
		return ResponseClaims{}, err
	}
	if claims["code"] != nil {
		if val, ok = claims["code"].(string); ok {
			responseClaims.Code = val
		}
	}
	if claims["state"] != nil {
		if val, ok = claims["state"].(string); ok {
			responseClaims.State = val
		}
	}
	if claims["error"] != nil {
		if val, ok = claims["error"].(string); ok {
			responseClaims.Error = val
		}
	}
	if claims["error_description"] != nil {
		if val, ok = claims["error_description"].(string); ok {
			responseClaims.ErrorDescription = val
		}
	}

	return responseClaims, nil
}
