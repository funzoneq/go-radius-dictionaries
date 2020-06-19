// Code generated by radius-dict-gen. DO NOT EDIT.

package nomadix

import (
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Nomadix_VendorID = 3309
)

func _Nomadix_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Nomadix_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Nomadix_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Nomadix_VendorID {
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

func _Nomadix_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Nomadix_VendorID {
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

func _Nomadix_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Nomadix_VendorID {
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
	return _Nomadix_AddVendor(p, typ, attr)
}

func _Nomadix_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Nomadix_VendorID {
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

type NomadixBwUp uint32

var NomadixBwUp_Strings = map[NomadixBwUp]string{}

func (a NomadixBwUp) String() string {
	if str, ok := NomadixBwUp_Strings[a]; ok {
		return str
	}
	return "NomadixBwUp(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func NomadixBwUp_Add(p *radius.Packet, value NomadixBwUp) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_AddVendor(p, 1, a)
}

func NomadixBwUp_Get(p *radius.Packet) (value NomadixBwUp) {
	value, _ = NomadixBwUp_Lookup(p)
	return
}

func NomadixBwUp_Gets(p *radius.Packet) (values []NomadixBwUp, err error) {
	var i uint32
	for _, attr := range _Nomadix_GetsVendor(p, 1) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, NomadixBwUp(i))
	}
	return
}

func NomadixBwUp_Lookup(p *radius.Packet) (value NomadixBwUp, err error) {
	a, ok := _Nomadix_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = NomadixBwUp(i)
	return
}

func NomadixBwUp_Set(p *radius.Packet, value NomadixBwUp) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_SetVendor(p, 1, a)
}

func NomadixBwUp_Del(p *radius.Packet) {
	_Nomadix_DelVendor(p, 1)
}

type NomadixBwDown uint32

var NomadixBwDown_Strings = map[NomadixBwDown]string{}

func (a NomadixBwDown) String() string {
	if str, ok := NomadixBwDown_Strings[a]; ok {
		return str
	}
	return "NomadixBwDown(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func NomadixBwDown_Add(p *radius.Packet, value NomadixBwDown) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_AddVendor(p, 2, a)
}

func NomadixBwDown_Get(p *radius.Packet) (value NomadixBwDown) {
	value, _ = NomadixBwDown_Lookup(p)
	return
}

func NomadixBwDown_Gets(p *radius.Packet) (values []NomadixBwDown, err error) {
	var i uint32
	for _, attr := range _Nomadix_GetsVendor(p, 2) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, NomadixBwDown(i))
	}
	return
}

func NomadixBwDown_Lookup(p *radius.Packet) (value NomadixBwDown, err error) {
	a, ok := _Nomadix_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = NomadixBwDown(i)
	return
}

func NomadixBwDown_Set(p *radius.Packet, value NomadixBwDown) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_SetVendor(p, 2, a)
}

func NomadixBwDown_Del(p *radius.Packet) {
	_Nomadix_DelVendor(p, 2)
}

func NomadixURLRedirection_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Nomadix_AddVendor(p, 3, a)
}

func NomadixURLRedirection_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Nomadix_AddVendor(p, 3, a)
}

func NomadixURLRedirection_Get(p *radius.Packet) (value []byte) {
	value, _ = NomadixURLRedirection_Lookup(p)
	return
}

func NomadixURLRedirection_GetString(p *radius.Packet) (value string) {
	value, _ = NomadixURLRedirection_LookupString(p)
	return
}

func NomadixURLRedirection_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Nomadix_GetsVendor(p, 3) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NomadixURLRedirection_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Nomadix_GetsVendor(p, 3) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NomadixURLRedirection_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Nomadix_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func NomadixURLRedirection_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Nomadix_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func NomadixURLRedirection_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Nomadix_SetVendor(p, 3, a)
}

func NomadixURLRedirection_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Nomadix_SetVendor(p, 3, a)
}

func NomadixURLRedirection_Del(p *radius.Packet) {
	_Nomadix_DelVendor(p, 3)
}

type NomadixIPUpsell uint32

