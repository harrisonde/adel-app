package middleware

import (
	"myapp/data"

	"github.com/harrisonde/adele"
)

type Middleware struct {
	App    *adele.Adele
	Models data.Models
}
