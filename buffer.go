package go_binutils

import (
	"fmt"
	"io"
)

// Buffer Buffer类，封装[]byte，用于读取或者写入基础数据
// 方法分为两大类：Append***BE/LE 以及 Read***BE/LE
// Append会自动扩展[]byte，如果大小不足的时候
// Read如果缓冲区不足的时候，会返回error
// BE表示大端，LE表示小端
type Buffer struct {
	buf   []byte
	index int
}

// NewBuffer 新建Buffer对象
func NewBuffer(buf []byte) *Buffer {
	return &Buffer{
		buf:   buf,
		index: 0,
	}
}

func (buffer *Buffer) Init(buf []byte) {
	buffer.buf = buf
	buffer.index = 0
}

// Reset 复位
func (buffer *Buffer) Reset() {
	buffer.buf = buffer.buf[0:0]
	buffer.index = 0
}

// SetBuf 重新设置
func (buffer *Buffer) SetBuf(s []byte) {
	buffer.buf = s
	buffer.index = 0
}

// Bytes 获取内部的buffer
func (buffer *Buffer) Bytes() []byte {
	return buffer.buf
}

//------------- 8bit

// AppendByte 添加byte
func (buffer *Buffer) AppendByte(b byte) error {
	buffer.buf = append(buffer.buf, b)
	return nil
}

// ReadByte 读取Byte
func (buffer *Buffer) ReadByte() (v byte, err error) {
	i := buffer.index + 1
	if i < 0 || i > len(buffer.buf) {
		err = io.ErrUnexpectedEOF
		return
	}

	v = buffer.buf[buffer.index]
	buffer.index = i

	return
}

// AppendInt8 添加Int8
func (buffer *Buffer) AppendInt8(v int8) error {
	buffer.buf = append(buffer.buf, uint8(v))
	return nil
}

// ReadInt8 读取Int8
func (buffer *Buffer) ReadInt8() (v int8, err error) {
	i := buffer.index + 1
	if i < 0 || i > len(buffer.buf) {
		err = io.ErrUnexpectedEOF
		return
	}

	v = int8(buffer.buf[buffer.index])
	buffer.index = i

	return
}

// AppendUint8 添加uint8
func (buffer *Buffer) AppendUint8(v uint8) error {
	buffer.buf = append(buffer.buf, v)
	return nil
}

// ReadUint8 读取uint8
func (buffer *Buffer) ReadUint8() (v uint8, err error) {
	return buffer.ReadByte()
}

//------------- 16bit

// AppendUint16BE 大端方式添加uint16
func (buffer *Buffer) AppendUint16BE(v uint16) error {
	buffer.buf = append(buffer.buf, uint8(v>>8), uint8(v))
	return nil
}

// ReadUint16BE 大端方式读取uint16
func (buffer *Buffer) ReadUint16BE() (v uint16, err error) {
	i := buffer.index + 2
	if i < 0 || i > len(buffer.buf) {
		err = io.ErrUnexpectedEOF
		return
	}
	buffer.index = i

	v = uint16(buffer.buf[i-1])
	v |= uint16(buffer.buf[i-2]) << 8

	return
}

// AppendInt16BE 大端方式添加int16
func (buffer *Buffer) AppendInt16BE(v int16) error {
	buffer.buf = append(buffer.buf, uint8(v>>8), uint8(v))
	return nil
}

// ReadInt16BE 大端方式读取int16
func (buffer *Buffer) ReadInt16BE() (v int16, err error) {
	i := buffer.index + 2
	if i < 0 || i > len(buffer.buf) {
		err = io.ErrUnexpectedEOF
		return
	}
	buffer.index = i

	v = int16(buffer.buf[i-1])
	v |= int16(buffer.buf[i-2]) << 8

	return
}

// AppendUint16LE 小端方式添加uint16
func (buffer *Buffer) AppendUint16LE(v uint16) error {
	buffer.buf = append(buffer.buf, uint8(v), uint8(v>>8))
	return nil
}

// ReadUint16LE 小端方式读取uint16
func (buffer *Buffer) ReadUint16LE() (v uint16, err error) {
	i := buffer.index + 2
	if i < 0 || i > len(buffer.buf) {
		err = io.ErrUnexpectedEOF
		return
	}
	buffer.index = i

	v = uint16(buffer.buf[i-2])
	v |= uint16(buffer.buf[i-1]) << 8

	return
}

// AppendInt16LE 小端方式添加int16
func (buffer *Buffer) AppendInt16LE(v int16) error {
	buffer.buf = append(buffer.buf, uint8(v>>8), uint8(v))
	return nil
}

