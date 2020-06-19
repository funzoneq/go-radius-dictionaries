// Code generated by radius-dict-gen. DO NOT EDIT.

package trapeze

import (
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Trapeze_VendorID = 14525
)

func _Trapeze_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Trapeze_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Trapeze_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Trapeze_VendorID {
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

func _Trapeze_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Trapeze_VendorID {
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

func _Trapeze_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Trapeze_VendorID {
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
	return _Trapeze_AddVendor(p, typ, attr)
}

func _Trapeze_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Trapeze_VendorID {
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

func TrapezeVLANName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 1, a)
}

func TrapezeVLANName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 1, a)
}

func TrapezeVLANName_Get(p *radius.Packet) (value []byte) {
	value, _ = TrapezeVLANName_Lookup(p)
	return
}

func TrapezeVLANName_GetString(p *radius.Packet) (value string) {
	value, _ = TrapezeVLANName_LookupString(p)
	return
}

func TrapezeVLANName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Trapeze_GetsVendor(p, 1) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeVLANName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Trapeze_GetsVendor(p, 1) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeVLANName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Trapeze_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func TrapezeVLANName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Trapeze_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func TrapezeVLANName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 1, a)
}

func TrapezeVLANName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 1, a)
}

func TrapezeVLANName_Del(p *radius.Packet) {
	_Trapeze_DelVendor(p, 1)
}

func TrapezeMobilityProfile_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 2, a)
}

func TrapezeMobilityProfile_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 2, a)
}

func TrapezeMobilityProfile_Get(p *radius.Packet) (value []byte) {
	value, _ = TrapezeMobilityProfile_Lookup(p)
	return
}

func TrapezeMobilityProfile_GetString(p *radius.Packet) (value string) {
	value, _ = TrapezeMobilityProfile_LookupString(p)
	return
}

func TrapezeMobilityProfile_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Trapeze_GetsVendor(p, 2) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeMobilityProfile_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Trapeze_GetsVendor(p, 2) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeMobilityProfile_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Trapeze_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func TrapezeMobilityProfile_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Trapeze_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func TrapezeMobilityProfile_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 2, a)
}

func TrapezeMobilityProfile_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 2, a)
}

func TrapezeMobilityProfile_Del(p *radius.Packet) {
	_Trapeze_DelVendor(p, 2)
}

func TrapezeEncryptionType_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 3, a)
}

func TrapezeEncryptionType_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 3, a)
}

func TrapezeEncryptionType_Get(p *radius.Packet) (value []byte) {
	value, _ = TrapezeEncryptionType_Lookup(p)
	return
}

func TrapezeEncryptionType_GetString(p *radius.Packet) (value string) {
	value, _ = TrapezeEncryptionType_LookupString(p)
	return
}

func TrapezeEncryptionType_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Trapeze_GetsVendor(p, 3) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeEncryptionType_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Trapeze_GetsVendor(p, 3) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeEncryptionType_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Trapeze_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func TrapezeEncryptionType_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Trapeze_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func TrapezeEncryptionType_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 3, a)
}

func TrapezeEncryptionType_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 3, a)
}

func TrapezeEncryptionType_Del(p *radius.Packet) {
	_Trapeze_DelVendor(p, 3)
}

func TrapezeTimeOfDay_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 4, a)
}

func TrapezeTimeOfDay_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 4, a)
}

func TrapezeTimeOfDay_Get(p *radius.Packet) (value []byte) {
	value, _ = TrapezeTimeOfDay_Lookup(p)
	return
}

func TrapezeTimeOfDay_GetString(p *radius.Packet) (value string) {
	value, _ = TrapezeTimeOfDay_LookupString(p)
	return
}

func TrapezeTimeOfDay_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Trapeze_GetsVendor(p, 4) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeTimeOfDay_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Trapeze_GetsVendor(p, 4) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeTimeOfDay_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Trapeze_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func TrapezeTimeOfDay_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Trapeze_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func TrapezeTimeOfDay_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 4, a)
}

func TrapezeTimeOfDay_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 4, a)
}

func TrapezeTimeOfDay_Del(p *radius.Packet) {
	_Trapeze_DelVendor(p, 4)
}

