package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type User struct {
	ip_addr   string
	conn      net.Conn
	namespace string
	// connection_id string
	user_id string
}

type Header struct {
	Connection_type  string
	Namespace        string
	Date_established string
	User_id          string
}

const PORT = ":8080"

var user_storage = make(map[string]User)

func main() {
	var app = set_server_socket()
	start_server(app)

}

func set_server_socket() net.Listener {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("Unable to create a listener ")
	}
	return listener
}

func start_server(server net.Listener) {

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Unable to create a tcp connection")
		}
		handle_tcp_connection(conn)
	}

}

func handle_tcp_connection(conn net.Conn) {
	buffer := make([]byte, 1024)
	bytes_read, err := conn.Read(buffer)

	if err != nil {
		fmt.Println("Cant Reaaaad")
	}

	header_bytes := []byte(buffer[:bytes_read])
	client_request_header := Header{}
	decode_err := json.Unmarshal(header_bytes, &client_request_header)
	if decode_err != nil {
		fmt.Println("Unable to decode client request header")
	}
	fmt.Println(client_request_header)
	new_user := User{ip_addr: conn.RemoteAddr().String(), user_id: client_request_header.User_id, conn: conn, namespace: client_request_header.Namespace}
	create_user(new_user)
}

func create_user(new_user User) {
	user_storage[new_user.ip_addr] = new_user
	fmt.Println(new_user)
	fmt.Println(user_storage)
}
