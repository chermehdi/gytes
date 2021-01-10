package gytes

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

// JClassReader is the main type to read a classfile into a `JavaClass` type.
type JClassReader interface {
	ReadClass(reader io.Reader) (*JavaClass, error)
}

// Component that is responsible of reading a sequence of bytes
// provides by a io.Reader, and converting it into a class representation
// in memory according to the spec defined in the JVM spec:
// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.10.2.2
type ClassReader struct {
	CurrentIndex int
	// Contains next entry index in the class's Constant pool
	PoolItems []int
	// Cache for the stings found in the ConstantPool
	PoolStr []string
	// Pointer to the offset where the class header starts
	HeadStart int
}

func (c *ClassReader) ReadClass(reader io.Reader) (*JavaClass, error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	magic := readMagic(bytes)
	if magic != MAGIC {
		return nil, errors.New(fmt.Sprintf("Invalid class file, expected magic bit found %v", magic))
	}
	jclass := &JavaClass{}
	jclass.MinorVersion = readUnsignedShort(bytes, 4)
	jclass.MajorVersion = readUnsignedShort(bytes, 6)
	jclass.PoolCount = readUnsignedShort(bytes, 8)

	if err := c.fillPoolItems(bytes, jclass.PoolCount); err != nil {
		return nil, err
	}

	jclass.Access = readUnsignedShort(bytes, c.HeadStart)
	jclass.Name = c.readClass(bytes, c.HeadStart+2)
	jclass.SuperName = c.readClass(bytes, c.HeadStart+4)
	interfaceCount := int(readUnsignedShort(bytes, c.HeadStart+6))
	jclass.Interfaces = make([]string, interfaceCount)
	cur := c.HeadStart + 8
	for i := 0; i < interfaceCount; i++ {
		jclass.Interfaces[i] = c.readClass(bytes, cur)
		cur += 2
	}
	f := cur
	// Skip fields
	fieldsCount := int(readUnsignedShort(bytes, f))
	f += 2
	jclass.Fields = make([]JavaField, fieldsCount)
	for i := 0; i < fieldsCount; i++ {
		jclass.Fields[i].Modifiers = readUnsignedShort(bytes, f)
		jclass.Fields[i].Name = c.readStr(bytes, f+2)
		jclass.Fields[i].Descriptor = c.readStr(bytes, f+4)
		attrCount := int(readUnsignedShort(bytes, f+6))
		// Skip access_flags, name, descriptor, and count
		f += 8
		// skip attributes
		for ; attrCount > 0; attrCount-- {
			attrName := c.readStr(bytes, f)
			// TODO: handle more attributes
			if attrName == "Synthetic" {
				jclass.Fields[i].Modifiers |= ACC_SYNTHETIC
			}
			attrLen := int(readInt(bytes, f+2))
			f += 6 + attrLen
		}
	}
	// Skip methods
	m := f
	methCount := int(readUnsignedShort(bytes, m))
	jclass.Methods = make([]JavaMethod, methCount)
	m += 2
	for i := 0; i < methCount; i++ {
		jclass.Methods[i].Modifiers = readUnsignedShort(bytes, m)
		jclass.Methods[i].Name = c.readStr(bytes, m+2)
		jclass.Methods[i].Descriptor = c.readStr(bytes, m+4)
		attrCount := int(readUnsignedShort(bytes, m+6))
		m += 8
		for ; attrCount > 0; attrCount-- {
			attrName := c.readStr(bytes, m)
			attrLen := int(readInt(bytes, m+2))
			m += 6
			// TODO: handle more attributes
			fmt.Printf("Found attribute %s with length %d\n", attrName, attrLen)
			if attrName == "Synthetic" {
				jclass.Methods[i].Modifiers |= ACC_SYNTHETIC
			} else if attrName == "Code" {
				jclass.Methods[i].BodyOffset = m
				jclass.Methods[i].MaxStack = readUnsignedShort(bytes, m)
				jclass.Methods[i].MaxLocals = readUnsignedShort(bytes, m+2)
				codeLen := readUnsignedInt(bytes, m+4)
				body, err := readCode(bytes, m+8, codeLen)
				if err != nil {
					return nil, err
				}
				jclass.Methods[i].Body = body
			}
			m += attrLen
		}
	}
	attrCount := int(readUnsignedShort(bytes, m))
	m += 2
	for i := 0; i < attrCount; i++ {
		attrName := c.readStr(bytes, m)
		// TODO: handle more attributes
		if attrName == "SourceFile" {
			jclass.SourceName = c.readStr(bytes, m+6)
		}
		attrLen := int(readInt(bytes, m+2))
		m += 6 + attrLen
	}
	return jclass, nil
}

