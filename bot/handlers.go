package bot

import (
	"StudyTgServer/api"
	"StudyTgServer/utils"
	"fmt"
	"strconv"

	"gopkg.in/telebot.v3"
)

func (b *Bot) RegisterHandlers() {
	b.Handle("/start", b.startHandler)
	b.Handle("/create", b.createHandler)
	b.Handle(&telebot.Btn{Unique: "create_sex"}, b.createSexHandler)
	b.Handle("/get", b.getHandler)
	b.Handle("/get_all", b.getAllHandler)
	b.Handle("/update", b.updateHandler)
	b.Handle("/delete", b.deleteHandler)
	b.Handle(&telebot.Btn{Unique: "delete_yes"}, b.deleteYesHandler)
	b.Handle(&telebot.Btn{Unique: "delete_no"}, b.deleteNoHandler)
	b.Handle(telebot.OnText, b.textHandler)
}

// TEXT

func (b *Bot) textHandler(c telebot.Context) error {
	state, exists := b.states[c.Chat().ID]
	if !exists {
		return c.Send("Команда не найдена.")
	}

	switch state {
	// Create handlers
	case StudentCreateNameState:
		return b.createNameHandler(c)
	case StudentCreateAgeState:
		return b.createAgeHandler(c)
	case StudentCreateCourseState:
		return b.createCourseHandler(c)
	// Get handlers
	case StudentGetIdState:
		return b.getIdHandler(c)
	// Delete handlers
	case StudentDeleteIdState:
		return b.deleteIdHandler(c)
	}
	return nil
}

// COMMAND START

func (b *Bot) startHandler(c telebot.Context) error {
	return c.Send(
		utils.MultiLine(
			"Добро пожаловать в <b>StudyTgBot</b>",
			"",
			"Вот список команд для использования бота",
			"/start - Запустить бота",
			"/create - Добавить нового студента",
			"/get - Получить студента по ID",
			"/get_all - Получить всех студентов",
			"/update - Обновить студента по ID",
			"/delete - Удалить студента по ID",
		),
		&telebot.SendOptions{
			DisableWebPagePreview: false,
			ParseMode:             telebot.ModeHTML,
		},
	)
}

// COMMAND CREATE

const (
	StudentCreateNameState   States = "student_create_name"
	StudentCreateAgeState    States = "student_create_age"
	StudentCreateCourseState States = "student_create_course"
)

func (b *Bot) createHandler(c telebot.Context) error {
	b.states[c.Chat().ID] = StudentCreateNameState
	b.data[c.Chat().ID] = api.StudyApiStudent{}
	return c.Send(
		utils.MultiLine(
			"Создание студента",
			"",
			"Пришлите <b>Имя</b> студента",
		),
		&telebot.SendOptions{
			DisableWebPagePreview: false,
			ParseMode:             telebot.ModeHTML,
		},
	)
}

func (b *Bot) createNameHandler(c telebot.Context) error {
	studentName := c.Text()
	if studentName == "" {
		return c.Send(
			"Текст не найден, отправьте ещё раз",
			&telebot.SendOptions{
				DisableWebPagePreview: false,
				ParseMode:             telebot.ModeHTML,
			},
		)
	}

	student := b.data[c.Chat().ID]
	student.Name = studentName
	b.data[c.Chat().ID] = student

	return c.Send(
		utils.MultiLine(
			"Создание студента",
			"",
			"Выберите <b>Пол</b> студента",
		),
		&telebot.SendOptions{
			DisableWebPagePreview: false,
			ParseMode:             telebot.ModeHTML,
			ReplyMarkup:           createSexKeyboard(),
		},
	)
}

func (b *Bot) createSexHandler(c telebot.Context) error {
	studentSex := c.Callback().Data

	if studentSex != "man" && studentSex != "woman" {
		return c.Send(
			"Произошла ошибка, обратитесь в тех. Поддержку",
			&telebot.SendOptions{
				DisableWebPagePreview: false,
				ParseMode:             telebot.ModeHTML,
			},
		)
	}

	student := b.data[c.Chat().ID]
	student.Sex = studentSex
	b.data[c.Chat().ID] = student
	b.states[c.Chat().ID] = StudentCreateAgeState
	return c.Edit(
		utils.MultiLine(
			"Создание студента",
			"",
			"Пришлите <b>Возраст</b> студента",
		),
		&telebot.SendOptions{
			DisableWebPagePreview: false,
			ParseMode:             telebot.ModeHTML,
		},
	)
}

