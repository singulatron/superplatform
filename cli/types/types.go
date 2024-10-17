package types

type Config struct {
	Environments map[string]*Environment `json:"environments"`
}

type Environment struct {
	URL         string  `json:"url"`
	ShortName   string  `json:"shortname"`
	Description string  `json:"description"`
	Users       []*User `json:"users"`
}

type User struct {
	Slug  string `json:"slug"`
	Token string `json:"token"`
}
