package config

type ErrorCode string

const (
	NameLengthError        ErrorCode = "name length is 1 to 100 characters"
	NameExistsError        ErrorCode = "todo name existed"
	TodoListNullError      ErrorCode = "todo list is null"
	TodoNameNotExistsError ErrorCode = "todo name is not exists"
)
