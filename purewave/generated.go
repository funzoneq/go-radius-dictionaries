// Code generated by radius-dict-gen. DO NOT EDIT.

package purewave

import (
	"net"
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Purewave_VendorID = 21074
)

func _Purewave_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Purewave_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Purewave_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Purewave_VendorID {
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

func _Purewave_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Purewave_VendorID {
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

func _Purewave_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Purewave_VendorID {
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
	return _Purewave_AddVendor(p, typ, attr)
}

func _Purewave_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Purewave_VendorID {
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

type PurewaveClientProfile uint32

var PurewaveClientProfile_Strings = map[PurewaveClientProfile]string{}

func (a PurewaveClientProfile) String() string {
	if str, ok := PurewaveClientProfile_Strings[a]; ok {
		return str
	}
	return "PurewaveClientProfile(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func PurewaveClientProfile_Add(p *radius.Packet, value PurewaveClientProfile) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Purewave_AddVendor(p, 1, a)
}

func PurewaveClientProfile_Get(p *radius.Packet) (value PurewaveClientProfile) {
	value, _ = PurewaveClientProfile_Lookup(p)
	return
}

func PurewaveClientProfile_Gets(p *radius.Packet) (values []PurewaveClientProfile, err error) {
	var i uint32
	for _, attr := range _Purewave_GetsVendor(p, 1) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, PurewaveClientProfile(i))
	}
	return
}

func PurewaveClientProfile_Lookup(p *radius.Packet) (value PurewaveClientProfile, err error) {
	a, ok := _Purewave_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = PurewaveClientProfile(i)
	return
}

func PurewaveClientProfile_Set(p *radius.Packet, value PurewaveClientProfile) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Purewave_SetVendor(p, 1, a)
}

func PurewaveClientProfile_Del(p *radius.Packet) {
	_Purewave_DelVendor(p, 1)
}

type PurewaveCSType uint32

const (
	PurewaveCSType_Value_IPV4CS     PurewaveCSType = 0
	PurewaveCSType_Value_ETHERNETCS PurewaveCSType = 1
)

var PurewaveCSType_Strings = map[PurewaveCSType]string{
	PurewaveCSType_Value_IPV4CS:     "IPV4_CS",
	PurewaveCSType_Value_ETHERNETCS: "ETHERNET_CS",
}

func (a PurewaveCSType) String() string {
	if str, ok := PurewaveCSType_Strings[a]; ok {
		return str
	}
	return "PurewaveCSType(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func PurewaveCSType_Add(p *radius.Packet, value PurewaveCSType) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Purewave_AddVendor(p, 2, a)
}

func PurewaveCSType_Get(p *radius.Packet) (value PurewaveCSType) {
	value, _ = PurewaveCSType_Lookup(p)
	return
}

func PurewaveCSType_Gets(p *radius.Packet) (values []PurewaveCSType, err error) {
	var i uint32
	for _, attr := range _Purewave_GetsVendor(p, 2) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, PurewaveCSType(i))
	}
	return
}

func PurewaveCSType_Lookup(p *radius.Packet) (value PurewaveCSType, err error) {
	a, ok := _Purewave_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = PurewaveCSType(i)
	return
}

func PurewaveCSType_Set(p *radius.Packet, value PurewaveCSType) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Purewave_SetVendor(p, 2, a)
}

func PurewaveCSType_Del(p *radius.Packet) {
	_Purewave_DelVendor(p, 2)
}

type PurewaveMaxDownlinkRate uint32

const (
	PurewaveMaxDownlinkRate_Value_QPSK12  PurewaveMaxDownlinkRate = 3
	PurewaveMaxDownlinkRate_Value_QPSK34  PurewaveMaxDownlinkRate = 4
	PurewaveMaxDownlinkRate_Value_QAM1612 PurewaveMaxDownlinkRate = 5
	PurewaveMaxDownlinkRate_Value_QAM1634 PurewaveMaxDownlinkRate = 6
	PurewaveMaxDownlinkRate_Value_QAM6412 PurewaveMaxDownlinkRate = 7
	PurewaveMaxDownlinkRate_Value_QAM6423 PurewaveMaxDownlinkRate = 8
	PurewaveMaxDownlinkRate_Value_QAM6434 PurewaveMaxDownlinkRate = 9
	PurewaveMaxDownlinkRate_Value_QAM6456 PurewaveMaxDownlinkRate = 10
)

