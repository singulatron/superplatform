package user_svc

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UserId string `json:"sui"` // `sui`: singulatron user ids
	Slug   string `json:"slu"` // `sri`: singulatron slug
	// Contacts []Contact // `sco`: singulatron contacts
	RoleIds []string `json:"sri"` // `sri`: singulatron role ids
	jwt.RegisteredClaims
}
