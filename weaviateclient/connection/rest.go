package connection

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

const apiVersion = "v1"

type Connection struct {
	basePath string
	httpClient *http.Client
}

func NewConnection(scheme string, host string) *Connection {
	return &Connection{
		basePath: scheme + "://" + host + "/" + apiVersion,
		httpClient: &http.Client{},
	}
}

func (con *Connection) addHeaderToRequest(request *http.Request) {
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
}

func (con *Connection) marshalBody(body interface{}) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}
	jsonBody, err := json.Marshal(body)  // Create the JSON body
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(jsonBody), nil
}

func (con *Connection) createRequest(ctx context.Context, path string, restMethod string, body interface{}) (*http.Request, error) {
	url := con.basePath+path  // Create the URL

	jsonBody, err := con.marshalBody(body)
	if err != nil{
		return nil, err
	}

	request, err := http.NewRequest(restMethod, url, jsonBody)
	if err != nil {
		return nil, err
	}
	con.addHeaderToRequest(request)
	request.WithContext(ctx)
	return request, nil
}

// RunREST executes a http request
// path: expects a resource path e.g. `/schema/things`
// restMethod: as they are defined in constants in the *http* package
// Returns:
//  a response that may be parsed into a struct after the fact
//  error if there was a network issue
func (con *Connection) RunREST(ctx context.Context, path string, restMethod string, requestBody interface{}) (*ResponseData, error) {
	request, requestErr := con.createRequest(ctx, path, restMethod, requestBody)
	if requestErr != nil {
		return nil, requestErr
	}
	response, responseErr := con.httpClient.Do(request)
	if responseErr != nil {
		return nil, responseErr
	}

	defer response.Body.Close()
	body, bodyErr := ioutil.ReadAll(response.Body)
	if bodyErr != nil {
		return nil, bodyErr
	}

	return &ResponseData{
		Body:       body,
		StatusCode: response.StatusCode,
	}, nil
}

// ResponseData encapsulation of the http request body and status
type ResponseData struct {
	Body []byte
	StatusCode int
}

// DecodeBodyIntoTarget unmarshall body into target var
// successful if err is nil
func (rd *ResponseData) DecodeBodyIntoTarget(target interface{}) error {
	err := json.Unmarshal(rd.Body, target)
	if err != nil {
		return err
	}
	return nil
}

