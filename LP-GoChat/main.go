package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()//Cria o server
	go s.run()//manda o server rodar as funções da lista de comandos
//Faz o server esperar pelo Client
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Incapaz de iniciar o servidor: %s", err.Error())
	}
	defer listener.Close()
	log.Printf("Servidor iniciado na porta :8888")
	for {
		//For para sempre que fica esperando o Cliente conectar
		conn, err := listener.Accept()//Cliente connectando
		if err != nil {
			log.Printf("Falha para aceitar a conexao : %s", err.Error())
			continue
		}

		go s.newClient(conn) // PARTE PRINCIPAL DO CODIGO
	}
}
