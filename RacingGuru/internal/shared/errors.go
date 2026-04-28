package shared

type Error string

func (e Error) Error() string {
	return string(e)
}

const ErrorNotFound = Error("not found")
const ErrorPermissionDenied = Error("permission denied")
const ErrorInvalidToken = Error("invalid token")
const ErrorTokenExpired = Error("token expired")
const ErrorUserAlreadyExists = Error("user already exists")
const ErrorAdminAlreadyExists = Error("admin already exists")
const ErrorAdminSecretAlreadyIssued = Error("admin secret already issued")
const ErrorInvalidAdminSecret = Error("invalid admin secret")
const ErrorIncorrectPassword = Error("incorrect password")
const ErrorInvalidData = Error("invalid data")
