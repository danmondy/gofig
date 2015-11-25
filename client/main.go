package main

import(
	"net"
	"fmt"
	"io/ioutil"
)

const PORT string = ":24601"

func main(){

	tcpAddr, err := net.ResolveTCPAddr("tcp4", PORT)
	if err != nil{ panic(err) }

	fmt.Println(tcpAddr)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil{ fmt.Println(err)}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Write([]byte("fail"))
			continue
		}
		handleClient(conn)
	}
}

func handleClient(conn net.Conn){
	defer conn.Close()

	fmt.Println("Connection established")
	data, err := ioutil.ReadAll(conn);
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
	err = ioutil.WriteFile("serverconfig.json", data, 0666)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Writing file")
	conn.Write([]byte("ok"))
}
