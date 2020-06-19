// Code generated by radius-dict-gen. DO NOT EDIT.

package brocade

import (
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Brocade_VendorID = 1588
)

func _Brocade_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Brocade_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Brocade_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Brocade_VendorID {
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

func _Brocade_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Brocade_VendorID {
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

func _Brocade_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Brocade_VendorID {
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
	return _Brocade_AddVendor(p, typ, attr)
}

func _Brocade_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Brocade_VendorID {
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

func BrocadeAuthRole_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Brocade_AddVendor(p, 1, a)
}

func BrocadeAuthRole_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Brocade_AddVendor(p, 1, a)
}

func BrocadeAuthRole_Get(p *radius.Packet) (value []byte) {
	value, _ = BrocadeAuthRole_Lookup(p)
	return
}

func BrocadeAuthRole_GetString(p *radius.Packet) (value string) {
	value, _ = BrocadeAuthRole_LookupString(p)
	return
}

func BrocadeAuthRole_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Brocade_GetsVendor(p, 1) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func BrocadeAuthRole_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Brocade_GetsVendor(p, 1) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func BrocadeAuthRole_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Brocade_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func BrocadeAuthRole_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Brocade_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func BrocadeAuthRole_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Brocade_SetVendor(p, 1, a)
}

func BrocadeAuthRole_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Brocade_SetVendor(p, 1, a)
}

func BrocadeAuthRole_Del(p *radius.Packet) {
	_Brocade_DelVendor(p, 1)
}

func BrocadeAVPairs1_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Brocade_AddVendor(p, 2, a)
}

func BrocadeAVPairs1_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Brocade_AddVendor(p, 2, a)
}

func BrocadeAVPairs1_Get(p *radius.Packet) (value []byte) {
	value, _ = BrocadeAVPairs1_Lookup(p)
	return
}

func BrocadeAVPairs1_GetString(p *radius.Packet) (value string) {
	value, _ = BrocadeAVPairs1_LookupString(p)
	return
}

func BrocadeAVPairs1_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Brocade_GetsVendor(p, 2) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func BrocadeAVPairs1_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Brocade_GetsVendor(p, 2) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func BrocadeAVPairs1_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Brocade_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func BrocadeAVPairs1_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Brocade_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func BrocadeAVPairs1_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Brocade_SetVendor(p, 2, a)
}

func BrocadeAVPairs1_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Brocade_SetVendor(p, 2, a)
}

func BrocadeAVPairs1_Del(p *radius.Packet) {
	_Brocade_DelVendor(p, 2)
}

func BrocadeAVPairs2_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Brocade_AddVendor(p, 3, a)
}

func BrocadeAVPairs2_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Brocade_AddVendor(p, 3, a)
}

func BrocadeAVPairs2_Get(p *radius.Packet) (value []byte) {
	value, _ = BrocadeAVPairs2_Lookup(p)
	return
}

func BrocadeAVPairs2_GetString(p *radius.Packet) (value string) {
	value, _ = BrocadeAVPairs2_LookupString(p)
	return
}

func BrocadeAVPairs2_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Brocade_GetsVendor(p, 3) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func BrocadeAVPairs2_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Brocade_GetsVendor(p, 3) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func BrocadeAVPairs2_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Brocade_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func BrocadeAVPairs2_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Brocade_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func BrocadeAVPairs2_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Brocade_SetVendor(p, 3, a)
}

func BrocadeAVPairs2_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Brocade_SetVendor(p, 3, a)
}

func BrocadeAVPairs2_Del(p *radius.Packet) {
	_Brocade_DelVendor(p, 3)
}

func BrocadeAVPairs3_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Brocade_AddVendor(p, 4, a)
}

func BrocadeAVPairs3_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Brocade_AddVendor(p, 4, a)
}

func BrocadeAVPairs3_Get(p *radius.Packet) (value []byte) {
	value, _ = BrocadeAVPairs3_Lookup(p)
	return
}

func BrocadeAVPairs3_GetString(p *radius.Packet) (value string) {
	value, _ = BrocadeAVPairs3_LookupString(p)
	return
}

