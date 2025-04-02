package queue

type Queue[E any] struct {
	buf []E
}

func New[E any](cap ...int) *Queue[E] {
	if len(cap) > 0 {
		return &Queue[E]{
			buf: make([]E, 0, cap[0]),
		}
	}
	return &Queue[E]{
		buf: make([]E, 0),
	}
}

func (q *Queue[E]) Offer(e E) {
	q.buf = append(q.buf, e)
}

func (q *Queue[E]) Poll() E {
	if q.IsEmpty() {
		var zero E
		return zero
	}
	re := q.buf[0]
	q.buf = q.buf[1:]
	return re
}

func (q *Queue[E]) Peek() E {
	if q.IsEmpty() {
		var zero E
		return zero
	}
	return q.buf[0]
}

func (q *Queue[E]) Len() int {
	return len(q.buf)
}

func (q *Queue[E]) IsEmpty() bool {
	return q.Len() == 0
}