const (
	NomadixIPUpsell_Value_PrivatePool NomadixIPUpsell = 0
	NomadixIPUpsell_Value_PublicPool  NomadixIPUpsell = 1
)

var NomadixIPUpsell_Strings = map[NomadixIPUpsell]string{
	NomadixIPUpsell_Value_PrivatePool: "PrivatePool",
	NomadixIPUpsell_Value_PublicPool:  "PublicPool",
}

func (a NomadixIPUpsell) String() string {
	if str, ok := NomadixIPUpsell_Strings[a]; ok {
		return str
	}
	return "NomadixIPUpsell(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func NomadixIPUpsell_Add(p *radius.Packet, value NomadixIPUpsell) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_AddVendor(p, 4, a)
}

func NomadixIPUpsell_Get(p *radius.Packet) (value NomadixIPUpsell) {
	value, _ = NomadixIPUpsell_Lookup(p)
	return
}

func NomadixIPUpsell_Gets(p *radius.Packet) (values []NomadixIPUpsell, err error) {
	var i uint32
	for _, attr := range _Nomadix_GetsVendor(p, 4) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, NomadixIPUpsell(i))
	}
	return
}

func NomadixIPUpsell_Lookup(p *radius.Packet) (value NomadixIPUpsell, err error) {
	a, ok := _Nomadix_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = NomadixIPUpsell(i)
	return
}

func NomadixIPUpsell_Set(p *radius.Packet, value NomadixIPUpsell) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_SetVendor(p, 4, a)
}

func NomadixIPUpsell_Del(p *radius.Packet) {
	_Nomadix_DelVendor(p, 4)
}

func NomadixExpiration_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Nomadix_AddVendor(p, 5, a)
}

func NomadixExpiration_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Nomadix_AddVendor(p, 5, a)
}

func NomadixExpiration_Get(p *radius.Packet) (value []byte) {
	value, _ = NomadixExpiration_Lookup(p)
	return
}

func NomadixExpiration_GetString(p *radius.Packet) (value string) {
	value, _ = NomadixExpiration_LookupString(p)
	return
}

func NomadixExpiration_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Nomadix_GetsVendor(p, 5) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NomadixExpiration_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Nomadix_GetsVendor(p, 5) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NomadixExpiration_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Nomadix_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func NomadixExpiration_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Nomadix_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func NomadixExpiration_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Nomadix_SetVendor(p, 5, a)
}

func NomadixExpiration_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Nomadix_SetVendor(p, 5, a)
}

func NomadixExpiration_Del(p *radius.Packet) {
	_Nomadix_DelVendor(p, 5)
}

func NomadixSubnet_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Nomadix_AddVendor(p, 6, a)
}

func NomadixSubnet_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Nomadix_AddVendor(p, 6, a)
}

func NomadixSubnet_Get(p *radius.Packet) (value []byte) {
	value, _ = NomadixSubnet_Lookup(p)
	return
}

func NomadixSubnet_GetString(p *radius.Packet) (value string) {
	value, _ = NomadixSubnet_LookupString(p)
	return
}

func NomadixSubnet_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Nomadix_GetsVendor(p, 6) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NomadixSubnet_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Nomadix_GetsVendor(p, 6) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NomadixSubnet_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Nomadix_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func NomadixSubnet_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Nomadix_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func NomadixSubnet_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Nomadix_SetVendor(p, 6, a)
}

func NomadixSubnet_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Nomadix_SetVendor(p, 6, a)
}

func NomadixSubnet_Del(p *radius.Packet) {
	_Nomadix_DelVendor(p, 6)
}

type NomadixMaxBytesUp uint32

var NomadixMaxBytesUp_Strings = map[NomadixMaxBytesUp]string{}

func (a NomadixMaxBytesUp) String() string {
	if str, ok := NomadixMaxBytesUp_Strings[a]; ok {
		return str
	}
	return "NomadixMaxBytesUp(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func NomadixMaxBytesUp_Add(p *radius.Packet, value NomadixMaxBytesUp) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_AddVendor(p, 7, a)
}

func NomadixMaxBytesUp_Get(p *radius.Packet) (value NomadixMaxBytesUp) {
	value, _ = NomadixMaxBytesUp_Lookup(p)
	return
}

