package api

type Account struct {
	ID      int    `json:"-"`
	Type    string `json:"type"`
	Expires string `json:"expires"`
}
