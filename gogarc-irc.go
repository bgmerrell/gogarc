package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/bgmerrell/gogarc/lib/command"
	"github.com/bgmerrell/gogarc/lib/dispatch"
	"github.com/bgmerrell/gogarc/lib/game"
	irc "github.com/fluffle/goirc/client"
	"github.com/fluffle/goirc/logging/glog"

	_ "github.com/bgmerrell/gogarc/lib/handlers/join"
	_ "github.com/bgmerrell/gogarc/lib/handlers/start"
	_ "github.com/bgmerrell/gogarc/lib/handlers/stats"
)

var host *string = flag.String("host", "irc.frws.com", "IRC server")
var channel *string = flag.String("channel", "#gogarc", "IRC channel")

type privMsgHandler struct {
	ircGame *game.Game
}

func (p *privMsgHandler) Handle(conn *irc.Conn, line *irc.Line) {
	var cmd string
	var args string
	// command consists of a command and its optional argument(s)
	const maxCmdLen = 2
	const cmdIdx = 0
	const argIdx = 1
	ed := dispatch.NewEventDispatcher()
	lineText := line.Text()
	// TODO: make command prefix configurable
	// Ignore messages not for gogarc
	if !strings.HasPrefix(lineText, ".") {
		return
	}
	lineSplit := strings.SplitN(lineText, " ", maxCmdLen)
	// Chop off the command prefix
	cmd = lineSplit[cmdIdx][1:]
	if cmd == "" {
		return
	}
	if len(lineSplit) == maxCmdLen {
		args = lineSplit[argIdx]
	}
	go ed.Send(p.ircGame, &command.Command{line.Nick, cmd, args})
	for msg := range ed.Output {
		conn.Privmsg("#gogarc", msg+"\r\n")
	}
}

func main() {
	flag.Parse()
	glog.Init()

	// create new IRC connection
	c := irc.SimpleClient("Gogarc", "gogarc")
	c.EnableStateTracking()
	c.HandleFunc("connected",
		func(conn *irc.Conn, line *irc.Line) { conn.Join(*channel) })

	// Set up a handler to notify of disconnect events.
	quit := make(chan bool)
	c.HandleFunc("disconnected",
		func(conn *irc.Conn, line *irc.Line) { quit <- true })

	// Set up a handler to read messages
	c.HandleBG(irc.PRIVMSG, &privMsgHandler{game.NewGame()})

	// set up a goroutine to read commands from stdin
	in := make(chan string, 4)
	reallyquit := false
	go func() {
		con := bufio.NewReader(os.Stdin)
		for {
			s, err := con.ReadString('\n')
			if err != nil {
				// wha?, maybe ctrl-D...
				close(in)
				break
			}
			// no point in sending empty lines down the channel
			if len(s) > 2 {
				in <- s[0 : len(s)-1]
			}
		}
	}()

	// set up a goroutine to do parsey things with the stuff from stdin
	go func() {
		for cmd := range in {
			if cmd[0] == ':' {
				switch idx := strings.Index(cmd, " "); {
				case cmd[1] == 'd':
					fmt.Printf(c.String())
				case cmd[1] == 'f':
					if len(cmd) > 2 && cmd[2] == 'e' {
						// enable flooding
						c.Config().Flood = true
					} else if len(cmd) > 2 && cmd[2] == 'd' {
						// disable flooding
						c.Config().Flood = false
					}
					for i := 0; i < 20; i++ {
						c.Privmsg("#", "flood test!")
					}
				case idx == -1:
					continue
				case cmd[1] == 'q':
					reallyquit = true
					c.Quit(cmd[idx+1 : len(cmd)])
				case cmd[1] == 'j':
					c.Join(cmd[idx+1 : len(cmd)])
				case cmd[1] == 'p':
					c.Part(cmd[idx+1 : len(cmd)])
				case cmd[1] == 'm':
					c.Privmsg("#gogarc", cmd[idx+1:len(cmd)])
				}
			} else {
				c.Raw(cmd)
			}
		}
	}()

	for !reallyquit {
		// connect to server
		if err := c.ConnectTo(*host); err != nil {
			fmt.Printf("Connection error: %s\n", err)
			return
		}

		// wait on quit channel
		<-quit
	}
}
