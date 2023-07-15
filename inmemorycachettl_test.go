package main

import (
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	c := NewCache()
	c.Set("test", 1, time.Second*5)
}

func TestGetEmpty(t *testing.T) {
	c := NewCache()
	value, err := c.Get("test")
	if err == nil {
		t.Errorf("Getting empty value from cache must return not found error")
	}
	if value != nil {
		t.Errorf("Getting empty value from cache must return nil value. Returned: %v", value)
	}
}

func TestGetExisting(t *testing.T) {
	expectedValue := 1
	key := "test"

	c := NewCache()
	c.Set(key, expectedValue, time.Second*5)
	value, err := c.Get(key)
	if err != nil {
		t.Errorf("Getting existing value must not return error: %s", err.Error())
	}
	if value != expectedValue {
		t.Errorf("Expected %d. Returned value: %v", expectedValue, value)
	}
}

func TestDeleteEmpty(t *testing.T) {
	key := "test"

	c := NewCache()
	c.Delete(key)
	value, err := c.Get(key)
	if err == nil {
		t.Errorf("Getting deleted value must return not found error")
	}
	if value != nil {
		t.Errorf("Getting deleted value must return nil value")
	}
}

func TestDelete(t *testing.T) {
	key := "test"

	c := NewCache()
	c.Set(key, 1, time.Second*5)
	c.Delete(key)
	value, err := c.Get(key)
	if err == nil {
		t.Errorf("Getting deleted value must return not found error")
	}
	if value != nil {
		t.Errorf("Getting deleted value must return nil value")
	}
	if len(c.storage) != 0 {
		t.Errorf("After deleting a single value from storage, storage must be empty")
	}
}

func TestExpired(t *testing.T) {
	key := "test"

	c := NewCache()
	c.Set(key, 1, time.Second)

	time.Sleep(time.Second * 2)

	value, err := c.Get(key)
	if err == nil {
		t.Errorf("Getting expired value must return not found error")
	}
	if value != nil {
		t.Errorf("Getting expired value must return nil value")
	}
	if len(c.storage) != 0 {
		t.Errorf("After expiration of a single value from storage, storage must be empty")
	}
}
