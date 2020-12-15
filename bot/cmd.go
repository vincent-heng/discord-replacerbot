package bot

import (
	"github.com/bwmarrin/discordgo"
)

func (b *Bot) handleUpStats(s *discordgo.Session, m *discordgo.MessageCreate, userID uint, stat string) _Response {

	return simpleResponse("Répartition effectuée !")
}
