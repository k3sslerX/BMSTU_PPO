package shared

type Error string

func (e Error) Error() string {
	return string(e)
}

const ErrorNotFound = Error("not found")
const ErrorPermissionDenied = Error("permission denied")
const ErrorInvalidToken = Error("invalid token")
const ErrorUserAlreadyExists = Error("user already exists")
const ErrorIncorrectPassword = Error("incorrect password")
