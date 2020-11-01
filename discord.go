package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
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

func main() {
	token, ok := os.LookupEnv("SAMCU_TOKEN")
	if !ok {
		panic(fmt.Sprintf("Need token in env var SAMCU_TOKEN"))
	}
	discord, err := discordgo.New("Bot " + token)
	must(err)
	discord.ShouldReconnectOnError = true

	must(discord.Open())
	must(discord.UpdateListeningStatus("tinju’i toi"))
	discord.AddHandler(handleMessage)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
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
    fmt.Println("→", msg)
    _, err := s.ChannelMessageSend(m.Message.ChannelID, msg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
	}
	var errmsg = func() {
		respond("<#" + helpChannel + ">")
	}

	fields := strings.Fields(stripped)
	if len(fields) == 0 {
		errmsg()
		return
	}
  cmd := strings.TrimSuffix(fields[0], ":")
  fields = fields[1:]

  fmt.Println(cmd, fields)

	switch cmd {
	case "lujvo":
		jvozba(respond, strings.Join(fields, " "))
	default:
		errmsg()
	}
}
