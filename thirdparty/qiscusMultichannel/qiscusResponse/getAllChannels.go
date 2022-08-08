package qiscusResponse

import (
	"encoding/json"
	"errors"
	"fmt"
)

type GetAllChannels struct {
	Data map[string]interface{} `json:"data"`
}

type CustomChannelData struct {
	Id                  int    `json:"id"`
	WebhookUrl          string `json:"webhook_url"`
	LogoUrl             string `json:"logo_url"`
	IdentifierKey       string `json:"identifier_key"`
	Name                string `json:"name"`
	IsActive            bool   `json:"is_active"`
	UseChannelResponder bool   `json:"use_channel_responder"`
}

func (r *GetAllChannels) ChannelsIdIsAvailable(listChannel map[string]interface{}) error {
	var err error
	for channelKey, data := range listChannel {
		listChannelId, ok1 := data.([]interface{})
		if ok1 {
			for qiscusChannelKey, qiscusListChannel := range r.Data {
				if channelKey == qiscusChannelKey {
					listChannelIdQiscus, ok2 := qiscusListChannel.([]interface{})
					if ok2 {
						for _, i1 := range listChannelId {
							isExist := false
							notExitChannelId := i1.(string)
							for _, i2 := range listChannelIdQiscus {
								data, ok3 := i2.(map[string]interface{})
								if ok3 {
									a := fmt.Sprint(i1)
									b := fmt.Sprint(data["id"])
									if a == b {
										isExist = true
									}
								}
							}
							if !isExist {
								err = errors.New("channel Id " + fmt.Sprint(notExitChannelId) + " in " + channelKey + " is not available")
								break
							}
						}
					}
				}
			}
		} else {
			err = errors.New("format active channel: " + fmt.Sprint(data) + " not valid")
			break
		}
	}

	return err
}

func (r *GetAllChannels) GetDataCustomChannel(identifierKey string) (*CustomChannelData, error) {
	var customChannelData *CustomChannelData
	for key, data := range r.Data {
		if key == "custom_channels" {
			listChannel := new([]CustomChannelData)
			d, err := json.Marshal(data)
			if err == nil {
				err = json.Unmarshal(d, &listChannel)
				for _, dc := range *listChannel {
					if dc.IdentifierKey == identifierKey {
						customChannelData = &dc
						break
					}
				}
			}
		}
	}

	if customChannelData == nil {
		return nil, fmt.Errorf("custom channel with identifier_key=%v not found in your multichannel account", identifierKey)
	} else {
		return customChannelData, nil
	}
}