var PurewaveMaxDownlinkRate_Strings = map[PurewaveMaxDownlinkRate]string{
	PurewaveMaxDownlinkRate_Value_QPSK12:  "QPSK_1/2",
	PurewaveMaxDownlinkRate_Value_QPSK34:  "QPSK_3/4",
	PurewaveMaxDownlinkRate_Value_QAM1612: "QAM16_1/2",
	PurewaveMaxDownlinkRate_Value_QAM1634: "QAM16_3/4",
	PurewaveMaxDownlinkRate_Value_QAM6412: "QAM64_1/2",
	PurewaveMaxDownlinkRate_Value_QAM6423: "QAM64_2/3",
	PurewaveMaxDownlinkRate_Value_QAM6434: "QAM64_3/4",
	PurewaveMaxDownlinkRate_Value_QAM6456: "QAM64_5/6",
}

func (a PurewaveMaxDownlinkRate) String() string {
	if str, ok := PurewaveMaxDownlinkRate_Strings[a]; ok {
		return str
	}
	return "PurewaveMaxDownlinkRate(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func PurewaveMaxDownlinkRate_Add(p *radius.Packet, value PurewaveMaxDownlinkRate) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Purewave_AddVendor(p, 3, a)
}

func PurewaveMaxDownlinkRate_Get(p *radius.Packet) (value PurewaveMaxDownlinkRate) {
	value, _ = PurewaveMaxDownlinkRate_Lookup(p)
	return
}

func PurewaveMaxDownlinkRate_Gets(p *radius.Packet) (values []PurewaveMaxDownlinkRate, err error) {
	var i uint32
	for _, attr := range _Purewave_GetsVendor(p, 3) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, PurewaveMaxDownlinkRate(i))
	}
	return
}

func PurewaveMaxDownlinkRate_Lookup(p *radius.Packet) (value PurewaveMaxDownlinkRate, err error) {
	a, ok := _Purewave_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = PurewaveMaxDownlinkRate(i)
	return
}

func PurewaveMaxDownlinkRate_Set(p *radius.Packet, value PurewaveMaxDownlinkRate) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Purewave_SetVendor(p, 3, a)
}

func PurewaveMaxDownlinkRate_Del(p *radius.Packet) {
	_Purewave_DelVendor(p, 3)
}

type PurewaveMaxUplinkRate uint32

const (
	PurewaveMaxUplinkRate_Value_QPSK12  PurewaveMaxUplinkRate = 3
	PurewaveMaxUplinkRate_Value_QPSK34  PurewaveMaxUplinkRate = 4
	PurewaveMaxUplinkRate_Value_QAM1612 PurewaveMaxUplinkRate = 5
	PurewaveMaxUplinkRate_Value_QAM1634 PurewaveMaxUplinkRate = 6
	PurewaveMaxUplinkRate_Value_QAM6412 PurewaveMaxUplinkRate = 7
	PurewaveMaxUplinkRate_Value_QAM6423 PurewaveMaxUplinkRate = 8
	PurewaveMaxUplinkRate_Value_QAM6434 PurewaveMaxUplinkRate = 9
	PurewaveMaxUplinkRate_Value_QAM6456 PurewaveMaxUplinkRate = 10
)

var PurewaveMaxUplinkRate_Strings = map[PurewaveMaxUplinkRate]string{
	PurewaveMaxUplinkRate_Value_QPSK12:  "QPSK_1/2",
	PurewaveMaxUplinkRate_Value_QPSK34:  "QPSK_3/4",
	PurewaveMaxUplinkRate_Value_QAM1612: "QAM16_1/2",
	PurewaveMaxUplinkRate_Value_QAM1634: "QAM16_3/4",
	PurewaveMaxUplinkRate_Value_QAM6412: "QAM64_1/2",
	PurewaveMaxUplinkRate_Value_QAM6423: "QAM64_2/3",
	PurewaveMaxUplinkRate_Value_QAM6434: "QAM64_3/4",
	PurewaveMaxUplinkRate_Value_QAM6456: "QAM64_5/6",
}

func (a PurewaveMaxUplinkRate) String() string {
	if str, ok := PurewaveMaxUplinkRate_Strings[a]; ok {
		return str
	}
	return "PurewaveMaxUplinkRate(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func PurewaveMaxUplinkRate_Add(p *radius.Packet, value PurewaveMaxUplinkRate) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Purewave_AddVendor(p, 4, a)
}

