package snowflake

import (
	"math/rand"
	"sync"
	"time"
)

const seqMax = (1 << 12) - 1

type idWorker struct {
	last  int64
	start int64
	seq   int64
	mid   int64

	mutex           *sync.Mutex
	rollBackHandler func(*idWorker) error
}

type Option func(w *idWorker)

func NewIDGenerator(ops ...Option) IDGenerator {
	i := &idWorker{
		last:            0,
		start:           ms(),
		seq:             0,
		mid:             rand.Int63n(1 << 10),
		mutex:           &sync.Mutex{},
		rollBackHandler: Spin,
	}
	for _, op := range ops {
		op(i)
	}
	return i
}

func (w *idWorker) geneId() (ret int64) {
	ret = w.last << 22
	ret |= w.mid << 12
	ret |= w.seq
	return
}

func (w *idWorker) NextID() int64 {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	now := ms()
	// if now > w.last, it means the time has changed, we can reset the sequence
	if now > w.last {
		w.last = now
		w.seq = 0
		return w.geneId()
	}
	// if now == w.last, it means the time has not changed, we need to increase the sequence
	if now == w.last {
		if w.seq+1 < seqMax {
			w.seq += 1
			return w.geneId()
		}
	}

	// if now < w.last, it means the time has not changed, we need to increase the sequence
	err := w.rollBackHandler(w)
	if err != nil {
		return -1
	}
	return w.geneId()
}

func ms() int64 {
	return time.Now().UTC().UnixMilli()
}

func WithEpoch(start int64) Option {
	return func(w *idWorker) {
		w.start = start
	}
}

func WithMid(mid int64) Option {
	return func(w *idWorker) {
		w.mid = mid
	}
}

func WithSpin() Option {
	return func(w *idWorker) {
		w.rollBackHandler = Spin
	}
}

func WithRent() Option {
	return func(w *idWorker) {
		w.rollBackHandler = Rent
	}
}
