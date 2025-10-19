package middleware

import (
	"os"
	"regexp"
	"strings"
)

type FireSPMiddleware struct {
}

func (mid *FireSPMiddleware) AuthEnvs() (*CredClientEnv, error) {
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	var envUser string = "FIRESIMPLE_DEFAULT_USER"
	var envPassword string = "FIRESIMPLE_DEFAULT_PASSWORD"

	user := os.Getenv(envUser)
	if user == "" {
		return nil, &NotFoundEnv{Field: envUser}
	}
	isAuth := re.MatchString(user)
	if !isAuth {
		return nil, &InvalidFormaterEnv{Field: envUser}
	}

	password := os.Getenv(envPassword)
	if password == "" {
		return nil, &NotFoundEnv{Field: envPassword}
	}
	isAuth = re.MatchString(password)
	if !isAuth {
		return nil, &InvalidFormaterEnv{Field: envPassword}
	}

	return &CredClientEnv{
		Password: password,
		User:     user,
	}, nil
}

// url default: frsp://user:password@firesimple:5412/
func (mid *FireSPMiddleware) AuthClient(cred *CredClientEnv, url string) error {

	withoutPrefix := strings.TrimPrefix(url, "frsp://")

	parts := strings.SplitN(withoutPrefix, "@", 2)

	userPass := parts[0]

	credURL := strings.SplitN(userPass, ":", 2)

	if credURL[0] != cred.User {
		return &NotAuthClient{
			Msg: "User different default",
		}
	}

	if credURL[1] != cred.Password {
		return &NotAuthClient{
			Msg: "Password different default",
		}
	}

	return nil
}