func PurewaveMaxUplinkRate_Get(p *radius.Packet) (value PurewaveMaxUplinkRate) {
	value, _ = PurewaveMaxUplinkRate_Lookup(p)
	return
}

func PurewaveMaxUplinkRate_Gets(p *radius.Packet) (values []PurewaveMaxUplinkRate, err error) {
	var i uint32
	for _, attr := range _Purewave_GetsVendor(p, 4) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, PurewaveMaxUplinkRate(i))
	}
	return
}

func PurewaveMaxUplinkRate_Lookup(p *radius.Packet) (value PurewaveMaxUplinkRate, err error) {
	a, ok := _Purewave_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = PurewaveMaxUplinkRate(i)
	return
}

func PurewaveMaxUplinkRate_Set(p *radius.Packet, value PurewaveMaxUplinkRate) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Purewave_SetVendor(p, 4, a)
}

func PurewaveMaxUplinkRate_Del(p *radius.Packet) {
	_Purewave_DelVendor(p, 4)
}

func PurewaveIPAddress_Add(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Purewave_AddVendor(p, 5, a)
}

func PurewaveIPAddress_Get(p *radius.Packet) (value net.IP) {
	value, _ = PurewaveIPAddress_Lookup(p)
	return
}

func PurewaveIPAddress_Gets(p *radius.Packet) (values []net.IP, err error) {
	var i net.IP
	for _, attr := range _Purewave_GetsVendor(p, 5) {
		i, err = radius.IPAddr(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func PurewaveIPAddress_Lookup(p *radius.Packet) (value net.IP, err error) {
	a, ok := _Purewave_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value, err = radius.IPAddr(a)
	return
}

func PurewaveIPAddress_Set(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Purewave_SetVendor(p, 5, a)
}

func PurewaveIPAddress_Del(p *radius.Packet) {
	_Purewave_DelVendor(p, 5)
}

func PurewaveIPNetmask_Add(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Purewave_AddVendor(p, 6, a)
}

func PurewaveIPNetmask_Get(p *radius.Packet) (value net.IP) {
	value, _ = PurewaveIPNetmask_Lookup(p)
	return
}

func PurewaveIPNetmask_Gets(p *radius.Packet) (values []net.IP, err error) {
	var i net.IP
	for _, attr := range _Purewave_GetsVendor(p, 6) {
		i, err = radius.IPAddr(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func PurewaveIPNetmask_Lookup(p *radius.Packet) (value net.IP, err error) {
	a, ok := _Purewave_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value, err = radius.IPAddr(a)
	return
}

func PurewaveIPNetmask_Set(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Purewave_SetVendor(p, 6, a)
}

func PurewaveIPNetmask_Del(p *radius.Packet) {
	_Purewave_DelVendor(p, 6)
}

type PurewaveServiceEnable uint32

var PurewaveServiceEnable_Strings = map[PurewaveServiceEnable]string{}

func (a PurewaveServiceEnable) String() string {
	if str, ok := PurewaveServiceEnable_Strings[a]; ok {
		return str
	}
	return "PurewaveServiceEnable(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func PurewaveServiceEnable_Add(p *radius.Packet, value PurewaveServiceEnable) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Purewave_AddVendor(p, 7, a)
}

func PurewaveServiceEnable_Get(p *radius.Packet) (value PurewaveServiceEnable) {
	value, _ = PurewaveServiceEnable_Lookup(p)
	return
}

func PurewaveServiceEnable_Gets(p *radius.Packet) (values []PurewaveServiceEnable, err error) {
	var i uint32
	for _, attr := range _Purewave_GetsVendor(p, 7) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, PurewaveServiceEnable(i))
	}
	return
}

func PurewaveServiceEnable_Lookup(p *radius.Packet) (value PurewaveServiceEnable, err error) {
	a, ok := _Purewave_LookupVendor(p, 7)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = PurewaveServiceEnable(i)
	return
}

func PurewaveServiceEnable_Set(p *radius.Packet, value PurewaveServiceEnable) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Purewave_SetVendor(p, 7, a)
}

func PurewaveServiceEnable_Del(p *radius.Packet) {
	_Purewave_DelVendor(p, 7)
}
