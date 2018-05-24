package linkedHashMap

import (
	"math"
	"fmt"
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
	isRUL    bool
}

func NewLinkedHashMap() *LinkedHashMap {
	return &LinkedHashMap{hashSize: HASH_TABLE_SIZE}
}

func NewEntity(key int, value interface{}) *entry {
	return &entry{key, value, nil, nil, nil}
}

func (this *LinkedHashMap) Init(initialCapacity int, isRUL bool) {
	this.hashSize = int(math.Pow(2, math.Round(math.Log2(float64(initialCapacity)))))
	this.table = []*entry{}
	this.isRUL=isRUL
	for i := 0; i < this.hashSize; i++ {
		this.table = append(this.table, nil)
	}
	this.base = &entry{}
}

func (this *LinkedHashMap) Put(key int, value interface{}) {
	e, _ := this.bindEntry(key, value)
	this.linkSize++
	if this.linkSize!=1{
	this.base.before.after=e
	e.before = this.base.before
	e.after = this.base
		this.RecordAccess(this, e)
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

func (this *LinkedHashMap) Remove(key int) (interface{}, bool) {
	var exists bool
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
				pointer.next = item.next
				break
			}
			pointer = item
			item = item.next
			this.linkSize--
			exists = true
		}
		return r, exists
	}
}

func (this *LinkedHashMap) RecordAccess(lhm *LinkedHashMap, e *entry) {
	if (this.isRUL) {
		fmt.Println(this.table)
		if _, ok := lhm.Remove(e.key); ok {
			lhm.AddBefore(lhm.base.after.key, e)
		}
	}
}

func (this *LinkedHashMap) AddBefore(key int, e *entry, ) (bool) {
	if this.linkSize == 0 {
		return false
	}
	e, exits := this.bindEntry(e.key, e.value)
	if ct, ok := this.GetEntry(key); ok && !exits {
		e.after = ct
		ct.before.after = e
		e = ct.before
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
		this.table[i]=e
		if (this.base.after == nil) {
			this.base.after = e
			this.base.before = e
			e.before = this.base
			e.after = this.base
		}
		return e, false
	} else {
		for item.next != nil {
			item = item.next
		}
		if (item.key == key) {
			return item, true
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

