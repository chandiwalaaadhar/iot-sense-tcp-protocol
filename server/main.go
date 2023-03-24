package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
)

// Packet Types
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
	// Start server
	startServer()
}

func startServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept() // accept a connection
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		packet := &IoTSensePacket{}
		packet.PacketType, _ = reader.ReadByte()                  // reading one byte since PacketType is a byte
		binary.Read(reader, binary.BigEndian, &packet.SensorID)   // reading two bytes since SensorID is a uint16
		binary.Read(reader, binary.BigEndian, &packet.DataLength) // reading two bytes since DataLength is a uint16

		packet.Data = make([]byte, packet.DataLength) // make a slice of bytes with length DataLength
		_, err := reader.Read(packet.Data)            // read DataLength bytes into the Data slice
		if err != nil {
			fmt.Println("Error reading data:", err)
			break
		}

		if packet.PacketType == DATA {
			fmt.Println("Data:", string(packet.Data)) // print the data
		}

		if packet.PacketType == TERMINATE {
			break
		}
	}
}
