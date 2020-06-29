package almaerror

// SystemError .
type SystemError struct {
	Err         error
	StatusCode  int32
	MessageCode string
	Params      []interface{}
}

// Error .
func (se *SystemError) Error() string {
	return se.Err.Error()
}

// LogicError .
type LogicError struct {
	StatusCode  int32
	MessageCode string
	Params      []interface{}
}

// Error .
func (le *LogicError) Error() string {
	return le.MessageCode
}
