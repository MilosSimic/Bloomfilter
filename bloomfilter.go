package main

type BloomFilter interface {
	Add(key string)
	Test(key string) bool
}
