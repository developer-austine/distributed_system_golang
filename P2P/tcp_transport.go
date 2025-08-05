package p2p

import "net"

type TCPTransport struct {
	listenAddress string
	listener 	  net.Listener
	mu 			  sync.RWMutex
	peers 		  map net.Addr.Peer
}

func NewTCPransport(ListenerAddr string) Transport {
	return &TCPTransport {
		listenAddress: listerAddr
	}
} 

func Test () {
	//The rest of the code goes here.
}