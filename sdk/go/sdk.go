package sdk

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/sony/sonyflake"
)

type Claims struct {
	UserId  string   `json:"sui"` // `sui`: singulatron user ids
	Slug    string   `json:"slu"` // `slu`: singulatron slug
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

var sonyFlake *sonyflake.Sonyflake

func init() {
	sonyFlake = sonyflake.NewSonyflake(sonyflake.Settings{})
	if sonyFlake == nil {
		panic("Sonyflake not created")
	}
}

const base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Id(prefix string) string {
	number, err := sonyFlake.NextID()
	if err != nil {
		panic(err)
	}

	if number == 0 {
		return string(base62[0])
	}

	b := make([]byte, 0)
	for number > 0 {
		remainder := number % 62
		number = number / 62
		b = append([]byte{base62[remainder]}, b...)
	}

	return prefix + "_" + string(b)
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

func Marshal(value any) *string {
	jsonBytes, _ := json.Marshal(value)

	v := string(jsonBytes)
	return &v
}