func NomadixMaxBytesUp_Gets(p *radius.Packet) (values []NomadixMaxBytesUp, err error) {
	var i uint32
	for _, attr := range _Nomadix_GetsVendor(p, 7) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, NomadixMaxBytesUp(i))
	}
	return
}

func NomadixMaxBytesUp_Lookup(p *radius.Packet) (value NomadixMaxBytesUp, err error) {
	a, ok := _Nomadix_LookupVendor(p, 7)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = NomadixMaxBytesUp(i)
	return
}

func NomadixMaxBytesUp_Set(p *radius.Packet, value NomadixMaxBytesUp) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_SetVendor(p, 7, a)
}

func NomadixMaxBytesUp_Del(p *radius.Packet) {
	_Nomadix_DelVendor(p, 7)
}

type NomadixMaxBytesDown uint32

var NomadixMaxBytesDown_Strings = map[NomadixMaxBytesDown]string{}

func (a NomadixMaxBytesDown) String() string {
	if str, ok := NomadixMaxBytesDown_Strings[a]; ok {
		return str
	}
	return "NomadixMaxBytesDown(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func NomadixMaxBytesDown_Add(p *radius.Packet, value NomadixMaxBytesDown) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_AddVendor(p, 8, a)
}

func NomadixMaxBytesDown_Get(p *radius.Packet) (value NomadixMaxBytesDown) {
	value, _ = NomadixMaxBytesDown_Lookup(p)
	return
}

func NomadixMaxBytesDown_Gets(p *radius.Packet) (values []NomadixMaxBytesDown, err error) {
	var i uint32
	for _, attr := range _Nomadix_GetsVendor(p, 8) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, NomadixMaxBytesDown(i))
	}
	return
}

func NomadixMaxBytesDown_Lookup(p *radius.Packet) (value NomadixMaxBytesDown, err error) {
	a, ok := _Nomadix_LookupVendor(p, 8)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = NomadixMaxBytesDown(i)
	return
}

func NomadixMaxBytesDown_Set(p *radius.Packet, value NomadixMaxBytesDown) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_SetVendor(p, 8, a)
}

func NomadixMaxBytesDown_Del(p *radius.Packet) {
	_Nomadix_DelVendor(p, 8)
}

type NomadixEndofSession uint32

var NomadixEndofSession_Strings = map[NomadixEndofSession]string{}

func (a NomadixEndofSession) String() string {
	if str, ok := NomadixEndofSession_Strings[a]; ok {
		return str
	}
	return "NomadixEndofSession(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func NomadixEndofSession_Add(p *radius.Packet, value NomadixEndofSession) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_AddVendor(p, 9, a)
}

func NomadixEndofSession_Get(p *radius.Packet) (value NomadixEndofSession) {
	value, _ = NomadixEndofSession_Lookup(p)
	return
}

func NomadixEndofSession_Gets(p *radius.Packet) (values []NomadixEndofSession, err error) {
	var i uint32
	for _, attr := range _Nomadix_GetsVendor(p, 9) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, NomadixEndofSession(i))
	}
	return
}

func NomadixEndofSession_Lookup(p *radius.Packet) (value NomadixEndofSession, err error) {
	a, ok := _Nomadix_LookupVendor(p, 9)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = NomadixEndofSession(i)
	return
}

func NomadixEndofSession_Set(p *radius.Packet, value NomadixEndofSession) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_SetVendor(p, 9, a)
}

func NomadixEndofSession_Del(p *radius.Packet) {
	_Nomadix_DelVendor(p, 9)
}

func NomadixLogoffURL_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Nomadix_AddVendor(p, 10, a)
}

func NomadixLogoffURL_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Nomadix_AddVendor(p, 10, a)
}

func NomadixLogoffURL_Get(p *radius.Packet) (value []byte) {
	value, _ = NomadixLogoffURL_Lookup(p)
	return
}

func NomadixLogoffURL_GetString(p *radius.Packet) (value string) {
	value, _ = NomadixLogoffURL_LookupString(p)
	return
}

