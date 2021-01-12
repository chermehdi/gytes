package gytes

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func readClassFile(className string) io.Reader {
	file, err := os.Open(className)
	if err != nil {
		panic(err)
	}
	return bufio.NewReader(file)
}

func TestCanReadJavaClass(t *testing.T) {
	jclass, err := readClass("testdata/compiled/Hello.class")
	assert.Nil(t, err)

	expected := &JavaClass{
		Name:         "Hello",
		MinorVersion: 0,
		MajorVersion: 52,
		SuperName:    "java/lang/Object",
		Access:       uint16(ACC_PUBLIC | ACC_SUPER),
		Interfaces:   []string{"java/io/Serializable"},
		SourceName:   "Hello.java",
		Fields: []JavaField{
			{
				Name:       "MAGIC",
				Modifiers:  uint16(ACC_PUBLIC | ACC_STATIC | ACC_FINAL),
				Descriptor: "I",
			},
			{
				Name:       "message",
				Modifiers:  uint16(ACC_PRIVATE),
				Descriptor: "Ljava/lang/String;",
			},
		},
		Methods: []JavaMethod{
			{
				Name:       "<init>",
				Modifiers:  uint16(ACC_PUBLIC),
				Descriptor: "(Ljava/lang/String;)V",
			},
			{
				Name:       "main",
				Modifiers:  uint16(ACC_PUBLIC | ACC_STATIC),
				Descriptor: "([Ljava/lang/String;)V",
			},
		},
	}

	AssertClass(t, expected, jclass)
}

func TestCanReadJavaClassWithException(t *testing.T) {
	jclass, err := readClass("testdata/compiled/HelloJavaException.class")
	assert.Nil(t, err)
	assert.NotNil(t, jclass)

	expected := &JavaClass{
		Name:         "HelloJavaException",
		MinorVersion: 0,
		MajorVersion: 52,
		SuperName:    "java/lang/Object",
		Access:       uint16(ACC_PUBLIC | ACC_SUPER),
		Interfaces:   []string{},
		SourceName:   "HelloJavaException.java",
		Fields:       []JavaField{},
		Methods: []JavaMethod{
			{
				Name:       "<init>",
				Modifiers:  0,
				Descriptor: "()V",
			},
			{
				Name:       "methodWithException",
				Modifiers:  0,
				Descriptor: "()V",
				Exceptions: []string{"java/lang/Exception"},
			},
		},
	}

	AssertClass(t, expected, jclass)
}

func readClass(src string) (*JavaClass, error) {
	reader := readClassFile(src)
	classReader := &ClassReader{}
	return classReader.ReadClass(reader)
}

func AssertClass(t *testing.T, expected, got *JavaClass) {
	assert.NotNil(t, got)
	assert.Equal(t, expected.Name, got.Name)
	assert.Equal(t, expected.SuperName, got.SuperName)
	assert.Equal(t, expected.Interfaces, got.Interfaces)
	assert.True(t, 48 < got.MajorVersion && got.MajorVersion < 100)

	assert.Equal(t, len(expected.Fields), len(got.Fields))
	// Fields should be in the same order
	for i := range expected.Fields {
		AssertField(t, &expected.Fields[i], &got.Fields[i])
	}

	assert.Equal(t, len(expected.Methods), len(got.Methods))
	// Fields should be in the same order
	for i := range expected.Methods {
		AssertMethod(t, &expected.Methods[i], &got.Methods[i])
	}
}

func AssertField(t *testing.T, expected, got *JavaField) {
	assert.Equal(t, expected.Name, got.Name)
	assert.Equal(t, expected.Modifiers, got.Modifiers)
	assert.Equal(t, expected.Descriptor, got.Descriptor)
}

func AssertMethod(t *testing.T, expected, got *JavaMethod) {
	assert.Equal(t, expected.Name, got.Name)
	assert.Equal(t, expected.Modifiers, got.Modifiers)
	assert.Equal(t, expected.Descriptor, got.Descriptor)
	assert.Equal(t, expected.Exceptions, got.Exceptions)

	AssertCode(t, expected.Body, got.Body)
}

func AssertCode(t *testing.T, expected, got []BytesBlock) {
	// TODO
}
