package gcc

import (
	"fmt"
	"time"
)

// Acknowledgment holds information about a packet and if/when it has been
// sent/received.
type Acknowledgment struct {
	TLCC      uint16
	Size      int
	Departure time.Time
	Arrival   time.Time
	RTT       time.Duration
}

func (a Acknowledgment) String() string {
	s := "ACK:\n"
	s += fmt.Sprintf("\tTLCC:\t%v\n", a.TLCC)
	s += fmt.Sprintf("\tSIZE:\t%v\n", a.Size)
	s += fmt.Sprintf("\tDEPARTURE:\t%v\n", a.Departure)
	s += fmt.Sprintf("\tARRIVAL:\t%v\n", a.Arrival)
	s += fmt.Sprintf("\tRTT:\t%v\n", a.RTT)
	return s
}
