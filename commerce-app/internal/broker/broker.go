package broker

import (
	"sync"
	"log"
)


type Broker struct {
    clients map[chan string]bool
    mu      sync.Mutex
}

func NewBroker() *Broker {
    return &Broker{
        clients: make(map[chan string]bool),
    }
}

func (b *Broker) AddClient(ch chan string) {
    b.mu.Lock()
    b.clients[ch] = true
    b.mu.Unlock()
}

func (b *Broker) RemoveClient(ch chan string) {
    b.mu.Lock()
    delete(b.clients, ch)
    close(ch)
    b.mu.Unlock()
}

func (b *Broker) Broadcast(msg string) {
	log.Println("Broadcasting:", msg)
    b.mu.Lock()
    for ch := range b.clients {
        ch <- msg
    }
    b.mu.Unlock()
}
var Brokerk = NewBroker()
