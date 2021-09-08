package stream

import "reflect"

type Stream chan interface{}

type KenFunc func(element interface{}) interface{}

// From 构建数据源
func (s Stream) From(src interface{}) Stream {
	v := reflect.ValueOf(src)
	switch k := v.Kind(); k {
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		go func() {
			defer close(s)
			for i := 0; i < v.Len(); i++ {
				s <- v.Index(i).Interface()
			}
		}()
	default:
		panic("get a non-slice kind source")
	}
	return s
}

// NewStream 新建一个流
func NewStream() Stream {
	return make(chan interface{})
}

// ToList 导出流数据为一个列表
func (s Stream) ToList() (retArr []interface{}) {
	for i := range s {
		retArr = append(retArr, i)
	}
	return retArr
}
