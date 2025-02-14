package configfetcher

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GithubConfig struct {
	Pi   string `json:"pi"`
	Port string `json:"port"`
}

func fetchGithubConfig() (*GithubConfig, error) {

	resp, err := http.Get("https://raw.githubusercontent.com/palash117/Globals/refs/heads/main/config.json")

	if err != nil {
		fmt.Println("http err is ", err)
		return nil, err
	}

	infoResp, err := getFormattedGithubConfig(resp)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("err is ", err)
		return nil, err
	}
	return infoResp, err
}

func GetPiIP() (string, error) {
	infoResp, err := fetchGithubConfig()
	if err != nil {
		return "", fmt.Errorf("Github config err %+v", err)
	}
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
func GetPiPort() (string, error) {
	infoResp, err := fetchGithubConfig()
	if err != nil {
		return "", fmt.Errorf("Github config err %+v", err)
	}
	return infoResp.Port, nil

}
