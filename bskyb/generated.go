// Code generated by radius-dict-gen. DO NOT EDIT.

package bskyb

import (
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_BSkyB_VendorID = 16924
)

func _BSkyB_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_BSkyB_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _BSkyB_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, attr := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _BSkyB_VendorID {
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

func _BSkyB_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, a := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(a)
		if err != nil || vendorID != _BSkyB_VendorID {
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

func _BSkyB_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		vendorID, vsa, err := radius.VendorSpecific(p.Attributes[rfc2865.VendorSpecific_Type][i])
		if err != nil || vendorID != _BSkyB_VendorID {
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
	return _BSkyB_AddVendor(p, typ, attr)
}

func _BSkyB_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		attr := p.Attributes[rfc2865.VendorSpecific_Type][i]
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _BSkyB_VendorID {
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

type SkyWifiAPID uint32

var SkyWifiAPID_Strings = map[SkyWifiAPID]string{}

func (a SkyWifiAPID) String() string {
	if str, ok := SkyWifiAPID_Strings[a]; ok {
		return str
	}
	return "SkyWifiAPID(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func SkyWifiAPID_Add(p *radius.Packet, value SkyWifiAPID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _BSkyB_AddVendor(p, 1, a)
}

func SkyWifiAPID_Get(p *radius.Packet) (value SkyWifiAPID) {
	value, _ = SkyWifiAPID_Lookup(p)
	return
}

func SkyWifiAPID_Gets(p *radius.Packet) (values []SkyWifiAPID, err error) {
	var i uint32
	for _, attr := range _BSkyB_GetsVendor(p, 1) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, SkyWifiAPID(i))
	}
	return
}

func SkyWifiAPID_Lookup(p *radius.Packet) (value SkyWifiAPID, err error) {
	a, ok := _BSkyB_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = SkyWifiAPID(i)
	return
}

func SkyWifiAPID_Set(p *radius.Packet, value SkyWifiAPID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _BSkyB_SetVendor(p, 1, a)
}

func SkyWifiAPID_Del(p *radius.Packet) {
	_BSkyB_DelVendor(p, 1)
}

type SkyWifiServiceID uint32

var SkyWifiServiceID_Strings = map[SkyWifiServiceID]string{}

func (a SkyWifiServiceID) String() string {
	if str, ok := SkyWifiServiceID_Strings[a]; ok {
		return str
	}
	return "SkyWifiServiceID(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func SkyWifiServiceID_Add(p *radius.Packet, value SkyWifiServiceID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _BSkyB_AddVendor(p, 2, a)
}

func SkyWifiServiceID_Get(p *radius.Packet) (value SkyWifiServiceID) {
	value, _ = SkyWifiServiceID_Lookup(p)
	return
}

func SkyWifiServiceID_Gets(p *radius.Packet) (values []SkyWifiServiceID, err error) {
	var i uint32
	for _, attr := range _BSkyB_GetsVendor(p, 2) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, SkyWifiServiceID(i))
	}
	return
}

func SkyWifiServiceID_Lookup(p *radius.Packet) (value SkyWifiServiceID, err error) {
	a, ok := _BSkyB_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = SkyWifiServiceID(i)
	return
}

func SkyWifiServiceID_Set(p *radius.Packet, value SkyWifiServiceID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _BSkyB_SetVendor(p, 2, a)
}

func SkyWifiServiceID_Del(p *radius.Packet) {
	_BSkyB_DelVendor(p, 2)
}

func SkyWifiFilterProfile_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _BSkyB_AddVendor(p, 3, a)
}

func SkyWifiFilterProfile_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _BSkyB_AddVendor(p, 3, a)
}

func SkyWifiFilterProfile_Get(p *radius.Packet) (value []byte) {
	value, _ = SkyWifiFilterProfile_Lookup(p)
	return
}

func SkyWifiFilterProfile_GetString(p *radius.Packet) (value string) {
	value, _ = SkyWifiFilterProfile_LookupString(p)
	return
}

