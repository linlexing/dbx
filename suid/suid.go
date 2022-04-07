package suid

import (
	"encoding/base32"
	"encoding/binary"

	"github.com/sony/sonyflake"
)

var sf = sonyflake.NewSonyflake(sonyflake.Settings{})

func Next() (string, error) {
	id, err := sf.NextID()
	if err != nil {
		return "", err
	}
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, id)
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(b), nil
}
