package main

import "fmt"

// Generic interface
type Cache[K comparable, V any] interface {
	Put(key K, val V)
	Get(key K) (V, error)
	Remove(key K) error
}

// Generic Struct
type MyCache[K comparable, V any] struct {
	Storage map[K]V
}

// Generic function
func InitCache[K comparable, V any]() Cache[K, V] {
	return &MyCache[K, V]{
		Storage: make(map[K]V),
	}
}

func (m *MyCache[K, V]) Put(key K, val V) {
	m.Storage[key] = val
}

func (m *MyCache[K, V]) Get(key K) (V, error) {
	var val V
	if _, ok := m.Storage[key]; ok {
		val = m.Storage[key]
	} else {
		return val, fmt.Errorf("key not found")
	}
	return val, nil
}

func (m *MyCache[K, V]) Remove(key K) error {
	if _, ok := m.Storage[key]; ok {
		delete(m.Storage, key)
	} else {
		return fmt.Errorf("key not found")
	}
	return nil
}

func main() {

	cache1 := InitCache[string, int]()
	cache1.Put("hi", 0)
	cache1.Put("golang", 1)
	cache1.Put("hello", 2)

	//Uncommenting below statement will result in COMPILE-TIME ERROR since cache[string,int] can't take a string value. Hence generics allow TYPE-SAFETY.
	//cache1.Put("hello", "world")
	val, err := cache1.Get("hi")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Key = ", "hi", ", val = ", val)

	val, err = cache1.Get("golang")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Key = ", "golang", ", val = ", val)
}

/*
Output:
Key =  hi , val =  0
Key =  golang , val =  1
*/
