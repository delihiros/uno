package discord

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"uno/pkg/analysis/maps"
	"uno/pkg/entities"
	"uno/pkg/jsonutil"
	"uno/pkg/proxy"

	"github.com/spf13/viper"

	"github.com/bwmarrin/discordgo"
)

var (
	usageText = `Usage: !wardell|:wardell:|@wardell [cmd]
supporting cmd:
  elo playername#tagLine
  history playername#tagLine
  killmap matchID`
)

type Wardell struct {
	matcher *regexp.Regexp
	p       *proxy.Proxy
	s       *discordgo.Session
}

func New(token string) (*Wardell, error) {
	p, err := proxy.New()
	if err != nil {
		return nil, err
	}
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	return &Wardell{
		matcher: regexp.MustCompile(viper.Get("discord_mension_string").(string)),
		p:       p,
		s:       dg,
	}, nil
}

func (w *Wardell) Execute() error {
	w.s.AddHandler(w.messageCreate)
	w.s.Identify.Intents = discordgo.IntentsGuildMessages
	err := w.s.Open()
	if err != nil {
		log.Println("error opening connection,", err)
		return err
	}
	log.Println("Wardell is shouting.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	w.s.Close()
	return nil
}

func (w *Wardell) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == w.s.State.User.ID {
		return
	}
	if w.matcher.MatchString(m.Content) {
		cmdString := m.Content[w.matcher.FindStringIndex(m.Content)[1]:]
		cmdString = strings.TrimSpace(cmdString)
		cmdSlice := strings.Split(cmdString, " ")
		switch cmdSlice[0] {
		case "help":
			_, err := w.s.ChannelMessageSend(m.ChannelID, usageText)
			if err != nil {
				log.Println(m, err)
			}
		case "elo":
			if len(cmdSlice) < 2 {
				w.s.ChannelMessageSend(m.ChannelID, usageText)
				return
			}
			name, tag := parseNameTag(cmdSlice[1])
			message, err := w.elo(name, tag)
			if err != nil {
				log.Println(cmdSlice, err)
			}
			_, err = w.s.ChannelMessageSend(m.ChannelID, message)
			if err != nil {
				log.Println(message, err)
			}
		case "history":
			if len(cmdSlice) < 2 {
				w.s.ChannelMessageSend(m.ChannelID, usageText)
				return
			}
			name, tag := parseNameTag(cmdSlice[1])
			message, err := w.history(name, tag)
			if err != nil {
				log.Println(cmdSlice, err)
			}
			_, err = w.s.ChannelMessageSend(m.ChannelID, message)
			if err != nil {
				log.Println(message, err)
			}
		case "killmap":
			log.Println("killmap of", cmdSlice)
			if len(cmdSlice) < 2 {
				w.s.ChannelMessageSend(m.ChannelID, usageText)
				return
			}
			matchID := cmdSlice[1]
			filename, err := w.createKillMap(matchID)
			if err != nil {
				log.Println(cmdSlice, err)
			}
			f, err := os.Open(filename)
			if err != nil {
				log.Println(cmdSlice, err)
			}
			_, err = w.s.ChannelFileSend(m.ChannelID, cmdSlice[1]+".png", f)
			if err != nil {
				log.Println(cmdSlice, err)
			}
		default:
			_, err := w.s.ChannelMessageSend(m.ChannelID, usageText)
			if err != nil {
				log.Println(m, err)
			}
		}
	}
}

func (w *Wardell) elo(name, tag string) (string, error) {
	log.Println("Elo of", name, tag)
	mmr, err := w.p.GetMMRDataByNameTag("ap", name, tag)
	if err != nil {
		return "", err
	}
	return jsonutil.FormatJSON(mmr.CurrentData, true)
}

func parseNameTag(nameTag string) (string, string) {
	s := strings.Split(nameTag, "#")
	return s[0], s[1]
}

func (w *Wardell) history(name, tag string) (string, error) {
	log.Println("History of", name+"#"+tag)
	// history, err := w.p.GetMatchHistory("ap", name, tag, "competitive")
	history, err := w.p.GetMatchHistory("ap", name, tag, "")
	if err != nil {
		return "", err
	}
	summarized := []string{}
	for _, match := range history {
		summary, err := summarizeMatch(match, name, tag)
		if err != nil {
			log.Println(err)
		}
		summarized = append(summarized, summary)
	}
	return strings.Join(summarized, "\n"), nil
}

func summarizeMatch(m *entities.Match, name, tag string) (string, error) {
	player, err := m.FindPlayer(name, tag)
	if err != nil {
		return "", err
	}
	result := ""
	if player.Team == "Red" {
		result = fmt.Sprintf("%d-%d", m.Teams.Red.RoundsWon, m.Teams.Red.RoundsLost)
	} else {
		result = fmt.Sprintf("%d-%d", m.Teams.Red.RoundsLost, m.Teams.Red.RoundsWon)
	}
	return strings.Join([]string{
		m.Metadata.Matchid,
		m.Metadata.Map,
		m.Metadata.Mode,
		result,
	}, " "), nil
}

func (w *Wardell) createKillMap(matchID string) (string, error) {
	match, err := w.p.GetMatchByID(matchID)
	if err != nil {
		return "something went wrong", err
	}
	var m *maps.Map
	switch match.Metadata.Map {
	case "Ascent":
		m = maps.NewAscent()
	case "Haven":
		m = maps.NewHaven()
	case "Split":
		m = maps.NewSplit()
	case "Breeze":
		m = maps.NewBreeze()
	case "Bind":
		m = maps.NewBind()
	case "Icebox":
		m = maps.NewIcebox()
	case "Fracture":
		m = maps.NewFracture()
	default:
		return "not supported!", fmt.Errorf("map not supported")
	}
	for _, round := range match.Rounds {
		for _, status := range round.PlayerStats {
			for _, event := range status.KillEvents {
				victimLocation := event.VictimDeathLocation
				killerLocation := event.FindKillerLocation()
				m.DrawCircle(float64(victimLocation.X), float64(victimLocation.Y), 3, 1, 0, 0)
				// will be nil in DeathMatch
				if killerLocation != nil {
					m.DrawCircle(float64(killerLocation.X), float64(killerLocation.Y), 3, 0, 0, 1)
					m.DrawLine(float64(victimLocation.X), float64(victimLocation.Y), float64(killerLocation.X), float64(killerLocation.Y), 2, 0, 0.5, 0.5)
				} else {
					_, err := jsonutil.FormatJSON(event, true)
					if err != nil {
						panic(err)
					}
				}
			}
		}
	}
	// TODO
	m.SaveImage("_temporary/death.png")
	return "_temporary/death.png", nil
}
