package gytes

const MAGIC = 0xCAFEBABE

const MIN_VERSION = 45
const MAJ_VERSION = 58

/*
Source: https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4
-----------------------------------
Constant type                |Value|
-----------------------------------
CONSTANT_Class	 			       |  7 |
CONSTANT_Fieldref	           |  9 |
CONSTANT_Methodref	         | 10 |
CONSTANT_InterfaceMethodref	 | 11 |
CONSTANT_String		           |  8 |
CONSTANT_Integer	           |  3 |
CONSTANT_Float	             |  4 |
CONSTANT_Long	               |  5 |
CONSTANT_Double	             |  6 |
CONSTANT_NameAndType	       | 12 |
CONSTANT_Utf8	               |  1 |
CONSTANT_MethodHandle	       | 15 |
CONSTANT_MethodType	         | 16 |
CONSTANT_InvokeDynamic	     | 18 |
-----------------------------------
*/

const (
	ConstClass              = 7
	ConstFieldref           = 9
	ConstMethodref          = 10
	ConstInterfaceMethodref = 11
	ConstString             = 8
	ConstInteger            = 3
	ConstFloat              = 4
	ConstLong               = 5
	ConstDouble             = 6
	ConstNameAndType        = 12
	ConstUtf8               = 1
	ConstMethodHandle       = 15
	ConstMethodType         = 16
	ConstInvokeDynamic      = 18
)

var ConstSizeMap = map[int]int{
	ConstClass:              3,
	ConstFieldref:           5,
	ConstMethodref:          5,
	ConstInterfaceMethodref: 5,
	ConstString:             3,
	ConstInteger:            5,
	ConstFloat:              5,
	ConstLong:               9,
	ConstDouble:             9,
	ConstNameAndType:        5,
	ConstMethodHandle:       4,
	ConstMethodType:         3,
	ConstInvokeDynamic:      5,
}
