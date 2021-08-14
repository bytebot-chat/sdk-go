package bytebot

import (
	"github.com/satori/go.uuid"
)

type Metadata struct {
	Source string
	Dest   string
	ID     uuid.UUID
}
