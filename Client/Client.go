package Client

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
)

var mute = make(chan bool)

const SockAddr = "/home/ilya/OS/echo.sock"

type Client struct {
	c  net.Conn
	mu sync.Mutex
}

func NewClient() *Client {
	c, err := net.Dial("unix", SockAddr)
	if err != nil {
		panic(err)
	}
	return &Client{c: c}
}

func (c *Client) reader(r io.Reader) {
	buf := make([]byte, 4096)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			println(err)
			return
		}
		println("Client got:\"", string(buf[0:n]), "\"")
		mute <- true
	}
}

func (c *Client) ListenAndServe() {
	defer (*c).c.Close()

	go (*c).reader((*c).c)
	for {
		fmt.Println("Enter your message:")
		var text string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			text = scanner.Text()

		}
		_, err := (*c).c.Write([]byte(text))
		if err != nil {
			log.Fatal("write error:", err)
		}
		<-mute
	}
}
