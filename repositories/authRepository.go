package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/lorezi/golang-bank-app/logger"
)

func NewAuthRepository() RemoteAuthRepository {
	return RemoteAuthRepository{}
}

type RemoteAuthRepository struct {
}

func (r RemoteAuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) bool {
	u := buildVerifyURL(token, routeName, vars)

	res, err := http.Get(u)

	if err != nil {
		logger.Error("error while sending... " + err.Error())
		return false
	}

	m := map[string]bool{}

	if err := json.NewDecoder(res.Body).Decode(&m); err != nil {
		logger.Error("Error while decoding response from auth server:" + err.Error())
		return false
	}
	return m["isAuthorized"]

}

/*
  This will generate a url for token verification in the below format
  /auth/verify?token={token string}
              &routeName={current route name}
              &customer_id={customer id from the current route}
              &account_id={account id from current route if available}
  Sample: /auth/verify?token=aaaa.bbbb.cccc&routeName=MakeTransaction&customer_id=2000&account_id=95470
*/
func buildVerifyURL(token string, routeName string, vars map[string]string) string {

	server := os.Getenv("AUTH_SERVER")
	port := os.Getenv("AUTH_SERVER_PORT")
	addr := fmt.Sprintf("%s:%s", server, port)
	u := url.URL{Host: addr, Path: "/auth/verify", Scheme: "http"}

	q := u.Query()
	q.Add("token", token)
	q.Add("routName", routeName)
	for i, v := range vars {
		q.Add(i, v)
	}
	u.RawQuery = q.Encode()

	return u.String()

}
