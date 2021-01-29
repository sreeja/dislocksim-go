package main

import (
	"testing"
	// "time"
)

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
	got, err := getlocks("test", "opa", "1", "1", "1")
	want := []Lock{Lock{"opa_opb", "shared", LockType{"opa_opb", "p", "cent"}}}
	if err != nil || len(got) != len(want) {
		t.Errorf("got %v want %v", got, want)
	} else {
		if got[0] != want[0] {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestGetlocksSorting(t *testing.T) {
	got, err := getlocks("test", "opb", "1", "1", "1")
	want := []Lock{Lock{"opa_opb", "exclusive", LockType{"opa_opb", "p", "cent"}}, Lock{"opd_opb", "exclusive", LockType{"opd_opb", "p", "cent"}}}
	if err != nil || len(got) != len(want) {
		t.Errorf("got %v want %v", got, want)
	} else {
		if got[0] != want[0] {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestGetlocksInvalidOplockfile(t *testing.T) {
	_, err := getlocks("test", "opa", "1", "2", "1")
	if err == nil {
		t.Errorf("should throw an error: no such file or directory")
	}
}

func TestGetlocksInvalidLocktypefile(t *testing.T) {
	_, err := getlocks("test", "opa", "1", "1", "2")
	if err == nil {
		t.Errorf("should throw an error: no such file or directory")
	}
}

// func TestExecute(t *testing.T) {
// 	start := time.Now()
// 	err := execute("test", "opa", "1", "1", "1")
// 	if err != nil {
// 		t.Errorf("threw error %v", err)
// 	}
// 	end := time.Now()
// 	elapsed := end.Sub(start)
// 	if elapsed < time.Duration(5)*time.Millisecond {
// 		t.Errorf("not executed, time taken is %v", elapsed)
// 	}
// }
