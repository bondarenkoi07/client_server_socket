package Server

import (
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
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
		prefix := make([]byte, 256) // big buffer
		_, err := Client.Read(prefix)
		if err != nil {
			log.Fatal("Write: ", err)
		}
		log.Print("string", string(prefix))
		matched, err := regexp.Match(`^(\d):$`, prefix)
		if err != nil {
			log.Println("regexp: ", err)
			break
		}
		if matched {
			count, err := strconv.Atoi(string(prefix[:len(prefix)-1]))
			if err != nil {
				log.Println("strconv : ", err)
				break
			}
			buffer := make([]byte, 4096)
			n := 0

			for readCount := 0; readCount < count; n++ {
				tmp := make([]byte, 256)

				n, err = Client.Read(tmp)
				if err != nil {
					log.Fatal("Write: ", err)
				}
				n += 1
				buffer = append(buffer, tmp...)
			}
			data := buffer[:count]
			message := []byte("Server:")
			message = append(message, data...)
			_, err = Client.Write(data)
			if err != nil {
				log.Fatal("Write: ", err)
			}
		}
	}
}
