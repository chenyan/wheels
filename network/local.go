package network

import (
	"errors"
	"net"
)

var (
	ErrorNoNetwork = errors.New("no network")
)

// GetOutboundIP returns the local IP address of the machine that initiates the connection to the internet.
// It does so by dialing a UDP connection to "8.8.8.8" address and port, and then returns the local IP address of the connection.
func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, ErrorNoNetwork
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}
