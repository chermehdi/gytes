package gytes

import (
	"bufio"
	"fmt"
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
	reader := readClassFile("testdata/compiled/Hello.class")
	classReader := &ClassReader{}

	jclass, err := classReader.ReadClass(reader)

	assert.Nil(t, err)
	assert.NotNil(t, jclass)
	assert.Equal(t, uint16(0), jclass.MinorVersion)
	assert.True(t, 48 < jclass.MajorVersion && jclass.MajorVersion < 100)
	assert.Equal(t, uint16(39), jclass.PoolCount)

	assert.Equal(t, uint16(ACC_PUBLIC|ACC_SUPER), jclass.Access)
	assert.Equal(t, "Hello", jclass.Name)
	assert.Equal(t, "java/lang/Object", jclass.SuperName)
	assert.Equal(t, []string{"java/io/Serializable"}, jclass.Interfaces)
	assert.Equal(t, "Hello.java", jclass.SourceName)

	// Field assertions
	assert.Equal(t, 2, len(jclass.Fields))
	assert.Equal(t, "MAGIC", jclass.Fields[0].Name)
	assert.Equal(t, uint16(ACC_PUBLIC|ACC_STATIC|ACC_FINAL), jclass.Fields[0].Modifiers)
	assert.Equal(t, "I", jclass.Fields[0].Descriptor)

	assert.Equal(t, "message", jclass.Fields[1].Name)
	assert.Equal(t, uint16(ACC_PRIVATE), jclass.Fields[1].Modifiers)
	assert.Equal(t, "Ljava/lang/String;", jclass.Fields[1].Descriptor)

	// Method assertions
	assert.Equal(t, 2, len(jclass.Methods))
	assert.Equal(t, "<init>", jclass.Methods[0].Name)
	assert.Equal(t, uint16(ACC_PUBLIC), jclass.Methods[0].Modifiers)
	assert.Equal(t, "(Ljava/lang/String;)V", jclass.Methods[0].Descriptor)

	assert.Equal(t, "main", jclass.Methods[1].Name)
	assert.Equal(t, uint16(ACC_PUBLIC|ACC_STATIC), jclass.Methods[1].Modifiers)
	assert.Equal(t, "([Ljava/lang/String;)V", jclass.Methods[1].Descriptor)

	// compiled with a recent java version

	fmt.Printf("%v\n", jclass)
}
