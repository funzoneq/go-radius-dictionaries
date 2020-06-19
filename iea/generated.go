// Code generated by radius-dict-gen. DO NOT EDIT.

package iea

import (
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_IEASoftware_VendorID = 24023
)

func _IEASoftware_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_IEASoftware_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _IEASoftware_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _IEASoftware_VendorID {
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

func _IEASoftware_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _IEASoftware_VendorID {
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

func _IEASoftware_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _IEASoftware_VendorID {
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
	return _IEASoftware_AddVendor(p, typ, attr)
}

func _IEASoftware_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _IEASoftware_VendorID {
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

func AMInterruptHTMLFile_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _IEASoftware_AddVendor(p, 1, a)
}

func AMInterruptHTMLFile_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _IEASoftware_AddVendor(p, 1, a)
}

func AMInterruptHTMLFile_Get(p *radius.Packet) (value []byte) {
	value, _ = AMInterruptHTMLFile_Lookup(p)
	return
}

func AMInterruptHTMLFile_GetString(p *radius.Packet) (value string) {
	value, _ = AMInterruptHTMLFile_LookupString(p)
	return
}

func AMInterruptHTMLFile_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _IEASoftware_GetsVendor(p, 1) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func AMInterruptHTMLFile_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _IEASoftware_GetsVendor(p, 1) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func AMInterruptHTMLFile_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _IEASoftware_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func AMInterruptHTMLFile_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _IEASoftware_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func AMInterruptHTMLFile_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _IEASoftware_SetVendor(p, 1, a)
}

func AMInterruptHTMLFile_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _IEASoftware_SetVendor(p, 1, a)
}

func AMInterruptHTMLFile_Del(p *radius.Packet) {
	_IEASoftware_DelVendor(p, 1)
}

type AMInterruptInterval uint32

var AMInterruptInterval_Strings = map[AMInterruptInterval]string{}

func (a AMInterruptInterval) String() string {
	if str, ok := AMInterruptInterval_Strings[a]; ok {
		return str
	}
	return "AMInterruptInterval(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func AMInterruptInterval_Add(p *radius.Packet, value AMInterruptInterval) (err error) {
	a := radius.NewInteger(uint32(value))
	return _IEASoftware_AddVendor(p, 2, a)
}

func AMInterruptInterval_Get(p *radius.Packet) (value AMInterruptInterval) {
	value, _ = AMInterruptInterval_Lookup(p)
	return
}

func AMInterruptInterval_Gets(p *radius.Packet) (values []AMInterruptInterval, err error) {
	var i uint32
	for _, attr := range _IEASoftware_GetsVendor(p, 2) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AMInterruptInterval(i))
	}
	return
}

func AMInterruptInterval_Lookup(p *radius.Packet) (value AMInterruptInterval, err error) {
	a, ok := _IEASoftware_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AMInterruptInterval(i)
	return
}

func AMInterruptInterval_Set(p *radius.Packet, value AMInterruptInterval) (err error) {
	a := radius.NewInteger(uint32(value))
	return _IEASoftware_SetVendor(p, 2, a)
}

func AMInterruptInterval_Del(p *radius.Packet) {
	_IEASoftware_DelVendor(p, 2)
}

type AMInterruptTimeout uint32

var AMInterruptTimeout_Strings = map[AMInterruptTimeout]string{}

func (a AMInterruptTimeout) String() string {
	if str, ok := AMInterruptTimeout_Strings[a]; ok {
		return str
	}
	return "AMInterruptTimeout(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func AMInterruptTimeout_Add(p *radius.Packet, value AMInterruptTimeout) (err error) {
	a := radius.NewInteger(uint32(value))
	return _IEASoftware_AddVendor(p, 3, a)
}

func AMInterruptTimeout_Get(p *radius.Packet) (value AMInterruptTimeout) {
	value, _ = AMInterruptTimeout_Lookup(p)
	return
}

func AMInterruptTimeout_Gets(p *radius.Packet) (values []AMInterruptTimeout, err error) {
	var i uint32
	for _, attr := range _IEASoftware_GetsVendor(p, 3) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AMInterruptTimeout(i))
	}
	return
}

func AMInterruptTimeout_Lookup(p *radius.Packet) (value AMInterruptTimeout, err error) {
	a, ok := _IEASoftware_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AMInterruptTimeout(i)
	return
}

func AMInterruptTimeout_Set(p *radius.Packet, value AMInterruptTimeout) (err error) {
	a := radius.NewInteger(uint32(value))
	return _IEASoftware_SetVendor(p, 3, a)
}

func AMInterruptTimeout_Del(p *radius.Packet) {
	_IEASoftware_DelVendor(p, 3)
}

func AMStatusHTMLFile_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _IEASoftware_AddVendor(p, 4, a)
}

func AMStatusHTMLFile_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _IEASoftware_AddVendor(p, 4, a)
}

func AMStatusHTMLFile_Get(p *radius.Packet) (value []byte) {
	value, _ = AMStatusHTMLFile_Lookup(p)
	return
}

func AMStatusHTMLFile_GetString(p *radius.Packet) (value string) {
	value, _ = AMStatusHTMLFile_LookupString(p)
	return
}

func AMStatusHTMLFile_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _IEASoftware_GetsVendor(p, 4) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func AMStatusHTMLFile_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _IEASoftware_GetsVendor(p, 4) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func AMStatusHTMLFile_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _IEASoftware_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func AMStatusHTMLFile_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _IEASoftware_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func AMStatusHTMLFile_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _IEASoftware_SetVendor(p, 4, a)
}

func AMStatusHTMLFile_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _IEASoftware_SetVendor(p, 4, a)
}

func AMStatusHTMLFile_Del(p *radius.Packet) {
	_IEASoftware_DelVendor(p, 4)
}

type AMHTTPProxyPort uint32

var AMHTTPProxyPort_Strings = map[AMHTTPProxyPort]string{}

func (a AMHTTPProxyPort) String() string {
	if str, ok := AMHTTPProxyPort_Strings[a]; ok {
		return str
	}
	return "AMHTTPProxyPort(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func AMHTTPProxyPort_Add(p *radius.Packet, value AMHTTPProxyPort) (err error) {
	a := radius.NewInteger(uint32(value))
	return _IEASoftware_AddVendor(p, 5, a)
}

func AMHTTPProxyPort_Get(p *radius.Packet) (value AMHTTPProxyPort) {
	value, _ = AMHTTPProxyPort_Lookup(p)
	return
}

func AMHTTPProxyPort_Gets(p *radius.Packet) (values []AMHTTPProxyPort, err error) {
	var i uint32
	for _, attr := range _IEASoftware_GetsVendor(p, 5) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AMHTTPProxyPort(i))
	}
	return
}

func AMHTTPProxyPort_Lookup(p *radius.Packet) (value AMHTTPProxyPort, err error) {
	a, ok := _IEASoftware_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AMHTTPProxyPort(i)
	return
}

func AMHTTPProxyPort_Set(p *radius.Packet, value AMHTTPProxyPort) (err error) {
	a := radius.NewInteger(uint32(value))
	return _IEASoftware_SetVendor(p, 5, a)
}

func AMHTTPProxyPort_Del(p *radius.Packet) {
	_IEASoftware_DelVendor(p, 5)
}
