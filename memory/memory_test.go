package memoryStore

import (
	"testing"
	"time"

	"github.com/gofiber/utils"
)

func Test_Set(t *testing.T) {

	store := New()

	id := "hello"
	value := []byte("Hi there!")

	store.Set(id, value, 0)

	utils.AssertEqual(t, dataPoint{value, 0}, store.data[id])

}

func Test_SetExpiry(t *testing.T) {

	store := New()

	id := "hello"
	value := []byte("Hi there!")
	expiry := time.Second * 10

	store.Set(id, value, expiry)

	now := time.Now().Unix()
	fromStore, found := store.data[id]
	utils.AssertEqual(t, true, found)

	delta := fromStore.expiry - now
	upperBound := int64(expiry.Seconds())
	lowerBound := upperBound - 2

	if !(delta <= upperBound && delta > lowerBound) {
		t.Fatalf("Test_SetExpiry: expiry delta out of bounds (is %d, must be %d<x<=%d)", delta, lowerBound, upperBound)
	}

}

func Test_GC(t *testing.T) {

	// New() isn't being used here so the gcInterval can be set low
	store := &MemoryStore{
		data:       make(map[string]dataPoint),
		gcInterval: time.Second * 1,
	}
	go store.gc()

	id := "hello"
	value := []byte("Hi there!")

	expireAt := time.Now().Add(time.Second * 2).Unix()

	store.data[id] = dataPoint{value, expireAt}

	time.Sleep(time.Second * 4) // The purpose of the long delay is to ensure the GC has time to run and delete the value

	_, found := store.data[id]
	utils.AssertEqual(t, false, found)

}

func Test_Get(t *testing.T) {

	store := New()

	id := "hello"
	value := []byte("Hi there!")

	store.data[id] = dataPoint{value, 0}

	returnedValue, err := store.Get(id)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, value, returnedValue)

}

func Test_Delete(t *testing.T) {

	store := New()

	id := "hello"
	value := []byte("Hi there!")

	store.data[id] = dataPoint{value, 0}

	err := store.Delete(id)
	utils.AssertEqual(t, nil, err)

	_, found := store.data[id]
	utils.AssertEqual(t, false, found)

}

func Test_Clear(t *testing.T) {

	store := New()

	id := "hello"
	value := []byte("Hi there!")

	store.data[id] = dataPoint{value, 0}

	err := store.Clear()
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, make(map[string]dataPoint), store.data)

}

func Benchmark_Set(b *testing.B) {

	store := New()

	id := "hello"
	value := []byte("Hi there!")
	expiry := time.Duration(0)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		store.Set(id, value, expiry)
	}

}

func Benchmark_Get(b *testing.B) {

	store := New()

	id := "hello"
	value := []byte("Hi there!")

	store.data[id] = dataPoint{value, 0}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		store.Get(id)
	}
}