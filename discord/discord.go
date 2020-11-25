package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/uakci/samcu"
)

func must(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	myRole      = "772142260128710719"
	helpChannel = "772167961771245578"
)

var helpMessage = "<#" + helpChannel + ">"

func updateStatus(discord *discordgo.Session, quit <-chan struct{}) {
	updaterStatuser(discord)
	ticker := time.NewTicker(time.Hour)
	for {
		select {
		case <-ticker.C:
			updaterStatuser(discord)
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func updaterStatuser(discord *discordgo.Session) {
	e := discord.UpdateListeningStatus("tinju’i toi")
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v", e)
	}
}

func main() {
	must(samcu.LoadJVS())

	token, ok := os.LookupEnv("SAMCU_TOKEN")
	if !ok {
		panic(fmt.Sprintf("Need token in env var SAMCU_TOKEN"))
	}
	discord, err := discordgo.New("Bot " + token)
	must(err)
	discord.ShouldReconnectOnError = true

	must(discord.Open())
	discord.AddHandler(handleMessage)
	quitter := make(chan struct{}, 1)
	go updateStatus(discord, quitter)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	quitter <- struct{}{}

	discord.Close()
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if s.State.User.ID == m.Author.ID {
		return
	}
	stripped := strings.TrimPrefix(m.Message.Content, "<@&"+myRole+">")
	if len(m.GuildID) > 0 && stripped == m.Message.Content {
		return
	}

	var respond = func(msg string) {
		for i := 0; i < len(msg); i += 1918 {
			_, err := s.ChannelMessageSend(m.Message.ChannelID, msg[i:min(i+1918, len(msg))])
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
			}
		}
	}

	response, ok := samcu.Respond(stripped)
	if ok {
		respond(response)
	} else {
		respond(helpMessage)
	}
}
