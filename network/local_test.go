package network

import (
	"log"
	"testing"
)

func TestGetOutboundIP(t *testing.T) {
	ip, err := GetOutboundIP()
	if err != nil {
		t.Fatalf("GetOutboundIP returned an error: %v", err)
	}

	if ip == nil {
		t.Fatalf("GetOutboundIP returned a nil IP")
	}

	if ip.To4() == nil {
		t.Fatalf("GetOutboundIP returned a non-IPv4 address")
	}

	if ip.IsLoopback() {
		t.Fatalf("GetOutboundIP returned a loopback address")
	}

	if ip.IsUnspecified() {
		t.Fatalf("GetOutboundIP returned an unspecified address")
	}

	if ip.IsMulticast() {
		t.Fatalf("GetOutboundIP returned a multicast address")
	}
	log.Println(ip.String())
}