func (c *ClassReader) readClass(b []byte, offset int) string {
	return c.readStr(b, c.PoolItems[readUnsignedShort(b, offset)])
}

func (c *ClassReader) fillPoolItems(b []byte, poolSize uint16) error {
	c.PoolItems = make([]int, poolSize)
	c.PoolStr = make([]string, poolSize)
	ptr := 10
	for i := uint16(1); i < poolSize; i += 1 {
		c.PoolItems[i] = ptr + 1
		curIndex := int(b[ptr])
		curSize, exists := ConstSizeMap[curIndex]
		if !exists {
			if curIndex == ConstUtf8 {
				curSize = 3 + int(readUnsignedShort(b, ptr+1))
			} else {
				return errors.New(fmt.Sprintf("Could not determine pool item type %v\n", curIndex))
			}
		}
		ptr += curSize
	}
	c.CurrentIndex = ptr
	c.HeadStart = ptr
	return nil
}

func readMagic(bytes []byte) uint32 {
	return readUnsignedInt(bytes, 0)
}

func readUnsignedLong(b []byte, offset int) uint64 {
	var c = (uint64(b[offset]) & uint64(0xFF)) << 56
	c = c | (uint64(b[offset+1])&uint64(0xFF))<<48
	c = c | (uint64(b[offset+2])&uint64(0xFF))<<40
	c = c | (uint64(b[offset+3])&uint64(0xFF))<<32
	c = c | (uint64(b[offset+4])&uint64(0xFF))<<24
	c = c | (uint64(b[offset+5])&uint64(0xFF))<<16
	c = c | (uint64(b[offset+6])&uint64(0xFF))<<8
	c = c | (uint64(b[offset+7]) & uint64(0xFF))
	return c
}

func readLong(b []byte, offset int) int64 {
	var c = (int64(b[offset]) & int64(0xFF)) << 56
	c = c | (int64(b[offset+1])&int64(0xFF))<<48
	c = c | (int64(b[offset+2])&int64(0xFF))<<40
	c = c | (int64(b[offset+3])&int64(0xFF))<<32
	c = c | (int64(b[offset+4])&int64(0xFF))<<24
	c = c | (int64(b[offset+5])&int64(0xFF))<<16
	c = c | (int64(b[offset+6])&int64(0xFF))<<8
	c = c | (int64(b[offset+7]) & int64(0xFF))
	return c
}

func readUnsignedInt(b []byte, offset int) uint32 {
	var c = (uint32(b[offset]) & uint32(0xFF)) << 24
	c = c | (uint32(b[offset+1])&uint32(0xFF))<<16
	c = c | (uint32(b[offset+2])&uint32(0xFF))<<8
	c = c | (uint32(b[offset+3]) & uint32(0xFF))
	return c
}

func readInt(b []byte, offset int) int32 {
	var c = (int32(b[offset]) & int32(0xFF)) << 24
	c = c | (int32(b[offset+1])&int32(0xFF))<<16
	c = c | (int32(b[offset+2])&int32(0xFF))<<8
	c = c | (int32(b[offset+3]) & int32(0xFF))
	return c
}

func readUnsignedShort(b []byte, offset int) uint16 {
	var c = (uint16(b[offset]) & uint16(0xFF)) << 8
	c = c | (uint16(b[offset+1]) & uint16(0xFF))
	return c
}

func (c *ClassReader) readStr(b []byte, offset int) string {
	item := readUnsignedShort(b, offset)
	index := c.PoolItems[item]
	length := int(readUnsignedShort(b, index))
	index += 2
	// TODO: strings in java are UTF-8 encoded, it should be handled accordingly
	endIndex := index + length
	return string(b[index:endIndex])
}

// Read the code of the method
func readCode(b []byte, offset int, length uint32) ([]BytesBlock, error) {
	blocks := make([]BytesBlock, 0)
	end := uint32(offset) + length
	curBlock := NewByteBlock()
	for i := uint32(offset); i < end; i = i + 1 {
		// TODO: curBlock may change when dealing with labels or control flow
		bc, err := curBlock.Add(b[i])
		if err != nil {
			return nil, err
		}
		i = i + uint32(bc.Size)
	}
	blocks = append(blocks, curBlock)
	return blocks, nil
}