// ReadInt16LE 小端方式读取int16
func (buffer *Buffer) ReadInt16LE() (v int16, err error) {
	i := buffer.index + 2
	if i < 0 || i > len(buffer.buf) {
		err = io.ErrUnexpectedEOF
		return
	}
	buffer.index = i

	v = int16(buffer.buf[i-2])
	v |= int16(buffer.buf[i-1]) << 8

	return
}

//------------- 32bit

// AppendUint32BE 大端方式添加uint32
func (buffer *Buffer) AppendUint32BE(v uint32) error {
	buffer.buf = append(buffer.buf,
		uint8(v>>24),
		uint8(v>>16),
		uint8(v>>8),
		uint8(v))
	return nil
}

// ReadUint32BE 大端方式读取uint32
func (buffer *Buffer) ReadUint32BE() (v uint32, err error) {
	i := buffer.index + 4
	if i < 0 || i > len(buffer.buf) {
		err = io.ErrUnexpectedEOF
		return
	}
	buffer.index = i

	v = uint32(buffer.buf[i-1])
	v |= uint32(buffer.buf[i-2]) << 8
	v |= uint32(buffer.buf[i-3]) << 16
	v |= uint32(buffer.buf[i-4]) << 24

	return
}

// AppendInt32BE 大端方式添加int32
func (buffer *Buffer) AppendInt32BE(v int32) error {
	buffer.buf = append(buffer.buf,
		uint8(v>>24),
		uint8(v>>16),
		uint8(v>>8),
		uint8(v))
	return nil
}

// ReadInt32BE 大端方式读取int32
func (buffer *Buffer) ReadInt32BE() (v int32, err error) {
	i := buffer.index + 4
	if i < 0 || i > len(buffer.buf) {
		err = io.ErrUnexpectedEOF
		return
	}
	buffer.index = i

	v = int32(buffer.buf[i-1])
	v |= int32(buffer.buf[i-2]) << 8
	v |= int32(buffer.buf[i-3]) << 16
	v |= int32(buffer.buf[i-4]) << 24

	return
}

// AppendUint32LE 小端方式添加uint32
func (buffer *Buffer) AppendUint32LE(v uint32) error {
	buffer.buf = append(buffer.buf,
		uint8(v),
		uint8(v>>8),
		uint8(v>>16),
		uint8(v>>24))
	return nil
}

// ReadUint32LE 小端方式添加uint32
func (buffer *Buffer) ReadUint32LE() (v uint32, err error) {
	i := buffer.index + 4
	if i < 0 || i > len(buffer.buf) {
		err = io.ErrUnexpectedEOF
		return
	}
	buffer.index = i

	v = uint32(buffer.buf[i-4]) | uint32(buffer.buf[i-3])<<8 | uint32(buffer.buf[i-2])<<16 | uint32(buffer.buf[i-1])<<24

	return
}

// AppendInt32LE 小端方式添加int32
func (buffer *Buffer) AppendInt32LE(v int32) error {
	buffer.buf = append(buffer.buf,
		uint8(v>>24),
		uint8(v>>16),
		uint8(v>>8),
		uint8(v))
	return nil
}

// ReadInt32LE 小端方式添加int32
func (buffer *Buffer) ReadInt32LE() (v int32, err error) {
	i := buffer.index + 4
	if i < 0 || i > len(buffer.buf) {
		err = io.ErrUnexpectedEOF
		return
	}
	buffer.index = i

	v = int32(buffer.buf[i-4])
	v |= int32(buffer.buf[i-3]) << 8
	v |= int32(buffer.buf[i-2]) << 16
	v |= int32(buffer.buf[i-1]) << 24

	return
}

//------------- 64bit

// AppendUint64BE 以大端方式添加uint64
func (buffer *Buffer) AppendUint64BE(v uint64) error {
	buffer.buf = append(buffer.buf,
		uint8(v>>56),
		uint8(v>>48),
		uint8(v>>40),
		uint8(v>>32),
		uint8(v>>24),
		uint8(v>>16),
		uint8(v>>8),
		uint8(v))
	return nil
}

// ReadUint64BE 大端方式读取uint64
func (buffer *Buffer) ReadUint64BE() (v uint64, err error) {
	i := buffer.index + 8
	if i < 0 || i > len(buffer.buf) {
		err = io.ErrUnexpectedEOF
		return
	}
	buffer.index = i

	v = uint64(buffer.buf[i-1])
	v |= uint64(buffer.buf[i-2]) << 8
	v |= uint64(buffer.buf[i-3]) << 16
	v |= uint64(buffer.buf[i-4]) << 24
	v |= uint64(buffer.buf[i-5]) << 32
	v |= uint64(buffer.buf[i-6]) << 40
	v |= uint64(buffer.buf[i-7]) << 48
	v |= uint64(buffer.buf[i-8]) << 56
	return
}

