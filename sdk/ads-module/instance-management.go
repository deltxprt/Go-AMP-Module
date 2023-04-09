package ads_module

import (
	"bytes"
	"encoding/json"
	"github.com/deltxprt/Go-AMP-Module/internal/data"
	"io"
	"net/http"
)

func GetInstance(url, instanceId string, loginData *data.ReceiveLogin) (*Instance, error) {
	instanceUrl := url + "/API/ADSModule/GetInstance"

	type InstanceBody struct {
		InstanceId string `json:"instanceId"`
		SESSIONID  string `json:"SESSIONID"`
	}

	body := &InstanceBody{
		InstanceId: instanceId,
		SESSIONID:  loginData.SessionId,
	}

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(body)

	request, err := http.NewRequest("POST", instanceUrl, buffer)

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
