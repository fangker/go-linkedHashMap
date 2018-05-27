## GO LinkedHashMap

LinkedHashMap supports dynamically capacity increasing.It's  It's very simple to build
an least recently used cache.

###  install
```go
go get github.com/fangker/go-linkedHashMap
```

### Usage

#### import

```go
import "github.com/fangker/go-linkedHashMap"
```

#### Create a linkedHashMap
```go
lhm := NewLinkedHashMap()
```

#### Initialize the LHM
```go
lhm.Init(<int> initialCapacity, <bool> useLRU)
```
The initial capacity parameter can suggest the length of hash arry

#### Put Hash data
```go
lhm.Put(<int> key, <interface> data)
```
#### Get Hash data
```go
lhm.Get(<int> key, <interface> data)
```

#### Get Hash data
```go
lhm.Get(<int> key, <interface> data)
```

#### Get Base item
```go
lhm.Base()
```

#### implement LRU Algorithm
You can override "RecordAccess" function to implement LRU Algorithm.
The default "RecordAccess" method "MoveTo" will be invoked.
