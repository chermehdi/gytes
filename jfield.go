package gytes

// A Java field representation
//
// field_info {
//   u2             access_flags;
//   u2             name_index;
//   u2             descriptor_index;
//   u2             attributes_count;
//   attribute_info attributes[attributes_count];
// }
type JavaField struct {
	Name       string
	Modifiers  uint16
	Descriptor string
}
