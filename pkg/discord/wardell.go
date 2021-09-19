package discord

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"uno/pkg/client"

	"github.com/spf13/viper"

	"github.com/bwmarrin/discordgo"
)

var (
	usageText = `Usage: !wardell|:wardell:|@wardell [cmd]
supporting cmd:
  elo playername:tagLine
  history playername:tagLine`
)

func Wardell(token string) error {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return err
	}
	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		log.Println("error opening connection,", err)
		return err
	}

	log.Println("Wardell is shouting.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
	return nil
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	wardellMatch := regexp.MustCompile(viper.Get("discord_mension_string").(string))
	if wardellMatch.MatchString(m.Content) {
		cmdString := m.Content[wardellMatch.FindStringIndex(m.Content)[1]:]
		cmdString = strings.TrimSpace(cmdString)
		cmdSlice := strings.Split(cmdString, " ")
		switch cmdSlice[0] {
		case "help":
			_, err := s.ChannelMessageSend(m.ChannelID, usageText)
			if err != nil {
				log.Println(m, err)
			}
		case "elo":
			message, err := elo(m, cmdSlice)
			if err != nil {
				log.Println(cmdSlice, err)
			}
			_, err = s.ChannelMessageSend(m.ChannelID, message)
			if err != nil {
				log.Println(message, err)
			}
		case "history":
			message, err := history(m, cmdSlice)
			if err != nil {
				log.Println(cmdSlice, err)
			}
			_, err = s.ChannelMessageSend(m.ChannelID, message)
			if err != nil {
				log.Println(message, err)
			}
		default:
			_, err := s.ChannelMessageSend(m.ChannelID, usageText)
			if err != nil {
				log.Println(m, err)
			}
		}
	}
}

func elo(m *discordgo.MessageCreate, cmdSlice []string) (string, error) {
	if len(cmdSlice) < 2 {
		return "", fmt.Errorf("not enough argument: %v", cmdSlice)
	}
	name, tag := parseNameTag(cmdSlice[1])
	log.Println("Elo of ", name, tag)
	c := client.New()
	mmr, err := c.GetMMRDataByNameTag("ap", name, tag)
	if err != nil {
		return "", err
	}
	return client.FormatJSON(mmr.CurrentData, true)
}

func parseNameTag(nameTag string) (string, string) {
	s := strings.Split(nameTag, "#")
	return s[0], s[1]
}

func history(m *discordgo.MessageCreate, cmdSlice []string) (string, error) {
	if len(cmdSlice) < 2 {
		return "", fmt.Errorf("not enough argument: %v", cmdSlice)
	}
	name, tag := parseNameTag(cmdSlice[1])
	log.Println("History of ", name, tag)
	c := client.New()
	history, err := c.GetMatchHistory("ap", name, tag, "competitive")
	if err != nil {
		return "", err
	}
	summarized := []string{}
	for _, match := range history {
		summarized = append(summarized, summarizeMatch(match))
	}
	return strings.Join(summarized, "\n"), nil
}

func summarizeMatch(m *client.Match) string {
	return strings.Join([]string{
		m.Metadata.Matchid,
		m.Metadata.Map,
	}, " ")
}
