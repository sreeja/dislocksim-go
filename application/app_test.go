package main

import "testing"

func TestGetexectime(t *testing.T) {
	got, err := getexectime("test", "opa")
	want := 5
	if err != nil || got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestGetexectimeInvalidfile(t *testing.T) {
	_, err := getexectime("testing", "opa")
	if err == nil {
		t.Errorf("should throw an error: no such file or directory")
	}
}

func TestGetexectimeInvalidop(t *testing.T) {
	_, err := getexectime("test", "opz")
	if err == nil {
		t.Errorf("should throw an error: Operation not found")
	}
}

func TestGetlocks(t *testing.T) {
	got, err := getlocks("test", "opa", "1", "1")
	want := []Lock{Lock{"opa_opb", "shared", LockType{"opa_opb", "p", "cent"}}}
	if err != nil || len(got) != len(want) {
		t.Errorf("got %v want %v", got, want)
	} else {
		if got[0] != want[0] {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestGetlocksInvalidOplockfile(t *testing.T) {
	_, err := getlocks("test", "opa", "2", "1")
	if err == nil {
		t.Errorf("should throw an error: no such file or directory")
	}
}

func TestGetlocksInvalidLocktypefile(t *testing.T) {
	_, err := getlocks("test", "opa", "1", "2")
	if err == nil {
		t.Errorf("should throw an error: no such file or directory")
	}
}
