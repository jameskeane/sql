package sql

import (
	"http"
)

type Driver interface {
	Connect(url *http.URL) (Connection, Error)
}


