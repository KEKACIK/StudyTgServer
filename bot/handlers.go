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
	b.Handle(&telebot.Btn{Unique: "get_update"}, b.getUpdateHandler)
	b.Handle(&telebot.Btn{Unique: "get_update_sex"}, b.getUpdateSexHandler)
	b.Handle("/get_all", b.getAllHandler)
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
	case CreateNameState:
		return b.createNameHandler(c)
	case CreateAgeState:
		return b.createAgeHandler(c)
	case CreateCourseState:
		return b.createCourseHandler(c)
	// Get handlers
	case GetIdState:
		return b.getIdHandler(c)
	case GetUpdateNameState:
		return b.getUpdateNameHandler(c)
	case GetUpdateAgeState:
		return b.getUpdateAgeHandler(c)
	case GetUpdateCourseState:
		return b.getUpdateCourseHandler(c)
	// Delete handlers
	case DeleteIdState:
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
	CreateNameState   States = "create_name"
	CreateAgeState    States = "create_age"
	CreateCourseState States = "create_course"
)

func (b *Bot) createHandler(c telebot.Context) error {
	b.states[c.Chat().ID] = CreateNameState
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
	studentName, err := studentNameValidation(c.Text())
	if err != nil {
		return c.Send(
			err.Error(),
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
			ReplyMarkup:           selectSexKeyboard("create_sex"),
		},
	)
}

