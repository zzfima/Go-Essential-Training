package part8

import (
	"encoding/json"
	"net/http"
)

// UserInfo describes github user
type UserInfo struct {
	Login         string `json:"login"`
	Name          string `json:"name"`
	NumberOfRepos int    `json:"public_repos"`
}

// GetUserInfo return UserInfo by url to profile in github
func GetUserInfo(url string) (UserInfo, error) {
	r, e := http.Get(url)
	if e != nil {
		return UserInfo{}, e
	}
	defer r.Body.Close()

	var user UserInfo
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&user)
	return user, nil
}
