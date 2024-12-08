package util

type LoopingIndex struct {
	index int
	max   int
}

func NewLoopingIndex(max int) *LoopingIndex {
	return &LoopingIndex{0, max}
}

func (l *LoopingIndex) Current() int {
	return l.index
}

func (l *LoopingIndex) Next() int {
	l.index++
	if l.index > l.max {
		l.Reset()
	}
	return l.index
}

func (l *LoopingIndex) Set(i int) {
	l.index = i
}

func (l *LoopingIndex) Reset() {
	l.index = 0
}
