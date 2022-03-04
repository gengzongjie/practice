package tcp

import (
	"fmt"
	"net"
	"testing"
)

func Test_Server(t *testing.T)  {
	ip := "127.0.0.1"
	port := 8080

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		t.Error(err)
		return
	}
	defer listener.Close()

	for {
		t.Log("server listening ...")
		conn, err := listener.Accept()
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("server listened accept.")

		for {
			var msg = make([]byte, 1024)
			n, err := conn.Read(msg)
			if err != nil {
				t.Error(err)
				return
			}
			if n > 0 {
				t.Logf("%d:%s", n, msg)
			}

			_, err = conn.Write([]byte("server recved"))
			if err != nil {
				t.Error(err)
			}
		}
	}
}
