package bot

import "gopkg.in/telebot.v3"

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
			Text:        "update",
			Description: "Обновить студента по ID",
		},
		{
			Text:        "delete",
			Description: "Удалить студента по ID",
		},
	}
	return b.Bot.SetCommands(commands)
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
