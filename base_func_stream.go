package stream

// Function Function<T,R> 一进一出
type Function func(element interface{}) interface{}

// Consumer Consumer<T> 只进不出
type Consumer func(element interface{})

// Supplier Supplier<T> 只出不进
type Supplier func() interface{}

// Predicate Predicate<T> 只出布尔值
type Predicate func(element interface{}) bool

// Sink  只出流
type Sink func(element interface{}) Stream
