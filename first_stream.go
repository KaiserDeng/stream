package stream

// First 获取第一个元素
func (s Stream) First() interface{} {
	list := s.ToList()
	return list[0]
}

// Last 获取第一个元素
func (s Stream) Last() interface{} {
	list := s.ToList()
	return list[len(list)-1]
}
