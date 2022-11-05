package domain

import (
	"encoding/json"
	"fmt"
	"github.com/y-sugiyama654/banking-lib/logger"
	"net/http"
	"net/url"
)

type AuthRepository interface {
	IsAuthorized(string, string, map[string]string) bool
}

type RemoteAuthRepository struct {
}

func (r RemoteAuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) bool {
	u := buildVerifyURL(token, routeName, vars)
	fmt.Println(u)

	if response, err := http.Get(u); err != nil {
		logger.Error("Error while sending...." + err.Error())
		return false
	} else {
		m := map[string]bool{}
		if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
			logger.Error("Error while decoding response from server: " + err.Error())
			return false
		}
		return m["isAuthorized"]
	}
}

func buildVerifyURL(token string, routeName string, vars map[string]string) string {
	u := url.URL{
		Host:   "localhost:8101",
		Path:   "/auth/verify",
		Scheme: "http",
	}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	return u.String()
}

func NewAuthRepository() RemoteAuthRepository {
	return RemoteAuthRepository{}
}
