package middleware

import (
	"myapp/data"

	"git.int.86labs.cloud/harrisonde/adele-framework"
)

type Middleware struct {
	App    *adele.Adele
	Models data.Models
}
