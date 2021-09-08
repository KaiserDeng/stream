package stream

// FlatMap 压平流
func (s Stream) FlatMap(fun Sink) Stream {
	ns := NewStream()
	// 将通道数据转换后写入到新流中
	go func() {
		defer close(ns)
		for i := range s {
			ss := fun(i)
			for iv := range ss {
				ns <- iv
			}
		}
	}()
	return ns
}
