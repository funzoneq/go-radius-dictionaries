// Code generated by radius-dict-gen. DO NOT EDIT.

package zeus

import (
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Zeus_VendorID = 7146
)

func _Zeus_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Zeus_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Zeus_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Zeus_VendorID {
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

func _Zeus_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Zeus_VendorID {
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

func _Zeus_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Zeus_VendorID {
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
	return _Zeus_AddVendor(p, typ, attr)
}

func _Zeus_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Zeus_VendorID {
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

func ZeusZXTMGroup_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Zeus_AddVendor(p, 1, a)
}

func ZeusZXTMGroup_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Zeus_AddVendor(p, 1, a)
}

func ZeusZXTMGroup_Get(p *radius.Packet) (value []byte) {
	value, _ = ZeusZXTMGroup_Lookup(p)
	return
}

func ZeusZXTMGroup_GetString(p *radius.Packet) (value string) {
	value, _ = ZeusZXTMGroup_LookupString(p)
	return
}

func ZeusZXTMGroup_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Zeus_GetsVendor(p, 1) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ZeusZXTMGroup_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Zeus_GetsVendor(p, 1) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ZeusZXTMGroup_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Zeus_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func ZeusZXTMGroup_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Zeus_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func ZeusZXTMGroup_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Zeus_SetVendor(p, 1, a)
}

func ZeusZXTMGroup_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Zeus_SetVendor(p, 1, a)
}

func ZeusZXTMGroup_Del(p *radius.Packet) {
	_Zeus_DelVendor(p, 1)
}
