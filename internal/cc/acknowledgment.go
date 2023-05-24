// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package cc

import (
	"fmt"
	"time"

	"github.com/pion/rtcp"
)

// Acknowledgment holds information about a packet and if/when it has been
// sent/received.
type Acknowledgment struct {
	SequenceNumber uint16 // Either RTP SequenceNumber or TWCC
	SSRC           uint32
	Size           int
	Departure      time.Time
	Arrival        time.Time
	ECN            rtcp.ECN
}

func (a Acknowledgment) String() string {
	return fmt.Sprintf("ssrc:%d sn:%d size:%d departure:%s arrival:%s rtt:%s",
		a.SSRC,
		a.SequenceNumber,
		a.Size,
		a.Departure.String(), a.Arrival.String(),
		a.Arrival.Sub(a.Departure).String(),
	)
}
