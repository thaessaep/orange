package models

import "time"

type News struct {
	Date              time.Time `json:"date""`
	Rate              int64     `json:"rate""`
	CompaniesAffected []string  `json:"companiesAffected"`
}
