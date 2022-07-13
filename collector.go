package collector

import (
	"crypto/tls"
	"errors"
	"golang.org/x/text/encoding/charmap"
	"io/ioutil"
	"net/http"
)

type Collector interface {
	GetContent(url string) (string, error)
}

type collector struct {
	client HTTPClient
}

func NewCollector() Collector {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}
	c := &http.Client{Transport: tr}
	return &collector{client: c}
}

func (c *collector) GetContent(url string) (string, error) {

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", errors.New(requestCreationFailed)
	}

	response, err := c.client.Do(request)

	if c.isError(response, err) {
		return "", err
		// return "", errors.New(requestExecutionFailed)
	}

	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.New(contentReadingFailed)
	}
	bodyString := string(bodyBytes)
	decoder := charmap.Windows1252.NewDecoder()
	content, _ := decoder.String(bodyString)
	return content, nil
}

func (c *collector) isError(response *http.Response, err error) bool {
	return err != nil || response.StatusCode == http.StatusGatewayTimeout
}