func BrocadeAVPairs3_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Brocade_GetsVendor(p, 4) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func BrocadeAVPairs3_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Brocade_GetsVendor(p, 4) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func BrocadeAVPairs3_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Brocade_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func BrocadeAVPairs3_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Brocade_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func BrocadeAVPairs3_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Brocade_SetVendor(p, 4, a)
}

func BrocadeAVPairs3_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Brocade_SetVendor(p, 4, a)
}

func BrocadeAVPairs3_Del(p *radius.Packet) {
	_Brocade_DelVendor(p, 4)
}

func BrocadeAVPairs4_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Brocade_AddVendor(p, 5, a)
}

func BrocadeAVPairs4_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Brocade_AddVendor(p, 5, a)
}

func BrocadeAVPairs4_Get(p *radius.Packet) (value []byte) {
	value, _ = BrocadeAVPairs4_Lookup(p)
	return
}

func BrocadeAVPairs4_GetString(p *radius.Packet) (value string) {
	value, _ = BrocadeAVPairs4_LookupString(p)
	return
}

func BrocadeAVPairs4_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Brocade_GetsVendor(p, 5) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func BrocadeAVPairs4_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Brocade_GetsVendor(p, 5) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func BrocadeAVPairs4_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Brocade_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func BrocadeAVPairs4_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Brocade_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func BrocadeAVPairs4_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Brocade_SetVendor(p, 5, a)
}

func BrocadeAVPairs4_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Brocade_SetVendor(p, 5, a)
}

func BrocadeAVPairs4_Del(p *radius.Packet) {
	_Brocade_DelVendor(p, 5)
}

func BrocadePasswdExpiryDate_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Brocade_AddVendor(p, 6, a)
}

func BrocadePasswdExpiryDate_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Brocade_AddVendor(p, 6, a)
}

func BrocadePasswdExpiryDate_Get(p *radius.Packet) (value []byte) {
	value, _ = BrocadePasswdExpiryDate_Lookup(p)
	return
}

func BrocadePasswdExpiryDate_GetString(p *radius.Packet) (value string) {
	value, _ = BrocadePasswdExpiryDate_LookupString(p)
	return
}

func BrocadePasswdExpiryDate_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Brocade_GetsVendor(p, 6) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func BrocadePasswdExpiryDate_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Brocade_GetsVendor(p, 6) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func BrocadePasswdExpiryDate_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Brocade_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func BrocadePasswdExpiryDate_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Brocade_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func BrocadePasswdExpiryDate_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Brocade_SetVendor(p, 6, a)
}

func BrocadePasswdExpiryDate_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Brocade_SetVendor(p, 6, a)
}

func BrocadePasswdExpiryDate_Del(p *radius.Packet) {
	_Brocade_DelVendor(p, 6)
}

func BrocadePasswdWarnPeriod_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Brocade_AddVendor(p, 7, a)
}

func BrocadePasswdWarnPeriod_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Brocade_AddVendor(p, 7, a)
}

func BrocadePasswdWarnPeriod_Get(p *radius.Packet) (value []byte) {
	value, _ = BrocadePasswdWarnPeriod_Lookup(p)
	return
}

func BrocadePasswdWarnPeriod_GetString(p *radius.Packet) (value string) {
	value, _ = BrocadePasswdWarnPeriod_LookupString(p)
	return
}

func BrocadePasswdWarnPeriod_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Brocade_GetsVendor(p, 7) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func BrocadePasswdWarnPeriod_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Brocade_GetsVendor(p, 7) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func BrocadePasswdWarnPeriod_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Brocade_LookupVendor(p, 7)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func BrocadePasswdWarnPeriod_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Brocade_LookupVendor(p, 7)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func BrocadePasswdWarnPeriod_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Brocade_SetVendor(p, 7, a)
}

func BrocadePasswdWarnPeriod_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Brocade_SetVendor(p, 7, a)
}

func BrocadePasswdWarnPeriod_Del(p *radius.Packet) {
	_Brocade_DelVendor(p, 7)
}
