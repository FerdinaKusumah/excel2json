package excel2json

import (
	"fmt"
	"github.com/gojektech/heimdall/v6/hystrix"
	"io/ioutil"
	"net/http"
	"time"
)

// getFileUrl fetch file body with http client
// fileUrl string url
func getFileUrl(fileUrl string) ([]byte, error) {
	var (
		body []byte
		res  = new(http.Response)
		err  error
	)
	// Create a new hystrix-wrapped HTTP client with the command name, along with other required options
	client := hystrix.NewClient(
		hystrix.WithHTTPTimeout(5*time.Second),
		hystrix.WithCommandName(fmt.Sprintf(`%s||%s`, fileUrl, "GetFileUrl")),
		hystrix.WithHystrixTimeout(10*time.Second),
		hystrix.WithMaxConcurrentRequests(30),
		hystrix.WithErrorPercentThreshold(20),
	)
	// Create an http.Request instance
	if res, err = client.Get(fileUrl, nil); err != nil {
		return nil, err
	}
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}
	return body, nil
}

// getFilePath fetch file body with http client
// filePath string url
func getFilePath(filePath string) ([]byte, error) {
	var (
		body []byte
		err  error
	)
	// just pass the file name
	if body, err = ioutil.ReadFile(filePath); err != nil {
		return nil, err
	}
	return body, nil
}
