package marisa

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	marisa_go_test_error_helper(1)
	os.Exit(m.Run())
}

func TestTrieIO(t *testing.T) {
	emptyBuf := bytes.NewBuffer(nil)
	emptyS := "1aa6c451104c2c1b24ecb66ecb84bde2403c49b1" // marisa-build </dev/null | sha1sum -

	normalWd := []string{"asd", "bnm", "cvb", "dfg"} // for n in asd bnm cvb dfg; do echo $n; done | marisa-build | sha1sum -
	normalBuf := bytes.NewBuffer(nil)
	normalS := "bdf9be48216379734fa0256263467ba6ab2e0931"

	t.Run("WriteAll", func(t *testing.T) {
		t.Run("Error", func(t *testing.T) {
			err := WriteAll(bytes.NewBuffer(nil), []string{""})
			if v := "c++ runtime error: go_test_error"; err == nil || err.Error() != v {
				t.Errorf("expected err to be `%v`, got `%v`", v, err)
			}
		})
		t.Run("WriteError", func(t *testing.T) {
			err := WriteAll(new(errIO), normalWd)
			if v := "go_test_error"; err == nil || err.Error() != v {
				t.Errorf("expected err to be `%v`, got `%v`", v, err)
			}
		})
		t.Run("Empty", func(t *testing.T) {
			ss := sha1.New()
			if err := WriteAll(io.MultiWriter(emptyBuf, ss), nil); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if v := hex.EncodeToString(ss.Sum(nil)); v != emptyS {
				t.Errorf("output sha1 mismatch: expected %s, got %s", emptyS, v)
			}
		})
		t.Run("Normal", func(t *testing.T) {
			ss := sha1.New()
			if err := WriteAll(io.MultiWriter(normalBuf, ss), normalWd); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if v := hex.EncodeToString(ss.Sum(nil)); v != normalS {
				t.Errorf("output sha1 mismatch: expected %s, got %s", normalS, v)
			}
		})
	})
	t.Run("ReadAll", func(t *testing.T) {
		t.Run("Error", func(t *testing.T) {
			wd, err := ReadAll(bytes.NewReader([]byte{0}))
			if v := "c++ runtime error: go_test_error"; err == nil || err.Error() != v {
				t.Errorf("expected err to be `%v`, got `%v`", v, err)
			}
			if wd != nil {
				t.Errorf("expected returned slice to be nil, got %#v", wd)
			}
		})
		t.Run("ReadError", func(t *testing.T) {
			wd, err := ReadAll(new(errIO))
			if v := "go_test_error"; err == nil || err.Error() != v {
				t.Errorf("expected err to be `%v`, got `%v`", v, err)
			}
			if wd != nil {
				t.Errorf("expected returned slice to be nil, got %#v", wd)
			}
		})
		t.Run("Empty", func(t *testing.T) {
			wd, err := ReadAll(emptyBuf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if len(wd) != 0 {
				t.Errorf("expected no words to be returned")
			}
		})
		t.Run("Normal", func(t *testing.T) {
			wd, err := ReadAll(normalBuf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(wd, normalWd) {
				t.Errorf("expected %#v, got %#v", normalWd, wd)
			}
		})
	})
}

type errIO struct{}

func (*errIO) Write([]byte) (int, error) { return 0, errors.New("go_test_error") }
func (*errIO) Read([]byte) (int, error) { return 0, errors.New("go_test_error") }
