package constants

import "errors"

var (
	ErrorNotFound       = errors.New("Error Record Not Found")
	ErrorRecordExists   = errors.New("Error Record Already Exists")
	ErrorOrderIdEmpty   = errors.New("Error Order Id Empty")
	ErrorArticleIdEmpty = errors.New("Error Article Id Empty")
)
