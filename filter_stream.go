package stream

// Filter 过滤
func (s Stream) Filter(predicate Predicate) Stream {
	ns := NewStream()
	go func() {
		defer close(ns)
		for v := range s {
			if predicate(v) {
				ns <- v
			}
		}
	}()
	return ns
}
