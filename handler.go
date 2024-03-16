package snowflake

func Spin(w *idWorker) error {
	var t int64 = 0
	for {
		t = ms()
		if t > w.last {
			break
		}
	}
	w.seq = 0
	w.last = t
	return nil
}

func Rent(w *idWorker) error {
	if w.seq+1 < seqMax {
		w.seq += 1
		return nil
	}
	w.last += 1
	w.seq = 0
	return nil
}

func WithCustomRollBackHandler(f func(*idWorker) error) Option {
	return func(w *idWorker) {
		w.rollBackHandler = f
	}
}
