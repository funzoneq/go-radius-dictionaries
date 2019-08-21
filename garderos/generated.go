// Code generated by radius-dict-gen. DO NOT EDIT.

package garderos

import (
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Garderos_VendorID = 16108
)

func _Garderos_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Garderos_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Garderos_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, attr := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Garderos_VendorID {
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

func _Garderos_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, a := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(a)
		if err != nil || vendorID != _Garderos_VendorID {
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

func _Garderos_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		vendorID, vsa, err := radius.VendorSpecific(p.Attributes[rfc2865.VendorSpecific_Type][i])
		if err != nil || vendorID != _Garderos_VendorID {
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
	return _Garderos_AddVendor(p, typ, attr)
}

func _Garderos_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		attr := p.Attributes[rfc2865.VendorSpecific_Type][i]
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Garderos_VendorID {
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

func GarderosLocationName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Garderos_AddVendor(p, 1, a)
}

func GarderosLocationName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Garderos_AddVendor(p, 1, a)
}

func GarderosLocationName_Get(p *radius.Packet) (value []byte) {
	value, _ = GarderosLocationName_Lookup(p)
	return
}

func GarderosLocationName_GetString(p *radius.Packet) (value string) {
	value, _ = GarderosLocationName_LookupString(p)
	return
}

func GarderosLocationName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Garderos_GetsVendor(p, 1) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func GarderosLocationName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Garderos_GetsVendor(p, 1) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func GarderosLocationName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Garderos_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func GarderosLocationName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Garderos_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func GarderosLocationName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Garderos_SetVendor(p, 1, a)
}

func GarderosLocationName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Garderos_SetVendor(p, 1, a)
}

func GarderosLocationName_Del(p *radius.Packet) {
	_Garderos_DelVendor(p, 1)
}

func GarderosServiceName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Garderos_AddVendor(p, 2, a)
}

func GarderosServiceName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Garderos_AddVendor(p, 2, a)
}

func GarderosServiceName_Get(p *radius.Packet) (value []byte) {
	value, _ = GarderosServiceName_Lookup(p)
	return
}

func GarderosServiceName_GetString(p *radius.Packet) (value string) {
	value, _ = GarderosServiceName_LookupString(p)
	return
}

func GarderosServiceName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Garderos_GetsVendor(p, 2) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func GarderosServiceName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Garderos_GetsVendor(p, 2) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func GarderosServiceName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Garderos_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func GarderosServiceName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Garderos_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func GarderosServiceName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Garderos_SetVendor(p, 2, a)
}

func GarderosServiceName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Garderos_SetVendor(p, 2, a)
}

func GarderosServiceName_Del(p *radius.Packet) {
	_Garderos_DelVendor(p, 2)
}

func GarderosMSISDN_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Garderos_AddVendor(p, 3, a)
}

func GarderosMSISDN_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Garderos_AddVendor(p, 3, a)
}

func GarderosMSISDN_Get(p *radius.Packet) (value []byte) {
	value, _ = GarderosMSISDN_Lookup(p)
	return
}

func GarderosMSISDN_GetString(p *radius.Packet) (value string) {
	value, _ = GarderosMSISDN_LookupString(p)
	return
}

func GarderosMSISDN_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Garderos_GetsVendor(p, 3) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func GarderosMSISDN_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Garderos_GetsVendor(p, 3) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func GarderosMSISDN_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Garderos_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func GarderosMSISDN_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Garderos_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func GarderosMSISDN_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Garderos_SetVendor(p, 3, a)
}

func GarderosMSISDN_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Garderos_SetVendor(p, 3, a)
}

func GarderosMSISDN_Del(p *radius.Packet) {
	_Garderos_DelVendor(p, 3)
}

func GarderosProxy_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Garderos_AddVendor(p, 4, a)
}

func GarderosProxy_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Garderos_AddVendor(p, 4, a)
}

func GarderosProxy_Get(p *radius.Packet) (value []byte) {
	value, _ = GarderosProxy_Lookup(p)
	return
}

func GarderosProxy_GetString(p *radius.Packet) (value string) {
	value, _ = GarderosProxy_LookupString(p)
	return
}

func GarderosProxy_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Garderos_GetsVendor(p, 4) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func GarderosProxy_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Garderos_GetsVendor(p, 4) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func GarderosProxy_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Garderos_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func GarderosProxy_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Garderos_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func GarderosProxy_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Garderos_SetVendor(p, 4, a)
}

func GarderosProxy_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Garderos_SetVendor(p, 4, a)
}

func GarderosProxy_Del(p *radius.Packet) {
	_Garderos_DelVendor(p, 4)
}