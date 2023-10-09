package zero

import (
    "github.com/kelvinkuo/crud/db"
    "github.com/kelvinkuo/crud/protocol"
)

type ZeroApiConverter struct {
}

func (z *ZeroApiConverter) CreateMetaMessage(table db.Table) (protocol.Message, error) {
    // TODO implement me
    panic("implement me")
}
