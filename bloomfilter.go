package bloomfilter

type BloomFilter interface {
	Add(key string)
	Test(key string) bool
	Data() []byte
}
