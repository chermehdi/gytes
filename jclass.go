package gytes

import "io"

// A Java class file representation.
// The class representation follows the definition found in the JVM specification
//
//    ClassFile {
//       u4             magic;
//       u2             minor_version;
//       u2             major_version;
//       u2             constant_pool_count;
//       cp_info        constant_pool[constant_pool_count-1];
//       u2             access_flags;
//       u2             this_class;
//       u2             super_class;
//       u2             interfaces_count;
//       u2             interfaces[interfaces_count];
//       u2             fields_count;
//       field_info     fields[fields_count];
//       u2             methods_count;
//       method_info    methods[methods_count];
//       u2             attributes_count;
//       attribute_info attributes[attributes_count];
//   }
//
type JavaClass struct {
	Name         string
	PoolCount    uint16
	CPool        ConstantPool
	SuperName    string
	MinorVersion uint16
	MajorVersion uint16
	Interfaces   []string
	Fields       []JavaField
	Methods      []JavaMethod
	Access       uint16
	SourceName   string
}

func NewJavaClass(name string) *JavaClass {
	return &JavaClass{
		Name:         name,
		SuperName:    "java.lang.Object",
		MinorVersion: MIN_VERSION,
		MajorVersion: MAJ_VERSION,
		Interfaces:   make([]string, 0),
		Fields:       make([]JavaField, 0),
		Methods:      make([]JavaMethod, 0),
		Access:       0,
	}
}

func (jc *JavaClass) SuperClass(className string) *JavaClass {
	jc.SuperName = className
	return jc
}

func (jc *JavaClass) Visibility(accessMask uint16) *JavaClass {
	jc.Access = accessMask
	return jc
}

func (jc *JavaClass) Implements(interfaces []string) *JavaClass {
	jc.Interfaces = interfaces
	return jc
}

func (jc *JavaClass) AddFields(fields []JavaField) *JavaClass {
	jc.Fields = fields
	return jc
}

func (jc *JavaClass) AddMethods(methods []JavaMethod) *JavaClass {
	jc.Methods = methods
	return jc
}

func (jc *JavaClass) Write(writer io.Writer) {
}

