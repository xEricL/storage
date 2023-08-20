package couchbase

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSetCouchbase_ShouldReturnNoError(t *testing.T) {
	testStorage := New(Config{
		Username: "admin",
		Password: "123456",
		Host:     "127.0.0.1:8091",
		Bucket:   "fiber_storage",
	})

	err := testStorage.Set("test", []byte("test"), 0)

	require.Nil(t, err)
}

func TestGetCouchbase_ShouldReturnNil_WhenDocumentNotFound(t *testing.T) {
	testStorage := New(Config{
		Username: "admin",
		Password: "123456",
		Host:     "127.0.0.1:8091",
		Bucket:   "fiber_storage",
	})

	val, err := testStorage.Get("not_found_key")

	require.Nil(t, err)
	require.Zero(t, len(val))
}

func TestSetAndGet_GetShouldReturn_SettedValueWithoutError(t *testing.T) {
	testStorage := New(Config{
		Username: "admin",
		Password: "123456",
		Host:     "127.0.0.1:8091",
		Bucket:   "fiber_storage",
	})

	err := testStorage.Set("test", []byte("fiber_test_value"), 0)
	require.Nil(t, err)

	val, err := testStorage.Get("test")

	require.Nil(t, err)
	require.Equal(t, val, []byte("fiber_test_value"))
}

func TestSetAndGet_GetShouldReturnNil_WhenTTLExpired(t *testing.T) {
	testStorage := New(Config{
		Username: "admin",
		Password: "123456",
		Host:     "127.0.0.1:8091",
		Bucket:   "fiber_storage",
	})

	err := testStorage.Set("test", []byte("fiber_test_value"), 3*time.Second)
	require.Nil(t, err)

	time.Sleep(6 * time.Second)

	val, err := testStorage.Get("test")

	require.Nil(t, err)
	require.Zero(t, len(val))
}

func TestSetAndDelete_DeleteShouldReturn_NoError(t *testing.T) {
	testStorage := New(Config{
		Username: "admin",
		Password: "123456",
		Host:     "127.0.0.1:8091",
		Bucket:   "fiber_storage",
	})

	err := testStorage.Set("test", []byte("fiber_test_value"), 0)
	require.Nil(t, err)

	err = testStorage.Delete("test")
	require.Nil(t, err)

	_, err = testStorage.Get("test")
	require.Nil(t, err)
}

func TestSetAndReset_ResetShouldReturn_NoError(t *testing.T) {
	testStorage := New(Config{
		Username: "admin",
		Password: "123456",
		Host:     "127.0.0.1:8091",
		Bucket:   "fiber_storage",
	})

	err := testStorage.Set("test", []byte("fiber_test_value"), 0)
	require.Nil(t, err)

	err = testStorage.Reset()
	require.Nil(t, err)

	_, err = testStorage.Get("test")
	require.Nil(t, err)
}

func TestClose_CloseShouldReturn_NoError(t *testing.T) {
	testStorage := New(Config{
		Username: "admin",
		Password: "123456",
		Host:     "127.0.0.1:8091",
		Bucket:   "fiber_storage",
	})

	err := testStorage.Close()
	require.Nil(t, err)
}

func TestGetConn_ReturnsNotNill(t *testing.T) {
	testStorage := New(Config{
		Username: "admin",
		Password: "123456",
		Host:     "127.0.0.1:8091",
		Bucket:   "fiber_storage",
	})
	require.True(t, testStorage.Conn() != nil)
}
