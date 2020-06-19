// Code generated by radius-dict-gen. DO NOT EDIT.

package manzara

import (
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Manzara_VendorID = 19382
)

func _Manzara_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Manzara_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Manzara_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Manzara_VendorID {
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

func _Manzara_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Manzara_VendorID {
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

func _Manzara_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Manzara_VendorID {
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
	return _Manzara_AddVendor(p, typ, attr)
}

func _Manzara_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Manzara_VendorID {
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

type ManzaraUserUID uint32

var ManzaraUserUID_Strings = map[ManzaraUserUID]string{}

func (a ManzaraUserUID) String() string {
	if str, ok := ManzaraUserUID_Strings[a]; ok {
		return str
	}
	return "ManzaraUserUID(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func ManzaraUserUID_Add(p *radius.Packet, value ManzaraUserUID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Manzara_AddVendor(p, 1, a)
}

func ManzaraUserUID_Get(p *radius.Packet) (value ManzaraUserUID) {
	value, _ = ManzaraUserUID_Lookup(p)
	return
}

func ManzaraUserUID_Gets(p *radius.Packet) (values []ManzaraUserUID, err error) {
	var i uint32
	for _, attr := range _Manzara_GetsVendor(p, 1) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, ManzaraUserUID(i))
	}
	return
}

func ManzaraUserUID_Lookup(p *radius.Packet) (value ManzaraUserUID, err error) {
	a, ok := _Manzara_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = ManzaraUserUID(i)
	return
}

func ManzaraUserUID_Set(p *radius.Packet, value ManzaraUserUID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Manzara_SetVendor(p, 1, a)
}

func ManzaraUserUID_Del(p *radius.Packet) {
	_Manzara_DelVendor(p, 1)
}

type ManzaraUserGID uint32

var ManzaraUserGID_Strings = map[ManzaraUserGID]string{}

func (a ManzaraUserGID) String() string {
	if str, ok := ManzaraUserGID_Strings[a]; ok {
		return str
	}
	return "ManzaraUserGID(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func ManzaraUserGID_Add(p *radius.Packet, value ManzaraUserGID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Manzara_AddVendor(p, 2, a)
}

func ManzaraUserGID_Get(p *radius.Packet) (value ManzaraUserGID) {
	value, _ = ManzaraUserGID_Lookup(p)
	return
}

func ManzaraUserGID_Gets(p *radius.Packet) (values []ManzaraUserGID, err error) {
	var i uint32
	for _, attr := range _Manzara_GetsVendor(p, 2) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, ManzaraUserGID(i))
	}
	return
}

func ManzaraUserGID_Lookup(p *radius.Packet) (value ManzaraUserGID, err error) {
	a, ok := _Manzara_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = ManzaraUserGID(i)
	return
}

func ManzaraUserGID_Set(p *radius.Packet, value ManzaraUserGID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Manzara_SetVendor(p, 2, a)
}

func ManzaraUserGID_Del(p *radius.Packet) {
	_Manzara_DelVendor(p, 2)
}

func ManzaraUserHome_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Manzara_AddVendor(p, 3, a)
}

func ManzaraUserHome_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Manzara_AddVendor(p, 3, a)
}

func ManzaraUserHome_Get(p *radius.Packet) (value []byte) {
	value, _ = ManzaraUserHome_Lookup(p)
	return
}

func ManzaraUserHome_GetString(p *radius.Packet) (value string) {
	value, _ = ManzaraUserHome_LookupString(p)
	return
}

func ManzaraUserHome_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Manzara_GetsVendor(p, 3) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManzaraUserHome_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Manzara_GetsVendor(p, 3) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManzaraUserHome_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Manzara_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func ManzaraUserHome_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Manzara_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func ManzaraUserHome_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Manzara_SetVendor(p, 3, a)
}

func ManzaraUserHome_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Manzara_SetVendor(p, 3, a)
}

func ManzaraUserHome_Del(p *radius.Packet) {
	_Manzara_DelVendor(p, 3)
}

func ManzaraUserShell_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Manzara_AddVendor(p, 4, a)
}

func ManzaraUserShell_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Manzara_AddVendor(p, 4, a)
}

