package delete

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

type TestStruct struct {
	Path   string
	Result bool
}

func TestMatch(t *testing.T) {
	r := matchFile(createName(time.Now()))
	if r {
		t.Fatalf("error")
	}
	r = matchFile(createName(time.Now().AddDate(0, -3, 0)))
	if !r {
		t.Fatalf("error")
	}
}
func createName(date time.Time) string {
	return fmt.Sprintf("おぎやはぎのメガネびいき_%04d-%02d-%02d-%02d_%02d.mp3",
		date.Year(),
		date.Month(),
		date.Day(),
		date.Hour(),
		date.Minute(),
	)
}

func TestExampleSuccess(t *testing.T) {
	searchDir := "hoge"
	dummyList := []TestStruct{
		{createName(time.Now().AddDate(0, -1, -10)), false},
		{createName(time.Now().AddDate(0, 0, -5)), true},
		{createName(time.Now().AddDate(0, -2, -5)), false},
		{"asfa8asf.mp3", true},
	}
	os.RemoveAll(searchDir)
	err := Delete(searchDir)
	if err == nil {
		t.Fatalf("failed test %#v", err)
	}
	os.Mkdir(searchDir, 0777)

	for _, v := range dummyList {
		p := filepath.Join(searchDir, v.Path)
		if err = createFile(p); err != nil {
			t.Fatalf("%#v", err)
		}
	}
	err = Delete(searchDir)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	for _, v := range dummyList {
		p := filepath.Join(searchDir, v.Path)
		if Exists(p) != v.Result {
			t.Fatalf("%#v", p)
		}
	}
}

func createFile(path string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()
	return err
}
func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
