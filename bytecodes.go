package gytes

import (
	"errors"
	"fmt"
)

type ByteCode struct {
	Name  string
	Value uint8
	Size  int
}

func (bb ByteCode) String() string {
	return fmt.Sprintf("(%d, %s)", bb.Value, bb.Name)
}

var UnknownByteCodeError = errors.New("Unknown bytecode number")

var ByteCodes = []ByteCode{
	{"nop", 0, 0},
	{"aconst_null", 1, 0},
	{"iconst_m1", 2, 0},
	{"iconst_0", 3, 0},
	{"iconst_1", 4, 0},
	{"iconst_2", 5, 0},
	{"iconst_3", 6, 0},
	{"iconst_4", 7, 0},
	{"iconst_5", 8, 0},
	{"lconst_0", 9, 0},
	{"lconst_1", 10, 0},
	{"fconst_0", 11, 0},
	{"fconst_1", 12, 0},
	{"fconst_2", 13, 0},
	{"dconst_0", 14, 0},
	{"dconst_1", 15, 0},
	{"bipush", 16, 1},
	{"sipush", 17, 2},
	{"ldc", 18, 1},
	{"ldc_w", 19, 2},
	{"ldc2_w", 20, 2},
	{"iload", 21, 1},
	{"lload", 22, 1},
	{"fload", 23, 1},
	{"dload", 24, 1},
	{"aload", 25, 1},
	{"iload_0", 26, 0},
	{"iload_1", 27, 0},
	{"iload_2", 28, 0},
	{"iload_3", 29, 0},
	{"lload_0", 30, 0},
	{"lload_1", 31, 0},
	{"lload_2", 32, 0},
	{"lload_3", 33, 0},
	{"fload_0", 34, 0},
	{"fload_1", 35, 0},
	{"fload_2", 36, 0},
	{"fload_3", 37, 0},
	{"dload_0", 38, 0},
	{"dload_1", 39, 0},
	{"dload_2", 40, 0},
	{"dload_3", 41, 0},
	{"aload_0", 42, 0},
	{"aload_1", 43, 0},
	{"aload_2", 44, 0},
	{"aload_3", 45, 0},
	{"iaload", 46, 0},
	{"laload", 47, 0},
	{"faload", 48, 0},
	{"daload", 49, 0},
	{"aaload", 50, 0},
	{"baload", 51, 0},
	{"caload", 52, 0},
	{"saload", 53, 0},
	{"istore", 54, 0},
	{"lstore", 55, 0},
	{"fstore", 56, 0},
	{"dstore", 57, 0},
	{"astore", 58, 0},
	{"istore_0", 59, 0},
	{"istore_1", 60, 0},
	{"istore_2", 61, 0},
	{"istore_3", 62, 0},
	{"lstore_0", 63, 0},
	{"lstore_1", 64, 0},
	{"lstore_2", 65, 0},
	{"lstore_3", 66, 0},
	{"fstore_0", 67, 0},
	{"fstore_1", 68, 0},
	{"fstore_2", 69, 0},
	{"fstore_3", 70, 0},
	{"dstore_0", 71, 0},
	{"dstore_1", 72, 0},
	{"dstore_2", 73, 0},
	{"dstore_3", 74, 0},
	{"astore_0", 75, 0},
	{"astore_1", 76, 0},
	{"astore_2", 77, 0},
	{"astore_3", 78, 0},
	{"iastore", 79, 0},
	{"lastore", 80, 0},
	{"fastore", 81, 0},
	{"dastore", 82, 0},
	{"aastore", 83, 0},
	{"bastore", 84, 0},
	{"castore", 85, 0},
	{"sastore", 86, 0},
	{"pop", 87, 0},
	{"pop2", 88, 0},
	{"dup", 89, 0},
	{"dup_x1", 90, 0},
	{"dup_x2", 91, 0},
	{"dup2", 92, 0},
	{"dup2_x1", 93, 0},
	{"dup2_x2", 94, 0},
	{"swap", 95, 0},
	{"iadd", 96, 0},
	{"ladd", 97, 0},
	{"fadd", 98, 0},
	{"dadd", 99, 0},
	{"isub", 100, 0},
	{"lsub", 101, 0},
	{"fsub", 102, 0},
	{"dsub", 103, 0},
	{"imul", 104, 0},
	{"lmul", 105, 0},
	{"fmul", 106, 0},
	{"dmul", 107, 0},
	{"idiv", 108, 0},
	{"ldiv", 109, 0},
	{"fdiv", 110, 0},
	{"ddiv", 111, 0},
	{"irem", 112, 0},
	{"lrem", 113, 0},
	{"frem", 114, 0},
	{"drem", 115, 0},
	{"ineg", 116, 0},
	{"lneg", 117, 0},
	{"fneg", 118, 0},
	{"dneg", 119, 0},
	{"ishl", 120, 0},
	{"lshl", 121, 0},
	{"ishr", 122, 0},
	{"lshr", 123, 0},
	{"iushr", 124, 0},
	{"lushr", 125, 0},
	{"iand", 126, 0},
	{"land", 127, 0},
	{"ior", 128, 0},
	{"lor", 129, 0},
	{"ixor", 130, 0},
	{"lxor", 131, 0},
	{"iinc", 132, 0},
	{"i2l", 133, 0},
	{"i2f", 134, 0},
	{"i2d", 135, 0},
	{"l2i", 136, 0},
	{"l2f", 137, 0},
	{"l2d", 138, 0},
	{"f2i", 139, 0},
	{"f2l", 140, 0},
	{"f2d", 141, 0},
	{"d2i", 142, 0},
	{"d2l", 143, 0},
	{"d2f", 144, 0},
	{"i2b", 145, 0},
	{"i2c", 146, 0},
	{"i2s", 147, 0},
	{"lcmp", 148, 0},
	{"fcmpl", 149, 0},
	{"fcmpg", 150, 0},
	{"dcmpl", 151, 0},
	{"dcmpg", 152, 0},
	{"ifeq", 153, 2},
	{"ifne", 154, 2},
	{"iflt", 155, 2},
	{"ifge", 156, 2},
	{"ifgt", 157, 2},
	{"ifle", 158, 2},
	{"if_icmpeq", 159, 2},
	{"if_icmpne", 160, 2},
	{"if_icmplt", 161, 2},
	{"if_icmpge", 162, 2},
	{"if_icmpgt", 163, 2},
	{"if_icmple", 164, 2},
	{"if_acmpeq", 165, 2},
	{"if_acmpne", 166, 2},
	{"goto", 167, 2},
	{"jsr", 168, 2},
	{"ret", 169, 1},
	// TODO: This is tricky, let's handle it later
	{"tableswitch", 170, 0},
	// TODO: This is tricky, let's handle it later
	{"lookupswitch", 171, 0},
	{"ireturn", 172, 0},
	{"lreturn", 173, 0},
	{"freturn", 174, 0},
	{"dreturn", 175, 0},
	{"areturn", 176, 0},
	{"return", 177, 0},
	{"getstatic", 178, 2},
	{"putstatic", 179, 2},
	{"getfield", 180, 2},
	{"putfield", 181, 2},
	{"invokevirtual", 182, 2},
	{"invokespecial", 183, 2},
	{"invokestatic", 184, 2},
	{"invokeinterface", 185, 4},
	{"invokedynamic", 186, 4},
	{"new", 187, 2},
	{"newarray", 188, 1},
	{"anewarray", 189, 2},
	{"arraylength", 190, 0},
	{"athrow", 191, 0},
	{"checkcast", 192, 2},
	{"instanceof", 193, 2},
	{"monitorenter", 194, 0},
	{"monitorexit", 195, 0},
	// TODO: this is tricky, let's handle it later
	{"wide", 196, 0},
	{"multianewarray", 197, 3},
	{"ifnull", 198, 2},
	{"ifnonnull", 199, 2},
	{"goto_w", 200, 4},
	{"jsr_w", 201, 4},
	{"breakpoint", 202, 0},
}

// CreateByteCode creates a ByteCode from a byteCode number
func CreateByteCode(byteCode uint8) (*ByteCode, error) {
	if byteCode < 0 || byteCode > 202 {
		return nil, UnknownByteCodeError
	}
	return &ByteCodes[byteCode], nil
}
