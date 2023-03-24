package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

const (
	DATA = iota
	TERMINATE
)

// IoTSensePacket represents a packet in the IoTSense protocol
type IoTSensePacket struct {
	PacketType byte
	SensorID   uint16
	DataLength uint16
	Data       []byte
}

func main() {
	startClient()
}

func startClient() {
	for { // loop forever to allow the user to send multiple messages
		fmt.Println("Enter data to send to IOT Server: ")
		reader := bufio.NewReader(os.Stdin)     // create a reader to read from the console
		message, err := reader.ReadString('\n') // read until newline character is encountered
		if err != nil {
			fmt.Println("Error reading message:", err)
			continue // continue to the next iteration of the loop
		}
		message = message[:len(message)-1] // remove newline character

		conn, err := net.Dial("tcp", "localhost:8080") // connect to the server
		if err != nil {
			fmt.Println("Error connecting to server:", err)
			return
		}
		defer conn.Close() // close the connection when the function returns

		// Example data packet
		dataPacket := &IoTSensePacket{
			PacketType: DATA,
			SensorID:   1,
			DataLength: uint16(len(message)),
			Data:       []byte(message),
		}

		// Termination packet
		terminationPacket := &IoTSensePacket{
			PacketType: TERMINATE,
			SensorID:   1,
			DataLength: 0,
			Data:       []byte{},
		}

		sendPacket(conn, dataPacket)
		sendPacket(conn, terminationPacket)
	}
}

func sendPacket(conn net.Conn, packet *IoTSensePacket) {
	conn.Write([]byte{packet.PacketType})                   // write one byte since PacketType is a byte
	binary.Write(conn, binary.BigEndian, packet.SensorID)   // write two bytes since SensorID is a uint16
	binary.Write(conn, binary.BigEndian, packet.DataLength) // write two bytes since DataLength is a uint16
	conn.Write(packet.Data)                                 // write DataLength bytes since Data is a slice of bytes
}
