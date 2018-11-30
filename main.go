package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/spy16/parens"
)

// Variables used for command line parameters
var (
	Token	string
	lisp	*parens.Interpreter
	root	*parens.Interpreter
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	scn := initScope()
	lisp = parens.New(scn)
	scr := parens.NewScope(scn)
	scr.Bind("say", func(args ...interface{}) string{
		return fmt.Sprint(args...)
	})
	root = parens.New(scr)

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	var(
		val	interface{}
		err	error
	)

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content[0] != '('{
		return
	}
	if m.Content[:6] == "(sudo " { 
		if !isSudoer(m.Author.ID){
			s.ChannelMessageSend(m.ChannelID, "User not in sudoers file. This incident **will be reported**")
			return
		}
		val, err = root.Execute(m.Content[6:len(m.Content)-1])
	} else {
		err = s.ChannelTyping(m.ChannelID)
		val, err = lisp.Execute(m.Content)
	}

	if err != nil{
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("An error occured:```\n%s```\n", err))
		return
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprint(val))
}

func isSudoer(name string) bool {
	return name == "268379122651103233"
}

