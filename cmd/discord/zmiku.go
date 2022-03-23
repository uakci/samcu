package main

import (
	"fmt"
  "log"
	"os"
	"os/signal"
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
	e := discord.UpdateListeningStatus("tinju'i toi")
	if e != nil {
		log.Print(e)
	}
}

func optionFor(argopt samcu.CommandOption, required bool) discordgo.ApplicationCommandOption {
  var choices []*discordgo.ApplicationCommandOptionChoice
  if argopt.Values != nil {
    choices = make([]*discordgo.ApplicationCommandOptionChoice, 0, len(argopt.Values))
    for _, pair := range argopt.Values {
      choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
        Name: pair[1],
        Value: pair[0],
      })
    }
  }

  var ty discordgo.ApplicationCommandOptionType
  switch argopt.Type {
  case samcu.StringType:
    ty = discordgo.ApplicationCommandOptionString
  case samcu.BoolType:
    ty = discordgo.ApplicationCommandOptionBoolean
  }

  return discordgo.ApplicationCommandOption{
    Type: ty,
    Name: argopt.Name,
    Description: argopt.Desc,
    Required: required,
    Choices: choices,
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
	discord.AddHandler(handleInteraction)
  log.Print("connected!")

  generatedCommands := make([]*discordgo.ApplicationCommand, 0, len(samcu.AllCommands))
  for _, cmd := range samcu.AllCommands {
    options := make([]*discordgo.ApplicationCommandOption, 0, len(cmd.Args) + len(cmd.Opts))
    for _, arg := range cmd.Args {
      option := optionFor(arg, true)
      options = append(options, &option)
    }
    for _, opt := range cmd.Opts {
      option := optionFor(opt, false)
      options = append(options, &option)
    }
    generatedCommands = append(generatedCommands, &discordgo.ApplicationCommand{
      Name: cmd.Name,
      Description: cmd.Desc,
      Options: options,
    })
  }
  _, err = discord.ApplicationCommandBulkOverwrite(discord.State.User.ID, "", generatedCommands)
  if err != nil {
    log.Panic(err)
  }

	quitter := make(chan struct{}, 1)
	go updateStatus(discord, quitter)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	quitter <- struct{}{}

  log.Print("quitting!")
	discord.Close()
}

func handleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
  data := i.Interaction.Data
  handler, ok := samcu.AllCommands[data.Name]
  if !ok { return }

  args := map[string]any{}
  for _, option := range data.Options {
    args[option.Name] = option.Value
  }
  out, err := handler.Func(args)

  var resp string
  if err != nil {
    resp = fmt.Sprintf("⚠ %v", err)
  } else {
    resp = out
  }

  s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
    Type: discordgo.InteractionResponseChannelMessageWithSource,
    Data: &discordgo.InteractionApplicationCommandResponseData{
      Content: resp,
    },
  })
}
