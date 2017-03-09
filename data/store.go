package data

import (
	"strconv"

	"github.com/hugorut/butter/sys"

	"bytes"
	"encoding/base64"
	"encoding/gob"
	"errors"

	"fmt"

	"regexp"

	"encoding/json"

	"time"

	"os"

	"github.com/garyburd/redigo/redis"
)

var ErrorNoValue = errors.New("key used has nil value")

// Store is an interface that defines a store that can be used for
// key value based operations
type Store interface {
	ChangeExpiration(time.Duration) Store
	Set(string, string) error
	SetEx(string, time.Duration, string) error
	Get(string) (StoreValue, error)
	Del(string) error
	Keys(string) (StoreValue, error)
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

// ChangeExpiration is not supported in in memory store
func (i *InMemoryStore) ChangeExpiration(exp time.Duration) Store {
	return i
}

// SetEx defaults to Set as in memory store does not support expiring keys
func (i *InMemoryStore) SetEx(key string, exp time.Duration, val string) error {
	return i.Set(key, val)
}

// Get the keys using a regular expression selection
func (i *InMemoryStore) Keys(s string) (StoreValue, error) {
	reg := regexp.MustCompile(s)
	var keys []string

	for k := range i.Mem {
		if reg.MatchString(k) {
			keys = append(keys, k)
		}
	}

	b, _ := json.Marshal(keys)
	return &InMemoryValue{string(b)}, nil
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
	Pool        *redis.Pool
	expires     time.Duration
	neverExpire bool
}

// NewPool returns a redis client with a a max number of pool workers set
func NewPool(port string, maxIdle int) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", port)

			if err != nil {
				return nil, err
			}

			return c, err
		},
		MaxIdle: maxIdle,
	}
}

// NewRedisStore returns a pointer to a RedisStore that uses env variables to define its connection
func NewRedisStore() *RedisStore {
	poolSize := sys.EnvOrDefault("REDIS_MAX_POOL_SIZE", "10")
	i, err := strconv.Atoi(poolSize)

	if err != nil {
		i = 10
	}

	url := fmt.Sprintf("%s:%s",
		sys.EnvOrDefault("REDIS_HOST", "localhost"),
		sys.EnvOrDefault("REDIS_PORT", "6379"),
	)

	d, err := time.ParseDuration(
		sys.EnvOrDefault("REDIS_DEFAULT_EXPIRY", "86400s"),
	)

	if err != nil {
		d, _ = time.ParseDuration("86400s")
	}

	var neverExpire bool
	v := os.Getenv("REDIS_NEVER_EXPIRE")
	if v != "" {
		neverExpire = true
	}

	return &RedisStore{
		Pool:        NewPool(url, i),
		expires:     d,
		neverExpire: neverExpire,
	}
}

// ChangeExpiration changes the global expiration of the key for redis
func (r *RedisStore) ChangeExpiration(expires time.Duration) Store {
	r.expires = expires

	return r
}

// Set stores a key for a given amount of time
func (r *RedisStore) SetEx(key string, exp time.Duration, val string) error {
	_, err := r.Pool.Get().Do("SETEX", key, int(exp.Seconds()), val)
	return err
}

// Set adds a value to the redis store
func (r *RedisStore) Set(key, val string) error {
	if r.neverExpire {
		_, err := r.Pool.Get().Do("SET", key, val)
		return err
	}

	return r.SetEx(key, r.expires, val)
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

// Get the keys from a string
func (r *RedisStore) Keys(key string) (StoreValue, error) {
	rep, err := r.Pool.Get().Do("KEYS", key)

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
	switch r.Reply.(type) {
	case []byte:
		return r.Reply.([]byte)
	case []interface{}:
		is := r.Reply.([]interface{})
		var vals []string
		for _, i := range is {
			switch i.(type) {
			case []byte:
				vals = append(vals, string(i.([]byte)))
			}
		}

		b, _ := json.Marshal(vals)
		return b
	}

	return make([]byte, 0)
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
