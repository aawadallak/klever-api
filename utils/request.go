package utils

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

func DoRequest(url string, method string, BodyInBytes *[]byte) (*http.Response, error) {
	var body io.Reader

	if BodyInBytes != nil {
		body = bytes.NewReader(*BodyInBytes)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	client := http.Client{
		Timeout: time.Second * 10,
	}

	return client.Do(req)
}
