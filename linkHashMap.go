package linkedHashMap

import (
	"math"
)

const HASH_TABLE_SIZE = 2;

type entry struct {
	key    int
	value  interface{}
	after  *entry
	before *entry
	next   *entry
}

type LinkedHashMap struct {
	base     *entry
	table    []*entry
	linkSize int
	hashSize int
	isLRU    bool
}

func NewLinkedHashMap() *LinkedHashMap {
	return &LinkedHashMap{hashSize: HASH_TABLE_SIZE}
}

func NewEntity(key int, value interface{}) *entry {
	return &entry{key, value, nil, nil, nil}
}

func (this *LinkedHashMap) Init(initialCapacity int, isLRU bool) {
	this.hashSize = int(math.Pow(2, math.Round(math.Log2(float64(initialCapacity)))))
	this.table = []*entry{}
	this.isLRU = isLRU
	for i := 0; i < this.hashSize; i++ {
		this.table = append(this.table, nil)
	}
	this.base = &entry{}
}

func (this *LinkedHashMap) Put(key int, value interface{}) {
	e, exist := this.bindEntry(key, value)
	if !exist {
		this.linkSize++
		if this.linkSize > 1 {
			this.base.before.after = e
			e.before = this.base.before
			e.after = this.base
			this.RecordAccess(this, e)
		}
	}
}

func (this *LinkedHashMap) Get(key int) interface{} {
	i := hashCode(key, this.hashSize)
	e := this.table[i]
	for e.next != nil {
		if e.key == key {
			this.RecordAccess(this, e)
			return e.value
		}
		e = e.next
	}
	return nil
}

func (this *LinkedHashMap) Remove(key int) (value interface{}, exists bool) {
	var isExists bool
	var r interface{}
	i := hashCode(key, this.hashSize)
	if item := this.table[i]; item == nil {
		return nil, exists
	} else {
		var pointer *entry
		for item != nil {
			if item.key == key {
				item.before.after = item.after
				item.after.before = item.before
				if item.next != nil {
					pointer.next = item.next
				}
				isExists = true
				break
			}
			pointer = item
			item = item.next
		}
		*item=entry{}
	}
	this.linkSize--
	return r, isExists
}

func (this *LinkedHashMap) RecordAccess(lhm *LinkedHashMap, e *entry) {
	tem:=*e
	if (this.isLRU) {
		if _, ok := lhm.Remove(e.key); ok {
			lhm.AddBefore(lhm.base.after.key, &tem)
		}
	}
}

func (this *LinkedHashMap) AddBefore(key int, e *entry, ) (bool) {
	if this.linkSize == 0 {
		return false
	}
	be, exits := this.bindEntry(e.key, e.value)
	if ct, ok := this.GetEntry(key); ok && !exits {
		//ct target,be ele
		be.before=ct.before
		be.after=ct
		be.before.after=be
		this.linkSize++
		return true
	} else {
		return false
	}
}

func hashCode(key int, len int) int {
	return key & (len - 1)
}

func (this *LinkedHashMap) bindEntry(key int, value interface{}) (e *entry, exits bool) {
	i := hashCode(key, this.hashSize)
	e = NewEntity(key, value)
	if item := this.table[i]; item == nil {
		this.table[i] = e
		if (this.base.after == nil) {
			this.base.after = e
			this.base.before = e
			e.before = this.base
			e.after = this.base
		}
		return e, false
	} else {
		for item.next != nil {
			if (item.key == key) {
				return item, true
			}
			item = item.next
		}
		item.next = e
		return e, false
	}
}

func (this *LinkedHashMap) GetEntry(key int) (*entry, bool) {
	i := hashCode(key, this.hashSize)
	e := this.table[i]
	for e != nil {
		if e.key == key {
			return e, true
		}
		e = e.next
	}
	return nil, false
}

func (this *LinkedHashMap) Base() *entry {
	return this.base
}

func (this *LinkedHashMap) LinkSize() int {
	return this.linkSize
}
