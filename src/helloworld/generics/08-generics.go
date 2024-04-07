// https://habr.com/ru/companies/otus/articles/782414/
package main

import "fmt"

// Cache - универсальная структура кэша. T - тип хранимых значений
type Cache[T any] struct {
	store map[string]T
}

// NewCache создает новый экземпляр Cache
func NewCache[T any]() *Cache[T] {
	return &Cache[T]{store: make(map[string]T)}
}

// Set добавляет значение в кэш
func (c *Cache[T]) Set(key string, value T) {
	c.store[key] = value
}

// Get возвращает значение из кэша
func (c *Cache[T]) Get(key string) (T, bool) {
	val, found := c.store[key]
	return val, found
}

func (c *Cache[T]) Print() {
	// ???
}

func main() {
	// Создаем кэш для int
	intCache := NewCache[int]()
	intCache.Set("key1", 10)
	fmt.Println(intCache.Get("key1"))

	stringCache := NewCache[string]()
	stringCache.Set("hello", "world")
	fmt.Println(stringCache.Get("hello"))
}
