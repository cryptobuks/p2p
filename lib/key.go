package ptp

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"
)

// NewKey function will return a new Key object with calculated id
// id must be unique inside a p2p environment, so function accepts list
// of existing ids to check if it's new or not
func NewKey(key string, start, expiration time.Time, idList []string) (*Key, error) {
	k := new(Key)
	k.key = key
	k.added = time.Now()
	k.starts = start
	k.expires = expiration

	// Checking expiration
	diff := expiration.Sub(k.added)
	if diff < 0 {
		return nil, fmt.Errorf("Can't create key: already expired")
	}

	if err := k.generateID(idList); err != nil {
		return nil, err
	}

	return k, nil
}

// Key represetns a crypto-key used by AES crypto subsystem
// id is a portion of MD5 hash of a key which must be unique
// inside a particular p2p environment
type Key struct {
	id        string    // Unique ID of the key
	key       string    // Crypto key
	added     time.Time // When key was added
	starts    time.Time // When key starts
	expires   time.Time // When key ends
	prolonged bool      // Whether key was prolonged due to missing other acceptable key
}

// Method will do 256 attempts to generate unique ID
// and return error if p2p fail to do this, which is unlikely
func (k *Key) generateID(idList []string) error {
	keyBase := k.key
	newID := ""

	for i := 0; i < 256; i++ {
		h := md5.New()
		io.WriteString(h, keyBase)
		newID = fmt.Sprintf("%x", h.Sum(nil))
		newID = newID[0:7]

		for _, id := range idList {
			if id == newID {
				keyBase = fmt.Sprintf("%s%d", keyBase, i)
				newID = ""
				break
			}
		}
		if newID != "" {
			k.id = newID
			return nil
		}
	}
	return fmt.Errorf("Failed to generate unique ID for key")
}

// IsExpired will return true if key is expired
func (k *Key) IsExpired() bool {
	diff := k.expires.Sub(time.Now())
	if diff <= 0 {
		return true
	}
	return false
}

// IsStarted will check if key timeframe has already veen started
func (k *Key) IsStarted() bool {
	diff := k.starts.Sub(time.Now())
	if diff <= 0 {
		return true
	}
	return false
}
