package gytes

// TODO: Might not need it, unless we will use it to do the label.
type BytesBlock struct {
	Instructions []*ByteCode
}

func NewByteBlock() BytesBlock {
	insts := make([]*ByteCode, 0)
	return BytesBlock{insts}
}

func (bb *BytesBlock) Add(byteCode uint8) (*ByteCode, error) {
	bc, err := CreateByteCode(byteCode)
	if err != nil {
		return nil, err
	}
	bb.Instructions = append(bb.Instructions, bc)
	return bc, nil
}
