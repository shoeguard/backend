package customErrors

type CustomError string

const (
	UnknownError         = CustomError("UnknownError")
	FormError            = CustomError("FormError")
	ParamterError        = CustomError("ParamterError")
	PhoneNumberDuplicate = CustomError("PhoneNumberDuplicate")
)
