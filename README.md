# Maps

## What are maps?

Maps, also called associative arrays or dictionaries, are an abstract data type
that can contain (key, value) data pairs. The association between a key and a
value is often known as a mapping.

The siplest real-life example of such a mapping is a phone book. It associates
names with phone numbers or maps names to phone numbers:

```txt
Jaroslavs : +41XXXXXXXX
Ksenija   : +371XXXXXXX
...
```

## Maps in Go

Source: https://go.dev/blog/maps

A Go map type looks like this:

```go
map[KeyType]ValueType
```

where `KeyType` may be any type that is comparable, and `ValueType` may be
absolutely any type including another map.

This example creates a variable `m` which is a map of string keys to int values:

```go
var m map[string]int
```

It is important to remember that maps are references (similarly to slices) and
must be initialized before being used:

```go
m = make(map[string]int)
```

### Working with maps

Using maps is similar to using slices.

Assigning:

```go
m["foo"] = 123
```

Reading:

```go
i := m["bar"]
```

Please note that if the requested key does not exist, the value type's zero
value is returned. The example above returns `0`.

The built-in `len` function returns the number of items in a map:

```go
n := len(m)
```

The built-in `delete` function removes an entry from the map:

```go
delete(m, "foo")
```

The `delete` function doesn't return anything and does nothing if the specified
key does not exist in the map.

A two-value assignment tests for the existence of a key:

```go
i, ok := m["bazz"]
```

In this statement, the first value `i` is assigned the value associated with the
key. If that key does not exist, `i` is the value type's zero value `0`. The
second value `ok` is a `bool` that is `true` if the key exists in the map, and
`false` if not.

To test an existence of a key without retrieving the value, use an underscore in
place of the first value:

```go
_, ok := m["bazz"]
```

To iterate over the contents of a map, use the `for-range` loop:

```go
for key, value := range m {
  fmt.Println("Key:", key, "Value:", value)
}
```

When iterating over a map with a range loop, the iteration order is not specified and is not guaranteed to be the same from one iteration to the next. If you require a stable iteration order you must maintain a separate data structure that specifies that order. This example uses a separate sorted slice of keys to print a map[int]string in key order:

```go
import "sort"

var m map[int]string
var keys []int
for k := range m {
  keys = append(keys, k)
}
sort.Ints(keys)
for _, k := range keys {
  fmt.Println("Key:", k, "Value:", m[k])
}
```

To initialize a map with some data, use a map literal:

```go
words := map[string]int{
  "foo": 5,
  "bar": 1,
}
```
