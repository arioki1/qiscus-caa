package qiscusMultichannel

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/arioki1/qiscus-caa/thirdparty/qiscusMultichannel/qiscusRequest"
	"github.com/arioki1/qiscus-caa/thirdparty/qiscusMultichannel/qiscusResponse"
	"github.com/google/go-querystring/query"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

type MultichannelConfig struct {
	AppID     string
	SecretKey string
	BaseUrl   string
}

type QiscusMultichannel interface {
	Login(ctx context.Context, req qiscusRequest.LoginQiscusRequest) (*qiscusResponse.LoginQiscusResponse, error)
	AssignAgentToChatRoom(ctx context.Context, req qiscusRequest.AssignAgentRequest) (*qiscusResponse.QiscusResponse, error)
	GetAgent(ctx context.Context, agentId int) (*qiscusResponse.GetAgent, error)
	GetRoomByRoomId(ctx context.Context, roomId string) (*qiscusResponse.GetRoomByRoomId, error)
	GetAllChannels(ctx context.Context) (*qiscusResponse.GetAllChannels, error)
	GetOfficeHours(ctx context.Context) (*qiscusResponse.GetOfficeHours, error)
	SendMessageText(ctx context.Context, req qiscusRequest.SendMessageTextRequest) error
	MarkAsResolveRoom(ctx context.Context, roomId string, lastCommentId string) error
	GetAgents(ctx context.Context, par qiscusRequest.GetAgents) (*qiscusResponse.GetAgents, error)
	GetUserInfo(ctx context.Context, roomId string) (*qiscusResponse.GetUserInfo, error)
	SetUserInfo(ctx context.Context, roomId string, req qiscusRequest.SetUserInfo) error
	AddOrUpdateUserInfo(ctx context.Context, roomId string, req qiscusRequest.SetUserInfo) error
	GetCustomerRoom(ctx context.Context, roomId string) (*qiscusResponse.GetCustomerRoom, error)
	UpdateCustomerRoom(ctx context.Context, contactId string, label string, value string) error
	UpdateMultipleCustomerRooms(ctx context.Context, roomId string, req qiscusRequest.UpdateMultipleCustomerRooms) error
	GetConversationHistory(ctx context.Context, roomId string) (*qiscusResponse.GetConversationHistory, error)
	CustomChannelSentMessageText(ctx context.Context, req qiscusRequest.CustomChannelSendMessageText) (*qiscusResponse.CustomChannelSendMessageText, error)
}

func NewQiscusMultichannelClient(conf MultichannelConfig) QiscusMultichannel {
	return &conf
}

func (q *MultichannelConfig) GetConversationHistory(ctx context.Context, roomId string) (*qiscusResponse.GetConversationHistory, error) {
	u := "/api/v1/export/conversations?room_id=" + roomId
	httpResp, err := q.sendRequest(ctx, u, http.MethodGet, "application/json", nil)
	if err != nil {
		return nil, err
	}

	resp := new(qiscusResponse.GetConversationHistory)
	if err := json.Unmarshal(httpResp, &resp); err != nil {
		return nil, err
	}

	if resp.Data.DownloadUrl == "" {
		return nil, errors.New("conversation history not found")
	}
	return resp, nil
}

