package link

import "sync"

type (
	Memory struct {
		mu    sync.RWMutex
		items map[Link][]byte
	}
)

func NewMemoryRepository() *Memory {
	return &Memory{
		mu:    sync.RWMutex{},
		items: make(map[Link][]byte),
	}
}

func (m *Memory) Find(l Link) (Mapped, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var item Mapped

	buf, ok := m.items[l]
	if !ok {
		return item, errLinkNotFound
	}

	return BytesToMapped(buf)
}

func (m *Memory) Create(item Mapped) error {
	if _, err := m.Find(item.Original); err == nil {
		// It already exists
		// Move on
		return nil
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.items[item.Original] = item.Bytes()
	return nil
}

func (m *Memory) len() int64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return int64(len(m.items))
}
