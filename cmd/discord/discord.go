package main

import (
	"fmt"
	"log"
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
	e := discord.UpdateListeningStatus("tinju'i toi")
	if e != nil {
		log.Print(e)
	}
}

func main() {
	samcu.Init()

	token, ok := os.LookupEnv("SAMCU_TOKEN")
	if !ok {
		log.Panicf("Need token in env var SAMCU_TOKEN")
	}
	discord, err := discordgo.New("Bot " + token)
	must(err)
	discord.ShouldReconnectOnError = true

	must(discord.Open())
	discord.AddHandler(handleMessage)
	log.Print("connected!")

	quitter := make(chan struct{}, 1)
	go updateStatus(discord, quitter)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	quitter <- struct{}{}

	log.Print("quitting!")
	discord.Close()
}

func handleMessage(s *discordgo.Session, e *discordgo.MessageCreate) {
	args := strings.Fields(e.Message.Content)
	if len(args) == 0 || !strings.HasPrefix(args[0], "-") {
		return
	}
	args[0] = args[0][1:]
	ok, msg, err := samcu.Respond(args)

	var resp string
	switch {
	case !ok:
		resp = fmt.Sprintf("❓ la'oi -%s mo", args[0])
	case err != nil:
		resp = fmt.Sprintf("⚠ %s", err.Error())
	default:
		resp = msg
	}

	s.ChannelMessageSend(e.Message.ChannelID, resp)
}
