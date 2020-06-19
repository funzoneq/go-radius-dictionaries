// Code generated by radius-dict-gen. DO NOT EDIT.

package compatible

import (
	"net"
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Compatible_VendorID = 255
)

func _Compatible_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Compatible_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Compatible_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Compatible_VendorID {
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

func _Compatible_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Compatible_VendorID {
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

func _Compatible_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Compatible_VendorID {
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
	return _Compatible_AddVendor(p, typ, attr)
}

func _Compatible_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Compatible_VendorID {
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

type CompatibleTunnelDelay uint32

var CompatibleTunnelDelay_Strings = map[CompatibleTunnelDelay]string{}

func (a CompatibleTunnelDelay) String() string {
	if str, ok := CompatibleTunnelDelay_Strings[a]; ok {
		return str
	}
	return "CompatibleTunnelDelay(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func CompatibleTunnelDelay_Add(p *radius.Packet, value CompatibleTunnelDelay) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Compatible_AddVendor(p, 0, a)
}

func CompatibleTunnelDelay_Get(p *radius.Packet) (value CompatibleTunnelDelay) {
	value, _ = CompatibleTunnelDelay_Lookup(p)
	return
}

func CompatibleTunnelDelay_Gets(p *radius.Packet) (values []CompatibleTunnelDelay, err error) {
	var i uint32
	for _, attr := range _Compatible_GetsVendor(p, 0) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, CompatibleTunnelDelay(i))
	}
	return
}

func CompatibleTunnelDelay_Lookup(p *radius.Packet) (value CompatibleTunnelDelay, err error) {
	a, ok := _Compatible_LookupVendor(p, 0)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = CompatibleTunnelDelay(i)
	return
}

func CompatibleTunnelDelay_Set(p *radius.Packet, value CompatibleTunnelDelay) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Compatible_SetVendor(p, 0, a)
}

func CompatibleTunnelDelay_Del(p *radius.Packet) {
	_Compatible_DelVendor(p, 0)
}

type CompatibleTunnelThroughput uint32

var CompatibleTunnelThroughput_Strings = map[CompatibleTunnelThroughput]string{}

func (a CompatibleTunnelThroughput) String() string {
	if str, ok := CompatibleTunnelThroughput_Strings[a]; ok {
		return str
	}
	return "CompatibleTunnelThroughput(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func CompatibleTunnelThroughput_Add(p *radius.Packet, value CompatibleTunnelThroughput) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Compatible_AddVendor(p, 1, a)
}

func CompatibleTunnelThroughput_Get(p *radius.Packet) (value CompatibleTunnelThroughput) {
	value, _ = CompatibleTunnelThroughput_Lookup(p)
	return
}

func CompatibleTunnelThroughput_Gets(p *radius.Packet) (values []CompatibleTunnelThroughput, err error) {
	var i uint32
	for _, attr := range _Compatible_GetsVendor(p, 1) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, CompatibleTunnelThroughput(i))
	}
	return
}

func CompatibleTunnelThroughput_Lookup(p *radius.Packet) (value CompatibleTunnelThroughput, err error) {
	a, ok := _Compatible_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = CompatibleTunnelThroughput(i)
	return
}

func CompatibleTunnelThroughput_Set(p *radius.Packet, value CompatibleTunnelThroughput) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Compatible_SetVendor(p, 1, a)
}

func CompatibleTunnelThroughput_Del(p *radius.Packet) {
	_Compatible_DelVendor(p, 1)
}

func CompatibleTunnelServerEndpoint_Add(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Compatible_AddVendor(p, 3, a)
}

func CompatibleTunnelServerEndpoint_Get(p *radius.Packet) (value net.IP) {
	value, _ = CompatibleTunnelServerEndpoint_Lookup(p)
	return
}

func CompatibleTunnelServerEndpoint_Gets(p *radius.Packet) (values []net.IP, err error) {
	var i net.IP
	for _, attr := range _Compatible_GetsVendor(p, 3) {
		i, err = radius.IPAddr(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func CompatibleTunnelServerEndpoint_Lookup(p *radius.Packet) (value net.IP, err error) {
	a, ok := _Compatible_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value, err = radius.IPAddr(a)
	return
}

func CompatibleTunnelServerEndpoint_Set(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Compatible_SetVendor(p, 3, a)
}

func CompatibleTunnelServerEndpoint_Del(p *radius.Packet) {
	_Compatible_DelVendor(p, 3)
}

func CompatibleTunnelGroupInfo_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Compatible_AddVendor(p, 4, a)
}

func CompatibleTunnelGroupInfo_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Compatible_AddVendor(p, 4, a)
}

func CompatibleTunnelGroupInfo_Get(p *radius.Packet) (value []byte) {
	value, _ = CompatibleTunnelGroupInfo_Lookup(p)
	return
}

func CompatibleTunnelGroupInfo_GetString(p *radius.Packet) (value string) {
	value, _ = CompatibleTunnelGroupInfo_LookupString(p)
	return
}

func CompatibleTunnelGroupInfo_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Compatible_GetsVendor(p, 4) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func CompatibleTunnelGroupInfo_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Compatible_GetsVendor(p, 4) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func CompatibleTunnelGroupInfo_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Compatible_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func CompatibleTunnelGroupInfo_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Compatible_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func CompatibleTunnelGroupInfo_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Compatible_SetVendor(p, 4, a)
}

