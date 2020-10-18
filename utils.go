package gytes

// For some reason Golang does not offer a Max function for uint32 in the standard
// Library, Weird!!
func max(a, b uint32) uint32 {
	if a < b {
		return b
	}
	return a
}

// Utility class used to hold written class bytes
type ByteVector struct {
	Data          []uint8
	Size          uint32
	currentLength uint32
}

func (bv *ByteVector) grow(minSize uint32) {
	doubleSize := bv.Size << 1
	newSize := max(doubleSize, bv.Size+minSize)
	newData := make([]uint8, newSize)
	copy(newData, bv.Data)
	bv.Size = newSize
}

func (bv *ByteVector) putByte(byteValue uint8) {
	if bv.currentLength+1 > bv.Size {
		bv.grow(1)
	}
	bv.Data[bv.currentLength] = byteValue
	bv.currentLength++
}

func (bv *ByteVector) put2Bytes(byteValue1, byteValue2 uint8) {
	if bv.currentLength+2 > bv.Size {
		bv.grow(2)
	}
	bv.Data[bv.currentLength] = byteValue1
	bv.currentLength++
	bv.Data[bv.currentLength] = byteValue2
	bv.currentLength++
}
