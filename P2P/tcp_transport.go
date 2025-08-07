package p2p

import (
	"net"
	"sync"
)

//TCPPeer represents the remote node over a TCP established connection
type TCPPeer struct {
	//conn is the underlying connection of the peer
	conn net.Conn
	//if we dial and retrieve a conn => outbound == true
	//if we acept and retrieve a conn => outbound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer[
		conn: conn,
		outboutbound: outbound,
	]
}

type TCPTransport struct {
	listenAddress string
	listener 	  net.Listener
	mu 			  sync.RWMutex
	peers 		  map [net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport {
		listenAddress: listenAddr,
	}
}

func (t*TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err := net.Listen("tcp", t.listenAddress)
	if err != nil (
		return err
	)
	go t.startAcceptLoop()

	return nil
}

func (t*TCPTransport) startAcceptLoop() {
	for (
		conn, error := t.listener.Accept()
		if err != nil (
			fmt.printf("TCP accept error: %s\n", err)
		)
		
		go t.handleConn(conn)
	)
}

func (t*TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	fmt.printf("new incoming connection %+v\n", peer)
}