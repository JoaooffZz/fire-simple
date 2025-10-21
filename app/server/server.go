package server

import (
	"bufio"
	"encoding/json"
	svc "firesimple/services"
	"fmt"
	"net"
)

type ServerTCP struct {
	services *svc.FireSPServices
}

func New(services *svc.FireSPServices) ServerTCP {
	return ServerTCP{
		services: services,
	}
}

func (s *ServerTCP) Run() error {
	listener, err := net.Listen("tcp", ":5252")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *ServerTCP) handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	fmt.Println("Cliente conectado:", conn.RemoteAddr())

	for scanner.Scan() {
		msg := scanner.Bytes()

		// var p Pessoa
		var msgData MessageData
		err := json.Unmarshal(msg, &msgData)
		if err != nil {
			// fmt.Println("Erro ao decodificar JSON:", err)
			// conn.Write([]byte("Erro: JSON inválido\n"))
			// continue
			return
		}

		switch msgData.TypeEvent {
		case "ip.guarde":
			handlerIPGuarde(s.services, msgData.TypeService, msgData.ClientIP, conn.Write)
			continue
		default:
			return
		}

		// fmt.Printf("Recebido de %s: %+v\n", conn.RemoteAddr(), p)
	}

	// Após o loop, verificar se o scanner terminou por erro
	if err := scanner.Err(); err != nil {
		fmt.Println("Erro na conexão:", err)
	} else {
		fmt.Println("Cliente desconectou:", conn.RemoteAddr())
	}
}