func TrapezeSSID_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 5, a)
}

func TrapezeSSID_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 5, a)
}

func TrapezeSSID_Get(p *radius.Packet) (value []byte) {
	value, _ = TrapezeSSID_Lookup(p)
	return
}

func TrapezeSSID_GetString(p *radius.Packet) (value string) {
	value, _ = TrapezeSSID_LookupString(p)
	return
}

func TrapezeSSID_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Trapeze_GetsVendor(p, 5) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeSSID_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Trapeze_GetsVendor(p, 5) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeSSID_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Trapeze_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func TrapezeSSID_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Trapeze_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func TrapezeSSID_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 5, a)
}

func TrapezeSSID_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 5, a)
}

func TrapezeSSID_Del(p *radius.Packet) {
	_Trapeze_DelVendor(p, 5)
}

func TrapezeEndDate_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 6, a)
}

func TrapezeEndDate_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 6, a)
}

func TrapezeEndDate_Get(p *radius.Packet) (value []byte) {
	value, _ = TrapezeEndDate_Lookup(p)
	return
}

func TrapezeEndDate_GetString(p *radius.Packet) (value string) {
	value, _ = TrapezeEndDate_LookupString(p)
	return
}

func TrapezeEndDate_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Trapeze_GetsVendor(p, 6) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeEndDate_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Trapeze_GetsVendor(p, 6) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeEndDate_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Trapeze_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func TrapezeEndDate_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Trapeze_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func TrapezeEndDate_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 6, a)
}

func TrapezeEndDate_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 6, a)
}

func TrapezeEndDate_Del(p *radius.Packet) {
	_Trapeze_DelVendor(p, 6)
}

func TrapezeStartDate_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 7, a)
}

func TrapezeStartDate_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 7, a)
}

func TrapezeStartDate_Get(p *radius.Packet) (value []byte) {
	value, _ = TrapezeStartDate_Lookup(p)
	return
}

func TrapezeStartDate_GetString(p *radius.Packet) (value string) {
	value, _ = TrapezeStartDate_LookupString(p)
	return
}

func TrapezeStartDate_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Trapeze_GetsVendor(p, 7) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeStartDate_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Trapeze_GetsVendor(p, 7) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeStartDate_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Trapeze_LookupVendor(p, 7)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func TrapezeStartDate_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Trapeze_LookupVendor(p, 7)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func TrapezeStartDate_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 7, a)
}

func TrapezeStartDate_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 7, a)
}

func TrapezeStartDate_Del(p *radius.Packet) {
	_Trapeze_DelVendor(p, 7)
}

func TrapezeURL_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 8, a)
}

func TrapezeURL_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 8, a)
}

func TrapezeURL_Get(p *radius.Packet) (value []byte) {
	value, _ = TrapezeURL_Lookup(p)
	return
}

func TrapezeURL_GetString(p *radius.Packet) (value string) {
	value, _ = TrapezeURL_LookupString(p)
	return
}

func TrapezeURL_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Trapeze_GetsVendor(p, 8) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeURL_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Trapeze_GetsVendor(p, 8) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeURL_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Trapeze_LookupVendor(p, 8)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func TrapezeURL_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Trapeze_LookupVendor(p, 8)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func TrapezeURL_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 8, a)
}

func TrapezeURL_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 8, a)
}

func TrapezeURL_Del(p *radius.Packet) {
	_Trapeze_DelVendor(p, 8)
}

func TrapezeUserGroupName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 9, a)
}

func TrapezeUserGroupName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 9, a)
}

func TrapezeUserGroupName_Get(p *radius.Packet) (value []byte) {
	value, _ = TrapezeUserGroupName_Lookup(p)
	return
}

func TrapezeUserGroupName_GetString(p *radius.Packet) (value string) {
	value, _ = TrapezeUserGroupName_LookupString(p)
	return
}

func TrapezeUserGroupName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Trapeze_GetsVendor(p, 9) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeUserGroupName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Trapeze_GetsVendor(p, 9) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeUserGroupName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Trapeze_LookupVendor(p, 9)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func TrapezeUserGroupName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Trapeze_LookupVendor(p, 9)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func TrapezeUserGroupName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 9, a)
}

