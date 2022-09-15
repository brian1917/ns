package ns

import (
	"encoding/json"
	"fmt"
)

type LoginReqBody struct {
	LoginInfo Login `json:"login"`
}
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRespBody struct {
	Errorcode int    `json:"errorcode"`
	Message   string `json:"message"`
	Severity  string `json:"severity"`
	Sessionid string `json:"sessionid"`
}

func (n *NetScaler) Login() (APIResp, error) {

	// Create the login request
	loginReq := LoginReqBody{LoginInfo: Login{Username: n.User, Password: n.Password}}
	loginReqBody, err := json.Marshal(loginReq)
	if err != nil {
		return APIResp{}, fmt.Errorf("error marshaling login request - %s", err)
	}

	// Login to the netscaler
	api, err := n.API("login", "POST", loginReqBody)
	if err != nil {
		return api, fmt.Errorf("error making POST to login - %s", err)
	}
	// Unmarshall the response
	var resp LoginRespBody
	json.Unmarshal([]byte(api.RespBody), &resp)

	// If we don't get a 201, return an error
	if api.StatusCode != 201 {
		return api, fmt.Errorf("expected 201. received %d - %s", api.StatusCode, api.RespBody)
	}

	// Set the session ID and return
	n.SessionID = resp.Sessionid
	return api, nil
}