func ManzaraUserShell_Get(p *radius.Packet) (value []byte) {
	value, _ = ManzaraUserShell_Lookup(p)
	return
}

func ManzaraUserShell_GetString(p *radius.Packet) (value string) {
	value, _ = ManzaraUserShell_LookupString(p)
	return
}

func ManzaraUserShell_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Manzara_GetsVendor(p, 4) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManzaraUserShell_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Manzara_GetsVendor(p, 4) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManzaraUserShell_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Manzara_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func ManzaraUserShell_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Manzara_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func ManzaraUserShell_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Manzara_SetVendor(p, 4, a)
}

func ManzaraUserShell_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Manzara_SetVendor(p, 4, a)
}

func ManzaraUserShell_Del(p *radius.Packet) {
	_Manzara_DelVendor(p, 4)
}

func ManzaraPPPAddrString_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Manzara_AddVendor(p, 5, a)
}

func ManzaraPPPAddrString_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Manzara_AddVendor(p, 5, a)
}

func ManzaraPPPAddrString_Get(p *radius.Packet) (value []byte) {
	value, _ = ManzaraPPPAddrString_Lookup(p)
	return
}

func ManzaraPPPAddrString_GetString(p *radius.Packet) (value string) {
	value, _ = ManzaraPPPAddrString_LookupString(p)
	return
}

func ManzaraPPPAddrString_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Manzara_GetsVendor(p, 5) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManzaraPPPAddrString_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Manzara_GetsVendor(p, 5) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManzaraPPPAddrString_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Manzara_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func ManzaraPPPAddrString_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Manzara_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func ManzaraPPPAddrString_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Manzara_SetVendor(p, 5, a)
}

func ManzaraPPPAddrString_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Manzara_SetVendor(p, 5, a)
}

func ManzaraPPPAddrString_Del(p *radius.Packet) {
	_Manzara_DelVendor(p, 5)
}

func ManzaraFullLoginString_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Manzara_AddVendor(p, 6, a)
}

func ManzaraFullLoginString_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Manzara_AddVendor(p, 6, a)
}

func ManzaraFullLoginString_Get(p *radius.Packet) (value []byte) {
	value, _ = ManzaraFullLoginString_Lookup(p)
	return
}

func ManzaraFullLoginString_GetString(p *radius.Packet) (value string) {
	value, _ = ManzaraFullLoginString_LookupString(p)
	return
}

func ManzaraFullLoginString_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Manzara_GetsVendor(p, 6) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManzaraFullLoginString_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Manzara_GetsVendor(p, 6) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManzaraFullLoginString_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Manzara_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func ManzaraFullLoginString_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Manzara_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func ManzaraFullLoginString_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Manzara_SetVendor(p, 6, a)
}

func ManzaraFullLoginString_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Manzara_SetVendor(p, 6, a)
}

func ManzaraFullLoginString_Del(p *radius.Packet) {
	_Manzara_DelVendor(p, 6)
}

type ManzaraTariffUnits uint32

var ManzaraTariffUnits_Strings = map[ManzaraTariffUnits]string{}

func (a ManzaraTariffUnits) String() string {
	if str, ok := ManzaraTariffUnits_Strings[a]; ok {
		return str
	}
	return "ManzaraTariffUnits(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func ManzaraTariffUnits_Add(p *radius.Packet, value ManzaraTariffUnits) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Manzara_AddVendor(p, 7, a)
}

func ManzaraTariffUnits_Get(p *radius.Packet) (value ManzaraTariffUnits) {
	value, _ = ManzaraTariffUnits_Lookup(p)
	return
}

func ManzaraTariffUnits_Gets(p *radius.Packet) (values []ManzaraTariffUnits, err error) {
	var i uint32
	for _, attr := range _Manzara_GetsVendor(p, 7) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, ManzaraTariffUnits(i))
	}
	return
}

func ManzaraTariffUnits_Lookup(p *radius.Packet) (value ManzaraTariffUnits, err error) {
	a, ok := _Manzara_LookupVendor(p, 7)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = ManzaraTariffUnits(i)
	return
}

