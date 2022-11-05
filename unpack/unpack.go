package unpack

import (
	"encoding/binary"
	"errors"
	"io"
)

const MsgHeader = "12345678"

// Encode
func Encode(bytesBuffer io.Writer, content string) error {
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(MsgHeader)); err != nil {
		return err
	}
	clen := int32(len([]byte(content)))
	if err := binary.Write(bytesBuffer, binary.BigEndian, clen); err != nil {
		return err
	}

	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(content)); err != nil {
		return err
	}
	return nil
}

// Decode
func Decode(bytesBuffer io.Reader) (bodyBuf []byte, err error) {
	headBuf := make([]byte, len(MsgHeader))
	if _, err := io.ReadFull(bytesBuffer, headBuf); err != nil {
		return nil, err
	}

	if string(headBuf) != MsgHeader {
		return nil, errors.New("data buff head invalid")
	}

	lbuf := make([]byte, 4)
	if _, err := io.ReadFull(bytesBuffer, lbuf); err != nil {
		return nil, err
	}

	l := binary.BigEndian.Uint32(lbuf)
	body := make([]byte, l)
	if _, err := io.ReadFull(bytesBuffer, body); err != nil {
		return nil, err
	}
	return body, err
}
