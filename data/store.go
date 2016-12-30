package data

import (
	"butter/sys"
	"reflect"
	"strconv"

	"bytes"
	"encoding/base64"
	"encoding/gob"
	"errors"

	"github.com/garyburd/redigo/redis"
)

var ErrorNoValue = errors.New("key used has nil value")

// Store is an interface that defines a store that can be used for
// key value based operations
type Store interface {
	Set(string, string) error
	Get(string) (StoreValue, error)
	Del(string) error
}

// StoreValue defines an interface for a value stored in the Store
type StoreValue interface {
	// Value returns the stored value as bytes
	// so that it can be marshaled to a data source
	Value() []byte
}

// in memory store in order to facilitate simple
// testing
type InMemoryStore struct {
	Mem map[string]string
}

// Set adds a value to the map
func (i *InMemoryStore) Set(key, val string) error {
	i.Mem[key] = val

	return nil
}

// Get returns a value from the map using a key
func (i *InMemoryStore) Get(key string) (StoreValue, error) {
	return &InMemoryValue{i.Mem[key]}, nil
}

// Del deletes from the map
func (i *InMemoryStore) Del(key string) error {
	if _, ok := i.Mem[key]; ok {
		delete(i.Mem, key)
	}

	return nil
}

type InMemoryValue struct {
	Val string
}

// Value returns the string value as bytes
func (i *InMemoryValue) Value() []byte {
	return []byte(i.Val)
}

// RedisStore
type RedisStore struct {
	Pool *redis.Pool
}

// NewPool returns a redis client with a a max number of pool workers set
func NewPool(port string, maxIdle int) *redis.Pool {
	return redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", port)

		if err != nil {
			return nil, err
		}

		return c, err
	}, maxIdle)
}

// NewRedisStore returns a pointer to a RedisStore that uses env variables to define its connection
func NewRedisStore() *RedisStore {
	poolSize := sys.EnvOrDefault("REDIS_MAX_POOL_SIZE", "10")
	i, err := strconv.Atoi(poolSize)

	if err != nil {
		i = 10
	}

	return &RedisStore{
		Pool: NewPool(sys.EnvOrDefault("REDIS_PORT", ":6379"), i),
	}
}

// Set adds a value to the redis store
func (r *RedisStore) Set(key, val string) error {
	_, err := r.Pool.Get().Do("SET", key, val)

	return err
}

// Get returns a value from the redis store
func (r *RedisStore) Get(key string) (StoreValue, error) {
	rep, err := r.Pool.Get().Do("GET", key)

	value := &RedisStoreValue{rep}

	if err != nil {
		return value, err
	}

	if rep == nil {
		return value, ErrorNoValue
	}

	return value, err
}

// Del deletes from a redis key
func (r *RedisStore) Del(key string) error {
	_, err := r.Pool.Get().Do("DEL", key)

	return err
}

type RedisStoreValue struct {
	Reply interface{}
}

// Value returns the bytes of the stored value at a specific key
func (r *RedisStoreValue) Value() []byte {
	v := reflect.ValueOf(r.Reply)

	if r.Reply == nil {
		return make([]byte, 0)
	}

	return v.Bytes()
}

func RegisterForSerialization(is ...interface{}) {
	for _, i := range is {
		gob.Register(i)
	}
}

// Serialize take an interface and attempts to stream it to a string
// by encoding using base64 and gobs
func Serialize(i interface{}) (string, error) {
	buffer := bytes.NewBuffer([]byte{})
	err := gob.NewEncoder(buffer).Encode(i)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}

// DeSerialize attempts to hydrate a interface with a given string which
// has been serialized via base64 encoding
func DeSerialize(src string, output interface{}) error {
	out, err := base64.StdEncoding.DecodeString(src)

	if err != nil {
		return err
	}

	buffer := bytes.NewBuffer(out)
	err = gob.NewDecoder(buffer).Decode(output)

	return err
}
