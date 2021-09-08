package stream

// ForEach 过滤
func (s Stream) ForEach(consumer Consumer) {
	for v := range s {
		consumer(v)
	}
}
