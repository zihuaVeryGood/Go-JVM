package constant_pool

import "jvm/class/class_file_commons"

/**
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/

type ConstantFieldrefInfo struct {
	TagInfo
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (c *ConstantFieldrefInfo) ReadInfo(reader class_file_commons.Reader) ConstantPoolInfo {
	c.ClassIndex = reader.ReadUint16()
	c.NameAndTypeIndex = reader.ReadUint16()
	return c
}
