package protocolfactory

import (
	"time"

	"github.com/kelvinkuo/crud/internal/consts"
	"github.com/kelvinkuo/crud/internal/protocol"
	"github.com/kelvinkuo/crud/internal/protocol/pb"
	"github.com/kelvinkuo/crud/internal/protocol/zero"
)

func NewProtocol(protocolType string, packageName string, goPackage string) protocol.Protocol {
	switch protocolType {
	case consts.ProtoBuf:
		return pb.NewProtocol("proto3", packageName, goPackage)
	case consts.ZeroApi:
		return zero.NewProtocol("v1", "crud", time.Now())
	}
	return nil
}

func NewMessage(protocolType string, name string) protocol.Message {
	switch protocolType {
	case consts.ProtoBuf:
		return pb.NewMessage(name)
	case consts.ZeroApi:
		return zero.NewMessage(name)
	}
	return nil
}