func (b *Bot) createAgeHandler(c telebot.Context) error {
	studentAge, err := strconv.Atoi(c.Text())
	if err != nil {
		return c.Send(
			"Возраст должен быть числовым, отправьте ещё раз",
			&telebot.SendOptions{
				DisableWebPagePreview: false,
				ParseMode:             telebot.ModeHTML,
			},
		)
	}
	if studentAge < 14 {
		return c.Send(
			"Возраст должен быть больше 14 лет, отправьте ещё раз",
			&telebot.SendOptions{
				DisableWebPagePreview: false,
				ParseMode:             telebot.ModeHTML,
			},
		)
	}
	if studentAge > 80 {
		return c.Send(
			"Возраст должен быть меньше 80 лет, отправьте ещё раз",
			&telebot.SendOptions{
				DisableWebPagePreview: false,
				ParseMode:             telebot.ModeHTML,
			},
		)
	}

	student := b.data[c.Chat().ID]
	student.Age = studentAge
	b.data[c.Chat().ID] = student
	b.states[c.Chat().ID] = StudentCreateCourseState

	return c.Send(
		utils.MultiLine(
			"Создание студента",
			"",
			"Пришлите <b>Курс</b> студента",
		),
		&telebot.SendOptions{
			DisableWebPagePreview: false,
			ParseMode:             telebot.ModeHTML,
		},
	)
}

func (b *Bot) createCourseHandler(c telebot.Context) error {
	studentCourse, err := strconv.Atoi(c.Text())
	if err != nil {
		return c.Send(
			"Курс должен быть числовым, отправьте ещё раз",
			&telebot.SendOptions{
				DisableWebPagePreview: false,
				ParseMode:             telebot.ModeHTML,
			},
		)
	}
	if studentCourse < 1 {
		return c.Send(
			"Курс должен быть больше 0, отправьте ещё раз",
			&telebot.SendOptions{
				DisableWebPagePreview: false,
				ParseMode:             telebot.ModeHTML,
			},
		)
	}
	if studentCourse > 6 {
		return c.Send(
			"Возраст должен быть меньше 6, отправьте ещё раз",
			&telebot.SendOptions{
				DisableWebPagePreview: false,
				ParseMode:             telebot.ModeHTML,
			},
		)
	}

	student := b.data[c.Chat().ID]
	studentID, err := b.api.Create(
		student.Name,
		student.Sex,
		student.Age,
		studentCourse,
	)
	if err != nil {
		return err
	}
	b.clear(c.Chat().ID)

	return c.Send(
		fmt.Sprintf("Студент с ID=%d успешно создан", studentID),
		&telebot.SendOptions{
			DisableWebPagePreview: false,
			ParseMode:             telebot.ModeHTML,
		},
	)
}

// COMMAND GET

const (
	StudentGetIdState States = "student_get_id"
)

func (b *Bot) getHandler(c telebot.Context) error {
	b.states[c.Chat().ID] = StudentGetIdState
	return c.Send(
		utils.MultiLine(
			"Получение студента по ID",
			"",
			"Пришлите ID студента",
		),
		&telebot.SendOptions{
			DisableWebPagePreview: false,
			ParseMode:             telebot.ModeHTML,
		},
	)
}

func (b *Bot) getIdHandler(c telebot.Context) error {
	studentID, err := strconv.Atoi(c.Text())
	if err != nil {
		return c.Send(
			"ID должен быть числовым, отправьте ещё раз",
			&telebot.SendOptions{
				DisableWebPagePreview: false,
				ParseMode:             telebot.ModeHTML,
			},
		)
	}
	student, err := b.api.Get(int64(studentID))
	if err != nil {
		c.Send(
			"Произошла ошибка, обратитесь в тех. Поддержку",
			&telebot.SendOptions{
				DisableWebPagePreview: false,
				ParseMode:             telebot.ModeHTML,
			},
		)
		return err
	}
	c.Send(
		utils.MultiLine(
			fmt.Sprintf("ID: <b><i>%d</i></b>", student.ID),
			fmt.Sprintf("Имя: <b><i>%s</i></b>", student.Name),
			fmt.Sprintf("Пол: <b><i>%s</i></b>", api.FormatSexToRu(student.Sex)),
			fmt.Sprintf("Возраст: <b><i>%d</i></b>", student.Age),
			fmt.Sprintf("Курс: <b><i>%d</i></b>", student.Course),
		),
		&telebot.SendOptions{
			DisableWebPagePreview: false,
			ParseMode:             telebot.ModeHTML,
		},
	)
	b.clear(c.Chat().ID)
	return nil
}

