package ptp

import (
	"fmt"
	"time"
)

// KeyManager holds security keys and provides
// fast access to the active one
type KeyManager struct {
	keys        []*Key
	activeKey   *Key
	running     bool
	initialized bool
}

func (m *KeyManager) init() error {
	m.initialized = true
	return nil
}

// Run checks keys every 100 milliseconds
func (m *KeyManager) run() {
	m.running = true
	for m.running {
		time.Sleep(time.Millisecond * 100)
		// Check if active key expired
		if m.activeKey != nil {

		}
		if len(m.keys) == 0 {
			continue
		}
		// Active key is not set
		m.SetActiveKey()
	}
}

func (m *KeyManager) close() error {
	if !m.running {
		return fmt.Errorf("Key Manager was already closed")
	}
	if !m.initialized {
		return fmt.Errorf("Key Manager wasn't initialized")
	}
	m.running = false
	return nil
}

// SetActiveKey will go through the list of available keys
// and pick the one with lowest expiration date and it's already
// started
func (m *KeyManager) SetActiveKey() error {
	if len(m.keys) == 0 {
		return fmt.Errorf("No keys available")
	}
	for _, k := range m.keys {
		if !k.IsStarted() {
			continue
		}

	}
	return nil
}
