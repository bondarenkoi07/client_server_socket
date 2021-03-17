package Client

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

const SockAddr = "/home/ilya/OS/echo.sock"

type Client struct {
	c net.Conn
}

func NewClient() *Client {
	c, err := net.Dial("unix", SockAddr)
	if err != nil {
		panic(err)
	}
	return &Client{c: c}
}

func reader(r io.Reader) {
	buf := make([]byte, 4096)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			println(err)
			return
		}
		println("Client got:", string(buf[0:n]))
	}
}

func (s *Client) ListenAndServe() {
	defer (*s).c.Close()

	go reader((*s).c)
	for {
		fmt.Println("Enter your message:")
		var text string
		_, err := fmt.Scanf("%s", &text)
		if err != nil {
			continue
		}
		message := string(rune(len(text))) + ": " + text
		_, err = (*s).c.Write([]byte(message))
		if err != nil {
			log.Fatal("write error:", err)
		}
		time.Sleep(1e9)
	}
}
