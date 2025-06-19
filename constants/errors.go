package constants

import "errors"

var (
	ErrorNotFound     = errors.New("Error Record Not Found")
	ErrorRecordExists = errors.New("Error Record Already Exists")
)
