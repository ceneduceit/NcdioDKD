package constant

import (
	"net"
)

type PlainContext interface {
	ID() uuid.UUID
}

type ConnContext interface {
	PlainContext
}

type PacketConnContext interface {
	PlainContext
	Metadata() *Metadata
	PacketConn() net.PacketConn
}
