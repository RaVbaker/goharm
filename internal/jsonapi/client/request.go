package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/google/jsonapi"
)

func Request(path, method string, inputResource, resource interface{}) error {
	in := bytes.NewBuffer(nil)
	jsonapi.MarshalOnePayloadEmbedded(in, inputResource)

	resp, err := makeRequest(method, path, in)
	if err != nil {
		return err
	}
	if err := jsonapi.UnmarshalPayload(resp.Body, resource); err != nil {
		return err
	}
	return nil
}

func makeRequest(method string, path string, in *bytes.Buffer) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, Host+path, in)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", jsonapi.MediaType)
	req.Header.Add("Content-Type", jsonapi.MediaType)
	if len(AccessToken) > 0 {
		req.Header.Add("Authorization", "Bearer "+AccessToken)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		reply, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("invalid response status code %d for %s - `%v`", resp.StatusCode, path, string(reply))
	}
	return resp, nil
}

func RequestMany(path string, inputResource interface{}, resultType reflect.Type) ([]interface{}, error) {
	in := bytes.NewBuffer(nil)
	if inputResource != nil {
		jsonapi.MarshalOnePayloadEmbedded(in, inputResource)
	}

	resp, err := makeRequest(http.MethodGet, path, in)
	if err != nil {
		return nil, err
	}
	records, err := jsonapi.UnmarshalManyPayload(resp.Body, resultType)
	if err != nil {
		return nil, err
	}
	return records, nil
}
