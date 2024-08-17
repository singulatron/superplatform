package sdk

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"

	usertypes "github.com/singulatron/singulatron/localtron/internal/services/user/types"
)

func DecodeJWT(tokenString string, publicKeyString string) (*usertypes.Claims, error) {
	publicKey, err := PublicKeyFromString(publicKeyString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get public key from string")
	}

	token, err := jwt.ParseWithClaims(tokenString, &usertypes.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT: %v", err)
	}

	if claims, ok := token.Claims.(*usertypes.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid JWT token")
}

func PublicKeyFromString(publicKeyPem string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyPem))
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	// Type assertion to convert from interface{} to *rsa.PublicKey
	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	return rsaPub, nil
}
