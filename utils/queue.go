package utils

type Queue[T any] struct {
	Push func(T)
	Pop  func() (T, bool)
	Peek func() (T, bool)
}

func NewQueue[T any]() Queue[T] {
	queue := make([]T, 0)
	length := 0
	return Queue[T]{
		Push: func(elem T) {
			queue = append(queue, elem)
			length++
		},
		Pop: func() (T, bool) {
			if length == 0 {
				return *new(T), false
			} else {
				elem := queue[0]
				queue = queue[1:]
				length--
				return elem, true
			}
		},
		Peek: func() (T, bool) {
			if length == 0 {
				return *new(T), false
			} else {
				return queue[0], true
			}
		},
	}
}
