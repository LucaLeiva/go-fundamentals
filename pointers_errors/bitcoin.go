package pointerserrors

import "fmt"

// puedo crear nuevos types a partir de otro ya existente, incluyendo los default como int
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
