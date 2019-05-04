package tarantool

import (
	"io"

	"github.com/vmihailenco/msgpack"
)

const mapElemsAllocLimit = 1e4

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func decodeStringMap(d *msgpack.Decoder) (interface{}, error) {
	n, err := d.DecodeMapLen()
	if err != nil {
		return nil, err
	}
	if n == -1 {
		return nil, nil
	}

	m := make(map[string]interface{}, min(n, mapElemsAllocLimit))
	for i := 0; i < n; i++ {
		mk, err := d.DecodeString()
		if err != nil {
			return nil, err
		}
		mv, err := d.DecodeInterface()
		if err != nil {
			return nil, err
		}
		m[mk] = mv
	}
	return m, nil
}

func decodeMap(d *msgpack.Decoder) (interface{}, error) {
	n, err := d.DecodeMapLen()
	if err != nil {
		return nil, err
	}
	if n == -1 {
		return nil, nil
	}

	m := make(map[interface{}]interface{}, min(n, mapElemsAllocLimit))
	for i := 0; i < n; i++ {
		mk, err := d.DecodeInterface()
		if err != nil {
			return nil, err
		}
		mv, err := d.DecodeInterface()
		if err != nil {
			return nil, err
		}
		m[mk] = mv
	}
	return m, nil
}

func newEncoder(w io.Writer) *msgpack.Encoder {
	e := msgpack.NewEncoder(w)
	// Compat encoding is necessary e.g. for correct unpacking of update
	// operations on tarantool side.
	//
	// Removing this will break things in unbeknownst places.
	e.UseCompactEncoding(true)
	// Using JSON tags by default, because it is common to duplicate them by
	// msgpack tags.
	e.UseJSONTag(true)
	return e
}

func newDecoder(r io.Reader) *msgpack.Decoder {
	d := msgpack.NewDecoder(r)
	d.UseJSONTag(true)
	return d
}
