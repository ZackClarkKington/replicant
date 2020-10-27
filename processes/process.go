package processes

import (
	"github.com/golang/protobuf/proto"
	"log"
)

type Process struct {
	Ballot int64
	Leader bool
	ID     string
}

func (p Process) CheckBallot(ballot int64, senderId string) {
	if ballot > p.Ballot {
		p.Ballot = ballot
		p.ID = senderId
	}
}

func (p Process) CallElection() {
	e := Election{
		ID:     p.ID,
		Ballot: p.Ballot,
	}
	_, err := proto.Marshal(&e)
	if err != nil {
		log.Fatalln("Failed to encode election message")
	}

}
