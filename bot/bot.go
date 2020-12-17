package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/vincent-heng/discord-replacerbot/config"
)

// Bot is the discord bot manager
type Bot struct {
	config.Config
}

type _Message struct {
	Channel string
	Message string
}

type _Response struct {
	msgs []_Message
	err  error
}

func simpleErr(err error, msg string) _Response {
	if msg == "" && err != nil {
		msg = err.Error()
	}

	return _Response{
		err: err,
		msgs: []_Message{
			{Message: msg},
		},
	}
}

func simpleResponse(msg string) _Response {
	return _Response{
		msgs: []_Message{
			{Message: msg},
		},
	}
}

// New instantiates a bot with config
func New(conf config.Config) (*Bot, error) {
	return &Bot{
		Config: conf,
	}, nil
}

// Handler for discord events
func (b *Bot) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.ChannelID == "788356138009755670" {
		return
	}

	newContent := strings.ReplaceAll(m.Content, "^^", ":blush:")
	if newContent == m.Content {
		return
	}

	newContentWithUsername := m.Message.Author.Username + ": " + newContent

	if _, e := s.ChannelMessageSend(m.ChannelID, newContentWithUsername); e != nil {
		log.Error().Err(e).Msg("Cannot create message")
		return
	}
	if e := s.ChannelMessageDelete(m.ChannelID, m.ID); e != nil {
		log.Error().Err(e).Msg("Cannot delete message")
		return
	}

	uuid := uuid.New().String()
	log.Info().
		Str("cmdID", uuid).
		Str("originMessage", m.Content).
		Str("editedMessage", newContentWithUsername).
		Msg("Replaced a message")
}
