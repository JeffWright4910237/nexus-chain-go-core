package main

import (
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"
)

type NexusP2PNode struct {
	NodeUID  string
	Host     string
	Port     int
	PeerList map[string]*NexusP2PNode
	Mutex    sync.RWMutex
	Online   bool
}

func NewNexusNode(host string, port int) *NexusP2PNode {
	return &NexusP2PNode{
		NodeUID:  fmt.Sprintf("NEXUS_NODE_%d", rand.Intn(999999)),
		Host:     host,
		Port:     port,
		PeerList: make(map[string]*NexusP2PNode),
		Online:   true,
	}
}

func (n *NexusP2PNode) RegisterPeer(peer *NexusP2PNode) {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()
	n.PeerList[peer.NodeUID] = peer
	fmt.Printf("[%s] Peer registered: %s\n", n.NodeUID, peer.NodeUID)
}

func (n *NexusP2PNode) SendGlobalMessage(msg string) {
	n.Mutex.RLock()
	defer n.Mutex.RUnlock()
	fmt.Printf("\n[%s] Broadcast: %s\n", n.NodeUID, msg)
	for uid := range n.PeerList {
		fmt.Printf("Sent to %s ✅\n", uid)
	}
}

func (n *NexusP2PNode) StartNode() {
	listen, _ := net.Listen("tcp", fmt.Sprintf("%s:%d", n.Host, n.Port))
	defer listen.Close()
	fmt.Printf("Nexus P2P Node Running: %s:%d\n", n.Host, n.Port)
	for n.Online {
		listen.Accept()
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	nodeA := NewNexusNode("127.0.0.1", 8765)
	nodeB := NewNexusNode("127.0.0.1", 8766)
	nodeC := NewNexusNode("127.0.0.1", 8767)

	nodeA.RegisterPeer(nodeB)
	nodeA.RegisterPeer(nodeC)
	go nodeA.StartNode()
	time.Sleep(1 * time.Second)
	nodeA.SendGlobalMessage("NEXUS_HEIGHT: 1999")
}
