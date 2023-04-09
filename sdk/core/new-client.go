package core

import (
	"bytes"
	"encoding/json"
	"github.com/deltxprt/Go-AMP-Module/internal/data"
	"io"
	"net/http"
)

func Login(url string, loginInfo *data.SendLogin) (*data.ReceiveLogin, error) {
	loginUrl := url + "/API/Core/Login"
	ReponseData := &data.ReceiveLogin{}

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(loginInfo)

	request, err := http.NewRequest("POST", loginUrl, buffer)
	request.Header.Set("accept", "application/json; charset=UTF-8")

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	err = json.Unmarshal(body, &ReponseData)

	if err != nil {
		return nil, err
	}

	return ReponseData, nil
}

func Logout(url string) error {
	logoutUrl := url + "/API/Core/Logout"

	request, err := http.NewRequest("POST", logoutUrl, nil)
	request.Header.Set("accept", "application/json; charset=UTF-8")

	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(request)

	if err != nil {
		return err
	}

	return nil
}
