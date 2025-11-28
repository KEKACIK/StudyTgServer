package bot

import (
	"StudyTgServer/internal/api"
	"errors"
	"strconv"
)

var (
	ErrStr2IntNotNumber = errors.New("invalid integer: not number")
)

func Str2IntValidation(numStr string) (int, error) {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, ErrStr2IntNotNumber
	}

	return num, nil
}

var (
	ErrNameEmpty    = errors.New("invalid name: empty")
	ErrNameTooSmall = errors.New("invalid name: too small")
	ErrNameTooBig   = errors.New("invalid name: too big")
)

func NameValidation(name string) error {
	if name == "" {
		return ErrNameEmpty
	}
	if len(name) < api.StudyStudentMinName {
		return ErrNameTooSmall
	}
	if len(name) > api.StudyStudentMaxName {
		return ErrNameTooBig
	}

	return nil
}

var (
	ErrValidationAgeTooSmall = errors.New("invalid age: too small")
	ErrValidationAgeTooBig   = errors.New("invalid age: too big")
)

func AgeValidation(age int) error {
	if age < api.StudyStudentMinAge {
		return ErrValidationAgeTooSmall
	}
	if age > api.StudyStudentMaxAge {
		return ErrValidationAgeTooBig
	}

	return nil
}

var (
	ErrValidationSexInvalid = errors.New("invalid sex")
)

func SexValidation(sex string) error {
	sexList := [2]string{
		api.StudyStudentSexMan,
		api.StudyStudentSexWoman,
	}
	sexInList := false
	for _, value := range sexList {
		if sex == value {
			sexInList = true
			break
		}
	}

	if !sexInList {
		return ErrValidationSexInvalid
	}

	return nil
}

var (
	ErrValidationCourseTooSmall = errors.New("invalid course: too small")
	ErrValidationCourseTooBig   = errors.New("invalid course: too big")
)

func CourseValidation(course int) error {
	if course < api.StudyStudentMinCourse {
		return ErrValidationCourseTooSmall
	}
	if course > api.StudyStudentMaxCourse {
		return ErrValidationCourseTooBig
	}
	return nil
}
