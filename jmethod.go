package gytes

// Java method representation
//
// method_info {
//   u2             access_flags;
//   u2             name_index;
//   u2             descriptor_index;
//   u2             attributes_count;
//   attribute_info attributes[attributes_count];
// }
type JavaMethod struct {
	Name       string
	Modifiers  uint16
	Descriptor string
	// The offset in the original class file at which the code of this class starts
	// This is computed at class read time by finding the Code attribute in the method's attribute list.
	BodyOffset int
	Body       []BytesBlock
}
