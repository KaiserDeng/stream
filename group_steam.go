package stream

type (
	Entry struct {
		Key   interface{}
		Value []interface{}
	}

	GroupedStream chan Entry
)

// GroupBy 分组
func (s Stream) GroupBy(f Function) GroupedStream {

	entries := make(chan Entry)
	cache := make(map[interface{}][]interface{})
	for i := range s {
		key := f(i)
		cache[key] = append(cache[key], i)
	}
	go func() {
		defer close(entries)
		for k, v := range cache {
			entries <- Entry{
				Key:   k,
				Value: v,
			}
		}
	}()
	return entries
}

// Collect 收集器
func (s GroupedStream) Collect() map[interface{}][]interface{} {
	retMap := make(map[interface{}][]interface{})
	for i := range s {
		retMap[i.Key] = i.Value
	}
	return retMap
}
