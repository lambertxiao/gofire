package core

type ErrTimeout struct{}

func (e ErrTimeout) Error() string {
	return "timeout"
}
