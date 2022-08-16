package cell

type Type string

const (
	EMPTY   Type = "0"
	BLOCKED Type = "X"
	START   Type = "S"
	GOAL    Type = "G"
	PATH    Type = "1"
)

func (c Type) String() string {
	return string(c)
}
