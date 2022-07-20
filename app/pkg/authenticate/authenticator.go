package authenticate

import(
	"echosample/pkg/db/repository"
)

type Result bool
const (
	Success Result = true
	Fail Result = false
)

func Login(name string, pw string) Result {
	user := repository.FindUserByName(name)

	if user.Password == pw {
		return Success
	}

	return Fail
}