func NomadixLogoffURL_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Nomadix_GetsVendor(p, 10) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NomadixLogoffURL_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Nomadix_GetsVendor(p, 10) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NomadixLogoffURL_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Nomadix_LookupVendor(p, 10)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func NomadixLogoffURL_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Nomadix_LookupVendor(p, 10)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func NomadixLogoffURL_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Nomadix_SetVendor(p, 10, a)
}

func NomadixLogoffURL_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Nomadix_SetVendor(p, 10, a)
}

func NomadixLogoffURL_Del(p *radius.Packet) {
	_Nomadix_DelVendor(p, 10)
}

type NomadixNetVLAN uint32

var NomadixNetVLAN_Strings = map[NomadixNetVLAN]string{}

func (a NomadixNetVLAN) String() string {
	if str, ok := NomadixNetVLAN_Strings[a]; ok {
		return str
	}
	return "NomadixNetVLAN(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func NomadixNetVLAN_Add(p *radius.Packet, value NomadixNetVLAN) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_AddVendor(p, 11, a)
}

func NomadixNetVLAN_Get(p *radius.Packet) (value NomadixNetVLAN) {
	value, _ = NomadixNetVLAN_Lookup(p)
	return
}

func NomadixNetVLAN_Gets(p *radius.Packet) (values []NomadixNetVLAN, err error) {
	var i uint32
	for _, attr := range _Nomadix_GetsVendor(p, 11) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, NomadixNetVLAN(i))
	}
	return
}

func NomadixNetVLAN_Lookup(p *radius.Packet) (value NomadixNetVLAN, err error) {
	a, ok := _Nomadix_LookupVendor(p, 11)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = NomadixNetVLAN(i)
	return
}

func NomadixNetVLAN_Set(p *radius.Packet, value NomadixNetVLAN) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_SetVendor(p, 11, a)
}

func NomadixNetVLAN_Del(p *radius.Packet) {
	_Nomadix_DelVendor(p, 11)
}

func NomadixConfigURL_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Nomadix_AddVendor(p, 12, a)
}

func NomadixConfigURL_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Nomadix_AddVendor(p, 12, a)
}

func NomadixConfigURL_Get(p *radius.Packet) (value []byte) {
	value, _ = NomadixConfigURL_Lookup(p)
	return
}

func NomadixConfigURL_GetString(p *radius.Packet) (value string) {
	value, _ = NomadixConfigURL_LookupString(p)
	return
}

func NomadixConfigURL_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Nomadix_GetsVendor(p, 12) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NomadixConfigURL_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Nomadix_GetsVendor(p, 12) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NomadixConfigURL_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Nomadix_LookupVendor(p, 12)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func NomadixConfigURL_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Nomadix_LookupVendor(p, 12)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func NomadixConfigURL_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Nomadix_SetVendor(p, 12, a)
}

func NomadixConfigURL_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Nomadix_SetVendor(p, 12, a)
}

func NomadixConfigURL_Del(p *radius.Packet) {
	_Nomadix_DelVendor(p, 12)
}

func NomadixGoodbyeURL_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Nomadix_AddVendor(p, 13, a)
}

func NomadixGoodbyeURL_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Nomadix_AddVendor(p, 13, a)
}

func NomadixGoodbyeURL_Get(p *radius.Packet) (value []byte) {
	value, _ = NomadixGoodbyeURL_Lookup(p)
	return
}

func NomadixGoodbyeURL_GetString(p *radius.Packet) (value string) {
	value, _ = NomadixGoodbyeURL_LookupString(p)
	return
}

func NomadixGoodbyeURL_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Nomadix_GetsVendor(p, 13) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NomadixGoodbyeURL_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Nomadix_GetsVendor(p, 13) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NomadixGoodbyeURL_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Nomadix_LookupVendor(p, 13)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func NomadixGoodbyeURL_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Nomadix_LookupVendor(p, 13)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func NomadixGoodbyeURL_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Nomadix_SetVendor(p, 13, a)
}

func NomadixGoodbyeURL_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Nomadix_SetVendor(p, 13, a)
}

func NomadixGoodbyeURL_Del(p *radius.Packet) {
	_Nomadix_DelVendor(p, 13)
}

