// Code generated by radius-dict-gen. DO NOT EDIT.

package iec

import (
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_IEC_VendorID = 41912
)

func _IEC_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_IEC_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _IEC_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _IEC_VendorID {
			continue
		}
		for len(vsa) >= 3 {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				break
			}
			if vsaTyp == typ {
				values = append(values, vsa[2:int(vsaLen)])
			}
			vsa = vsa[int(vsaLen):]
		}
	}
	return
}

func _IEC_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _IEC_VendorID {
			continue
		}
		for len(vsa) >= 3 {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				break
			}
			if vsaTyp == typ {
				return vsa[2:int(vsaLen)], true
			}
			vsa = vsa[int(vsaLen):]
		}
	}
	return
}

func _IEC_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _IEC_VendorID {
			i++
			continue
		}
		for j := 0; len(vsa[j:]) >= 3; {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa[j:]) || vsaLen < 3 {
				i++
				break
			}
			if vsaTyp == typ {
				vsa = append(vsa[:j], vsa[j+int(vsaLen):]...)
			}
			j += int(vsaLen)
		}
		if len(vsa) > 0 {
			copy(avp.Attribute[4:], vsa)
			i++
		} else {
			p.Attributes = append(p.Attributes[:i], p.Attributes[i+i:]...)
		}
	}
	return _IEC_AddVendor(p, typ, attr)
}

func _IEC_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _IEC_VendorID {
			i++
			continue
		}
		offset := 0
		for len(vsa[offset:]) >= 3 {
			vsaTyp, vsaLen := vsa[offset], vsa[offset+1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				continue vsaLoop
			}
			if vsaTyp == typ {
				copy(vsa[offset:], vsa[offset+int(vsaLen):])
				vsa = vsa[:len(vsa)-int(vsaLen)]
			} else {
				offset += int(vsaLen)
			}
		}
		if offset == 0 {
			p.Attributes = append(p.Attributes[:i], p.Attributes[i+1:]...)
		} else {
			i++
		}
	}
	return
}

type IEC623518RoleID uint32

var IEC623518RoleID_Strings = map[IEC623518RoleID]string{}

func (a IEC623518RoleID) String() string {
	if str, ok := IEC623518RoleID_Strings[a]; ok {
		return str
	}
	return "IEC623518RoleID(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func IEC623518RoleID_Add(p *radius.Packet, value IEC623518RoleID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _IEC_AddVendor(p, 1, a)
}

func IEC623518RoleID_Get(p *radius.Packet) (value IEC623518RoleID) {
	value, _ = IEC623518RoleID_Lookup(p)
	return
}

func IEC623518RoleID_Gets(p *radius.Packet) (values []IEC623518RoleID, err error) {
	var i uint32
	for _, attr := range _IEC_GetsVendor(p, 1) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, IEC623518RoleID(i))
	}
	return
}

func IEC623518RoleID_Lookup(p *radius.Packet) (value IEC623518RoleID, err error) {
	a, ok := _IEC_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = IEC623518RoleID(i)
	return
}

func IEC623518RoleID_Set(p *radius.Packet, value IEC623518RoleID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _IEC_SetVendor(p, 1, a)
}

func IEC623518RoleID_Del(p *radius.Packet) {
	_IEC_DelVendor(p, 1)
}

func IEC623518RoleDefinition_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _IEC_AddVendor(p, 2, a)
}

func IEC623518RoleDefinition_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _IEC_AddVendor(p, 2, a)
}

func IEC623518RoleDefinition_Get(p *radius.Packet) (value []byte) {
	value, _ = IEC623518RoleDefinition_Lookup(p)
	return
}

func IEC623518RoleDefinition_GetString(p *radius.Packet) (value string) {
	value, _ = IEC623518RoleDefinition_LookupString(p)
	return
}

func IEC623518RoleDefinition_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _IEC_GetsVendor(p, 2) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func IEC623518RoleDefinition_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _IEC_GetsVendor(p, 2) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func IEC623518RoleDefinition_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _IEC_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func IEC623518RoleDefinition_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _IEC_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func IEC623518RoleDefinition_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _IEC_SetVendor(p, 2, a)
}

func IEC623518RoleDefinition_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _IEC_SetVendor(p, 2, a)
}

func IEC623518RoleDefinition_Del(p *radius.Packet) {
	_IEC_DelVendor(p, 2)
}

