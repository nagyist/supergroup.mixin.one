package externals

import (
	"context"

	"github.com/MixinNetwork/bot-api-go-client"
	"github.com/MixinNetwork/supergroup.mixin.one/config"
)

func CreateConversation(ctx context.Context, category, participantId string) error {
	if config.AppConfig.Service.Environment == "test" {
		return nil
	}
	conversationId := bot.UniqueConversationId(config.AppConfig.Mixin.ClientId, participantId)
	participant := bot.Participant{
		UserId: participantId,
		Role:   "",
	}
	participants := []bot.Participant{
		participant,
	}
	mixin := config.AppConfig.Mixin
	_, err := bot.CreateConversation(ctx, category, conversationId, "", "", participants, mixin.ClientId, mixin.SessionId, mixin.SessionKey)
	if err != nil {
		return parseError(ctx, err.(bot.Error))
	}
	return nil
}
