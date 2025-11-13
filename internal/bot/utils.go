package bot

import (
	"StudyTgServer/internal/api"
	"StudyTgServer/tools"
	"fmt"

	"gopkg.in/telebot.v3"
)

func (b *Bot) SetCommands() error {
	commands := []telebot.Command{
		{
			Text:        "start",
			Description: "Запустить бота",
		},
		{
			Text:        "create",
			Description: "Добавить нового студента",
		},
		{
			Text:        "get",
			Description: "Получить студента по ID",
		},
		{
			Text:        "get_all",
			Description: "Получить всех студентов",
		},
		{
			Text:        "delete",
			Description: "Удалить студента по ID",
		},
	}
	return b.Bot.SetCommands(commands)
}

// Student

func (b *Bot) getStudentText(student api.StudyApiStudent) string {
	return tools.MultiLine(
		fmt.Sprintf("ID: <b><i>%d</i></b>", student.ID),
		fmt.Sprintf("Имя: <b><i>%s</i></b>", student.Name),
		fmt.Sprintf("Пол: <b><i>%s</i></b>", api.FormatSexToRu(student.Sex)),
		fmt.Sprintf("Возраст: <b><i>%d</i></b>", student.Age),
		fmt.Sprintf("Курс: <b><i>%d</i></b>", student.Course),
	)
}

// Clear
func (b *Bot) clearData(userId int64) {
	delete(b.data, userId)
}

func (b *Bot) clearState(userId int64) {
	delete(b.states, userId)
}

func (b *Bot) clear(userId int64) {
	b.clearData(userId)
	b.clearState(userId)
}
