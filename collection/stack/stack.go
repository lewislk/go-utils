package stack

type Stack[E any] struct {
	buf []E
}

func New[E any](cap ...int) *Stack[E] {
	if len(cap) > 0 {
		return &Stack[E]{
			buf: make([]E, 0, cap[0]),
		}
	}
	return &Stack[E]{
		buf: make([]E, 0),
	}
}

func (s *Stack[E]) Push(e E) {
	s.buf = append(s.buf, e)
}

func (s *Stack[E]) Pop() E {
	if s.IsEmpty() {
		var zero E
		return zero
	}
	re := s.buf[s.Len()-1]
	s.buf = s.buf[:s.Len()-1]
	return re
}

func (s *Stack[E]) Peek() E {
	if s.IsEmpty() {
		var zero E
		return zero
	}
	return s.buf[s.Len()-1]
}

func (s *Stack[E]) Len() int {
	return len(s.buf)
}

func (s *Stack[E]) IsEmpty() bool {
	return s.Len() == 0
}
