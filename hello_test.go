package main

import (
    "testing"
)

func TestHello(t *testing.T) {
    actual := hello()
    expected := "hello world"
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}
