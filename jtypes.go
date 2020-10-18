package gytes

type JType struct {
	Name  string
	VMRep string
}

var (
	JBool   = JType{"boolean", "Z"}
	JByte   = JType{"byte", "B"}
	JShort  = JType{"short", "S"}
	JChar   = JType{"char", "C"}
	JInt    = JType{"int", "I"}
	JLong   = JType{"long", "J"}
	JFloat  = JType{"float", "F"}
	JDouble = JType{"double", "D"}
	JVoid   = JType{"void", "V"}

	JABool   = JType{"boolean[]", "[Z"}
	JAByte   = JType{"byte[]", "[B"}
	JAShort  = JType{"short[]", "[S"}
	JAChar   = JType{"char[]", "[C"}
	JAInt    = JType{"int[]", "[I"}
	JALong   = JType{"long[]", "[J"}
	JAFloat  = JType{"float[]", "[F"}
	JADouble = JType{"double[]", "[D"}
)
