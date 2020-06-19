// Code generated by radius-dict-gen. DO NOT EDIT.

package springtide

import (
	"net"
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_SpringTide_VendorID = 3551
)

func _SpringTide_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_SpringTide_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _SpringTide_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _SpringTide_VendorID {
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

func _SpringTide_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _SpringTide_VendorID {
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

func _SpringTide_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _SpringTide_VendorID {
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
	return _SpringTide_AddVendor(p, typ, attr)
}

func _SpringTide_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _SpringTide_VendorID {
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

func STAcctVCConnectionID_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _SpringTide_AddVendor(p, 1, a)
}

func STAcctVCConnectionID_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _SpringTide_AddVendor(p, 1, a)
}

func STAcctVCConnectionID_Get(p *radius.Packet) (value []byte) {
	value, _ = STAcctVCConnectionID_Lookup(p)
	return
}

func STAcctVCConnectionID_GetString(p *radius.Packet) (value string) {
	value, _ = STAcctVCConnectionID_LookupString(p)
	return
}

func STAcctVCConnectionID_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _SpringTide_GetsVendor(p, 1) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func STAcctVCConnectionID_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _SpringTide_GetsVendor(p, 1) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func STAcctVCConnectionID_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _SpringTide_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func STAcctVCConnectionID_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _SpringTide_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func STAcctVCConnectionID_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _SpringTide_SetVendor(p, 1, a)
}

func STAcctVCConnectionID_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _SpringTide_SetVendor(p, 1, a)
}

func STAcctVCConnectionID_Del(p *radius.Packet) {
	_SpringTide_DelVendor(p, 1)
}

func STServiceName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _SpringTide_AddVendor(p, 2, a)
}

func STServiceName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _SpringTide_AddVendor(p, 2, a)
}

func STServiceName_Get(p *radius.Packet) (value []byte) {
	value, _ = STServiceName_Lookup(p)
	return
}

func STServiceName_GetString(p *radius.Packet) (value string) {
	value, _ = STServiceName_LookupString(p)
	return
}

func STServiceName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _SpringTide_GetsVendor(p, 2) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func STServiceName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _SpringTide_GetsVendor(p, 2) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func STServiceName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _SpringTide_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func STServiceName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _SpringTide_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func STServiceName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _SpringTide_SetVendor(p, 2, a)
}

func STServiceName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _SpringTide_SetVendor(p, 2, a)
}

func STServiceName_Del(p *radius.Packet) {
	_SpringTide_DelVendor(p, 2)
}

type STServiceDomain uint32

var STServiceDomain_Strings = map[STServiceDomain]string{}

func (a STServiceDomain) String() string {
	if str, ok := STServiceDomain_Strings[a]; ok {
		return str
	}
	return "STServiceDomain(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func STServiceDomain_Add(p *radius.Packet, value STServiceDomain) (err error) {
	a := radius.NewInteger(uint32(value))
	return _SpringTide_AddVendor(p, 3, a)
}

func STServiceDomain_Get(p *radius.Packet) (value STServiceDomain) {
	value, _ = STServiceDomain_Lookup(p)
	return
}

func STServiceDomain_Gets(p *radius.Packet) (values []STServiceDomain, err error) {
	var i uint32
	for _, attr := range _SpringTide_GetsVendor(p, 3) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, STServiceDomain(i))
	}
	return
}

func STServiceDomain_Lookup(p *radius.Packet) (value STServiceDomain, err error) {
	a, ok := _SpringTide_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = STServiceDomain(i)
	return
}

func STServiceDomain_Set(p *radius.Packet, value STServiceDomain) (err error) {
	a := radius.NewInteger(uint32(value))
	return _SpringTide_SetVendor(p, 3, a)
}

func STServiceDomain_Del(p *radius.Packet) {
	_SpringTide_DelVendor(p, 3)
}

func STPolicyName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _SpringTide_AddVendor(p, 4, a)
}

func STPolicyName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _SpringTide_AddVendor(p, 4, a)
}

func STPolicyName_Get(p *radius.Packet) (value []byte) {
	value, _ = STPolicyName_Lookup(p)
	return
}

func STPolicyName_GetString(p *radius.Packet) (value string) {
	value, _ = STPolicyName_LookupString(p)
	return
}

func STPolicyName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _SpringTide_GetsVendor(p, 4) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func STPolicyName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _SpringTide_GetsVendor(p, 4) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func STPolicyName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _SpringTide_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func STPolicyName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _SpringTide_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func STPolicyName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _SpringTide_SetVendor(p, 4, a)
}

func STPolicyName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _SpringTide_SetVendor(p, 4, a)
}

func STPolicyName_Del(p *radius.Packet) {
	_SpringTide_DelVendor(p, 4)
}

func STPrimaryDNSServer_Add(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _SpringTide_AddVendor(p, 5, a)
}

