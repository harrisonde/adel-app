package middleware

import (
	"myapp/data"

	"github.com/harrisonde/adel"
)

type Middleware struct {
	App    *adel.Adel
	Models data.Models
}
