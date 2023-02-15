package types

type Dimension int8

func (d Dimension) String() string {
	switch d {
	case Nether:
		return "Nether"
	case Overworld:
		return "Overworld"
	case End:
		return "End"
	default:
		return "Unknown"
	}
}

func (d Dimension) LevelType() string {
	switch d {
	case Nether:
		return "flat"
	case Overworld:
		return "default"
	case End:
		return "flat"
	default:
		return "unknown"
	}
}

const (
	Nether    Dimension = -1
	Overworld Dimension = 0
	End       Dimension = 1
)
