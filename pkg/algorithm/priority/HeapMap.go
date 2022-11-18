package priority

import "navigator/pkg/algorithm/priority/arry"

type Entry[KEY int | string, VAL any] struct {
	Key KEY
	Val VAL
	//
	//next *Entry[KEY,VAL]
	//pre *Entry[KEY,VAL]
}

func BuildEntry[KEY int | string, VAL any](key KEY, val VAL) *Entry[KEY, VAL] {
	return &Entry[KEY, VAL]{Key: key, Val: val}
}

type HeapMap[KEY int | string, VAL any] struct {
	EntryList arry.SortedArrayList[*Entry[KEY, VAL]]
	Map       map[KEY]*Entry[KEY, VAL]
}

func (m *HeapMap[KEY, VAL]) Init(compare func(one *Entry[KEY, VAL], other *Entry[KEY, VAL]) int) {
	m.EntryList = arry.SortedArrayList[*Entry[KEY, VAL]]{DoCompare: compare, Array: []*Entry[KEY, VAL]{}}
	m.Map = map[KEY]*Entry[KEY, VAL]{}
}

func (m *HeapMap[KEY, VAL]) Put(key KEY, val VAL) (ret VAL) {
	var entry = BuildEntry(key, val)
	if m.Map[key] != nil {
		ret = m.Map[key].Val
		//m.EntryList.Remove()
	}
	m.EntryList.Push(entry)
	m.Map[key] = entry
	return ret
}

func (m *HeapMap[KEY, VAL]) Get(key KEY) (ret VAL) {
	if m.Map[key] != nil {
		ret = m.Map[key].Val
	}
	return ret
}

func (m *HeapMap[KEY, VAL]) Values() (ret VAL) {
	//if m.Map[key] != nil {
	//	ret = m.Map[key].Val
	//}
	return ret
}
