package tcp

import (
	"net"
	"testing"
)

func Test_client(t *testing.T)  {
	serverAddr := "127.0.0.1:8080"

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	t.Log("client dialed")

	n, err := conn.Write([]byte("nihao"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("send %d", n)

	var recvData = make([]byte, 1024)
	l, err := conn.Read(recvData)
	if err != nil{
		t.Errorf("Read failed , err : %v\n", err)
		return
	}
	if l > 0 {
		t.Logf("%s", recvData)
	}

	return
}
