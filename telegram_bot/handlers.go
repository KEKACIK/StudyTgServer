package telegram_bot

import (
	"StudyTgServer/utils"
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// COMMAND START

func (tb *TgBot) startHandler(ctx context.Context, _ *bot.Bot, update *models.Update) {
	tb.client.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: utils.MultiLine(
			"Добро пожаловать в <b>StudyTgBot</b>",
			"",
			"Вот список комманд для использования бота",
			"/create - Добавить нового студента",
			"/get - Получить студента по ID",
			"/get_all - Получить всех студентов",
			"/update - Обновить студента по ID",
			"/delete - Удалить студента по ID",
		),
		ParseMode: models.ParseModeHTML,
	})
}

// COMMAND CREATE

func (tb *TgBot) createHandler(ctx context.Context, _ *bot.Bot, update *models.Update) {

}

// COMMAND GET

func (tb *TgBot) getHandler(ctx context.Context, _ *bot.Bot, update *models.Update) {

}

// COMMAND GET_ALL

func (tb *TgBot) getAllHandler(ctx context.Context, _ *bot.Bot, update *models.Update) {

}

// COMMAND UPDATE

func (tb *TgBot) updateHandler(ctx context.Context, _ *bot.Bot, update *models.Update) {

}

// COMMAND DELETE

func (tb *TgBot) deleteHandler(ctx context.Context, _ *bot.Bot, update *models.Update) {

}
