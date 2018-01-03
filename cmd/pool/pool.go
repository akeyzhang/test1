package pool

import (
	"sync"
	"io"
	"errors"
	"log"
)

var ErrPoolClosed = errors.New("資源池已關畢!")


type Pool struct {
	m       sync.Mutex
	res     chan io.Closer
	factory func() (io.Closer, error)
	closed  bool
}

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size太小!")
	}
	return &Pool{
		factory: fn,
		res:     make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.res:
		log.Println("Acquire:共響資源")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:生成新資源")
		return p.factory()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}
	p.closed = true
	close(p.res)

	for r := range p.res {
		r.Close()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}
	select {
	case p.res <- r:
		log.Println("資源釋放進池子了")
	default:
		log.Println("池子滿了,釋放該資源")
		r.Close()
	}
}
