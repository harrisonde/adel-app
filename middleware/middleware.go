package middleware

import (
	"myapp/data"

	"github.com/harrisonde/adele-framework"
)

type Middleware struct {
	App    *adele.Adele
	Models data.Models
}
