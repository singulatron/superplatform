package sdk

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

type Claims struct {
	UserId string `json:"sui"` // `sui`: singulatron user ids
	Slug   string `json:"slu"` // `sri`: singulatron slug
	// Contacts []Contact // `sco`: singulatron contacts
	RoleIds []string `json:"sri"` // `sri`: singulatron role ids
	jwt.RegisteredClaims
}

type Credential struct {
	Slug     string `json:"slug,omitempty"`
	Contact  string `json:"contact,omitempty"`
	Password string `json:"password,omitempty"`
}

func (c *Credential) GetId() string {
	return c.Contact
}

func DecodeJWT(tokenString string, publicKeyString string) (*Claims, error) {
	publicKey, err := PublicKeyFromString(publicKeyString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get public key from string")
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT: %v", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid JWT token")
}

func TokenFromRequest(r *http.Request) (string, bool) {
	authHeader := r.Header.Get("Authorization")
	authHeader = strings.Replace(authHeader, "Bearer ", "", 1)

	if authHeader == "" || authHeader == "Bearer" {
		return "", false
	}

	return authHeader, true
}

// OrganizationIDsFromRoleIDs extracts organization slugs from role IDs in the format "user-svc:org:{org-slug}:role"
func OrganizationSlugsFromRoleIDs(roleIDs []string) []string {
	ret := []string{}
	for _, roleID := range roleIDs {
		// Check if the roleID starts with the expected prefix
		if strings.HasPrefix(roleID, "user-svc:org:") {
			// Remove the prefix "user-svc:org:"
			trimmedRole := strings.TrimPrefix(roleID, "user-svc:org:")
			// Split the remaining string by ':'
			parts := strings.Split(trimmedRole, ":")
			// Ensure there are at least two parts: the org-slug and a role
			if len(parts) >= 2 {
				orgSlug := parts[0]
				// Strip the curly braces if they are part of the orgSlug
				orgSlug = strings.Trim(orgSlug, "{}")
				ret = append(ret, orgSlug)
			}
		}
	}
	return ret
}

func StaticRoles(roleIDs []string) []string {
	ret := []string{}
	for _, roleID := range roleIDs {
		if !strings.Contains(roleID, "{") && !strings.Contains(roleID, "}") {
			ret = append(ret, roleID)
		}
	}

	return ret
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
