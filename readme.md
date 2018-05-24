## GO LinkedHashMap

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
lhm.Init(<int> initialCapasity, <bool> useLRU)
```
The initial capasity parameter can suggest the length of hash arry

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
```

```
