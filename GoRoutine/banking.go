package GoRoutine

import (
	"sync"
)

type Banking struct {
	mutex   *sync.RWMutex
	balance int
}

func (m *Banking) addBalance(amount int) {

	m.mutex.Lock() // lock go-routine untuk menambahkan data ( menghindari race condition saat insert )
	m.balance = m.balance + amount
	m.mutex.Unlock() // un-lock kembali
}

func (m *Banking) getBalance() int {
	m.mutex.Lock() // lock go-routine untuk mengakses value ( menghindari race condition saat read )
	balance := m.balance
	m.mutex.Unlock() // un-lock kembali

	return balance
}

func NewBanking(mutex *sync.RWMutex) *Banking {
	return &Banking{
		mutex: mutex,
	}
}