func NomadixQosPolicy_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Nomadix_AddVendor(p, 14, a)
}

func NomadixQosPolicy_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Nomadix_AddVendor(p, 14, a)
}

func NomadixQosPolicy_Get(p *radius.Packet) (value []byte) {
	value, _ = NomadixQosPolicy_Lookup(p)
	return
}

func NomadixQosPolicy_GetString(p *radius.Packet) (value string) {
	value, _ = NomadixQosPolicy_LookupString(p)
	return
}

func NomadixQosPolicy_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Nomadix_GetsVendor(p, 14) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NomadixQosPolicy_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Nomadix_GetsVendor(p, 14) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NomadixQosPolicy_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Nomadix_LookupVendor(p, 14)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func NomadixQosPolicy_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Nomadix_LookupVendor(p, 14)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func NomadixQosPolicy_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Nomadix_SetVendor(p, 14, a)
}

func NomadixQosPolicy_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Nomadix_SetVendor(p, 14, a)
}

func NomadixQosPolicy_Del(p *radius.Packet) {
	_Nomadix_DelVendor(p, 14)
}

type NomadixSMTPRedirect uint32

var NomadixSMTPRedirect_Strings = map[NomadixSMTPRedirect]string{}

func (a NomadixSMTPRedirect) String() string {
	if str, ok := NomadixSMTPRedirect_Strings[a]; ok {
		return str
	}
	return "NomadixSMTPRedirect(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func NomadixSMTPRedirect_Add(p *radius.Packet, value NomadixSMTPRedirect) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_AddVendor(p, 17, a)
}

func NomadixSMTPRedirect_Get(p *radius.Packet) (value NomadixSMTPRedirect) {
	value, _ = NomadixSMTPRedirect_Lookup(p)
	return
}

func NomadixSMTPRedirect_Gets(p *radius.Packet) (values []NomadixSMTPRedirect, err error) {
	var i uint32
	for _, attr := range _Nomadix_GetsVendor(p, 17) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, NomadixSMTPRedirect(i))
	}
	return
}

func NomadixSMTPRedirect_Lookup(p *radius.Packet) (value NomadixSMTPRedirect, err error) {
	a, ok := _Nomadix_LookupVendor(p, 17)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = NomadixSMTPRedirect(i)
	return
}

func NomadixSMTPRedirect_Set(p *radius.Packet, value NomadixSMTPRedirect) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_SetVendor(p, 17, a)
}

func NomadixSMTPRedirect_Del(p *radius.Packet) {
	_Nomadix_DelVendor(p, 17)
}

func NomadixCentralizedMgmt_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Nomadix_AddVendor(p, 18, a)
}

func NomadixCentralizedMgmt_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Nomadix_AddVendor(p, 18, a)
}

func NomadixCentralizedMgmt_Get(p *radius.Packet) (value []byte) {
	value, _ = NomadixCentralizedMgmt_Lookup(p)
	return
}

func NomadixCentralizedMgmt_GetString(p *radius.Packet) (value string) {
	value, _ = NomadixCentralizedMgmt_LookupString(p)
	return
}

func NomadixCentralizedMgmt_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Nomadix_GetsVendor(p, 18) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NomadixCentralizedMgmt_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Nomadix_GetsVendor(p, 18) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NomadixCentralizedMgmt_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Nomadix_LookupVendor(p, 18)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func NomadixCentralizedMgmt_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Nomadix_LookupVendor(p, 18)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func NomadixCentralizedMgmt_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Nomadix_SetVendor(p, 18, a)
}

func NomadixCentralizedMgmt_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Nomadix_SetVendor(p, 18, a)
}

func NomadixCentralizedMgmt_Del(p *radius.Packet) {
	_Nomadix_DelVendor(p, 18)
}

type NomadixGroupPolicyID uint32

var NomadixGroupPolicyID_Strings = map[NomadixGroupPolicyID]string{}

func (a NomadixGroupPolicyID) String() string {
	if str, ok := NomadixGroupPolicyID_Strings[a]; ok {
		return str
	}
	return "NomadixGroupPolicyID(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func NomadixGroupPolicyID_Add(p *radius.Packet, value NomadixGroupPolicyID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_AddVendor(p, 19, a)
}

