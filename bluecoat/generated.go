// Code generated by radius-dict-gen. DO NOT EDIT.

package bluecoat

import (
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_BlueCoat_VendorID = 14501
)

func _BlueCoat_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_BlueCoat_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _BlueCoat_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, attr := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _BlueCoat_VendorID {
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

func _BlueCoat_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, a := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(a)
		if err != nil || vendorID != _BlueCoat_VendorID {
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

func _BlueCoat_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		vendorID, vsa, err := radius.VendorSpecific(p.Attributes[rfc2865.VendorSpecific_Type][i])
		if err != nil || vendorID != _BlueCoat_VendorID {
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
	return _BlueCoat_AddVendor(p, typ, attr)
}

func _BlueCoat_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		attr := p.Attributes[rfc2865.VendorSpecific_Type][i]
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _BlueCoat_VendorID {
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

func BlueCoatGroup_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _BlueCoat_AddVendor(p, 1, a)
}

func BlueCoatGroup_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _BlueCoat_AddVendor(p, 1, a)
}

func BlueCoatGroup_Get(p *radius.Packet) (value []byte) {
	value, _ = BlueCoatGroup_Lookup(p)
	return
}

func BlueCoatGroup_GetString(p *radius.Packet) (value string) {
	value, _ = BlueCoatGroup_LookupString(p)
	return
}

func BlueCoatGroup_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _BlueCoat_GetsVendor(p, 1) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func BlueCoatGroup_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _BlueCoat_GetsVendor(p, 1) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func BlueCoatGroup_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _BlueCoat_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func BlueCoatGroup_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _BlueCoat_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func BlueCoatGroup_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _BlueCoat_SetVendor(p, 1, a)
}

func BlueCoatGroup_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _BlueCoat_SetVendor(p, 1, a)
}

func BlueCoatGroup_Del(p *radius.Packet) {
	_BlueCoat_DelVendor(p, 1)
}

type BlueCoatAuthorization uint32

const (
	BlueCoatAuthorization_Value_NoAccess        BlueCoatAuthorization = 0
	BlueCoatAuthorization_Value_ReadOnlyAccess  BlueCoatAuthorization = 1
	BlueCoatAuthorization_Value_ReadWriteAccess BlueCoatAuthorization = 2
)

var BlueCoatAuthorization_Strings = map[BlueCoatAuthorization]string{
	BlueCoatAuthorization_Value_NoAccess:        "No-Access",
	BlueCoatAuthorization_Value_ReadOnlyAccess:  "Read-Only-Access",
	BlueCoatAuthorization_Value_ReadWriteAccess: "Read-Write-Access",
}

func (a BlueCoatAuthorization) String() string {
	if str, ok := BlueCoatAuthorization_Strings[a]; ok {
		return str
	}
	return "BlueCoatAuthorization(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func BlueCoatAuthorization_Add(p *radius.Packet, value BlueCoatAuthorization) (err error) {
	a := radius.NewInteger(uint32(value))
	return _BlueCoat_AddVendor(p, 2, a)
}

func BlueCoatAuthorization_Get(p *radius.Packet) (value BlueCoatAuthorization) {
	value, _ = BlueCoatAuthorization_Lookup(p)
	return
}

func BlueCoatAuthorization_Gets(p *radius.Packet) (values []BlueCoatAuthorization, err error) {
	var i uint32
	for _, attr := range _BlueCoat_GetsVendor(p, 2) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, BlueCoatAuthorization(i))
	}
	return
}

func BlueCoatAuthorization_Lookup(p *radius.Packet) (value BlueCoatAuthorization, err error) {
	a, ok := _BlueCoat_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = BlueCoatAuthorization(i)
	return
}

func BlueCoatAuthorization_Set(p *radius.Packet, value BlueCoatAuthorization) (err error) {
	a := radius.NewInteger(uint32(value))
	return _BlueCoat_SetVendor(p, 2, a)
}

func BlueCoatAuthorization_Del(p *radius.Packet) {
	_BlueCoat_DelVendor(p, 2)
}