package models

type VerificationCode struct {
	UserID int    `json:"UserID"`
	Code   string `json:"Code"`
}