func ManzaraTariffUnits_Set(p *radius.Packet, value ManzaraTariffUnits) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Manzara_SetVendor(p, 7, a)
}

func ManzaraTariffUnits_Del(p *radius.Packet) {
	_Manzara_DelVendor(p, 7)
}

type ManzaraTariffType uint32

const (
	ManzaraTariffType_Value_MMSPicture ManzaraTariffType = 1
	ManzaraTariffType_Value_Unused     ManzaraTariffType = 2
	ManzaraTariffType_Value_Internet   ManzaraTariffType = 3
)

var ManzaraTariffType_Strings = map[ManzaraTariffType]string{
	ManzaraTariffType_Value_MMSPicture: "MMS-Picture",
	ManzaraTariffType_Value_Unused:     "Unused",
	ManzaraTariffType_Value_Internet:   "Internet",
}

func (a ManzaraTariffType) String() string {
	if str, ok := ManzaraTariffType_Strings[a]; ok {
		return str
	}
	return "ManzaraTariffType(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func ManzaraTariffType_Add(p *radius.Packet, value ManzaraTariffType) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Manzara_AddVendor(p, 8, a)
}

func ManzaraTariffType_Get(p *radius.Packet) (value ManzaraTariffType) {
	value, _ = ManzaraTariffType_Lookup(p)
	return
}

func ManzaraTariffType_Gets(p *radius.Packet) (values []ManzaraTariffType, err error) {
	var i uint32
	for _, attr := range _Manzara_GetsVendor(p, 8) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, ManzaraTariffType(i))
	}
	return
}

func ManzaraTariffType_Lookup(p *radius.Packet) (value ManzaraTariffType, err error) {
	a, ok := _Manzara_LookupVendor(p, 8)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = ManzaraTariffType(i)
	return
}

func ManzaraTariffType_Set(p *radius.Packet, value ManzaraTariffType) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Manzara_SetVendor(p, 8, a)
}

func ManzaraTariffType_Del(p *radius.Packet) {
	_Manzara_DelVendor(p, 8)
}

func ManzaraECPSessionKey_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Manzara_AddVendor(p, 9, a)
}

func ManzaraECPSessionKey_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Manzara_AddVendor(p, 9, a)
}

func ManzaraECPSessionKey_Get(p *radius.Packet) (value []byte) {
	value, _ = ManzaraECPSessionKey_Lookup(p)
	return
}

func ManzaraECPSessionKey_GetString(p *radius.Packet) (value string) {
	value, _ = ManzaraECPSessionKey_LookupString(p)
	return
}

func ManzaraECPSessionKey_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Manzara_GetsVendor(p, 9) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManzaraECPSessionKey_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Manzara_GetsVendor(p, 9) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManzaraECPSessionKey_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Manzara_LookupVendor(p, 9)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func ManzaraECPSessionKey_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Manzara_LookupVendor(p, 9)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func ManzaraECPSessionKey_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Manzara_SetVendor(p, 9, a)
}

func ManzaraECPSessionKey_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Manzara_SetVendor(p, 9, a)
}

func ManzaraECPSessionKey_Del(p *radius.Packet) {
	_Manzara_DelVendor(p, 9)
}

func ManzaraMapName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Manzara_AddVendor(p, 10, a)
}

func ManzaraMapName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Manzara_AddVendor(p, 10, a)
}

func ManzaraMapName_Get(p *radius.Packet) (value []byte) {
	value, _ = ManzaraMapName_Lookup(p)
	return
}

func ManzaraMapName_GetString(p *radius.Packet) (value string) {
	value, _ = ManzaraMapName_LookupString(p)
	return
}

func ManzaraMapName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Manzara_GetsVendor(p, 10) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManzaraMapName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Manzara_GetsVendor(p, 10) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManzaraMapName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Manzara_LookupVendor(p, 10)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func ManzaraMapName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Manzara_LookupVendor(p, 10)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func ManzaraMapName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Manzara_SetVendor(p, 10, a)
}

func ManzaraMapName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Manzara_SetVendor(p, 10, a)
}

func ManzaraMapName_Del(p *radius.Packet) {
	_Manzara_DelVendor(p, 10)
}

