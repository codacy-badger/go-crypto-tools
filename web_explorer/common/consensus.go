package common

const (
	ConsensusPoS    = "PoS"
	ConsensusTPoS   = "TPoS"
	ConsensusPoW    = "PoW"
	ConsensusHybrid = "Hybrid"

	ConsensusLongPoS    = "Proof of Stake"
	ConsensusLongTPoS   = "Trustless Proof of Stake"
	ConsensusLongPoW    = "Proof of Work"
	ConsensusLongHybrid = "Hybrid"
)

type ConsensusType int

const (
	Undefined ConsensusType = 0
	PoW       ConsensusType = 1
	PoS       ConsensusType = 2
	TPoS      ConsensusType = 3
	Hybrid    ConsensusType = 4
)

func (c ConsensusType) String() string {
	list := [...]string{
		"undefined",
		ConsensusPoW,
		ConsensusPoS,
		ConsensusHybrid,
		ConsensusTPoS,
	}
	return list[c]
}

func (c ConsensusType) StringLong() string {
	list := [...]string{
		"undefined",
		ConsensusLongPoS,
		ConsensusLongTPoS,
		ConsensusLongPoW,
		ConsensusLongHybrid,
	}
	return list[c]
}

func GetConsensusType(consensusType string) ConsensusType {
	switch consensusType {
	case ConsensusPoW:
		return PoW
	case ConsensusPoS:
		return PoS
	case ConsensusTPoS:
		return TPoS
	case ConsensusHybrid:
		return Hybrid
	}
	return Undefined
}
