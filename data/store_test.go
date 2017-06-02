package data

import (
	"testing"

	"encoding/json"

	"os"
	"strings"

	"time"

	"github.com/alicebob/miniredis"
)

func TestInMemoryStore_Set_SetsValueInMap(t *testing.T) {
	store := NewInMemoryStore()

	store.Set("my-key", "val")
	val, _ := store.Get("my-key")

	b := val.Value()

	if string(b) != "val" {
		t.Errorf("failed asserting that value at key [my-key] is equal to 'val' instead got '%s'", string(b))
	}
}

type keysTest struct {
	vals         map[string]string
	regex        string
	s            string
	expectedKeys []string
}

var keysTests = []keysTest{
	{
		vals: map[string]string{
			"my-key":  "val",
			"my-key2": "val2",
		},
		regex: ".*",
		s:     "*",
		expectedKeys: []string{
			"my-key",
			"my-key2",
		},
	},
	{
		vals: map[string]string{
			"my-xey":  "val",
			"my-key2": "val2",
		},
		regex: "my-x",
		s:     "my-x*",
		expectedKeys: []string{
			"my-xey",
		},
	},
}

func TestInMemoryStore_Keys_ReturnsListOfMatchedKeys(t *testing.T) {
	for _, test := range keysTests {
		store := NewInMemoryStore()
		for k, v := range test.vals {
			store.Set(k, v)
		}

		b, _ := store.Keys(test.regex)
		var keys []string
		json.Unmarshal(b.Value(), &keys)

		if !equalValues(test.expectedKeys, keys) {
			t.Errorf("failed asserting keys were the same expected %+v go %+v", test.expectedKeys, keys)
		}
	}

}

func TestInMemoryStore_Del_RemovesValue(t *testing.T) {
	store := NewInMemoryStore()
	store.Set("test", "value")

	store.Del("test")
	val, err := store.Get("test")

	if err != ErrorNoValue {
		b := val.Value()
		t.Errorf("Found unexpected value in key [test] %s", b)
	}
}

func equalValues(a, b []string) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	valueMap := make(map[string]bool)

	for i := range b {
		valueMap[b[i]] = true
	}

	for i := range a {
		if _, ok := valueMap[a[i]]; !ok {
			return false
		}
	}

	return true
}

func createNewTestRedisStore() (*miniredis.Miniredis, *RedisStore) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	addressPieces := strings.Split(s.Addr(), ":")
	os.Setenv("REDIS_MAX_POOL_SIZE", "1")
	os.Setenv("REDIS_HOST", addressPieces[0])
	os.Setenv("REDIS_PORT", addressPieces[1])

	return s, NewRedisStore()
}

func TestRedisStore_Get_returnsStoredValue(t *testing.T) {
	s, mem := createNewTestRedisStore()
	defer s.Close()

	s.Set("test", "my value")
	val, _ := mem.Get("test")
	b := val.Value()

	if "my value" != string(b) {
		t.Errorf("Failed asserting that the redis store got the correct value.\nExpected: %s\nGot: %s", "my value", string(b))
	}

}

func TestRedisStore_Set_StoresValueCorrectly(t *testing.T) {
	s, mem := createNewTestRedisStore()
	defer s.Close()

	mem.Set("test", "my set value")

	val, _ := s.Get("test")

	if "my set value" != val {
		t.Errorf("Failed asserting that the redis store set the correct value.\nExpected: %s\nGot: %s", "my set value", val)
	}

}

func TestRedisStore_Keys_ReturnsListOfMatchedKeys(t *testing.T) {
	for _, test := range keysTests {
		s, mem := createNewTestRedisStore()

		for k, v := range test.vals {
			s.Set(k, v)
		}

		b, _ := mem.Keys(test.s)
		var keys []string
		json.Unmarshal(b.Value(), &keys)

		if !equalValues(test.expectedKeys, keys) {
			t.Errorf("failed asserting keys were the same expected %+v got %+v", test.expectedKeys, keys)
		}

		s.Close()
	}

}

func TestRedisStore_Del_RemovesKey(t *testing.T) {
	s, mem := createNewTestRedisStore()
	defer s.Close()

	s.Set("test", "value")
	mem.Del("test")

	if s.Exists("test") {
		t.Errorf("failed asserting that key [test] was deleted")
	}
}

func TestRedisStore_SetEx(t *testing.T) {
	s, mem := createNewTestRedisStore()
	defer s.Close()

	mem.SetEx("test", time.Second*12, "value")
	exp := s.TTL("test")

	if exp.Seconds() != 12 {
		t.Errorf("failed asserting key [test] was set for 12 seconds. Instead ttl was %v", exp.Seconds())
	}

}
