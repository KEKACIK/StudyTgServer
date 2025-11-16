package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func requestErrorUnpack(data []byte) (string, error) {
	var result StudyErrorResult
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		return "", err
	}
	return result.Message, nil
}

func requestErrorDetect(resp *http.Response, body []byte) error {
	if resp.StatusCode == 400 {
		errorMessage, err := requestErrorUnpack(body)
		if err != nil {
			return err
		}
		return fmt.Errorf("request error, status code %d: %s", resp.StatusCode, errorMessage)
	}
	if resp.StatusCode == 401 {
		errorMessage, err := requestErrorUnpack(body)
		if err != nil {
			return err
		}
		return fmt.Errorf("request auth error, status code %d: %s", resp.StatusCode, errorMessage)
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("request error, status code %d", resp.StatusCode)
	}
	return nil
}

func requestGet(url string, header http.Header) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}

	req.Header = header

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read error: %v", err)
	}

	err = requestErrorDetect(resp, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func requestPost(url string, header http.Header, data map[string]interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("data formatted to json error: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}

	req.Header = header

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read error: %v", err)
	}

	err = requestErrorDetect(resp, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func requestPut(url string, header http.Header, data map[string]interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("data formatted to json error: %v", err)
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}

	req.Header = header
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read error: %v", err)
	}

	err = requestErrorDetect(resp, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
func requestDelete(url string, header http.Header) ([]byte, error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}

	req.Header = header

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read error: %v", err)
	}

	err = requestErrorDetect(resp, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
