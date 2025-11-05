package telegram_bot

import (
	"StudyTgServer/api"
	"context"

	"github.com/go-telegram/bot"
)

type TgBot struct {
	client *bot.Bot
	api    *api.StudyApiServer
}

func NewTgBot(token string, api *api.StudyApiServer) (*TgBot, error) {
	client, err := bot.New(token)
	if err != nil {
		return nil, err
	}
	tgBot := TgBot{
		client: client,
		api:    api,
	}
	tgBot.registerHandlers()
	return &tgBot, nil
}

func (tb *TgBot) Start(ctx context.Context) {
	tb.client.Start(ctx)
}

func (tb *TgBot) registerHandlers() {
	tb.client.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, tb.startHandler)
	tb.client.RegisterHandler(bot.HandlerTypeMessageText, "/create", bot.MatchTypeExact, tb.createHandler)
	tb.client.RegisterHandler(bot.HandlerTypeMessageText, "/get", bot.MatchTypeExact, tb.getHandler)
	tb.client.RegisterHandler(bot.HandlerTypeMessageText, "/get_all", bot.MatchTypeExact, tb.getAllHandler)
	tb.client.RegisterHandler(bot.HandlerTypeMessageText, "/update", bot.MatchTypeExact, tb.updateHandler)
	tb.client.RegisterHandler(bot.HandlerTypeMessageText, "/delete", bot.MatchTypeExact, tb.deleteHandler)
}
