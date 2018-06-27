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
			m.ValidateKey()
			continue
		}
		if len(m.keys) == 0 {
			continue
		}
		// Active key is not set
		m.SetActiveKey()
	}
}

// Close will terminate KeyManager main loop
func (m *KeyManager) Close() error {
	if !m.running {
		return fmt.Errorf("Key Manager was already closed")
	}
	if !m.initialized {
		return fmt.Errorf("Key Manager wasn't initialized")
	}
	m.running = false
	return nil
}

// ValidateKey will check current active key
func (m *KeyManager) ValidateKey() {
	if m.activeKey.IsExpired() {
		if !m.activeKey.prolonged {
			Log(Info, "Key %s expired", m.activeKey.id)
		}
		if m.SwitchKey() != nil {
			err := m.ProlongKey()
			if err != nil {
				Log(Error, "Key switch failed: %s", err.Error())
			}
		}
	}
}

// SetActiveKey will go through the list of available keys
// and pick the one with lowest expiration date and it's already
// started
func (m *KeyManager) SetActiveKey() error {
	if len(m.keys) == 0 {
		return fmt.Errorf("No keys available")
	}
	k, err := m.FindFreshKey()
	if err != nil {
		return fmt.Errorf("Failed to set active key: %s", err.Error())
	}
	m.activeKey = k
	return nil
}

// SwitchKey will switch active key to another one if possible
// Another key can be accepted and became active if current date
// is equal or after it's start date, and it's end time isn't up
// If there is no such key available error will be returned
func (m *KeyManager) SwitchKey() error {
	if m.activeKey == nil {
		return fmt.Errorf("Active key is not set")
	}
	newKey, err := m.FindFreshKey()
	if err != nil {
		return fmt.Errorf("Failed to switch key: %s", err.Error())
	}
	Log(Info, "Switching key from %s to %s", m.activeKey.id, newKey.id)
	m.activeKey = newKey
	return nil
}

// ProlongKey will set `prolonged` flag of current key to true
// which means there was no acceptable key available to replace
// current active key
func (m *KeyManager) ProlongKey() error {
	if m.activeKey == nil {
		return fmt.Errorf("Failed to prolong key: no active key")
	}
	m.activeKey.prolonged = true
	return nil
}

// FindFreshKey will find a key with the lowest start date
// from keys that is started and not ended
func (m *KeyManager) FindFreshKey() (*Key, error) {
	var candidate *Key
	for _, k := range m.keys {
		if k == nil {
			continue
		}
		if !k.IsStarted() || k.IsExpired() {
			continue
		}
		if candidate == nil {
			candidate = k
			continue
		}
		if candidate.starts.Sub(k.starts) > 0 {
			candidate = k
		}
	}
	if candidate == nil {
		return nil, fmt.Errorf("No suitable key was found")
	}
	return candidate, nil
}
