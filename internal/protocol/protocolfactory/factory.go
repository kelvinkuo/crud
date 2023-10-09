package protocolfactory

import (
    "fmt"
    
    "github.com/kelvinkuo/crud/consts"
    "github.com/kelvinkuo/crud/protocol"
    "github.com/kelvinkuo/crud/protocol/pb"
)

func NewProtocol(protocolType string, packageName string) protocol.Protocol {
    switch protocolType {
    case consts.PROTOBUF:
        return pb.NewProtocol("proto3", packageName, fmt.Sprintf("./%s", packageName))
    }
    return nil
}

func NewMessage(protocolType string, name string) protocol.Message {
    switch protocolType {
    case consts.PROTOBUF:
        return pb.NewMessage(name)
    }
    return nil
}
