package bot

import (
	"StudyTgServer/api"
	"fmt"
)

func studentNameValidation(name string) (string, error) {
	if len(name) < 2 || len(name) > 32 {
		return "", fmt.Errorf("имя должно содержать 2-32 символа, повторите снова")
	}
	return name, nil
}

func studentSexValidation(sex string) (string, error) {
	if sex != api.StudyStudentSexMan && sex != api.StudyStudentSexWoman {
		return "", fmt.Errorf("sex field not '%s' or '%s' keyboard error.", api.StudyStudentSexMan, api.StudyStudentSexWoman)
	}
	return sex, nil
}
func studentAgeValidation(sex string) (int64, error) {
	return 0, nil
}
func studentCourseValidation(sex string) (int64, error) {
	return 0, nil
}
