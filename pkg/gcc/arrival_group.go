// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package gcc

import (
	"fmt"
	"time"

	"github.com/pion/interceptor/internal/cc"
)

type arrivalGroup struct {
	packets        []cc.Acknowledgment
	departure      time.Time
	arrival        time.Time
	firstDeparture time.Time
}

func (g *arrivalGroup) add(a cc.Acknowledgment) {
	if len(g.packets) == 0 {
		g.firstDeparture = a.Departure
	}
	g.packets = append(g.packets, a)
	g.arrival = a.Arrival
	if a.Departure.After(g.departure) {
		g.departure = a.Departure
	}
	if a.Departure.Before(g.firstDeparture) {
		g.firstDeparture = a.Departure
	}

}

func (g arrivalGroup) String() string {
	s := "ARRIVALGROUP:\n"
	s += fmt.Sprintf("\tARRIVAL:\t%v\n", int64(float64(g.arrival.UnixNano())/1e+6))
	s += fmt.Sprintf("\tDEPARTURE:\t%v\n", int64(float64(g.departure.UnixNano())/1e+6))
	s += fmt.Sprintf("\tPACKETS:\n%v\n", g.packets)
	return s
}
