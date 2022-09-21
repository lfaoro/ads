package browser

import "github.com/lfaoro/ads/datas/stack"

type History struct {
	backward stack.Stack[string]
	forward  stack.Stack[string]
}

func (b *History) Visit(v string) {
	b.backward.Push(v)
}

func (b *History) Back() string {
	v := b.backward.Pop()
	b.forward.Push(v)
	return v
}

func (b *History) Forward() string {
	v := b.forward.Pop()
	b.backward.Push(v)
	return v
}