func ManzaraMapKey_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Manzara_AddVendor(p, 11, a)
}

func ManzaraMapKey_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Manzara_AddVendor(p, 11, a)
}

func ManzaraMapKey_Get(p *radius.Packet) (value []byte) {
	value, _ = ManzaraMapKey_Lookup(p)
	return
}

func ManzaraMapKey_GetString(p *radius.Packet) (value string) {
	value, _ = ManzaraMapKey_LookupString(p)
	return
}

func ManzaraMapKey_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Manzara_GetsVendor(p, 11) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManzaraMapKey_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Manzara_GetsVendor(p, 11) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManzaraMapKey_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Manzara_LookupVendor(p, 11)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func ManzaraMapKey_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Manzara_LookupVendor(p, 11)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func ManzaraMapKey_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Manzara_SetVendor(p, 11, a)
}

func ManzaraMapKey_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Manzara_SetVendor(p, 11, a)
}

func ManzaraMapKey_Del(p *radius.Packet) {
	_Manzara_DelVendor(p, 11)
}

func ManzaraMapValue_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Manzara_AddVendor(p, 12, a)
}

func ManzaraMapValue_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Manzara_AddVendor(p, 12, a)
}

func ManzaraMapValue_Get(p *radius.Packet) (value []byte) {
	value, _ = ManzaraMapValue_Lookup(p)
	return
}

func ManzaraMapValue_GetString(p *radius.Packet) (value string) {
	value, _ = ManzaraMapValue_LookupString(p)
	return
}

func ManzaraMapValue_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Manzara_GetsVendor(p, 12) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManzaraMapValue_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Manzara_GetsVendor(p, 12) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManzaraMapValue_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Manzara_LookupVendor(p, 12)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func ManzaraMapValue_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Manzara_LookupVendor(p, 12)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func ManzaraMapValue_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Manzara_SetVendor(p, 12, a)
}

func ManzaraMapValue_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Manzara_SetVendor(p, 12, a)
}

func ManzaraMapValue_Del(p *radius.Packet) {
	_Manzara_DelVendor(p, 12)
}

func ManzaraMapError_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Manzara_AddVendor(p, 13, a)
}

func ManzaraMapError_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Manzara_AddVendor(p, 13, a)
}

func ManzaraMapError_Get(p *radius.Packet) (value []byte) {
	value, _ = ManzaraMapError_Lookup(p)
	return
}

func ManzaraMapError_GetString(p *radius.Packet) (value string) {
	value, _ = ManzaraMapError_LookupString(p)
	return
}

func ManzaraMapError_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Manzara_GetsVendor(p, 13) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManzaraMapError_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Manzara_GetsVendor(p, 13) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManzaraMapError_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Manzara_LookupVendor(p, 13)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func ManzaraMapError_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Manzara_LookupVendor(p, 13)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func ManzaraMapError_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Manzara_SetVendor(p, 13, a)
}

func ManzaraMapError_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Manzara_SetVendor(p, 13, a)
}

func ManzaraMapError_Del(p *radius.Packet) {
	_Manzara_DelVendor(p, 13)
}

func ManzaraServiceType_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Manzara_AddVendor(p, 14, a)
}

func ManzaraServiceType_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Manzara_AddVendor(p, 14, a)
}

func ManzaraServiceType_Get(p *radius.Packet) (value []byte) {
	value, _ = ManzaraServiceType_Lookup(p)
	return
}

func ManzaraServiceType_GetString(p *radius.Packet) (value string) {
	value, _ = ManzaraServiceType_LookupString(p)
	return
}

func ManzaraServiceType_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Manzara_GetsVendor(p, 14) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManzaraServiceType_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Manzara_GetsVendor(p, 14) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManzaraServiceType_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Manzara_LookupVendor(p, 14)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func ManzaraServiceType_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Manzara_LookupVendor(p, 14)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func ManzaraServiceType_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Manzara_SetVendor(p, 14, a)
}

func ManzaraServiceType_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Manzara_SetVendor(p, 14, a)
}

func ManzaraServiceType_Del(p *radius.Packet) {
	_Manzara_DelVendor(p, 14)
}