func STPrimaryDNSServer_Get(p *radius.Packet) (value net.IP) {
	value, _ = STPrimaryDNSServer_Lookup(p)
	return
}

func STPrimaryDNSServer_Gets(p *radius.Packet) (values []net.IP, err error) {
	var i net.IP
	for _, attr := range _SpringTide_GetsVendor(p, 5) {
		i, err = radius.IPAddr(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func STPrimaryDNSServer_Lookup(p *radius.Packet) (value net.IP, err error) {
	a, ok := _SpringTide_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value, err = radius.IPAddr(a)
	return
}

func STPrimaryDNSServer_Set(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _SpringTide_SetVendor(p, 5, a)
}

func STPrimaryDNSServer_Del(p *radius.Packet) {
	_SpringTide_DelVendor(p, 5)
}

func STSecondaryDNSServer_Add(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _SpringTide_AddVendor(p, 6, a)
}

func STSecondaryDNSServer_Get(p *radius.Packet) (value net.IP) {
	value, _ = STSecondaryDNSServer_Lookup(p)
	return
}

func STSecondaryDNSServer_Gets(p *radius.Packet) (values []net.IP, err error) {
	var i net.IP
	for _, attr := range _SpringTide_GetsVendor(p, 6) {
		i, err = radius.IPAddr(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func STSecondaryDNSServer_Lookup(p *radius.Packet) (value net.IP, err error) {
	a, ok := _SpringTide_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value, err = radius.IPAddr(a)
	return
}

func STSecondaryDNSServer_Set(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _SpringTide_SetVendor(p, 6, a)
}

func STSecondaryDNSServer_Del(p *radius.Packet) {
	_SpringTide_DelVendor(p, 6)
}

func STPrimaryNBNSServer_Add(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _SpringTide_AddVendor(p, 7, a)
}

func STPrimaryNBNSServer_Get(p *radius.Packet) (value net.IP) {
	value, _ = STPrimaryNBNSServer_Lookup(p)
	return
}

func STPrimaryNBNSServer_Gets(p *radius.Packet) (values []net.IP, err error) {
	var i net.IP
	for _, attr := range _SpringTide_GetsVendor(p, 7) {
		i, err = radius.IPAddr(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func STPrimaryNBNSServer_Lookup(p *radius.Packet) (value net.IP, err error) {
	a, ok := _SpringTide_LookupVendor(p, 7)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value, err = radius.IPAddr(a)
	return
}

func STPrimaryNBNSServer_Set(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _SpringTide_SetVendor(p, 7, a)
}

func STPrimaryNBNSServer_Del(p *radius.Packet) {
	_SpringTide_DelVendor(p, 7)
}

func STSecondaryNBNSServer_Add(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _SpringTide_AddVendor(p, 8, a)
}

func STSecondaryNBNSServer_Get(p *radius.Packet) (value net.IP) {
	value, _ = STSecondaryNBNSServer_Lookup(p)
	return
}

func STSecondaryNBNSServer_Gets(p *radius.Packet) (values []net.IP, err error) {
	var i net.IP
	for _, attr := range _SpringTide_GetsVendor(p, 8) {
		i, err = radius.IPAddr(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func STSecondaryNBNSServer_Lookup(p *radius.Packet) (value net.IP, err error) {
	a, ok := _SpringTide_LookupVendor(p, 8)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value, err = radius.IPAddr(a)
	return
}

func STSecondaryNBNSServer_Set(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _SpringTide_SetVendor(p, 8, a)
}

func STSecondaryNBNSServer_Del(p *radius.Packet) {
	_SpringTide_DelVendor(p, 8)
}

type STPhysicalPort uint32

var STPhysicalPort_Strings = map[STPhysicalPort]string{}

func (a STPhysicalPort) String() string {
	if str, ok := STPhysicalPort_Strings[a]; ok {
		return str
	}
	return "STPhysicalPort(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func STPhysicalPort_Add(p *radius.Packet, value STPhysicalPort) (err error) {
	a := radius.NewInteger(uint32(value))
	return _SpringTide_AddVendor(p, 9, a)
}

func STPhysicalPort_Get(p *radius.Packet) (value STPhysicalPort) {
	value, _ = STPhysicalPort_Lookup(p)
	return
}

func STPhysicalPort_Gets(p *radius.Packet) (values []STPhysicalPort, err error) {
	var i uint32
	for _, attr := range _SpringTide_GetsVendor(p, 9) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, STPhysicalPort(i))
	}
	return
}

func STPhysicalPort_Lookup(p *radius.Packet) (value STPhysicalPort, err error) {
	a, ok := _SpringTide_LookupVendor(p, 9)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = STPhysicalPort(i)
	return
}

func STPhysicalPort_Set(p *radius.Packet, value STPhysicalPort) (err error) {
	a := radius.NewInteger(uint32(value))
	return _SpringTide_SetVendor(p, 9, a)
}

func STPhysicalPort_Del(p *radius.Packet) {
	_SpringTide_DelVendor(p, 9)
}

type STPhysicalSlot uint32

var STPhysicalSlot_Strings = map[STPhysicalSlot]string{}

func (a STPhysicalSlot) String() string {
	if str, ok := STPhysicalSlot_Strings[a]; ok {
		return str
	}
	return "STPhysicalSlot(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func STPhysicalSlot_Add(p *radius.Packet, value STPhysicalSlot) (err error) {
	a := radius.NewInteger(uint32(value))
	return _SpringTide_AddVendor(p, 10, a)
}

func STPhysicalSlot_Get(p *radius.Packet) (value STPhysicalSlot) {
	value, _ = STPhysicalSlot_Lookup(p)
	return
}

func STPhysicalSlot_Gets(p *radius.Packet) (values []STPhysicalSlot, err error) {
	var i uint32
	for _, attr := range _SpringTide_GetsVendor(p, 10) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, STPhysicalSlot(i))
	}
	return
}

func STPhysicalSlot_Lookup(p *radius.Packet) (value STPhysicalSlot, err error) {
	a, ok := _SpringTide_LookupVendor(p, 10)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = STPhysicalSlot(i)
	return
}

func STPhysicalSlot_Set(p *radius.Packet, value STPhysicalSlot) (err error) {
	a := radius.NewInteger(uint32(value))
	return _SpringTide_SetVendor(p, 10, a)
}

func STPhysicalSlot_Del(p *radius.Packet) {
	_SpringTide_DelVendor(p, 10)
}

type STVirtualPathID uint32

var STVirtualPathID_Strings = map[STVirtualPathID]string{}

func (a STVirtualPathID) String() string {
	if str, ok := STVirtualPathID_Strings[a]; ok {
		return str
	}
	return "STVirtualPathID(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func STVirtualPathID_Add(p *radius.Packet, value STVirtualPathID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _SpringTide_AddVendor(p, 11, a)
}

func STVirtualPathID_Get(p *radius.Packet) (value STVirtualPathID) {
	value, _ = STVirtualPathID_Lookup(p)
	return
}

func STVirtualPathID_Gets(p *radius.Packet) (values []STVirtualPathID, err error) {
	var i uint32
	for _, attr := range _SpringTide_GetsVendor(p, 11) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, STVirtualPathID(i))
	}
	return
}

func STVirtualPathID_Lookup(p *radius.Packet) (value STVirtualPathID, err error) {
	a, ok := _SpringTide_LookupVendor(p, 11)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = STVirtualPathID(i)
	return
}

func STVirtualPathID_Set(p *radius.Packet, value STVirtualPathID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _SpringTide_SetVendor(p, 11, a)
}

func STVirtualPathID_Del(p *radius.Packet) {
	_SpringTide_DelVendor(p, 11)
}

type STVirtualCircuitID uint32

var STVirtualCircuitID_Strings = map[STVirtualCircuitID]string{}

func (a STVirtualCircuitID) String() string {
	if str, ok := STVirtualCircuitID_Strings[a]; ok {
		return str
	}
	return "STVirtualCircuitID(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func STVirtualCircuitID_Add(p *radius.Packet, value STVirtualCircuitID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _SpringTide_AddVendor(p, 12, a)
}

func STVirtualCircuitID_Get(p *radius.Packet) (value STVirtualCircuitID) {
	value, _ = STVirtualCircuitID_Lookup(p)
	return
}

func STVirtualCircuitID_Gets(p *radius.Packet) (values []STVirtualCircuitID, err error) {
	var i uint32
	for _, attr := range _SpringTide_GetsVendor(p, 12) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, STVirtualCircuitID(i))
	}
	return
}

func STVirtualCircuitID_Lookup(p *radius.Packet) (value STVirtualCircuitID, err error) {
	a, ok := _SpringTide_LookupVendor(p, 12)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = STVirtualCircuitID(i)
	return
}

func STVirtualCircuitID_Set(p *radius.Packet, value STVirtualCircuitID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _SpringTide_SetVendor(p, 12, a)
}

func STVirtualCircuitID_Del(p *radius.Packet) {
	_SpringTide_DelVendor(p, 12)
}

func STRealmName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _SpringTide_AddVendor(p, 13, a)
}

func STRealmName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _SpringTide_AddVendor(p, 13, a)
}

func STRealmName_Get(p *radius.Packet) (value []byte) {
	value, _ = STRealmName_Lookup(p)
	return
}

func STRealmName_GetString(p *radius.Packet) (value string) {
	value, _ = STRealmName_LookupString(p)
	return
}

func STRealmName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _SpringTide_GetsVendor(p, 13) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func STRealmName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _SpringTide_GetsVendor(p, 13) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func STRealmName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _SpringTide_LookupVendor(p, 13)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func STRealmName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _SpringTide_LookupVendor(p, 13)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func STRealmName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _SpringTide_SetVendor(p, 13, a)
}

func STRealmName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _SpringTide_SetVendor(p, 13, a)
}

