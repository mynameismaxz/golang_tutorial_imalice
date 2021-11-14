package entity

import (
	"fmt"
	"time"
)

type Minebase struct {
	Name       string
	Coin       int
	miner      map[string]int
	updateCart chan string
}

func NewBase(name string, coin int) Minebase {
	m := Minebase{
		Name:       name,
		Coin:       coin,
		miner:      make(map[string]int),
		updateCart: make(chan string),
	}
	return m
}

func (m *Minebase) GetName() string {
	return m.Name
}

func (m *Minebase) GetCoin() int {
	return m.Coin
}

func (m *Minebase) SetName(name string) {
	m.Name = name
}

func (m *Minebase) Spawn(minerName string) {
	m.miner[minerName] = 0
	// spawn and send to mining
	go m.SendToMining(minerName)
}

func (m *Minebase) GetMiner() map[string]int {
	return m.miner
}

func (m *Minebase) ListenMiner() {
	for {
		select {
		case minerName := <-m.updateCart:
			m.miner[minerName] += 100
			fmt.Printf("Havest from %s: received %d coin\n", minerName, m.miner[minerName])
		}
	}
}

func (m *Minebase) SendToMining(carName string) {
	for {
		m.updateCart <- carName
		time.Sleep(1 * time.Second)
	}
}
