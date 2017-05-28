package startrack

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	ProdBaseURL        = "https://digitalapi.auspost.com.au"
	TestbedBaseURL     = "https://digitalapi.auspost.com.au/testbed"
	TestbedBetaBaseURL = "https://digitalapi.auspost.com.au/testbedbeta"
)

type Client struct {
	username      string
	password      string
	accountNumber string
	BaseURL       string
	client        *http.Client
}

func New(username, password, accountNumber string) *Client {
	return &Client{
		username:      username,
		password:      password,
		accountNumber: accountNumber,
		BaseURL:       ProdBaseURL,
		client:        &http.Client{Timeout: 30 * time.Second},
	}
}

type errorResponseError struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

type errorResponse struct {
	Errors []errorResponseError `json:"errors"`
}

func (e *errorResponse) String() string {
	s := ""
	for _, e := range e.Errors {
		if s != "" {
			s += " "
		}
		s += e.Message
	}
	return s
}

func (c *Client) post(data interface{}, resource string, response interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	log.Print("POST body: " + string(body))

	req, err := http.NewRequest(http.MethodPost, c.BaseURL+resource, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	return c.do(req, response)
}

func (c *Client) get(resource string, response interface{}) error {
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+resource, nil)
	if err != nil {
		return err
	}
	return c.do(req, response)
}

//func (c *Client) do(respData interface{}, httpResp *http.Response) error {
func (c *Client) do(req *http.Request, respData interface{}) error {
	req.SetBasicAuth(c.username, c.password)
	req.Header.Add("Account-Number", c.accountNumber)
	req.Header.Add("Accept", "application/json")

	log.Printf("Sending %s request to %s", req.Method, req.URL.String())
	for k, vs := range req.Header {
		for _, v := range vs {
			log.Printf("  %s: %s", k, v)
		}
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("Response code: %d", resp.StatusCode)
	log.Printf("Response body: %s", string(body))

	er := errorResponse{}
	if err = json.Unmarshal(body, &er); err != nil {
		return errors.New("failed to read if there were any errors: " + err.Error())
	}

	if len(er.Errors) > 0 {
		return errors.New("API error from AusPost: " + er.String())
	}

	if err = json.Unmarshal(body, respData); err != nil {
		return errors.New("failed to unmarshal response: " + err.Error())
	}

	return nil
}
