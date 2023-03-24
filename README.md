# IoTSense: A Lightweight TCP Protocol for IoT Devices

IoTSense is a lightweight and efficient TCP-based protocol designed for transmitting sensor data between IoT devices. This repository contains a Go implementation of the IoTSense protocol with a client and server application.

## Features

- Simple and efficient packet structure
- Customizable packet types for data transmission and termination
- Easy-to-understand Go implementation
- Command-line input for sending sensor data

## Getting Started

### Prerequisites

- Go 1.14+ installed

### Installation

1. Clone the repository:

```bash
git clone https://github.com/chandiwalaaadhar/iot-sense-tcp-protocol.git
```

2. Change to the repository directory:

```bash
cd iot-sense-tcp-protocol
```

3. Compile the client and server applications:
```bash
go build -o client client/main.go
go build -o server server/main.go
```

### Usage

1. Run the IoTSense server:

```bash
/server/main 
```

2. In a separate terminal, run the IoTSense client:

```bash
/client/main 
```

3. Follow the prompts to enter the sensor data and sensor ID to send to the server. The server will display the received data and sensor ID in the Client Terminal.

```bash
this is a test message>1
```

##Understanding the Code

###Client
The client application (client/main.go) allows users to input sensor data and sensor ID from the command line. It connects to the server, sends the data using the IoTSense protocol, and sends a termination packet to signal the end of the transmission.

###Server
The server application (server/main.go) listens for incoming connections on a specified port and processes incoming IoTSense packets. When a data packet is received, the server prints the sensor data and sensor ID.

###Packet Structure
The `IoTSensePacket` structure represents packets in the IoTSense protocol:

```bash
type IoTSensePacket struct {
	PacketType byte
	SensorID   uint16
	DataLength uint16
	Data       []byte
}
```

##Contributing
1. Fork the repository
2. Create a new feature branch (`git checkout -b feature/new-feature`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature/new-feature`)
5. Create a new Pull Request

