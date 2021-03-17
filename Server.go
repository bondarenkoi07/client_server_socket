package Server

import (
	"log"
	"net"
	"os"
)

const SockAddr = "/home/ilya/OS/echo.sock"

type Server struct {
	l net.Listener
}

func NewServer() *Server {
	if err := os.RemoveAll(SockAddr); err != nil {

		log.Fatal(err)
	}
	conn, err := net.Listen("unix", SockAddr)
	if err != nil {
		panic(err)
	}
	return &Server{l: conn}
}

func (s *Server) ListenAndServe() {
	defer (*s).l.Close()

	for {
		client, err := (*s).l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		print(client.RemoteAddr().Network())
		go s.moderateMessage(client)
	}
}

func (s *Server) moderateMessage(Client net.Conn) {
	for {
		buffer := make([]byte, 4096) // big buffer
		count, err := Client.Read(buffer)
		print(string(buffer))
		if err != nil {
			log.Fatal("Write: ", err)
		}
		data := buffer[:count]
		message := []byte("Server:")
		message = append(message, data...)
		_, err = Client.Write(message)
		if err != nil {
			log.Fatal("Write: ", err)
		}
	}
}