// COMMAND GET_ALL

func (b *Bot) getAllHandler(c telebot.Context) error {
	students, err := b.api.GetAll()
	if err != nil {
		c.Send(
			"Произошла ошибка, обратитесь в тех. Поддержку",
			&telebot.SendOptions{
				DisableWebPagePreview: false,
				ParseMode:             telebot.ModeHTML,
			},
		)
		return err
	}

	text := "Список студентов:\n"
	for _, student := range students {
		text = utils.MultiLine(
			text,
			fmt.Sprintf(
				"ID: <b><i>%d</i></b>, Имя: <b><i>%s</i></b>, Пол: <b><i>%s</i></b>, Возраст: <b><i>%d</i></b>, Курс: <b><i>%d</i></b>",
				student.ID, student.Name, api.FormatSexToRu(student.Sex), student.Age, student.Course,
			),
		)
	}

	return c.Send(
		text,
		&telebot.SendOptions{
			DisableWebPagePreview: false,
			ParseMode:             telebot.ModeHTML,
		},
	)
}

// COMMAND UPDATE

const (
	StudentUpdateId     States = "student_update_id"
	StudentUpdateName   States = "student_update_name"
	StudentUpdateSex    States = "student_update_sex"
	StudentUpdateAge    States = "student_update_age"
	StudentUpdateCourse States = "student_update_course"
)

func (b *Bot) updateHandler(c telebot.Context) error {
	return nil
}

// COMMAND DELETE

const (
	StudentDeleteIdState States = "student_delete_id"
)

func (b *Bot) deleteHandler(c telebot.Context) error {
	b.states[c.Chat().ID] = StudentDeleteIdState
	b.data[c.Chat().ID] = api.StudyApiStudent{}
	return c.Send(
		utils.MultiLine(
			"Удаление студента по ID",
			"",
			"Пришлите ID студента",
		),
		&telebot.SendOptions{
			DisableWebPagePreview: false,
			ParseMode:             telebot.ModeHTML,
		},
	)
}

func (b *Bot) deleteIdHandler(c telebot.Context) error {
	studentID, err := strconv.Atoi(c.Text())
	if err != nil {
		return c.Send(
			"ID должен быть числовым, отправьте ещё раз",
			&telebot.SendOptions{
				DisableWebPagePreview: false,
				ParseMode:             telebot.ModeHTML,
			},
		)
	}
	student, err := b.api.Get(int64(studentID))
	if err != nil {
		c.Send(
			"Произошла ошибка, обратитесь в тех. Поддержку",
			&telebot.SendOptions{
				DisableWebPagePreview: false,
				ParseMode:             telebot.ModeHTML,
			},
		)
		return err
	}
	b.data[c.Chat().ID] = *student
	c.Send(
		utils.MultiLine(
			fmt.Sprintf("ID: <b><i>%d</i></b>", student.ID),
			fmt.Sprintf("Имя: <b><i>%s</i></b>", student.Name),
			fmt.Sprintf("Пол: <b><i>%s</i></b>", api.FormatSexToRu(student.Sex)),
			fmt.Sprintf("Возраст: <b><i>%d</i></b>", student.Age),
			fmt.Sprintf("Курс: <b><i>%d</i></b>", student.Course),
			"",
			"Вы уверены, что хотите удалить?",
		),
		&telebot.SendOptions{
			DisableWebPagePreview: false,
			ParseMode:             telebot.ModeHTML,
			ReplyMarkup:           deleteSuccessKeyboard(),
		},
	)
	return nil
}

func (b *Bot) deleteYesHandler(c telebot.Context) error {
	student := b.data[c.Chat().ID]
	err := b.api.Delete(student.ID)
	if err != nil {
		c.Send(
			"Произошла ошибка, обратитесь в тех. Поддержку",
			&telebot.SendOptions{
				DisableWebPagePreview: false,
				ParseMode:             telebot.ModeHTML,
			},
		)
		return err
	}
	b.clear(c.Chat().ID)
	return c.Send(
		fmt.Sprintf("Студент с ID=%d успешно удален", student.ID),
		&telebot.SendOptions{
			DisableWebPagePreview: false,
			ParseMode:             telebot.ModeHTML,
		},
	)
}

func (b *Bot) deleteNoHandler(c telebot.Context) error {
	b.clear(c.Chat().ID)
	return c.Delete()
}
