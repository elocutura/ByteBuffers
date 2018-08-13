package Server

import (
	"encoding/binary"
	"math"
	)

type ByteBuffer struct {
	Buffer 		 []byte
	bufferLength int
	Position	 int
}

func (bb *ByteBuffer) InitBuffer (bufferLength int) {
	bb.Buffer = make([]byte, bufferLength)
	bb.bufferLength = bufferLength
	bb.Position = 0
}

func (bb *ByteBuffer) Flush () {
	bb.Buffer = make([]byte, bb.bufferLength)
	bb.Position = 0
}

func (bb *ByteBuffer) ReadByte () byte {
	value := bb.Buffer[bb.Position]
	bb.Position ++

	return  value
}

func (bb *ByteBuffer) ReadBytes (index int, numOfBytes int) []byte {
	var value []byte
	value = make([]byte, numOfBytes)
	for i := 0; i < numOfBytes; i++ {
		value[i] = bb.Buffer[index + i]
	}
	bb.Position += numOfBytes
	return value
}
func (bb *ByteBuffer) WriteBytes (value []byte) {
	for i := 0; i < len(value); i++ {
		bb.Buffer[bb.Position + i] = value[i]
	}
	bb.Position += len(value)
}

func (bb *ByteBuffer) ReadUInt16 () uint16 {
	auxByteArray := bb.ReadBytes(bb.Position, 2)
	value := binary.LittleEndian.Uint16(auxByteArray)
	return value
}

func (bb *ByteBuffer) ReadInt16 () int16 {
	auxByteArray := bb.ReadBytes(bb.Position, 4)
	value := binary.LittleEndian.Uint32(auxByteArray)
	return int16(value)
}

func (bb *ByteBuffer) ReadUInt32 () uint32 {
	auxByteArray := bb.ReadBytes(bb.Position, 4)
	value := binary.LittleEndian.Uint32(auxByteArray)
	return value
}

func (bb *ByteBuffer) ReadInt32 () int32 {
	auxByteArray := bb.ReadBytes(bb.Position, 4)
	value := binary.LittleEndian.Uint32(auxByteArray)
	return int32(value)
}

func (bb *ByteBuffer) ReadFloat () float32 {
	auxByteArray := bb.ReadBytes(bb.Position, 4)
	value := binary.LittleEndian.Uint32(auxByteArray)
	return math.Float32frombits(value)
}

func (bb *ByteBuffer) ReadUInt64 () uint64 {
	auxByteArray := bb.ReadBytes(bb.Position, 8)
	value := binary.LittleEndian.Uint64(auxByteArray)
	return value
}

func (bb *ByteBuffer) ReadInt64 () int64 {
	auxByteArray := bb.ReadBytes(bb.Position, 8)
	value := binary.LittleEndian.Uint32(auxByteArray)
	return int64(value)
}

func (bb *ByteBuffer) ReadBoolean () bool {
	received := bb.ReadByte()
	if received == 01 {
		return true
	} else {
		return false
	}
}

func (bb *ByteBuffer) ReadString () string {

	auxByteArray := bb.ReadBytes(bb.Position, 2)
	stringLength := binary.LittleEndian.Uint16(auxByteArray)

	stringByteArray := bb.ReadBytes(bb.Position,int(stringLength))

	return string(stringByteArray[:stringLength])
}

func (bb *ByteBuffer) WriteByte (value byte) {
	bb.Buffer[bb.Position] = value
	bb.Position ++
}

func (bb *ByteBuffer) WriteUInt16 (value uint16) {
	toWrite := make([]byte, 2)
	binary.LittleEndian.PutUint16(toWrite, value)
	bb.WriteBytes(toWrite)
}

func (bb *ByteBuffer) WriteInt16 (value int16) {
	toWrite := make([]byte, 2)
	binary.LittleEndian.PutUint16(toWrite, uint16(value))
	bb.WriteBytes(toWrite)
}

func (bb *ByteBuffer) WriteUInt32 (value uint32) {
	toWrite := make([]byte, 4)
	binary.LittleEndian.PutUint32(toWrite, value)
	bb.WriteBytes(toWrite)
}

func (bb *ByteBuffer) WriteInt32 (value int32) {
	toWrite := make([]byte, 4)
	binary.LittleEndian.PutUint32(toWrite, uint32(value))
	bb.WriteBytes(toWrite)
}

func (bb *ByteBuffer) WriteFloat (value float32) {
	toWrite := make([]byte, 4)
	binary.LittleEndian.PutUint32(toWrite, math.Float32bits(value))
	bb.WriteBytes(toWrite)
}

func (bb *ByteBuffer) WriteUInt64 (value uint64) {
	toWrite := make([]byte, 8)
	binary.LittleEndian.PutUint64(toWrite, value)
	bb.WriteBytes(toWrite)
}

func (bb *ByteBuffer) WriteInt64 (value int64) {
	toWrite := make([]byte, 8)
	binary.LittleEndian.PutUint64(toWrite, uint64(value))
	bb.WriteBytes(toWrite)
}


func (bb *ByteBuffer) WriteBoolean (value bool) {
	toWrite := make([]byte, 1)
	if value{
		toWrite[0] = 01
	} else {
		toWrite[0]	= 00
	}
	bb.WriteByte(toWrite[0])
}

func (bb *ByteBuffer) WriteString (value string) {
	toWrite := make([]byte, 2 + len(value))
	binary.LittleEndian.PutUint16(toWrite, uint16(len(value)))
	byteString := []byte(value)

	for i := 0; i < len(value); i++ {
		toWrite[i+2] = byteString[i]
	}

	bb.WriteBytes(toWrite)
}