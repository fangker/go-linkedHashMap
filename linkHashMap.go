package linkedHashMap

import (
	"math"
)

const HASH_TABLE_SIZE = 2 << 7;

type entry struct {
	key    int
	value  interface{}
	after  *entry
	before *entry
	next   *entry
}

type LinkedHashMap struct {
	base        *entry
	table       []*entry
	linkSize    int
	hashSize    int
	isLRU       bool
	loadFactory float64
}

func NewEntity(key int, value interface{}) *entry {
	return &entry{key, value, nil, nil, nil}
}

func NewLinkedHashMap(initialCapacity int, loadFactor float64, isLRU bool) *LinkedHashMap {
	this := &LinkedHashMap{hashSize: HASH_TABLE_SIZE}
	this.hashSize = int(math.Pow(2, math.Round(math.Log2(float64(initialCapacity)))))
	this.table = []*entry{}
	this.isLRU = isLRU
	for i := 0; i < this.hashSize; i++ {
		this.table = append(this.table, nil)
	}
	this.base = &entry{}
	this.loadFactory = loadFactor
	return this
}

func (this *LinkedHashMap) Put(key int, value interface{}) {
	e, exist, lhm := this.bindEntry(key, value)
	if lhm != this {
		// extend
		for e := this.base.after; false == this.isBase(e); {
			lhm.Put(e.key, e.value)
			e = e.after
		}
		*this = *lhm
		this.Put(key, value)
		return
	}
	if !exist {
		this.linkSize++
		if this.linkSize > 1 {
			this.base.before.after = e
			e.before = this.base.before
			e.after = this.base
			this.base.before = e
			this.Base()
			this.RecordAccess(e)
		}
	} else if (this.linkSize >= 1) {
		this.RecordAccess(e)
	}
}

func (this *LinkedHashMap) Get(key int) interface{} {
	i := hashCode(key, this.hashSize)
	e := this.table[i]
	for e.next != nil {
		if e.key == key {
			this.RecordAccess(e)
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
					if (pointer == nil) {
						this.table[i] = item.next
					} else {
						pointer.next = item.next
					}
				} else {
					if (this.table[i].key == key) {
						this.table[i] = nil
					}
				}
				isExists = true
				break
			}
			pointer = item
			item = item.next
		}
		if (pointer != nil) {
			pointer.next = nil
		}
	}
	this.linkSize--
	return r, isExists
}

func (this *LinkedHashMap) RecordAccess( e *entry) {
	if (this.isLRU) {
		 this.MoveAfter(this.base,e)
	}
}

func (this *LinkedHashMap) AddBefore(key int, e *entry, ) (bool) {
	if this.linkSize == 0 {
		return false
	}
	be, exits, _ := this.bindEntry(e.key, e.value)
	if ct, ok := this.GetEntry(key); ok && !exits {
		//ct-> target,be -> ele

		be.before = ct.before
		be.after = ct
		be.before.after = be
		ct.before = be
		this.linkSize++
		return true
	} else {
		return false
	}
}

func hashCode(key int, len int) int {
	return key & (len - 1)
}

func (this *LinkedHashMap) bindEntry(key int, value interface{}) (e *entry, exits bool, lhm *LinkedHashMap) {
	i := hashCode(key, this.hashSize)
	e = NewEntity(key, value)
	if item := this.table[i]; item == nil {
		this.table[i] = e
		if (this.linkSize == 0) {
			this.base.before = e
		}
		if (this.base.after == nil) {
			this.base.after = e
			e.before = this.base
			e.after = this.base
		}
		return e, false, this
	} else {
		index := 0
		for item != nil {
			if (item.key == key) {
				return item, true, this
			}
			if item.next == nil {
				break
			}
			item = item.next
			index++
		}
		index++
		if index >= int(this.loadFactory*float64(this.hashSize)) {
			// need rehash and exten
			initalCapacity := int(math.Pow(2, math.Round(math.Log2(float64(this.hashSize*2)))))
			lhm := NewLinkedHashMap(initalCapacity, this.loadFactory, this.isLRU)
			return e, false, lhm
		}
		item.next = e
		return e, false, this
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

func (this *LinkedHashMap) isBase(e *entry) bool {
	return this.base == e
}

func (this *LinkedHashMap) MoveAfter(traget,e *entry){
	 e.before.after=e.after
	 e.after.before = e.before
	 traget.after.before = e
	 traget.after = e
	 e.after = traget.after.before
	 e.before =traget
}