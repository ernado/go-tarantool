package tarantool

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/vmihailenco/msgpack"
)

func TestOpEncode(t *testing.T) {
	buf := new(bytes.Buffer)
	e := msgpack.NewEncoder(buf)
	e.UseCompactEncoding(true)
	op := []interface{}{[]interface{}{"=", uint32(1), "bye"}, []interface{}{"#", uint32(2), uint32(1)}}
	if err := e.Encode(op); err != nil {
		t.Fatal(err)
	}
	data := buf.Bytes()
	d := msgpack.NewDecoder(buf)
	decoded := make([]Op, 0, 2)
	if err := d.Decode(&decoded); err != nil {
		t.Fatal(err)
	}
	if len(decoded) != 2 {
		t.Error("bad length")
	}
	newBuf := new(bytes.Buffer)
	e = msgpack.NewEncoder(newBuf)
	e.UseCompactEncoding(true)
	if err := e.Encode(decoded); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(newBuf.Bytes(), data) {
		t.Error("not equal")
		t.Log(hex.Dump(newBuf.Bytes()))
		t.Log(hex.Dump(data))
	}
}
