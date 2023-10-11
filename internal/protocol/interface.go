package protocol

// Protocol is interface of communication protocol
// proto3 or thrift or zero-api all of them are kinds of protocol
type Protocol interface {
	AddItem(item Item) error
	AddMessage(message Message) error
	String() string
}

// Item is an item of protocol just like a http api defined in zero api file or a rpc in pb file.
type Item interface {
	Name() string
	Comment() string
	Service() string
	Request() Message
	Response() Message
	String(indent int) string
}

type Message interface {
	Name() string
	AddField(field Field) error
	String(indent int) string
}

type Field interface {
	String(indent int) string
}