func IEC623518AoR_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _IEC_AddVendor(p, 3, a)
}

func IEC623518AoR_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _IEC_AddVendor(p, 3, a)
}

func IEC623518AoR_Get(p *radius.Packet) (value []byte) {
	value, _ = IEC623518AoR_Lookup(p)
	return
}

func IEC623518AoR_GetString(p *radius.Packet) (value string) {
	value, _ = IEC623518AoR_LookupString(p)
	return
}

func IEC623518AoR_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _IEC_GetsVendor(p, 3) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func IEC623518AoR_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _IEC_GetsVendor(p, 3) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func IEC623518AoR_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _IEC_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func IEC623518AoR_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _IEC_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func IEC623518AoR_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _IEC_SetVendor(p, 3, a)
}

func IEC623518AoR_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _IEC_SetVendor(p, 3, a)
}

func IEC623518AoR_Del(p *radius.Packet) {
	_IEC_DelVendor(p, 3)
}

type IEC623518Revision uint32

var IEC623518Revision_Strings = map[IEC623518Revision]string{}

func (a IEC623518Revision) String() string {
	if str, ok := IEC623518Revision_Strings[a]; ok {
		return str
	}
	return "IEC623518Revision(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func IEC623518Revision_Add(p *radius.Packet, value IEC623518Revision) (err error) {
	a := radius.NewInteger(uint32(value))
	return _IEC_AddVendor(p, 4, a)
}

func IEC623518Revision_Get(p *radius.Packet) (value IEC623518Revision) {
	value, _ = IEC623518Revision_Lookup(p)
	return
}

func IEC623518Revision_Gets(p *radius.Packet) (values []IEC623518Revision, err error) {
	var i uint32
	for _, attr := range _IEC_GetsVendor(p, 4) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, IEC623518Revision(i))
	}
	return
}

func IEC623518Revision_Lookup(p *radius.Packet) (value IEC623518Revision, err error) {
	a, ok := _IEC_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = IEC623518Revision(i)
	return
}

func IEC623518Revision_Set(p *radius.Packet, value IEC623518Revision) (err error) {
	a := radius.NewInteger(uint32(value))
	return _IEC_SetVendor(p, 4, a)
}

func IEC623518Revision_Del(p *radius.Packet) {
	_IEC_DelVendor(p, 4)
}

func IEC623518ValidFrom_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _IEC_AddVendor(p, 5, a)
}

func IEC623518ValidFrom_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _IEC_AddVendor(p, 5, a)
}

func IEC623518ValidFrom_Get(p *radius.Packet) (value []byte) {
	value, _ = IEC623518ValidFrom_Lookup(p)
	return
}

func IEC623518ValidFrom_GetString(p *radius.Packet) (value string) {
	value, _ = IEC623518ValidFrom_LookupString(p)
	return
}

func IEC623518ValidFrom_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _IEC_GetsVendor(p, 5) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func IEC623518ValidFrom_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _IEC_GetsVendor(p, 5) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func IEC623518ValidFrom_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _IEC_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func IEC623518ValidFrom_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _IEC_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func IEC623518ValidFrom_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _IEC_SetVendor(p, 5, a)
}

func IEC623518ValidFrom_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _IEC_SetVendor(p, 5, a)
}

func IEC623518ValidFrom_Del(p *radius.Packet) {
	_IEC_DelVendor(p, 5)
}

func IEC623518ValidTo_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _IEC_AddVendor(p, 6, a)
}

func IEC623518ValidTo_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _IEC_AddVendor(p, 6, a)
}

func IEC623518ValidTo_Get(p *radius.Packet) (value []byte) {
	value, _ = IEC623518ValidTo_Lookup(p)
	return
}

func IEC623518ValidTo_GetString(p *radius.Packet) (value string) {
	value, _ = IEC623518ValidTo_LookupString(p)
	return
}

func IEC623518ValidTo_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _IEC_GetsVendor(p, 6) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func IEC623518ValidTo_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _IEC_GetsVendor(p, 6) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func IEC623518ValidTo_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _IEC_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func IEC623518ValidTo_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _IEC_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func IEC623518ValidTo_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _IEC_SetVendor(p, 6, a)
}

func IEC623518ValidTo_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _IEC_SetVendor(p, 6, a)
}

func IEC623518ValidTo_Del(p *radius.Packet) {
	_IEC_DelVendor(p, 6)
}
