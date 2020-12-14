package main

import (
	"bufio"
	"net"
	"strings"
)

type client struct {
	conn net.Conn
	nick string
	commands chan<- command
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		//msg = strings.Trim(msg, "\r")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/nick":
			c.commands <- command{
				id:     CMD_NICK,
				client: c,
				args:   args,
			}
		case "":

		case "/quit":
			c.commands <- command{
				id:     CMD_QUIT,
				client: c,
			}
		default:
			var s []string
			s = append(s, cmd)
			for i := 0; i < len(args); i++ {
				s = append(s, args[i])
			}
			c.commands <- command{
				id:     CMD_MSG,
				client: c,
				args:   s,
			}
		}
	}
}

func (c *client) err(err error) {
	c.conn.Write([]byte("err: " + err.Error() + "\n"))
}

func (c *client) msg(msg string) {
	c.conn.Write([]byte("> " + msg + "\n"))
}