func NomadixGroupPolicyID_Get(p *radius.Packet) (value NomadixGroupPolicyID) {
	value, _ = NomadixGroupPolicyID_Lookup(p)
	return
}

func NomadixGroupPolicyID_Gets(p *radius.Packet) (values []NomadixGroupPolicyID, err error) {
	var i uint32
	for _, attr := range _Nomadix_GetsVendor(p, 19) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, NomadixGroupPolicyID(i))
	}
	return
}

func NomadixGroupPolicyID_Lookup(p *radius.Packet) (value NomadixGroupPolicyID, err error) {
	a, ok := _Nomadix_LookupVendor(p, 19)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = NomadixGroupPolicyID(i)
	return
}

func NomadixGroupPolicyID_Set(p *radius.Packet, value NomadixGroupPolicyID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_SetVendor(p, 19, a)
}

func NomadixGroupPolicyID_Del(p *radius.Packet) {
	_Nomadix_DelVendor(p, 19)
}

type NomadixGroupBwMaxUp uint32

var NomadixGroupBwMaxUp_Strings = map[NomadixGroupBwMaxUp]string{}

func (a NomadixGroupBwMaxUp) String() string {
	if str, ok := NomadixGroupBwMaxUp_Strings[a]; ok {
		return str
	}
	return "NomadixGroupBwMaxUp(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func NomadixGroupBwMaxUp_Add(p *radius.Packet, value NomadixGroupBwMaxUp) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_AddVendor(p, 20, a)
}

func NomadixGroupBwMaxUp_Get(p *radius.Packet) (value NomadixGroupBwMaxUp) {
	value, _ = NomadixGroupBwMaxUp_Lookup(p)
	return
}

func NomadixGroupBwMaxUp_Gets(p *radius.Packet) (values []NomadixGroupBwMaxUp, err error) {
	var i uint32
	for _, attr := range _Nomadix_GetsVendor(p, 20) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, NomadixGroupBwMaxUp(i))
	}
	return
}

func NomadixGroupBwMaxUp_Lookup(p *radius.Packet) (value NomadixGroupBwMaxUp, err error) {
	a, ok := _Nomadix_LookupVendor(p, 20)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = NomadixGroupBwMaxUp(i)
	return
}

func NomadixGroupBwMaxUp_Set(p *radius.Packet, value NomadixGroupBwMaxUp) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_SetVendor(p, 20, a)
}

func NomadixGroupBwMaxUp_Del(p *radius.Packet) {
	_Nomadix_DelVendor(p, 20)
}

type NomadixGroupBwMaxDown uint32

var NomadixGroupBwMaxDown_Strings = map[NomadixGroupBwMaxDown]string{}

func (a NomadixGroupBwMaxDown) String() string {
	if str, ok := NomadixGroupBwMaxDown_Strings[a]; ok {
		return str
	}
	return "NomadixGroupBwMaxDown(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func NomadixGroupBwMaxDown_Add(p *radius.Packet, value NomadixGroupBwMaxDown) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_AddVendor(p, 21, a)
}

func NomadixGroupBwMaxDown_Get(p *radius.Packet) (value NomadixGroupBwMaxDown) {
	value, _ = NomadixGroupBwMaxDown_Lookup(p)
	return
}

func NomadixGroupBwMaxDown_Gets(p *radius.Packet) (values []NomadixGroupBwMaxDown, err error) {
	var i uint32
	for _, attr := range _Nomadix_GetsVendor(p, 21) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, NomadixGroupBwMaxDown(i))
	}
	return
}

func NomadixGroupBwMaxDown_Lookup(p *radius.Packet) (value NomadixGroupBwMaxDown, err error) {
	a, ok := _Nomadix_LookupVendor(p, 21)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = NomadixGroupBwMaxDown(i)
	return
}

func NomadixGroupBwMaxDown_Set(p *radius.Packet, value NomadixGroupBwMaxDown) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Nomadix_SetVendor(p, 21, a)
}

func NomadixGroupBwMaxDown_Del(p *radius.Packet) {
	_Nomadix_DelVendor(p, 21)
}