func TrapezeUserGroupName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 9, a)
}

func TrapezeUserGroupName_Del(p *radius.Packet) {
	_Trapeze_DelVendor(p, 9)
}

func TrapezeQoSProfile_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 10, a)
}

func TrapezeQoSProfile_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 10, a)
}

func TrapezeQoSProfile_Get(p *radius.Packet) (value []byte) {
	value, _ = TrapezeQoSProfile_Lookup(p)
	return
}

func TrapezeQoSProfile_GetString(p *radius.Packet) (value string) {
	value, _ = TrapezeQoSProfile_LookupString(p)
	return
}

func TrapezeQoSProfile_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Trapeze_GetsVendor(p, 10) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeQoSProfile_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Trapeze_GetsVendor(p, 10) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeQoSProfile_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Trapeze_LookupVendor(p, 10)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func TrapezeQoSProfile_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Trapeze_LookupVendor(p, 10)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func TrapezeQoSProfile_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 10, a)
}

func TrapezeQoSProfile_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 10, a)
}

func TrapezeQoSProfile_Del(p *radius.Packet) {
	_Trapeze_DelVendor(p, 10)
}

func TrapezeSimultaneousLogins_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 11, a)
}

func TrapezeSimultaneousLogins_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 11, a)
}

func TrapezeSimultaneousLogins_Get(p *radius.Packet) (value []byte) {
	value, _ = TrapezeSimultaneousLogins_Lookup(p)
	return
}

func TrapezeSimultaneousLogins_GetString(p *radius.Packet) (value string) {
	value, _ = TrapezeSimultaneousLogins_LookupString(p)
	return
}

func TrapezeSimultaneousLogins_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Trapeze_GetsVendor(p, 11) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeSimultaneousLogins_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Trapeze_GetsVendor(p, 11) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeSimultaneousLogins_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Trapeze_LookupVendor(p, 11)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func TrapezeSimultaneousLogins_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Trapeze_LookupVendor(p, 11)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func TrapezeSimultaneousLogins_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 11, a)
}

func TrapezeSimultaneousLogins_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 11, a)
}

func TrapezeSimultaneousLogins_Del(p *radius.Packet) {
	_Trapeze_DelVendor(p, 11)
}

func TrapezeCoAUsername_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 12, a)
}

func TrapezeCoAUsername_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 12, a)
}

func TrapezeCoAUsername_Get(p *radius.Packet) (value []byte) {
	value, _ = TrapezeCoAUsername_Lookup(p)
	return
}

func TrapezeCoAUsername_GetString(p *radius.Packet) (value string) {
	value, _ = TrapezeCoAUsername_LookupString(p)
	return
}

func TrapezeCoAUsername_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Trapeze_GetsVendor(p, 12) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeCoAUsername_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Trapeze_GetsVendor(p, 12) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeCoAUsername_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Trapeze_LookupVendor(p, 12)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func TrapezeCoAUsername_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Trapeze_LookupVendor(p, 12)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func TrapezeCoAUsername_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 12, a)
}

func TrapezeCoAUsername_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 12, a)
}

func TrapezeCoAUsername_Del(p *radius.Packet) {
	_Trapeze_DelVendor(p, 12)
}

func TrapezeAudit_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 13, a)
}

func TrapezeAudit_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_AddVendor(p, 13, a)
}

func TrapezeAudit_Get(p *radius.Packet) (value []byte) {
	value, _ = TrapezeAudit_Lookup(p)
	return
}

func TrapezeAudit_GetString(p *radius.Packet) (value string) {
	value, _ = TrapezeAudit_LookupString(p)
	return
}

func TrapezeAudit_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Trapeze_GetsVendor(p, 13) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeAudit_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Trapeze_GetsVendor(p, 13) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TrapezeAudit_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Trapeze_LookupVendor(p, 13)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func TrapezeAudit_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Trapeze_LookupVendor(p, 13)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func TrapezeAudit_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 13, a)
}

func TrapezeAudit_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Trapeze_SetVendor(p, 13, a)
}

func TrapezeAudit_Del(p *radius.Packet) {
	_Trapeze_DelVendor(p, 13)
}
