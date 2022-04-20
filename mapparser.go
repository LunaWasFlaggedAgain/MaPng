package main

import (
	"io"

	"github.com/Tnze/go-mc/nbt"
)

type Map struct {
	Colors []byte `nbt:"colors"`
}

func ParseMap(r io.Reader) (Map, error) {
	var m struct {
		Data Map `nbt:"data"`
	}

	_, err := nbt.NewDecoder(r).Decode(&m)
	return m.Data, err
}
