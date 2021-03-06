package iowrap

import (
	"io"
	"io/ioutil"
	"testing"

	"github.com/minhjh/go-storage/v4/pkg/randbytes"
	"github.com/stretchr/testify/assert"
)

func TestPipe(t *testing.T) {
	cases := []struct {
		name string
		size int
	}{
		{"1B", 1},
		{"4k", 4 * 1024},
		{"16m", 16 * 1024 * 1024},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			expected := make([]byte, v.size)
			_, err := randbytes.NewRand().Read(expected)
			if err != nil {
				t.Error(err)
			}

			r, w := Pipe()
			io.Pipe()

			go func() {
				defer w.Close()

				_, _ = w.Write(expected)
			}()

			actual, err := ioutil.ReadAll(r)
			if err != nil {
				t.Error(err)
			}
			assert.EqualValues(t, expected, actual)
		})
	}
}
