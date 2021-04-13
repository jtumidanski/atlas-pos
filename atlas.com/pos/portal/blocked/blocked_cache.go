package blocked

import "sync"

type cache struct {
	blocks map[uint32][]string
	mutex  sync.RWMutex
}

var c *cache
var once sync.Once

func GetCache() *cache {
	once.Do(func() {
		c = &cache{
			blocks: make(map[uint32][]string, 0),
			mutex:  sync.RWMutex{},
		}
	})
	return c
}

func (b *cache) Blocked(characterId uint32, scriptName string) bool {
	var blocked = false
	b.mutex.RLock()
	if scripts, ok := b.blocks[characterId]; ok {
		for _, script := range scripts {
			if script == scriptName {
				blocked = true
			}
		}
	}
	b.mutex.RUnlock()
	return blocked
}

func (b *cache) AddBlockedPortal(characterId uint32, scriptName string) {
	b.mutex.Lock()
	var scripts []string
	var ok bool
	if scripts, ok = b.blocks[characterId]; !ok {
		scripts = make([]string, 0)
	}

	var contains = false
	for _, script := range scripts {
		if script == scriptName {
			contains = true
		}
	}

	if !contains {
		scripts = append(scripts, scriptName)
	}
	b.blocks[characterId] = scripts
	b.mutex.Unlock()
}
