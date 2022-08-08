package qiscusResponse

import (
	"io/ioutil"
	"net/http"
)

type GetConversationHistory struct {
	Data struct {
		DownloadUrl string `json:"download_url"`
	} `json:"data"`
}

func (c *GetConversationHistory) DownloadConversationHistory() (*string, error) {
	response, err := http.Get(c.Data.DownloadUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	responseString := string(responseData)

	return &responseString, nil
}
