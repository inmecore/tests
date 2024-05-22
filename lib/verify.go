package lib

var (
	UserVerify = Rules{
		"Username": {NotEmpty()},
		"Password": {NotEmpty()},
	}
	CodeVerify = Rules{
		"Code": {RegexpMatch("^\\d{6}$")},
	}
)
