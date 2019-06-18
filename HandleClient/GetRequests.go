package HandleClient

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

const (
	NewUser    			int = 0
	SetConf    			int = 1
	ControlBoiler   	int = 2
	HotWaterInMinutes	int = 3
)

func getRequest(conn net.Conn) {
	fmt.Println("New client connected to server!")
	//bufferFileName := make([]byte, 64)
	UserName := make([]byte, 64)
	Password := make([]byte, 64)
	RequestTypeBuff := make([]byte, 10)
	RequestType, error1 := strconv.Atoi(string(RequestTypeBuff))
	if error1 != nil {
		CloseConnection(conn, "Error reading RequestType from user:", error1)
		return
	}
	// Get User name
	_, err := conn.Read(UserName)
	if err != nil {
		CloseConnection(conn, "Error reading user name from user:", err)
		return
	}

	userName := strings.Split(string(UserName), "\n")[0]

	// Get Password
	_, err = conn.Read(Password)
	if err != nil {
		CloseConnection(conn, "Error reading password from user:", err)
		return
	}

	password := strings.Split(string(Password), "\n")[0]

	switch RequestType {
	case NewUser:
		HandleNewUser(conn, userName, password)
	case SetConf:
		SetConfiguration(conn, userName, password)
	case ControlBoiler:
		ControlBoilerConditions(conn, userName, password)
	default:
		conn.Close()
	}

	defer conn.Close()
}

func RecieveRequest(conn net.Conn) {
	fmt.Println("New client connected to server!")
	//bufferFileName := make([]byte, 64)
	UserName := make([]byte, 64)
	Password := make([]byte, 64)
	RequestTypeBuff := make([]byte, 10)
	RequestType, error1 := strconv.Atoi(string(RequestTypeBuff))
	if error1 != nil {
		CloseConnection(conn, "Error reading RequestType from user:", error1)
		return
	}
	// Get User name
	_, err := conn.Read(UserName)
	if err != nil {
		CloseConnection(conn, "Error reading user name from user:", err)
		return
	}

	userName := strings.Split(string(UserName), "\n")[0]

	// Get Password
	_, err = conn.Read(Password)
	if err != nil {
		CloseConnection(conn, "Error reading password from user:", err)
		return
	}

	password := strings.Split(string(Password), "\n")[0]

	switch RequestType {
	case NewUser:
		HandleNewUser(conn, userName, password)
	case SetConf:
		SetConfiguration(conn, userName, password)
	case ControlBoiler:
		ControlBoilerConditions(conn, userName, password)
	case HotWaterInMinutes:
		GetHotWaterInMinutes(conn, userName, password)
	default:
		conn.Close()
	}

	defer conn.Close()

}

func CloseConnection(conn net.Conn, errMsg string, err error) {
	fmt.Println(errMsg, err.Error())
	conn.Close()
}