func (b *Bot) createSexHandler(c telebot.Context) error {
	studentSex, err := studentSexValidation(c.Callback().Data)
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

	student := b.data[c.Chat().ID]
	student.Sex = studentSex
	b.data[c.Chat().ID] = student
	b.states[c.Chat().ID] = CreateAgeState
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
	studentAge, err := studentAgeValidation(c.Text())
	if err != nil {
		return c.Send(
			err.Error(),
			&telebot.SendOptions{
				DisableWebPagePreview: false,
				ParseMode:             telebot.ModeHTML,
			},
		)
	}

	student := b.data[c.Chat().ID]
	student.Age = studentAge
	b.data[c.Chat().ID] = student
	b.states[c.Chat().ID] = CreateCourseState

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
	studentCourse, err := studentAgeValidation(c.Text())
	if err != nil {
		return c.Send(
			err.Error(),
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

// COMMAND GET/UPDATE

const (
	GetIdState           States = "get_id"
	GetUpdateNameState   States = "get_update_name"
	GetUpdateAgeState    States = "get_update_age"
	GetUpdateCourseState States = "get_update_course"
)

func (b *Bot) getHandler(c telebot.Context) error {
	b.states[c.Chat().ID] = GetIdState
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
	b.data[c.Chat().ID] = *student

	c.Send(
		b.getStudentText(*student),
		&telebot.SendOptions{
			DisableWebPagePreview: false,
			ParseMode:             telebot.ModeHTML,
			ReplyMarkup:           getUpdateFieldKeyboard(),
		},
	)
	return nil
}

func (b *Bot) getUpdateHandler(c telebot.Context) error {
	student := b.data[c.Chat().ID]
	text := fmt.Sprintf("Обновление пользователя ID=%d", student.ID)
	replyMarkup := &telebot.ReplyMarkup{}

	switch c.Callback().Data {
	case "name":
		text = utils.MultiLine(
			text,
			fmt.Sprintf("Имя: <b><i>%s</i></b>", student.Name),
			"",
			"Пришлите новое имя",
		)
		b.states[c.Chat().ID] = GetUpdateNameState
	case "sex":
		text = utils.MultiLine(
			text,
			fmt.Sprintf("Пол: <b><i>%s</i></b>", api.FormatSexToRu(student.Sex)),
			"",
			"Выберите новый пол",
		)
		replyMarkup = selectSexKeyboard("get_update_sex")
	case "age":
		text = utils.MultiLine(
			text,
			fmt.Sprintf("Возраст: <b><i>%d</i></b>", student.Age),
			"",
			"Пришлите новый возраст",
		)
		b.states[c.Chat().ID] = GetUpdateAgeState
	case "course":
		text = utils.MultiLine(
			text,
			fmt.Sprintf("Курс: <b><i>%d</i></b>", student.Course),
			"",
			"Пришлите новый курс",
		)
		b.states[c.Chat().ID] = GetUpdateCourseState
	}

	return c.Send(
		text,
		&telebot.SendOptions{
			DisableWebPagePreview: false,
			ParseMode:             telebot.ModeHTML,
			ReplyMarkup:           replyMarkup,
		},
	)
}

func (b *Bot) getUpdateNameHandler(c telebot.Context) error {
	studentName, err := studentNameValidation(c.Text())
	if err != nil {
		return c.Send(
			err.Error(),
			&telebot.SendOptions{
				DisableWebPagePreview: false,
				ParseMode:             telebot.ModeHTML,
			},
		)
	}
	student := b.data[c.Chat().ID]
	studentNew, err := b.api.Update(
		student.ID,
		studentName,
		student.Sex,
		student.Age,
		student.Course,
	)
	if err != nil {
		return err
	}
	b.data[c.Chat().ID] = *studentNew
	c.Delete()
	return c.Send(
		b.getStudentText(*studentNew),
		&telebot.SendOptions{
			DisableWebPagePreview: false,
			ParseMode:             telebot.ModeHTML,
			ReplyMarkup:           getUpdateFieldKeyboard(),
		},
	)
}

func (b *Bot) getUpdateSexHandler(c telebot.Context) error {
	studentSex, err := studentSexValidation(c.Callback().Data)
	if err != nil {
		c.Edit(
			"Произошла ошибка, обратитесь в тех. Поддержку",
			&telebot.SendOptions{
				DisableWebPagePreview: false,
				ParseMode:             telebot.ModeHTML,
			},
		)
		return err
	}

	student := b.data[c.Chat().ID]
	studentNew, err := b.api.Update(
		student.ID,
		student.Name,
		studentSex,
		student.Age,
		student.Course,
	)
	if err != nil {
		return err
	}
	b.data[c.Chat().ID] = *studentNew
	return c.Edit(
		b.getStudentText(*studentNew),
		&telebot.SendOptions{
			DisableWebPagePreview: false,
			ParseMode:             telebot.ModeHTML,
			ReplyMarkup:           getUpdateFieldKeyboard(),
		},
	)
}

func (b *Bot) getUpdateAgeHandler(c telebot.Context) error {
	studentAge, err := studentAgeValidation(c.Text())
	if err != nil {
		return c.Send(
			err.Error(),
			&telebot.SendOptions{
				DisableWebPagePreview: false,
				ParseMode:             telebot.ModeHTML,
			},
		)
	}
	student := b.data[c.Chat().ID]
	studentNew, err := b.api.Update(
		student.ID,
		student.Name,
		student.Sex,
		studentAge,
		student.Course,
	)
	if err != nil {
		return err
	}
	b.data[c.Chat().ID] = *studentNew
	c.Delete()
	return c.Send(
		b.getStudentText(*studentNew),
		&telebot.SendOptions{
			DisableWebPagePreview: false,
			ParseMode:             telebot.ModeHTML,
			ReplyMarkup:           getUpdateFieldKeyboard(),
		},
	)
}

func (b *Bot) getUpdateCourseHandler(c telebot.Context) error {
	studentCourse, err := studentCourseValidation(c.Text())
	if err != nil {
		return c.Send(
			err.Error(),
			&telebot.SendOptions{
				DisableWebPagePreview: false,
				ParseMode:             telebot.ModeHTML,
			},
		)
	}
	student := b.data[c.Chat().ID]
	studentNew, err := b.api.Update(
		student.ID,
		student.Name,
		student.Sex,
		student.Age,
		studentCourse,
	)
	if err != nil {
		return err
	}
	b.data[c.Chat().ID] = *studentNew
	fmt.Println(student)
	fmt.Println(studentNew)
	c.Delete()
	return c.Send(
		b.getStudentText(*studentNew),
		&telebot.SendOptions{
			DisableWebPagePreview: false,
			ParseMode:             telebot.ModeHTML,
			ReplyMarkup:           getUpdateFieldKeyboard(),
		},
	)
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

// COMMAND DELETE

const (
	DeleteIdState States = "delete_id"
)

func (b *Bot) deleteHandler(c telebot.Context) error {
	b.states[c.Chat().ID] = DeleteIdState
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