func STRealmName_Del(p *radius.Packet) {
	_SpringTide_DelVendor(p, 13)
}

type STIPSecPfsGroup uint32

var STIPSecPfsGroup_Strings = map[STIPSecPfsGroup]string{}

func (a STIPSecPfsGroup) String() string {
	if str, ok := STIPSecPfsGroup_Strings[a]; ok {
		return str
	}
	return "STIPSecPfsGroup(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func STIPSecPfsGroup_Add(p *radius.Packet, value STIPSecPfsGroup) (err error) {
	a := radius.NewInteger(uint32(value))
	return _SpringTide_AddVendor(p, 14, a)
}

func STIPSecPfsGroup_Get(p *radius.Packet) (value STIPSecPfsGroup) {
	value, _ = STIPSecPfsGroup_Lookup(p)
	return
}

func STIPSecPfsGroup_Gets(p *radius.Packet) (values []STIPSecPfsGroup, err error) {
	var i uint32
	for _, attr := range _SpringTide_GetsVendor(p, 14) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, STIPSecPfsGroup(i))
	}
	return
}

func STIPSecPfsGroup_Lookup(p *radius.Packet) (value STIPSecPfsGroup, err error) {
	a, ok := _SpringTide_LookupVendor(p, 14)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = STIPSecPfsGroup(i)
	return
}

