package bot

import (
	"StudyTgServer/api"
	"fmt"
	"strconv"
)

func studentNameValidation(name string) (string, error) {
	if len(name) < 2 || len(name) > 32 {
		return "", fmt.Errorf("имя должно содержать 2-32 символа, повторите снова")
	}
	return name, nil
}

func studentSexValidation(sex string) (string, error) {
	if sex != api.StudyStudentSexMan && sex != api.StudyStudentSexWoman {
		return "", fmt.Errorf("keyboard error. sex field should be %s or %s", api.StudyStudentSexMan, api.StudyStudentSexWoman)
	}
	return sex, nil
}

func studentAgeValidation(ageStr string) (int, error) {
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return 0, fmt.Errorf("ID должен быть числовым, отправьте ещё раз")
	}
	if age < api.StudyStudentAgeMin {
		return 0, fmt.Errorf("возраст должен быть больше %d, отправьте ещё раз", api.StudyStudentAgeMin)
	}
	if age > api.StudyStudentAgeMax {
		return 0, fmt.Errorf("возраст должен быть меньше %d, отправьте ещё раз", api.StudyStudentAgeMax)
	}
	return age, nil
}

func studentCourseValidation(courseStr string) (int, error) {
	course, err := strconv.Atoi(courseStr)
	if err != nil {
		return 0, fmt.Errorf("курс должен быть числовым, отправьте ещё раз")
	}
	if course < api.StudyStudentCourseMin {
		return 0, fmt.Errorf("курс должен быть больше %d, отправьте ещё раз", api.StudyStudentCourseMin)
	}
	if course > api.StudyStudentCourseMax {
		return 0, fmt.Errorf("курс должен быть меньше %d, отправьте ещё раз", api.StudyStudentCourseMax)
	}
	return course, nil
}
