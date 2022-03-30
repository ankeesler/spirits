package runner

import "sync"

type syncBuffer struct {
	mu    sync.Mutex
	buf   []byte
	readI int
}

func (sb *syncBuffer) read() string {
	sb.mu.Lock()
	defer sb.mu.Unlock()

	s := string(sb.buf[sb.readI:len(sb.buf)])
	sb.readI = len(sb.buf)

	return s
}

func (sb *syncBuffer) Write(b []byte) (int, error) {
	sb.mu.Lock()
	defer sb.mu.Unlock()

	sb.buf = append(sb.buf, b...)

	return len(b), nil
}

func (sb *syncBuffer) reset() {
	sb.mu.Lock()
	defer sb.mu.Unlock()

	sb.buf = []byte{}
	sb.readI = 0
}
