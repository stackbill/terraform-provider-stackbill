package http

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"terraform-provider-stackbill/api"
)

// New http client
func NewHttpClient() HttpClient {
	return &httpClient{}
}

type HttpClient interface {
	Client() *http.Client
	Get(endPoint string) (string, error)
	PostJson(endPoint string, data interface{}) (string, error)
	PutJson(endPoint string, data interface{}) (string, error)
	Delete(endPoint string) (string, error)
}

type httpClient struct {
}

/*
@return custom http client
*/
func (c *httpClient) Client() *http.Client {
	// Custom Client
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	httpClient := &http.Client{Transport: customTransport}
	return httpClient
}

/*
@param endpoint - api end point url
@param apiKey
@param secretKey
@return string - http response
*/
func (c *httpClient) Get(endPoint string) (string, error) {
	// Make Get Request
	req, err := http.NewRequest("GET", endPoint, nil)
	// Check the error
	if err != nil {
		return "", err
	}
	// Set the authentication token
	req.Header.Add("apikey", api.API_KEY)
	req.Header.Add("secretKey", api.SECRET_KEY)
	// Create the client
	client := c.Client()
	// Make Request
	resp, err := client.Do(req)
	// Check the error
	if err != nil {
		return "", err
	}
	// Read Reponse
	body, err := ioutil.ReadAll(resp.Body)
	// Check the error
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// Check the status code
	if resp.StatusCode > 202 {
		return "", errors.New(string(body))
	}
	return string(body), nil
}

/*
@param endpoint - api end point url
@param apiKey
@param secretKey
@param data - Json data
@return string - http response
*/
func (c *httpClient) PostJson(endPoint string, data interface{}) (string, error) {
	// Convert struct to JSON
	exJson, _ := json.Marshal(data)
	// Make new post request
	jsonStr := []byte(string(exJson))
	req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(jsonStr))
	// Check the error
	if err != nil {
		return "", err
	}
	// Set the authentication token
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("apikey", api.API_KEY)
	req.Header.Add("secretKey", api.SECRET_KEY)
	// Get Custom Client
	client := c.Client()
	resp, err := client.Do(req)
	// Check the error
	if err != nil {
		return "", err
	}
	// Read the output from the response
	body, err := ioutil.ReadAll(resp.Body)
	// Check the error
	if err != nil {
		return "", err
	}
	// Check the status
	if resp.StatusCode > 202 {
		return "", errors.New(string(body))
	}
	defer resp.Body.Close()
	return string(body), nil
}

/*
@param endpoint - api end point url
@param apiKey
@param secretKey
@param data - Json data
@return string - http response
*/
func (c *httpClient) PutJson(endPoint string, data interface{}) (string, error) {
	// Convert struct to JSON
	exJson, _ := json.Marshal(data)
	// Make new post request
	jsonStr := []byte(string(exJson))
	req, err := http.NewRequest("PUT", endPoint, bytes.NewBuffer(jsonStr))
	// Check the error
	if err != nil {
		return "", err
	}
	// Set the authentication token
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("apikey", api.API_KEY)
	req.Header.Add("secretKey", api.SECRET_KEY)
	// Get Custom Client
	client := c.Client()
	resp, err := client.Do(req)
	// Check the error
	if err != nil {
		return "", err
	}
	// Read the output from the response
	body, err := ioutil.ReadAll(resp.Body)
	// Check the error
	if err != nil {
		return "", err
	}
	// Check the status
	if resp.StatusCode > 202 {
		return "", errors.New(string(body))
	}
	defer resp.Body.Close()
	return string(body), nil
}

/*
@param endpoint - api end point url
@param apiKey
@param secretKey
@return string - http response
*/
func (c *httpClient) Delete(endPoint string) (string, error) {
	// Make Get Request
	req, err := http.NewRequest("DELETE", endPoint, nil)
	// Check the error
	if err != nil {
		return "", err
	}
	// Set the authentication token
	req.Header.Add("apikey", api.API_KEY)
	req.Header.Add("secretKey", api.SECRET_KEY)
	// Create the client
	client := c.Client()
	// Make Request
	resp, err := client.Do(req)
	// Check the error
	if err != nil {
		return "", err
	}
	// Read Reponse
	body, err := ioutil.ReadAll(resp.Body)
	// Check the error
	if err != nil {
		return "", err
	}
	// Check the status
	if resp.StatusCode > 202 {
		return "", errors.New(string(body))
	}
	defer resp.Body.Close()
	return string(body), nil
}
