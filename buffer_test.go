package binary

import (
	"github.com/funny/utest"
	"math/rand"
	"testing"
)

func RandBytes(n int) []byte {
	n = rand.Intn(n) + 1
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = byte(rand.Intn(255))
	}
	return b
}

func Test_Buffer_ReadWrite(t *testing.T) {
	for i := 0; i < 10000; i++ {
		b := RandBytes(512)
		buf := Buffer{Data: make([]byte, len(b))}

		n, err := buf.Write(b)
		utest.IsNilNow(t, err)
		utest.EqualNow(t, n, len(b))

		c := make([]byte, len(b))
		n, err = buf.Read(c)
		utest.IsNilNow(t, err)
		utest.EqualNow(t, n, len(b))
		utest.EqualNow(t, b, c)
	}
}

func Test_Buffer_Uint8(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint8(rand.Intn(256))
	buf.WriteUint8(v1)

	v2 := buf.ReadUint8()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint16BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint16(rand.Intn(0xFFFF))
	buf.WriteUint16BE(v1)

	v2 := buf.ReadUint16BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint16LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint16(rand.Intn(0xFFFF))
	buf.WriteUint16LE(v1)

	v2 := buf.ReadUint16LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint24BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint32(rand.Intn(0xFFFFFF))
	buf.WriteUint24BE(v1)

	v2 := buf.ReadUint24BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint24LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint32(rand.Intn(0xFFFFFF))
	buf.WriteUint24LE(v1)

	v2 := buf.ReadUint24LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint32BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint32(rand.Intn(0xFFFFFFFF))
	buf.WriteUint32BE(v1)

	v2 := buf.ReadUint32BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint32LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint32(rand.Intn(0xFFFFFFFF))
	buf.WriteUint32LE(v1)

	v2 := buf.ReadUint32LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint40BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint64(rand.Intn(0xFFFFFFFFFF))
	buf.WriteUint40BE(v1)

	v2 := buf.ReadUint40BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint40LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint64(rand.Intn(0xFFFFFFFFFF))
	buf.WriteUint40LE(v1)

	v2 := buf.ReadUint40LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint48BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint64(rand.Intn(0xFFFFFFFFFFFF))
	buf.WriteUint48BE(v1)

	v2 := buf.ReadUint48BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint48LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint64(rand.Intn(0xFFFFFFFFFFFF))
	buf.WriteUint48LE(v1)

	v2 := buf.ReadUint48LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint56BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint64(rand.Intn(0xFFFFFFFFFFFFFF))
	buf.WriteUint56BE(v1)

	v2 := buf.ReadUint56BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint56LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint64(rand.Intn(0xFFFFFFFFFFFFFF))
	buf.WriteUint56LE(v1)

	v2 := buf.ReadUint56LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint64BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
	buf.WriteUint64BE(v1)

	v2 := buf.ReadUint64BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uint64LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
	buf.WriteUint64LE(v1)

	v2 := buf.ReadUint64LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Uvarint(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := uint64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
	buf.WriteUvarint(v1)

	v2 := buf.ReadUvarint()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Varint(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
	buf.WriteVarint(v1)

	v2 := buf.ReadVarint()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Float32BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := float32(rand.NormFloat64())
	buf.WriteFloat32BE(v1)

	v2 := buf.ReadFloat32BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Float32LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := float32(rand.NormFloat64())
	buf.WriteFloat32LE(v1)

	v2 := buf.ReadFloat32LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Float64BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := rand.NormFloat64()
	buf.WriteFloat64BE(v1)

	v2 := buf.ReadFloat64BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Float64LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := rand.NormFloat64()
	buf.WriteFloat64LE(v1)

	v2 := buf.ReadFloat64LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int8(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int8(rand.Intn(256))
	buf.WriteInt8(v1)

	v2 := buf.ReadInt8()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int16BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int16(rand.Intn(0xFFFF))
	buf.WriteInt16BE(v1)

	v2 := buf.ReadInt16BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int16LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int16(rand.Intn(0xFFFF))
	buf.WriteInt16LE(v1)

	v2 := buf.ReadInt16LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int24BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int32(rand.Intn(0xFFFFFF))
	buf.WriteInt24BE(v1)

	v2 := buf.ReadInt24BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int24LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int32(rand.Intn(0xFFFFFF))
	buf.WriteInt24LE(v1)

	v2 := buf.ReadInt24LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int32BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int32(rand.Intn(0xFFFFFFFF))
	buf.WriteInt32BE(v1)

	v2 := buf.ReadInt32BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int32LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int32(rand.Intn(0xFFFFFFFF))
	buf.WriteInt32LE(v1)

	v2 := buf.ReadInt32LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int40BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int64(rand.Intn(0xFFFFFFFFFF))
	buf.WriteInt40BE(v1)

	v2 := buf.ReadInt40BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int40LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int64(rand.Intn(0xFFFFFFFFFF))
	buf.WriteInt40LE(v1)

	v2 := buf.ReadInt40LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int48BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int64(rand.Intn(0xFFFFFFFFFFFF))
	buf.WriteInt48BE(v1)

	v2 := buf.ReadInt48BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int48LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int64(rand.Intn(0xFFFFFFFFFFFF))
	buf.WriteInt48LE(v1)

	v2 := buf.ReadInt48LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int56BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int64(rand.Intn(0xFFFFFFFFFFFFFF))
	buf.WriteInt56BE(v1)

	v2 := buf.ReadInt56BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int56LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int64(rand.Intn(0xFFFFFFFFFFFFFF))
	buf.WriteInt56LE(v1)

	v2 := buf.ReadInt56LE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int64BE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
	buf.WriteInt64BE(v1)

	v2 := buf.ReadInt64BE()
	utest.EqualNow(t, v1, v2)
}

func Test_Buffer_Int64LE(t *testing.T) {
	var buf = Buffer{Data: make([]byte, 10)}
	v1 := int64(rand.Int63n(0x7FFFFFFFFFFFFFFF))
	buf.WriteInt64LE(v1)

	v2 := buf.ReadInt64LE()
	utest.EqualNow(t, v1, v2)
}
