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

func requestGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request error: status code %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read error: %v", err)
	}
	return body, nil
}

func requestPost(url string, data map[string]interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("data formatted to json error: %v", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read error: %v", err)
	}

	if resp.StatusCode == 400 {
		errorMessage, err := requestErrorUnpack(body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("request error, status code %d: %s", resp.StatusCode, errorMessage)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request error, status code %d", resp.StatusCode)
	}

	return body, nil
}

func requestPut(url string, data map[string]interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("data formatted to json error: %v", err)
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read error: %v", err)
	}

	if resp.StatusCode == 400 {
		errorMessage, err := requestErrorUnpack(body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("request error, status code %d: %s", resp.StatusCode, errorMessage)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request error, status code %d", resp.StatusCode)
	}

	return body, nil
}
func requestDelete(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read error: %v", err)
	}

	if resp.StatusCode == 400 {
		errorMessage, err := requestErrorUnpack(body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("request error, status code %d: %s", resp.StatusCode, errorMessage)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request error, status code %d", resp.StatusCode)
	}

	return body, nil
}
