package configfetcher

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GithubConfig struct {
	Pi string `json:"pi"`
}

func GetPiIP() (string, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/palash117/Globals/refs/heads/main/config.json")

	if err != nil {
		fmt.Println("http err is ", err)
		return "", err
	}

	infoResp, err := getFormattedGithubConfig(resp)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("err is ", err)
		return "", err
	}
	printable, _ := json.MarshalIndent(infoResp, "", "")
	fmt.Println(fmt.Sprintf("%s", printable))
	return infoResp.Pi, nil
}

func getFormattedGithubConfig(httpResp *http.Response) (*GithubConfig, error) {
	body := httpResp.Body
	decoder := json.NewDecoder(body)

	resp := GithubConfig{}
	err := decoder.Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
