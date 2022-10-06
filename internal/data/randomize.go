// Filename: internal/data/randomize.go

package data

import (
	"time"

	"AWDquiz2.ryanarmstrong.net/internal/validator"
)

type Information struct {
	CreatedAt time.Time `json:"-"` // doesn't display to client
	Info      string    `json:"info"`
}

func ValidateInfo(v *validator.Validator, info *Information) {
	// Use the Check() method to execute our validation checks
	v.Check(info.Info != "", "info", "must be provided")
	v.Check(len(info.Info) <= 1000, "info", "must not be more than 1000 bytes long")
}