func CompatibleTunnelGroupInfo_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Compatible_SetVendor(p, 4, a)
}

func CompatibleTunnelGroupInfo_Del(p *radius.Packet) {
	_Compatible_DelVendor(p, 4)
}

func CompatibleTunnelPassword_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Compatible_AddVendor(p, 5, a)
}

func CompatibleTunnelPassword_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Compatible_AddVendor(p, 5, a)
}

func CompatibleTunnelPassword_Get(p *radius.Packet) (value []byte) {
	value, _ = CompatibleTunnelPassword_Lookup(p)
	return
}

func CompatibleTunnelPassword_GetString(p *radius.Packet) (value string) {
	value, _ = CompatibleTunnelPassword_LookupString(p)
	return
}

func CompatibleTunnelPassword_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Compatible_GetsVendor(p, 5) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func CompatibleTunnelPassword_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Compatible_GetsVendor(p, 5) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func CompatibleTunnelPassword_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Compatible_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func CompatibleTunnelPassword_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Compatible_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func CompatibleTunnelPassword_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Compatible_SetVendor(p, 5, a)
}

func CompatibleTunnelPassword_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Compatible_SetVendor(p, 5, a)
}

func CompatibleTunnelPassword_Del(p *radius.Packet) {
	_Compatible_DelVendor(p, 5)
}

type CompatibleEcho uint32

var CompatibleEcho_Strings = map[CompatibleEcho]string{}

func (a CompatibleEcho) String() string {
	if str, ok := CompatibleEcho_Strings[a]; ok {
		return str
	}
	return "CompatibleEcho(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func CompatibleEcho_Add(p *radius.Packet, value CompatibleEcho) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Compatible_AddVendor(p, 6, a)
}

func CompatibleEcho_Get(p *radius.Packet) (value CompatibleEcho) {
	value, _ = CompatibleEcho_Lookup(p)
	return
}

func CompatibleEcho_Gets(p *radius.Packet) (values []CompatibleEcho, err error) {
	var i uint32
	for _, attr := range _Compatible_GetsVendor(p, 6) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, CompatibleEcho(i))
	}
	return
}

func CompatibleEcho_Lookup(p *radius.Packet) (value CompatibleEcho, err error) {
	a, ok := _Compatible_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = CompatibleEcho(i)
	return
}

func CompatibleEcho_Set(p *radius.Packet, value CompatibleEcho) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Compatible_SetVendor(p, 6, a)
}

func CompatibleEcho_Del(p *radius.Packet) {
	_Compatible_DelVendor(p, 6)
}

type CompatibleTunnelIPX uint32

var CompatibleTunnelIPX_Strings = map[CompatibleTunnelIPX]string{}

func (a CompatibleTunnelIPX) String() string {
	if str, ok := CompatibleTunnelIPX_Strings[a]; ok {
		return str
	}
	return "CompatibleTunnelIPX(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func CompatibleTunnelIPX_Add(p *radius.Packet, value CompatibleTunnelIPX) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Compatible_AddVendor(p, 7, a)
}

func CompatibleTunnelIPX_Get(p *radius.Packet) (value CompatibleTunnelIPX) {
	value, _ = CompatibleTunnelIPX_Lookup(p)
	return
}

func CompatibleTunnelIPX_Gets(p *radius.Packet) (values []CompatibleTunnelIPX, err error) {
	var i uint32
	for _, attr := range _Compatible_GetsVendor(p, 7) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, CompatibleTunnelIPX(i))
	}
	return
}

func CompatibleTunnelIPX_Lookup(p *radius.Packet) (value CompatibleTunnelIPX, err error) {
	a, ok := _Compatible_LookupVendor(p, 7)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = CompatibleTunnelIPX(i)
	return
}

func CompatibleTunnelIPX_Set(p *radius.Packet, value CompatibleTunnelIPX) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Compatible_SetVendor(p, 7, a)
}

func CompatibleTunnelIPX_Del(p *radius.Packet) {
	_Compatible_DelVendor(p, 7)
}
