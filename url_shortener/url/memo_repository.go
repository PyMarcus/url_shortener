package url

type memoryRepository struct {
	urls map[string]*Url
}

func NewMemoryRepository() *memoryRepository {
	return &memoryRepository{urls: make(map[string]*Url)}
}

func (m *memoryRepository) ExistId(id string) bool {
	_, exist := m.urls[id]
	return exist
}

func (m *memoryRepository) SearchById(id string) *Url {
	return m.urls[id]
}

func (m *memoryRepository) SearchByUrl(url string) *Url {
	for _, u := range m.urls {
		if u.Dest == url {
			return u
		}
	}
	return nil
}

func (m *memoryRepository) Save(url Url) error {
	m.urls[url.Id] = &url
	return nil
}