func STIPSecPfsGroup_Set(p *radius.Packet, value STIPSecPfsGroup) (err error) {
	a := radius.NewInteger(uint32(value))
	return _SpringTide_SetVendor(p, 14, a)
}

func STIPSecPfsGroup_Del(p *radius.Packet) {
	_SpringTide_DelVendor(p, 14)
}

type STIPSecClientFirewall uint32

var STIPSecClientFirewall_Strings = map[STIPSecClientFirewall]string{}

func (a STIPSecClientFirewall) String() string {
	if str, ok := STIPSecClientFirewall_Strings[a]; ok {
		return str
	}
	return "STIPSecClientFirewall(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func STIPSecClientFirewall_Add(p *radius.Packet, value STIPSecClientFirewall) (err error) {
	a := radius.NewInteger(uint32(value))
	return _SpringTide_AddVendor(p, 15, a)
}

func STIPSecClientFirewall_Get(p *radius.Packet) (value STIPSecClientFirewall) {
	value, _ = STIPSecClientFirewall_Lookup(p)
	return
}

func STIPSecClientFirewall_Gets(p *radius.Packet) (values []STIPSecClientFirewall, err error) {
	var i uint32
	for _, attr := range _SpringTide_GetsVendor(p, 15) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, STIPSecClientFirewall(i))
	}
	return
}

func STIPSecClientFirewall_Lookup(p *radius.Packet) (value STIPSecClientFirewall, err error) {
	a, ok := _SpringTide_LookupVendor(p, 15)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = STIPSecClientFirewall(i)
	return
}

func STIPSecClientFirewall_Set(p *radius.Packet, value STIPSecClientFirewall) (err error) {
	a := radius.NewInteger(uint32(value))
	return _SpringTide_SetVendor(p, 15, a)
}

func STIPSecClientFirewall_Del(p *radius.Packet) {
	_SpringTide_DelVendor(p, 15)
}

func STIPSecClientSubnet_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _SpringTide_AddVendor(p, 16, a)
}

func STIPSecClientSubnet_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _SpringTide_AddVendor(p, 16, a)
}

func STIPSecClientSubnet_Get(p *radius.Packet) (value []byte) {
	value, _ = STIPSecClientSubnet_Lookup(p)
	return
}

func STIPSecClientSubnet_GetString(p *radius.Packet) (value string) {
	value, _ = STIPSecClientSubnet_LookupString(p)
	return
}

func STIPSecClientSubnet_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _SpringTide_GetsVendor(p, 16) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func STIPSecClientSubnet_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _SpringTide_GetsVendor(p, 16) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func STIPSecClientSubnet_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _SpringTide_LookupVendor(p, 16)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func STIPSecClientSubnet_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _SpringTide_LookupVendor(p, 16)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func STIPSecClientSubnet_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _SpringTide_SetVendor(p, 16, a)
}

func STIPSecClientSubnet_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _SpringTide_SetVendor(p, 16, a)
}

func STIPSecClientSubnet_Del(p *radius.Packet) {
	_SpringTide_DelVendor(p, 16)
}
