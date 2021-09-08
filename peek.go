package stream

// Peek 偷窥
func (s Stream) Peek(fun Function) Stream {
	ns := NewStream()
	go func() {
		defer close(ns)
		for v := range s {
			ns <- fun(v)
		}
	}()
	return ns
}
