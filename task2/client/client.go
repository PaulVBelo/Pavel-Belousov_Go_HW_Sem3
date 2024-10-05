package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"task2/dto"
	"time"
)

type Client struct {
	url string
}

func NewClient(url string) *Client {
	return &Client{url: url}
}

func (c *Client) GetVersion() ([]byte, error) {
	req, err := http.NewRequest("GET", c.url+"/version", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) GetHardOp() (bool, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", c.url+"/hard-op", nil)
	if err != nil {
		return false, 500, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return false, 500, nil
		}
		return false, 500, err
	}
	defer resp.Body.Close()

	return true, resp.StatusCode, nil
}

// Сохраняя мой вопрос про аннотации, пока что назову методы так. Если есть способы делать это покарсивей, буду рад узнать.
func (c *Client) PostDecode(inputString string) (string, error) {
	reqBody, err := json.Marshal(dto.DecodeRequest{InputString: inputString})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.url+"/decode", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var decoded dto.DecodeResponse
	err = json.Unmarshal(body, &decoded)
	if err != nil {
		return "", err
	}

	return decoded.OutputString, nil
}