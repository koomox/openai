package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/net/proxy"
	"io"
	"net/http"
)

type Client struct {
	apiKey       string
	organization string
}

func NewClient(apiKey, organization string) *Client {
	return &Client{
		apiKey:       apiKey,
		organization: organization,
	}
}

type Completion struct {
	Model       string      `json:"model"`
	Messages    interface{} `json:"messages"`
	Temperature float64     `json:"temperature"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (c *Client) Request(question string, dialer *http.Client) ([]byte, error) {
	b, err := json.Marshal(&Completion{
		Model:       "gpt-3.5-turbo",
		Messages:    []Message{{Role: "user", Content: question}},
		Temperature: 0.7,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:108.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	res, err := dialer.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request %s, status %s", s.resource, resp.Status)
	}

	return io.ReadAll(res.Body)
}

func NewProxyDialer(network, addr string) *http.Client {
	dialer, _ := proxy.SOCKS5(network, addr, nil, proxy.Direct)
	transport := &http.Transport{
		Dial: dialer.Dial,
	}
	return &http.Client{Transport: transport}
}
