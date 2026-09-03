package client

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	maxbot "github.com/max-messenger/max-bot-api-client-go/v2"
	"github.com/max-messenger/max-bot-api-client-go/v2/model"
)

type Bot struct {
	Api *maxbot.Api
}

func NewBot(token string) (*Bot, error) {
	opts := []maxbot.Opt{
		maxbot.WithHTTPClient(&http.Client{}),
	}
	api, err := maxbot.NewApi(token, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect %w", err)
	}
	return &Bot{Api: api}, nil
}

func (bot *Bot) EventProcessor() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	api := bot.Api
	info, err := api.Bots.GetMyInfo(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("info: %+v", info)

	handle := func(ctx context.Context, update model.Update) {
		fmt.Printf("Recieved: [%s] %#v\n", update.UpdateType, update)
		switch update.UpdateType {
		case model.UpdateMessageCreated:
			msg := maxbot.NewMessage().
				SetText("Нет, это ты " + strings.ToLower(update.Message.Body.Text)).
				SetChat(update.ChatID).
				SetUser(update.UserID)

			res, cErr := api.Messages.Send(ctx, msg)
			if cErr != nil {

			}

			log.Printf("%v\n", res)
		}
	}
	var updates []model.Update
	var marker int64
	for {
		select {
		case <-ctx.Done():
		default:
			updates, marker, err = api.Subscriptions.GetUpdates(ctx, marker)
			if err != nil {
				log.Println("GetUpdates: ", err)
				return
			}

			for _, update := range updates {
				handle(ctx, update)
			}
		}
	}

}
