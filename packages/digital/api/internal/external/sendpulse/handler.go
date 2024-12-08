package sendpulse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cyberpunkattack/environment"
	"io"
	"net/http"
	"time"
)

type SendPulseHandler struct {
	Now time.Time
}

type SendPulseUser struct {
	Email string `json:"email"`
	Name string `json:"name"`
}

type SendPulseMessage struct {
	Template int
	Subject string
	From string
	User *SendPulseUser
}

func sendToSendpulse(ctx context.Context, httpRequest *http.Request) (map[string]any, error) {
	tokenUrl := fmt.Sprintf("%s/oauth/access_token", environment.GEnv().GetVariable("SENDPULSE_BASE_URL"))

	dataType, err := json.Marshal(map[string]string{
		"grant_type": environment.GEnv().Get("SENDPULSE_GRANT_TYPE"),
		"client_id": environment.GEnv().Get("SENDPULSE_CLIENT_ID"),
		"client_secret": environment.GEnv().Get("SENDPULSE_SECRET"),
	})

	if err != nil {
		return map[string]any{}, err
	}

	req , err := http.NewRequestWithContext(ctx, "POST", tokenUrl, bytes.NewBuffer(dataType))

	if err != nil {
		return map[string]any{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return map[string]any{}, err
	}

	defer resp.Body.Close()

	reqBody, errResp := io.ReadAll(resp.Body)
	if errResp != nil {
		return map[string]any{}, errResp
	}

	res := map[string]any{}
	err = json.Unmarshal(reqBody, &res)
	if err != nil {
		return map[string]any{}, err
	}

	accessToken := res["access_token"].(string)

	httpRequest.Header.Set("Authorization", "Bearer " + accessToken)
	responseLast, err := client.Do(httpRequest)

	if err != nil {
		return map[string]any{}, err
	}

	resLastBody, errRespLast := io.ReadAll(responseLast.Body)
	if errRespLast != nil {
		return map[string]any{}, err
	}

	responseData := map[string]any{}

	err = json.Unmarshal(resLastBody, &responseData)
	if err != nil {
		return map[string]any{}, err
	}

	return responseData, nil
}

func (sp *SendPulseHandler) CreateUser(email, name string) *SendPulseUser {
	return &SendPulseUser{
        Email: email,
        Name:  name,
    }
}

func (us *SendPulseUser) AddUser(ctx context.Context, list string) error {
	addUserUrl := fmt.Sprintf("%s/addressbooks/%s/emails", environment.GEnv().GetVariable("SENDPULSE_BASE_URL"), list)

	user, err := json.Marshal(us)
	if err != nil {
		return err
	}

	httpRequest, err := http.NewRequestWithContext(ctx, "POST", addUserUrl, bytes.NewBuffer(user))
	if err != nil {
		return err
	}

	httpRequest.Header.Set("Content-Type", "application/json")
	_, err = sendToSendpulse(ctx, httpRequest)

	if err != nil {
		return err
	}

 	return nil
}

func (sp *SendPulseHandler) CreateMessage(template int, subject, from string, user *SendPulseUser) *SendPulseMessage {
	return &SendPulseMessage{
		Template: template,
		Subject:  subject,
		From:     from,
		User:     user,
	}
}

func (sp *SendPulseMessage) Send(ctx context.Context) error {
	sendEmailUrl := fmt.Sprintf("%s/smtp/emails", environment.GEnv().GetVariable("SENDPULSE_BASE_URL"))

	body := map[string]any{
		"email": map[string]any{
			"subject": sp.Subject,
			"from": map[string]string{
				"email": sp.From,
			},
			"template": map[string]any{
				"id": sp.Template,
				"variables": map[string]any{
					"name": sp.User.Name,
				},
			},
			"to": []map[string]string{
				{
					"name": sp.User.Name,
					"email": sp.User.Email,
				},
			},
		},
	}

	fmt.Println(body)

	bytesBody, err := json.Marshal(body)

	if err != nil {
		return err
	}

	httpRequest, err := http.NewRequestWithContext(ctx, "POST", sendEmailUrl, bytes.NewBuffer(bytesBody))
	if err != nil {
		return err
	}

	httpRequest.Header.Set("Content-Type", "application/json")
	_, err = sendToSendpulse(ctx, httpRequest)

	if err != nil {
		return err
	}

	return nil
}


func New() *SendPulseHandler {
	return &SendPulseHandler{
		Now: time.Now(),
	}
}


