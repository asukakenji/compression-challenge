package main

// BitWriter writes bits to the internal buffer in big-endian order.
// There is no gap between the last bit written and the current bit written.
type BitWriter struct {
	ByteIndex int
	BitIndex  uint
	Data      []byte
}

// NewBitWriter returns a new BitWriter.
func NewBitWriter() *BitWriter {
	return &BitWriter{
		Data: make([]byte, 0),
	}
}

// WriteBool writes a single bit to the buffer.
func (w *BitWriter) WriteBool(data bool) *BitWriter {
	if w.BitIndex == 8 {
		w.ByteIndex++
		w.BitIndex = 0
	}
	if w.BitIndex == 0 {
		w.Data = append(w.Data, 0)
	}
	if data {
		w.Data[w.ByteIndex] |= byte(1 << (7 - w.BitIndex))
	}
	w.BitIndex++
	return w
}

// WriteUint64 writes bitCount bits of data to the buffer.
func (w *BitWriter) WriteUint64(data uint64, bitCount uint) *BitWriter {
	for i := uint(0); i < bitCount; i++ {
		w.WriteBool((data>>(bitCount-i-1))&0x1 != 0)
	}
	return w
}

// WriteUint64 writes bitCount bits of data to the buffer.
// func (w *BitWriter) WriteUint64(data uint64, bitCount uint) *BitWriter {
// 	if bitCount > 64 {
// 		panic("bitCount > 64")
// 	}
// 	for bitCount > 0 {
// 		mask := uint64(1<<bitCount - 1)
// 		data &= mask
// 		if w.BitIndex == 8 {
// 			w.ByteIndex++
// 			w.BitIndex = 0
// 		}
// 		if w.BitIndex == 0 {
// 			w.Data = append(w.Data, 0)
// 		}
// 		bitCountAvailableInCurrentByte := 8 - w.BitIndex
// 		var bitCountToBeWritten uint
// 		if bitCount > bitCountAvailableInCurrentByte {
// 			bitCountToBeWritten = bitCountAvailableInCurrentByte
// 			bitsToWrite := byte(data >> (bitCount - bitCountAvailableInCurrentByte))
// 			w.Data[w.ByteIndex] |= bitsToWrite << (8 - w.BitIndex - bitCount)
// 		} else {
// 			bitCountToBeWritten = bitCount
// 			bitsToWrite := byte(data)
// 			w.Data[w.ByteIndex] |= bitsToWrite << (8 - w.BitIndex - bitCount)
// 		}
// 		bitCount -= bitCountToBeWritten
// 		w.BitIndex += bitCountToBeWritten
// 	}
// 	return w
// }

// BitReader reads bits from the internal buffer in big-endian order.
type BitReader struct {
	ByteIndex int
	BitIndex  uint
	Data      []byte
}

// NewBitReader returns a new BitReader.
func NewBitReader(data []byte) *BitReader {
	return &BitReader{
		Data: data,
	}
}

// ReadBool reads a single bit from the buffer.
func (r *BitReader) ReadBool() bool {
	if r.BitIndex == 8 {
		r.ByteIndex++
		r.BitIndex = 0
	}
	result := r.Data[r.ByteIndex]&(1<<(7-r.BitIndex)) != 0
	r.BitIndex++
	return result
}

// ReadUint64 writes bitCount bits of data from the buffer.
func (r *BitReader) ReadUint64(bitCount uint) uint64 {
	result := uint64(0)
	for i := uint(0); i < bitCount; i++ {
		result <<= 1
		if r.ReadBool() {
			result |= 0x1
		}
	}
	return result
}
