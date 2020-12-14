package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	members  map[net.Addr]*client
	commands chan command
}

func newServer() *server {
	return &server{
		members:  make(map[net.Addr]*client),
		commands: make(chan command),
	}
}
func (s *server) broadcast(sender *client, msg string) { // importante, envia pp todos ao msm tempo
	for addr, m := range s.members {
		if sender.conn.RemoteAddr() != addr {
			m.msg(msg)
		}
	}
}

func (s *server) run() { 
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args[1])
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT: //Modificar
			s.quit(cmd.client)
		}
	}
}

func (s *server) newClient(conn net.Conn) {
	//colorGreen := "\033[32m"
	//log.Printf("new client has joined: %s", conn.RemoteAddr().String())
	log.Printf("Cliente novo: %s", conn.RemoteAddr().String())
	//Criar funçao "qual o seu nick" e setar variavel temporario para input do usuario
	c := &client{
		conn:     conn,
		nick:     "anonymous", //variavel temporario
		commands: s.commands,
	}
	s.members[c.conn.RemoteAddr()] = c
	//VERIFICAR IMPLEMENTAÇÃO DA BIBLIOTECA DE IO
	c.msg("Commands: \n /nick Muda o nick \n /quit sai do servidor \n")
	c.readInput()
}

//Essa funçao nick ja cria o nick
func (s *server) nick(c *client, nick string) {
	//adicionar printf: qual o seu nick
	c.nick = nick
	c.msg(fmt.Sprintf("partiu, teu nome é: %s", nick))
}

func (s *server) msg(c *client, args []string) { //ENVIA MSG
	msg := strings.Join(args[1:len(args)], " ") //une a mesnagem
	s.broadcast(c, c.nick+": "+msg)
}

func (s *server) quit(c *client) {
	log.Printf("Cliente saiu: %s", c.conn.RemoteAddr().String())

	s.broadcast(c, "O "+c.nick+" saiu do chat ")
	c.msg("mas ja? =(")

	c.conn.Close()
}
