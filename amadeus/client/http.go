package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client interface {
	GetToken() error
	CheckToken() error
	GetBaseURL() (string, error)
	GetAuthorization() string
}

// Do send request to api
func Do(c Client, req Request, resp *Response, reqType string) error {

	var err error

	// check token
	err = c.CheckToken()
	if err != nil {

		fmt.Println(err)
		err = c.GetToken()
		if err != nil {
			return err
		}

		fmt.Println(err)
	}

	// get base api url
	baseURL, err := c.GetBaseURL()
	if err != nil {
		return err
	}

	// prepare request
	r, err := http.NewRequest(
		reqType,
		baseURL+req.GetURL(reqType),
		req.GetBody(reqType),
	)
	if err != nil {
		return err
	}

	// add headers
	r.Header.Add("Authorization", c.GetAuthorization())
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")

	// send request
	client := http.Client{}
	rsp, err := client.Do(r)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	// read body
	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	// check status code
	// return error if not 200

	fmt.Println("------------------")
	fmt.Println(rsp.StatusCode)
	fmt.Println(req.GetURL(reqType))
	fmt.Println(string(b))
	fmt.Println("------------------")

	// decode response to struct
	err = (*resp).Decode(b)

	if err != nil {
		return err
	}

	return nil

}
