package stream

import "reflect"

// Distinct 去重
func (s Stream) Distinct() Stream {
	list := s.ToList()
	var newList []interface{}
	// 先去重
	size := len(list)
	for i := 0; i < size; i++ {
		state := false
		for j := i + 1; j < size; j++ {
			if j > 0 && reflect.DeepEqual(list[i], list[j]) {
				state = true
				break
			}
		}
		if !state {
			newList = append(newList, list[i])
		}
	}
	ns := NewStream()
	go func() {
		defer close(ns)
		for _, value := range newList {
			ns <- value
		}
	}()
	return ns
}
