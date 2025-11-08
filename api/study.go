package api

import (
	"encoding/json"
	"fmt"
	"net/url"
)

const (
	StudyStudentSexMan   = "man"
	StudyStudentSexWoman = "woman"
)

type StudyErrorResult struct {
	Message string
}

type StudyCreateResult struct {
	ID int64
}

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

func (s *StudyApiServer) Create(name, sex string, age, course int) (int64, error) {
	url := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", s.Host, s.Port),
		Path:   "/student",
	}
	r, err := requestPost(
		url.String(),
		map[string]interface{}{
			"name":   name,
			"sex":    sex,
			"age":    age,
			"course": course,
		},
	)
	if err != nil {
		return 0, err
	}
	var createResult StudyCreateResult
	err = json.Unmarshal([]byte(r), &createResult)
	if err != nil {
		return 0, fmt.Errorf("json transformation error: %v", err)
	}
	return createResult.ID, nil
}

func (s *StudyApiServer) Get(id int64) (*StudyApiStudent, error) {
	url := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", s.Host, s.Port),
		Path:   fmt.Sprintf("/student/%d", id),
	}
	r, err := requestGet(url.String())
	if err != nil {
		return nil, err
	}
	var student StudyApiStudent
	err = json.Unmarshal([]byte(r), &student)
	if err != nil {
		return nil, fmt.Errorf("json transformation error: %v", err)
	}
	return &student, nil
}

func (s *StudyApiServer) GetAll() ([]StudyApiStudent, error) {
	url := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", s.Host, s.Port),
		Path:   "/student/list",
	}
	r, err := requestGet(url.String())
	if err != nil {
		return nil, err
	}
	var students []StudyApiStudent
	err = json.Unmarshal([]byte(r), &students)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (s *StudyApiServer) Update(id int64, name, sex string, age, course int) (*StudyApiStudent, error) {
	url := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", s.Host, s.Port),
		Path:   fmt.Sprintf("/student/%d", id),
	}
	r, err := requestPut(
		url.String(),
		map[string]interface{}{
			"name":   name,
			"sex":    sex,
			"age":    age,
			"course": course,
		},
	)
	if err != nil {
		return nil, err
	}

	var student StudyApiStudent
	err = json.Unmarshal([]byte(r), &student)
	if err != nil {
		return nil, fmt.Errorf("json transformation error: %v", err)
	}

	return &student, nil
}

func (s *StudyApiServer) Delete(id int64) error {
	url := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", s.Host, s.Port),
		Path:   fmt.Sprintf("/student/%d", id),
	}
	_, err := requestDelete(url.String())
	return err
}

func FormatSexToRu(sex string) string {
	switch sex {
	case StudyStudentSexMan:
		return "Мужской"
	case StudyStudentSexWoman:
		return "Женский"
	default:
		return "Не определен"
	}
}
