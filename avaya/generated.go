// Code generated by radius-dict-gen. DO NOT EDIT.

package avaya

import (
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_CajunP330_VendorID = 2167
)

func _CajunP330_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_CajunP330_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _CajunP330_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, attr := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _CajunP330_VendorID {
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

func _CajunP330_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, a := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(a)
		if err != nil || vendorID != _CajunP330_VendorID {
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

func _CajunP330_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		vendorID, vsa, err := radius.VendorSpecific(p.Attributes[rfc2865.VendorSpecific_Type][i])
		if err != nil || vendorID != _CajunP330_VendorID {
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
			copy(p.Attributes[rfc2865.VendorSpecific_Type][i][4:], vsa)
			i++
		} else {
			p.Attributes[rfc2865.VendorSpecific_Type] = append(p.Attributes[rfc2865.VendorSpecific_Type][:i], p.Attributes[rfc2865.VendorSpecific_Type][i+i:]...)
		}
	}
	return _CajunP330_AddVendor(p, typ, attr)
}

func _CajunP330_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		attr := p.Attributes[rfc2865.VendorSpecific_Type][i]
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _CajunP330_VendorID {
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
			p.Attributes[rfc2865.VendorSpecific_Type] = append(p.Attributes[rfc2865.VendorSpecific_Type][:i], p.Attributes[rfc2865.VendorSpecific_Type][i+1:]...)
		} else {
			i++
		}
	}
	return
}

type CajunServiceType uint32

const (
	CajunServiceType_Value_CajunReadOnlyUser  CajunServiceType = 1
	CajunServiceType_Value_CajunReadWriteUser CajunServiceType = 2
	CajunServiceType_Value_CajunAdminUser     CajunServiceType = 3
)

var CajunServiceType_Strings = map[CajunServiceType]string{
	CajunServiceType_Value_CajunReadOnlyUser:  "Cajun-Read-Only-User",
	CajunServiceType_Value_CajunReadWriteUser: "Cajun-Read-Write-User",
	CajunServiceType_Value_CajunAdminUser:     "Cajun-Admin-User",
}

func (a CajunServiceType) String() string {
	if str, ok := CajunServiceType_Strings[a]; ok {
		return str
	}
	return "CajunServiceType(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func CajunServiceType_Add(p *radius.Packet, value CajunServiceType) (err error) {
	a := radius.NewInteger(uint32(value))
	return _CajunP330_AddVendor(p, 1, a)
}

func CajunServiceType_Get(p *radius.Packet) (value CajunServiceType) {
	value, _ = CajunServiceType_Lookup(p)
	return
}

func CajunServiceType_Gets(p *radius.Packet) (values []CajunServiceType, err error) {
	var i uint32
	for _, attr := range _CajunP330_GetsVendor(p, 1) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, CajunServiceType(i))
	}
	return
}

func CajunServiceType_Lookup(p *radius.Packet) (value CajunServiceType, err error) {
	a, ok := _CajunP330_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = CajunServiceType(i)
	return
}

func CajunServiceType_Set(p *radius.Packet, value CajunServiceType) (err error) {
	a := radius.NewInteger(uint32(value))
	return _CajunP330_SetVendor(p, 1, a)
}

func CajunServiceType_Del(p *radius.Packet) {
	_CajunP330_DelVendor(p, 1)
}

func AvayaStaticVlanType_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _CajunP330_AddVendor(p, 12, a)
}

func AvayaStaticVlanType_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _CajunP330_AddVendor(p, 12, a)
}

func AvayaStaticVlanType_Get(p *radius.Packet) (value []byte) {
	value, _ = AvayaStaticVlanType_Lookup(p)
	return
}

func AvayaStaticVlanType_GetString(p *radius.Packet) (value string) {
	value, _ = AvayaStaticVlanType_LookupString(p)
	return
}

func AvayaStaticVlanType_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _CajunP330_GetsVendor(p, 12) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func AvayaStaticVlanType_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _CajunP330_GetsVendor(p, 12) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func AvayaStaticVlanType_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _CajunP330_LookupVendor(p, 12)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func AvayaStaticVlanType_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _CajunP330_LookupVendor(p, 12)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func AvayaStaticVlanType_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _CajunP330_SetVendor(p, 12, a)
}

func AvayaStaticVlanType_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _CajunP330_SetVendor(p, 12, a)
}

func AvayaStaticVlanType_Del(p *radius.Packet) {
	_CajunP330_DelVendor(p, 12)
}

type AvayaPortPriorityType uint32

const (
	AvayaPortPriorityType_Value_Type0 AvayaPortPriorityType = 0
	AvayaPortPriorityType_Value_Type1 AvayaPortPriorityType = 1
	AvayaPortPriorityType_Value_Type2 AvayaPortPriorityType = 2
	AvayaPortPriorityType_Value_Type3 AvayaPortPriorityType = 3
	AvayaPortPriorityType_Value_Type4 AvayaPortPriorityType = 4
	AvayaPortPriorityType_Value_Type5 AvayaPortPriorityType = 5
	AvayaPortPriorityType_Value_Type6 AvayaPortPriorityType = 6
	AvayaPortPriorityType_Value_Type7 AvayaPortPriorityType = 7
)

var AvayaPortPriorityType_Strings = map[AvayaPortPriorityType]string{
	AvayaPortPriorityType_Value_Type0: "Type-0",
	AvayaPortPriorityType_Value_Type1: "Type-1",
	AvayaPortPriorityType_Value_Type2: "Type-2",
	AvayaPortPriorityType_Value_Type3: "Type-3",
	AvayaPortPriorityType_Value_Type4: "Type-4",
	AvayaPortPriorityType_Value_Type5: "Type-5",
	AvayaPortPriorityType_Value_Type6: "Type-6",
	AvayaPortPriorityType_Value_Type7: "Type-7",
}

func (a AvayaPortPriorityType) String() string {
	if str, ok := AvayaPortPriorityType_Strings[a]; ok {
		return str
	}
	return "AvayaPortPriorityType(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func AvayaPortPriorityType_Add(p *radius.Packet, value AvayaPortPriorityType) (err error) {
	a := radius.NewInteger(uint32(value))
	return _CajunP330_AddVendor(p, 13, a)
}

func AvayaPortPriorityType_Get(p *radius.Packet) (value AvayaPortPriorityType) {
	value, _ = AvayaPortPriorityType_Lookup(p)
	return
}

func AvayaPortPriorityType_Gets(p *radius.Packet) (values []AvayaPortPriorityType, err error) {
	var i uint32
	for _, attr := range _CajunP330_GetsVendor(p, 13) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AvayaPortPriorityType(i))
	}
	return
}

func AvayaPortPriorityType_Lookup(p *radius.Packet) (value AvayaPortPriorityType, err error) {
	a, ok := _CajunP330_LookupVendor(p, 13)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AvayaPortPriorityType(i)
	return
}

func AvayaPortPriorityType_Set(p *radius.Packet, value AvayaPortPriorityType) (err error) {
	a := radius.NewInteger(uint32(value))
	return _CajunP330_SetVendor(p, 13, a)
}

func AvayaPortPriorityType_Del(p *radius.Packet) {
	_CajunP330_DelVendor(p, 13)
}