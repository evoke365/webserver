package auth

type Store interface {
	Get(key string, valuePtr interface{}) (uint64, error)
	Upsert(key string, value interface{}, expiry uint32) (uint64, error)
}
