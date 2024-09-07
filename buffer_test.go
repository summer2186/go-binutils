package go_binutils

import (
	"encoding/binary"
	"testing"
)

func TestBuffer_ReadUint16(t *testing.T) {
	data := make([]byte, 2)
	v := uint16(0x1234)
	binary.BigEndian.PutUint16(data, v)
	buffer1 := NewBuffer(data)
	v1, _ := buffer1.ReadUint16BE()
	if v != v1 {
		t.Fatalf("return is: %d", v1)
	}

	binary.LittleEndian.PutUint16(data, v)
	buffer2 := NewBuffer(data)
	v2, _ := buffer2.ReadUint16LE()
	if v != v2 {
		t.Fatalf("return is: %d", v2)
	}
}

func TestBuffer_ReadInt16(t *testing.T) {
	data := make([]byte, 2)
	v := int16(0x1234)
	binary.BigEndian.PutUint16(data, uint16(v))
	buffer1 := NewBuffer(data)
	v1, _ := buffer1.ReadInt16BE()
	if v != v1 {
		t.Fatalf("return is: %d", v1)
	}

	binary.LittleEndian.PutUint16(data, uint16(v))
	buffer2 := NewBuffer(data)
	v2, _ := buffer2.ReadInt16LE()
	if v != v2 {
		t.Fatalf("return is: %d", v2)
	}
}

func TestBuffer_ReadUint32(t *testing.T) {
	data := make([]byte, 4)
	v := uint32(0x12345678)
	binary.BigEndian.PutUint32(data, uint32(v))
	buffer1 := NewBuffer(data)
	v1, _ := buffer1.ReadUint32BE()
	if v != v1 {
		t.Fatalf("return is: %d", v1)
	}

	binary.LittleEndian.PutUint32(data, uint32(v))
	buffer2 := NewBuffer(data)
	v2, _ := buffer2.ReadUint32LE()
	if v != v2 {
		t.Fatalf("return is: %d", v2)
	}
}

func TestBuffer_ReadInt32(t *testing.T) {
	data := make([]byte, 4)
	v := int32(0x12345678)
	binary.BigEndian.PutUint32(data, uint32(v))
	buffer1 := NewBuffer(data)
	v1, _ := buffer1.ReadInt32BE()
	if v != v1 {
		t.Fatalf("return is: %d", v1)
	}

	binary.LittleEndian.PutUint32(data, uint32(v))
	buffer2 := NewBuffer(data)
	v2, _ := buffer2.ReadInt32LE()
	if v != v2 {
		t.Fatalf("return is: %d", v2)
	}
}

func TestBuffer_ReadInt64(t *testing.T) {
	data := make([]byte, 8)
	v := int64(0x12345678ABCDEF)
	binary.BigEndian.PutUint64(data, uint64(v))
	buffer1 := NewBuffer(data)
	v1, _ := buffer1.ReadInt64BE()
	if v != v1 {
		t.Fatalf("return is: %d", v1)
	}

	binary.LittleEndian.PutUint64(data, uint64(v))
	buffer2 := NewBuffer(data)
	v2, _ := buffer2.ReadInt64LE()
	if v != v2 {
		t.Fatalf("return is: %d", v2)
	}
}

func TestBuffer_ReadUint64(t *testing.T) {
	data := make([]byte, 8)
	v := uint64(0x12345678ABCDEF)
	binary.BigEndian.PutUint64(data, uint64(v))
	buffer1 := NewBuffer(data)
	v1, _ := buffer1.ReadUint64BE()
	if v != v1 {
		t.Fatalf("return is: %d", v1)
	}

	binary.LittleEndian.PutUint64(data, uint64(v))
	buffer2 := NewBuffer(data)
	v2, _ := buffer2.ReadUint64LE()
	if v != v2 {
		t.Fatalf("return is: %d", v2)
	}
}

func TestBuffer_Append(t *testing.T) {
	buffer := NewBuffer([]byte{})

	_ = buffer.AppendInt16BE(0x1234)
	if binary.BigEndian.Uint16(buffer.buf) != 0x1234 {
		t.Fatalf("return is: %d", binary.BigEndian.Uint16(buffer.buf))
	}

	buffer.Reset()
	_ = buffer.AppendUint16LE(0x1234)
	if binary.LittleEndian.Uint16(buffer.buf) != 0x1234 {
		t.Fatalf("return is: %d", binary.LittleEndian.Uint16(buffer.buf))
	}

	buffer.Reset()
	_ = buffer.AppendUint16BE(0x1234)
	if binary.BigEndian.Uint16(buffer.buf) != 0x1234 {
		t.Fatalf("return is: %d", binary.BigEndian.Uint16(buffer.buf))
	}

	buffer.Reset()
	_ = buffer.AppendUint16LE(0x1234)
	if binary.LittleEndian.Uint16(buffer.buf) != 0x1234 {
		t.Fatalf("return is: %d", binary.LittleEndian.Uint16(buffer.buf))
	}

	buffer.Reset()
	_ = buffer.AppendInt32BE(0x12345678)
	if binary.BigEndian.Uint32(buffer.buf) != 0x12345678 {
		t.Fatalf("return is: %d", binary.BigEndian.Uint32(buffer.buf))
	}

	buffer.Reset()
	_ = buffer.AppendUint32LE(0x12345678)
	if binary.LittleEndian.Uint32(buffer.buf) != 0x12345678 {
		t.Fatalf("return is: %d", binary.LittleEndian.Uint32(buffer.buf))
	}

	buffer.Reset()
	_ = buffer.AppendUint32BE(0x12345678)
	if binary.BigEndian.Uint32(buffer.buf) != 0x12345678 {
		t.Fatalf("return is: %d", binary.BigEndian.Uint32(buffer.buf))
	}

	buffer.Reset()
	_ = buffer.AppendUint32LE(0x12345678)
	if binary.LittleEndian.Uint32(buffer.buf) != 0x12345678 {
		t.Fatalf("return is: %d", binary.LittleEndian.Uint32(buffer.buf))
	}

	buffer.Reset()
	_ = buffer.AppendInt64BE(0x12345678ABCDEF)
	if binary.BigEndian.Uint64(buffer.buf) != 0x12345678ABCDEF {
		t.Fatalf("return is: %d", binary.BigEndian.Uint64(buffer.buf))
	}

	buffer.Reset()
	_ = buffer.AppendUint64LE(0x12345678ABCDEF)
	if binary.LittleEndian.Uint64(buffer.buf) != 0x12345678ABCDEF {
		t.Fatalf("return is: %d", binary.LittleEndian.Uint64(buffer.buf))
	}

	buffer.Reset()
	_ = buffer.AppendUint64BE(0x12345678ABCDEF)
	if binary.BigEndian.Uint64(buffer.buf) != 0x12345678ABCDEF {
		t.Fatalf("return is: %d", binary.BigEndian.Uint64(buffer.buf))
	}

	buffer.Reset()
	_ = buffer.AppendUint64LE(0x12345678ABCDEF)
	if binary.LittleEndian.Uint64(buffer.buf) != 0x12345678ABCDEF {
		t.Fatalf("return is: %d", binary.LittleEndian.Uint64(buffer.buf))
	}

}
