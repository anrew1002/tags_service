package models

type Tag struct {
	ID    string `json:"tag_id" db:"id"`
	Alias string `json:"alias" db:"alias"`
	Pass  string
}

type Key struct {
	Login  string `json:"login" db:"login"`
	APIKey string `json:"api_key" db:"apikey"`
}
