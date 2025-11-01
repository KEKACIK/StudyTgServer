package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type StudyApiStudent struct {
	ID        int64
	Name      string
	Sex       string
	Age       int
	Course    int
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

type StudyApiServer struct {
	Host string
	Port int64
}

func NewStudyApiServer(host string, port int64) *StudyApiServer {
	return &StudyApiServer{
		Host: host,
		Port: port,
	}
}

func request(method, url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Ошибка при выполнении запроса: Статус %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при чтении ответа: %v", err)
	}
	return body, nil
}

func (s *StudyApiServer) Get(id int64) (*StudyApiStudent, error) {
	url := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", s.Host, s.Port),
		Path:   fmt.Sprintf("/student/%d", id),
	}
	r, err := request("GET", url.String())
	if err != nil {
		return nil, err
	}
	var student StudyApiStudent
	err = json.Unmarshal([]byte(r), &student)
	if err != nil {
		return nil, fmt.Errorf("Ошибка преобразования JSON: %v", err)
	}
	return &student, nil
}
