package idgen

import (
	"hash/crc32"
	"os"
	"testing"
	"time"

	"github.com/chenyan/wheels/network"
)

func TestGenWorkerID(t *testing.T) {
	// Test with a valid IP address and PID
	ip, err := network.GetOutboundIP()
	if err != nil {
		return
	}
	expectedWorkerID := int64((crc32.ChecksumIEEE([]byte(ip)) + uint32(os.Getpid())) & uint32(workerIDMask))
	workerID, err := GenWorkerID()
	if err != nil {
		t.Errorf("GenWorkerID returned an error: %v", err)
	}
	if workerID != expectedWorkerID {
		t.Errorf("GenWorkerID returned %d, expected %d", workerID, expectedWorkerID)
	}
}
func TestSnowflake_GetID(t *testing.T) {
	sf := &Snowflake{
		dataCenterID: 1,
		workerID:     1,
		timestamp:    time.Now().UnixNano() / 1e6,
		sequence:     0,
	}

	id1 := sf.GetID()
	id2 := sf.GetID()

	if id1 == id2 {
		t.Errorf("IDs should be unique, but got %d and %d", id1, id2)
	}

	// Test that the sequence resets after maxSequence is reached
	for i := int64(0); i < maxSequence; i++ {
		sf.GetID()
	}
	id3 := sf.GetID()

	println(id3)

	if id3 == id2 {
		t.Errorf("Sequence should reset after reaching maxSequence, but got %d and %d", id2, id3)
	}
}