// AppendInt64BE 大端方式添加int64
func (buffer *Buffer) AppendInt64BE(v int64) error {
	buffer.buf = append(buffer.buf,
		uint8(v>>56),
		uint8(v>>48),
		uint8(v>>40),
		uint8(v>>32),
		uint8(v>>24),
		uint8(v>>16),
		uint8(v>>8),
		uint8(v))
	return nil
}

// ReadInt64BE 大端方式读取Int64
func (buffer *Buffer) ReadInt64BE() (v int64, err error) {
	i := buffer.index + 8
	if i < 0 || i > len(buffer.buf) {
		err = io.ErrUnexpectedEOF
		return
	}
	buffer.index = i

	v = int64(buffer.buf[i-1])
	v |= int64(buffer.buf[i-2]) << 8
	v |= int64(buffer.buf[i-3]) << 16
	v |= int64(buffer.buf[i-4]) << 24
	v |= int64(buffer.buf[i-5]) << 32
	v |= int64(buffer.buf[i-6]) << 40
	v |= int64(buffer.buf[i-7]) << 48
	v |= int64(buffer.buf[i-8]) << 56
	return
}

// AppendUint64LE 小端方式添加uint64
func (buffer *Buffer) AppendUint64LE(v uint64) error {
	buffer.buf = append(buffer.buf,
		uint8(v),
		uint8(v>>8),
		uint8(v>>16),
		uint8(v>>24),
		uint8(v>>32),
		uint8(v>>40),
		uint8(v>>48),
		uint8(v>>56))
	return nil
}

// ReadUint64LE 小端方式读取uint64
func (buffer *Buffer) ReadUint64LE() (v uint64, err error) {
	i := buffer.index + 8
	if i < 0 || i > len(buffer.buf) {
		err = io.ErrUnexpectedEOF
		return
	}
	buffer.index = i

	v = uint64(buffer.buf[i-8])
	v |= uint64(buffer.buf[i-7]) << 8
	v |= uint64(buffer.buf[i-6]) << 16
	v |= uint64(buffer.buf[i-5]) << 24
	v |= uint64(buffer.buf[i-4]) << 32
	v |= uint64(buffer.buf[i-3]) << 40
	v |= uint64(buffer.buf[i-2]) << 48
	v |= uint64(buffer.buf[i-1]) << 56
	return
}

// AppendInt64LE 小端方式添加int64
func (buffer *Buffer) AppendInt64LE(v int64) error {
	buffer.buf = append(buffer.buf,
		uint8(v>>56),
		uint8(v>>48),
		uint8(v>>40),
		uint8(v>>32),
		uint8(v>>24),
		uint8(v>>16),
		uint8(v>>8),
		uint8(v))
	return nil
}

// ReadInt64LE 小端方式读取int64
func (buffer *Buffer) ReadInt64LE() (v int64, err error) {
	i := buffer.index + 8
	if i < 0 || i > len(buffer.buf) {
		err = io.ErrUnexpectedEOF
		return
	}
	buffer.index = i

	v = int64(buffer.buf[i-8])
	v |= int64(buffer.buf[i-7]) << 8
	v |= int64(buffer.buf[i-6]) << 16
	v |= int64(buffer.buf[i-5]) << 24
	v |= int64(buffer.buf[i-4]) << 32
	v |= int64(buffer.buf[i-3]) << 40
	v |= int64(buffer.buf[i-2]) << 48
	v |= int64(buffer.buf[i-1]) << 56
	return
}

//------------- bytes

// AppendBytes 添加[]byte
func (buffer *Buffer) AppendBytes(v []byte) error {
	buffer.buf = append(buffer.buf, v...)
	return nil
}

// ReadBytes 读取指定的字节数
func (buffer *Buffer) ReadBytes(n int, alloc bool) (buf []byte, err error) {
	if n <= 0 {
		return nil, fmt.Errorf("buffer: bad byte length %d", n)
	}
	end := buffer.index + n
	if end < buffer.index || end > len(buffer.buf) {
		return nil, io.ErrUnexpectedEOF
	}

	if !alloc {
		// TODO: check if can get more uses of alloc=false
		buf = buffer.buf[buffer.index:end]
		buffer.index += n
		return
	}

	buf = make([]byte, n)
	copy(buf, buffer.buf[buffer.index:])
	buffer.index += n
	return
}
