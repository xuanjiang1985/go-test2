package singleton

import "sync"

var (
	p    *Pet
	once sync.Once
)

func init() {
	once.Do(func() {
		p = &Pet{}
	})
}

func GetInstance() *Pet {
	return p
}

type Pet struct {
	name string
	age  int
	mux  sync.Mutex
}

func (p *Pet) SetName(n string) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.name = n
}

func (p *Pet) GetName() string {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.name
}
