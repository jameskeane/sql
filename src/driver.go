package sql

type Driver interface {
	Connect(url *DSN) (Connection, Error)
}