func SkyWifiFilterProfile_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _BSkyB_GetsVendor(p, 3) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SkyWifiFilterProfile_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _BSkyB_GetsVendor(p, 3) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SkyWifiFilterProfile_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _BSkyB_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func SkyWifiFilterProfile_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _BSkyB_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func SkyWifiFilterProfile_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _BSkyB_SetVendor(p, 3, a)
}

func SkyWifiFilterProfile_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _BSkyB_SetVendor(p, 3, a)
}

func SkyWifiFilterProfile_Del(p *radius.Packet) {
	_BSkyB_DelVendor(p, 3)
}

func SkyWifiBillingClass_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _BSkyB_AddVendor(p, 4, a)
}

func SkyWifiBillingClass_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _BSkyB_AddVendor(p, 4, a)
}

func SkyWifiBillingClass_Get(p *radius.Packet) (value []byte) {
	value, _ = SkyWifiBillingClass_Lookup(p)
	return
}

func SkyWifiBillingClass_GetString(p *radius.Packet) (value string) {
	value, _ = SkyWifiBillingClass_LookupString(p)
	return
}

func SkyWifiBillingClass_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _BSkyB_GetsVendor(p, 4) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SkyWifiBillingClass_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _BSkyB_GetsVendor(p, 4) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SkyWifiBillingClass_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _BSkyB_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func SkyWifiBillingClass_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _BSkyB_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func SkyWifiBillingClass_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _BSkyB_SetVendor(p, 4, a)
}

func SkyWifiBillingClass_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _BSkyB_SetVendor(p, 4, a)
}

func SkyWifiBillingClass_Del(p *radius.Packet) {
	_BSkyB_DelVendor(p, 4)
}

type SkyWifiProviderID uint32

var SkyWifiProviderID_Strings = map[SkyWifiProviderID]string{}

func (a SkyWifiProviderID) String() string {
	if str, ok := SkyWifiProviderID_Strings[a]; ok {
		return str
	}
	return "SkyWifiProviderID(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func SkyWifiProviderID_Add(p *radius.Packet, value SkyWifiProviderID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _BSkyB_AddVendor(p, 5, a)
}

func SkyWifiProviderID_Get(p *radius.Packet) (value SkyWifiProviderID) {
	value, _ = SkyWifiProviderID_Lookup(p)
	return
}

func SkyWifiProviderID_Gets(p *radius.Packet) (values []SkyWifiProviderID, err error) {
	var i uint32
	for _, attr := range _BSkyB_GetsVendor(p, 5) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, SkyWifiProviderID(i))
	}
	return
}

func SkyWifiProviderID_Lookup(p *radius.Packet) (value SkyWifiProviderID, err error) {
	a, ok := _BSkyB_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = SkyWifiProviderID(i)
	return
}

func SkyWifiProviderID_Set(p *radius.Packet, value SkyWifiProviderID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _BSkyB_SetVendor(p, 5, a)
}

func SkyWifiProviderID_Del(p *radius.Packet) {
	_BSkyB_DelVendor(p, 5)
}

func SkyWifiCredentials_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _BSkyB_AddVendor(p, 6, a)
}

func SkyWifiCredentials_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _BSkyB_AddVendor(p, 6, a)
}

func SkyWifiCredentials_Get(p *radius.Packet) (value []byte) {
	value, _ = SkyWifiCredentials_Lookup(p)
	return
}

func SkyWifiCredentials_GetString(p *radius.Packet) (value string) {
	value, _ = SkyWifiCredentials_LookupString(p)
	return
}

func SkyWifiCredentials_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _BSkyB_GetsVendor(p, 6) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SkyWifiCredentials_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _BSkyB_GetsVendor(p, 6) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SkyWifiCredentials_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _BSkyB_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func SkyWifiCredentials_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _BSkyB_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func SkyWifiCredentials_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _BSkyB_SetVendor(p, 6, a)
}

func SkyWifiCredentials_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _BSkyB_SetVendor(p, 6, a)
}

func SkyWifiCredentials_Del(p *radius.Packet) {
	_BSkyB_DelVendor(p, 6)
}