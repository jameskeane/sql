package sql

type Error interface {
	String() string
}

type strError struct {
	msg string
}

func (self *strError) String() string {
	return self.msg
}

func NewError(msg string) Error {
	s := new(strError)
	s.msg = msg
	return s
}

var (
	Busy Error = NewError("Database connection is busy")
	InvalidBindType Error = NewError("Attempted to bind an unrecognized type")
)

