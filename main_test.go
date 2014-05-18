package main

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"
)

type testFile struct {
	path string
}

var (
	url = "https://api.github.com/repos/str1ngs/gotimer"
)

var testFiles = []testFile{
	{
		path: "testdata/gotimer.json",
	},
}

func TestReflect(t *testing.T) {
	for _, f := range testFiles {
		want, err := readWant(f.path + ".want")
		if err != nil {
			t.Fatal(err)
		}
		fd, err := os.Open(f.path)
		if err != nil {
			t.Error(err)
			continue
		}
		defer fd.Close()
		got := new(bytes.Buffer)
		err = read(fd, got)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(want, got.Bytes()) {
			t.Errorf("%s: want %d bytes got %d bytes", f.path, len(want), len(got.Bytes()))
		}
	}
}

func TestSliceType(t *testing.T) {
	ty := sliceType([]interface{}{})
	exp := "[]interface{}"
	if ty != exp {
		t.Fatalf("expected %s; got %s", exp, ty)
	}

	ty = sliceType([]interface{}{"a", "b"})
	exp = "[]string"
	if ty != exp {
		t.Fatalf("expected %s; got %s", exp, ty)
	}

	ty = sliceType([]interface{}{float64(1), float64(2)})
	exp = "[]int"
	if ty != exp {
		t.Fatalf("expected %s; got %s", exp, ty)
	}

	ty = sliceType([]interface{}{"a", 1})
	exp = "[]interface{}"
	if ty != exp {
		t.Fatalf("expected %s; got %s", exp, ty)
	}
}

func readWant(p string) ([]byte, error) {
	fd, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, fd)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