func (q *MultichannelConfig) UpdateMultipleCustomerRooms(ctx context.Context, roomId string, req qiscusRequest.UpdateMultipleCustomerRooms) error {
	customerRooms, err := q.GetCustomerRoom(ctx, roomId)
	if err != nil {
		return err
	}
	contactId := strconv.FormatInt(customerRooms.Data.CustomerRoom.ContactId, 10)
	for _, d := range req {

		err := q.UpdateCustomerRoom(ctx, contactId, d.Label, d.Value)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return nil
}

func (q *MultichannelConfig) GetCustomerRoom(ctx context.Context, roomId string) (*qiscusResponse.GetCustomerRoom, error) {
	u := "/api/v2/customer_rooms/" + roomId
	httpResp, err := q.sendRequest(ctx, u, http.MethodGet, "application/json", nil)
	if err != nil {
		return nil, err
	}

	resp := new(qiscusResponse.GetCustomerRoom)
	if err := json.Unmarshal(httpResp, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (q *MultichannelConfig) UpdateCustomerRoom(ctx context.Context, contactId string, label string, value string) error {
	u := "/api/v2/contacts/properties/" + contactId

	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	w.WriteField("label", label)
	w.WriteField("value", value)
	w.Close()

	_, err2 := q.sendRequest(ctx, u, http.MethodPost, w.FormDataContentType(), body)
	if err2 != nil {
		return err2
	}
	return nil
}

func (q *MultichannelConfig) AddOrUpdateUserInfo(ctx context.Context, roomId string, req qiscusRequest.SetUserInfo) error {
	if len(req.UserProperties) == 0 {
		return errors.New("no user info to update")
	}

	var result = &qiscusRequest.SetUserInfo{}
	userInfo, _ := q.GetUserInfo(ctx, roomId)

	if userInfo != nil {
		if len(userInfo.Data.Extras.UserProperties) > 0 {
			result = &qiscusRequest.SetUserInfo{
				UserProperties: userInfo.Data.Extras.UserProperties,
			}
		}
	}

	for index, userProperty := range req.UserProperties {
		found := false
		for _, userPropertyResult := range result.UserProperties {
			if userPropertyResult.Key == userProperty.Key {
				found = true
				result.UserProperties[index] = userProperty

			}
		}
		if !found {
			result.UserProperties = append(result.UserProperties, userProperty)
		}

	}

	if err := q.SetUserInfo(ctx, roomId, *result); err != nil {
		return err
	}

	return nil
}

func (q *MultichannelConfig) SetUserInfo(ctx context.Context, roomId string, req qiscusRequest.SetUserInfo) error {
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}
	u := "/api/v1/qiscus/room/" + roomId + "/user_info"
	_, err2 := q.sendRequest(ctx, u, http.MethodPost, "application/json", strings.NewReader(string(b)))
	if err2 != nil {
		return err
	}
	return nil
}

func (q *MultichannelConfig) GetUserInfo(ctx context.Context, roomId string) (*qiscusResponse.GetUserInfo, error) {
	u := "/api/v1/qiscus/room/" + roomId + "/user_info"
	httpResp, err := q.sendRequest(ctx, u, http.MethodGet, "application/json", nil)
	if err != nil {
		return nil, err
	}

	resp := new(qiscusResponse.GetUserInfo)
	if err := json.Unmarshal(httpResp, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (q *MultichannelConfig) GetAgents(ctx context.Context, par qiscusRequest.GetAgents) (*qiscusResponse.GetAgents, error) {
	v, err := query.Values(par)
	if err != nil {
		return nil, err
	}

	u := "/api/v2/admin/agents?" + v.Encode()
	httpResp, err := q.sendRequest(ctx, u, http.MethodGet, "application/json", nil)
	if err != nil {
		return nil, err
	}
	resp := new(qiscusResponse.GetAgents)
	if err := json.Unmarshal(httpResp, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *MultichannelConfig) Login(ctx context.Context, req qiscusRequest.LoginQiscusRequest) (*qiscusResponse.LoginQiscusResponse, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpResp, err := q.sendRequest(ctx, q.AppID+"/api/v1/auth", http.MethodPost, "application/json", strings.NewReader(string(b)))
	if err != nil {
		return nil, err
	}
	resp := new(qiscusResponse.LoginQiscusResponse)
	if err := json.Unmarshal(httpResp, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (q *MultichannelConfig) AssignAgentToChatRoom(ctx context.Context, req qiscusRequest.AssignAgentRequest) (*qiscusResponse.QiscusResponse, error) {

	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpResp, err := q.sendRequest(ctx, "/api/v1/admin/service/assign_agent", http.MethodPost, "application/json", strings.NewReader(string(b)))
	if err != nil {
		return nil, err
	}

	resp := new(qiscusResponse.QiscusResponse)
	if err := json.Unmarshal(httpResp, &resp); err != nil {
		return nil, err
	}

	resp.Message = "success send to multichannel"

	return resp, nil
}

func (q *MultichannelConfig) GetAgent(ctx context.Context, agentId int) (*qiscusResponse.GetAgent, error) {
	u := "/api/v2/admin/agent/" + strconv.Itoa(agentId)
	httpResp, err := q.sendRequest(ctx, u, http.MethodGet, "application/json", nil)
	if err != nil {
		return nil, err
	}

	resp := new(qiscusResponse.GetAgent)
	if err := json.Unmarshal(httpResp, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *MultichannelConfig) GetRoomByRoomId(ctx context.Context, roomId string) (*qiscusResponse.GetRoomByRoomId, error) {
	u := "/api/v2/customer_rooms/" + roomId
	httpResp, err := q.sendRequest(ctx, u, http.MethodGet, "application/json", nil)
	if err != nil {
		return nil, err
	}

	resp := new(qiscusResponse.GetRoomByRoomId)
	if err := json.Unmarshal(httpResp, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *MultichannelConfig) GetAllChannels(ctx context.Context) (*qiscusResponse.GetAllChannels, error) {
	u := "/api/v2/channels"
	httpResp, err := q.sendRequest(ctx, u, http.MethodGet, "application/json", nil)
	if err != nil {
		return nil, err
	}

	resp := new(qiscusResponse.GetAllChannels)
	if err := json.Unmarshal(httpResp, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *MultichannelConfig) GetOfficeHours(ctx context.Context) (*qiscusResponse.GetOfficeHours, error) {
	u := "/api/v1/admin/office_hours"
	httpResp, err := q.sendRequest(ctx, u, http.MethodGet, "application/json", nil)
	if err != nil {
		return nil, err
	}

	resp := new(qiscusResponse.GetOfficeHours)
	if err := json.Unmarshal(httpResp, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *MultichannelConfig) SendMessageText(ctx context.Context, req qiscusRequest.SendMessageTextRequest) error {
	b, err := json.Marshal(req)
	_, err = q.sendRequest(ctx, "/"+q.AppID+"/bot", http.MethodPost, "application/json", strings.NewReader(string(b)))
	if err != nil {
		return err
	}

	return nil
}

func (q *MultichannelConfig) MarkAsResolveRoom(ctx context.Context, roomId string, lastCommentId string) error {
	requestBody := qiscusRequest.MarkAsResolved{
		RoomId:        roomId,
		LastCommentId: lastCommentId,
	}
	b, _ := json.Marshal(requestBody)
	u := "/api/v1/admin/service/mark_as_resolved"
	_, err2 := q.sendRequest(ctx, u, http.MethodPost, "application/json", strings.NewReader(string(b)))

	if err2 != nil {
		return err2
	}
	return nil
}

func (q *MultichannelConfig) CustomChannelSentMessageText(ctx context.Context, req qiscusRequest.CustomChannelSendMessageText) (*qiscusResponse.CustomChannelSendMessageText, error) {
	b, _ := json.Marshal(req)
	u := fmt.Sprintf("/%v/api/v2/custom_channel/send", q.AppID)
	httpResp, err := q.sendRequest(ctx, u, http.MethodPost, "application/json", strings.NewReader(string(b)))
	if err != nil {
		return nil, err
	}

	resp := new(qiscusResponse.CustomChannelSendMessageText)
	if err := json.Unmarshal(httpResp, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *MultichannelConfig) sendRequest(ctx context.Context, uri, method, contentType string, payload io.Reader) ([]byte, error) {
	var data io.Reader
	if payload == nil {
		data = nil
	} else {
		data = payload
	}
	u := fmt.Sprintf("%s%s", q.BaseUrl, uri)
	r, err := http.NewRequest(method, u, data)
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", contentType)
	r.Header.Add("Qiscus-App-Id", q.AppID)
	r.Header.Add("Qiscus-Secret-Key", q.SecretKey)

	r.WithContext(ctx)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("failed to perform http request: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %v ", err)
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(responseBody))
	}

	return responseBody, nil
}
