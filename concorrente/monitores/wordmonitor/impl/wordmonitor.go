package impl

import (
	"sync"
)

type WordMonitor interface {
	Wait()
	Signal()
	GetData() []string
	SetData(string)
}

type Words struct { // all variable are private, i.e., non-capital letter
	mutex         *sync.Mutex
	wordsArray    []string
	isInitialized bool
}

func (m *Words) Init() {
	m.mutex = &sync.Mutex{}
	m.wordsArray = []string{}
	m.isInitialized = true
}

func (m *Words) Wait() {
	if m.isInitialized {
		m.mutex.Lock()
	}
}
func (m *Words) Signal() {
	if m.isInitialized {
		m.mutex.Unlock()
	}
}
func (m *Words) GetData() []string {
	return m.wordsArray
}

func (m *Words) SetData(word string) {
	m.Wait()
	// critical section
	m.wordsArray = append(m.wordsArray, word)
	// critical section done
	m.Signal()
}
