package tarantool

import (
	"io/ioutil"
	"testing"
)

func TestBodyDecode(t *testing.T) {
	data, err := ioutil.ReadFile("_testdata/badbody.bin")
	if err != nil {
		t.Fatal(err)
	}
	r := &Response{
		buf:        smallBuf{b: data},
		newDecoder: newDecoder,
	}
	if err = r.decodeBody(); err != nil {
		t.Error(err)
	}
}
