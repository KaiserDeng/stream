package stream

// Map 转换
func (s Stream) Map(fun Function) Stream {
	ns := NewStream()
	// 将通道数据转换后写入到新流中
	go func() {
		defer close(ns)
		for i := range s {
			ns <- fun(i)
		}
	}()
	return ns
}
