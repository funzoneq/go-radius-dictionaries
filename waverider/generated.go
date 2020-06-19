// Code generated by radius-dict-gen. DO NOT EDIT.

package waverider

import (
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Waverider_VendorID = 2979
)

func _Waverider_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Waverider_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Waverider_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Waverider_VendorID {
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

func _Waverider_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Waverider_VendorID {
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

func _Waverider_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Waverider_VendorID {
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
	return _Waverider_AddVendor(p, typ, attr)
}

func _Waverider_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Waverider_VendorID {
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

type WaveriderGradeOfService uint32

const (
	WaveriderGradeOfService_Value_Be     WaveriderGradeOfService = 1
	WaveriderGradeOfService_Value_Bronze WaveriderGradeOfService = 2
	WaveriderGradeOfService_Value_Silver WaveriderGradeOfService = 3
	WaveriderGradeOfService_Value_Gold   WaveriderGradeOfService = 4
)

var WaveriderGradeOfService_Strings = map[WaveriderGradeOfService]string{
	WaveriderGradeOfService_Value_Be:     "be",
	WaveriderGradeOfService_Value_Bronze: "bronze",
	WaveriderGradeOfService_Value_Silver: "silver",
	WaveriderGradeOfService_Value_Gold:   "gold",
}

func (a WaveriderGradeOfService) String() string {
	if str, ok := WaveriderGradeOfService_Strings[a]; ok {
		return str
	}
	return "WaveriderGradeOfService(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func WaveriderGradeOfService_Add(p *radius.Packet, value WaveriderGradeOfService) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Waverider_AddVendor(p, 1, a)
}

func WaveriderGradeOfService_Get(p *radius.Packet) (value WaveriderGradeOfService) {
	value, _ = WaveriderGradeOfService_Lookup(p)
	return
}

func WaveriderGradeOfService_Gets(p *radius.Packet) (values []WaveriderGradeOfService, err error) {
	var i uint32
	for _, attr := range _Waverider_GetsVendor(p, 1) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, WaveriderGradeOfService(i))
	}
	return
}

func WaveriderGradeOfService_Lookup(p *radius.Packet) (value WaveriderGradeOfService, err error) {
	a, ok := _Waverider_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = WaveriderGradeOfService(i)
	return
}

func WaveriderGradeOfService_Set(p *radius.Packet, value WaveriderGradeOfService) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Waverider_SetVendor(p, 1, a)
}

func WaveriderGradeOfService_Del(p *radius.Packet) {
	_Waverider_DelVendor(p, 1)
}

type WaveriderPriorityEnabled uint32

const (
	WaveriderPriorityEnabled_Value_Disabled WaveriderPriorityEnabled = 0
	WaveriderPriorityEnabled_Value_Enabled  WaveriderPriorityEnabled = 1
)

var WaveriderPriorityEnabled_Strings = map[WaveriderPriorityEnabled]string{
	WaveriderPriorityEnabled_Value_Disabled: "disabled",
	WaveriderPriorityEnabled_Value_Enabled:  "enabled",
}

func (a WaveriderPriorityEnabled) String() string {
	if str, ok := WaveriderPriorityEnabled_Strings[a]; ok {
		return str
	}
	return "WaveriderPriorityEnabled(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func WaveriderPriorityEnabled_Add(p *radius.Packet, value WaveriderPriorityEnabled) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Waverider_AddVendor(p, 2, a)
}

func WaveriderPriorityEnabled_Get(p *radius.Packet) (value WaveriderPriorityEnabled) {
	value, _ = WaveriderPriorityEnabled_Lookup(p)
	return
}

func WaveriderPriorityEnabled_Gets(p *radius.Packet) (values []WaveriderPriorityEnabled, err error) {
	var i uint32
	for _, attr := range _Waverider_GetsVendor(p, 2) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, WaveriderPriorityEnabled(i))
	}
	return
}

func WaveriderPriorityEnabled_Lookup(p *radius.Packet) (value WaveriderPriorityEnabled, err error) {
	a, ok := _Waverider_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = WaveriderPriorityEnabled(i)
	return
}

func WaveriderPriorityEnabled_Set(p *radius.Packet, value WaveriderPriorityEnabled) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Waverider_SetVendor(p, 2, a)
}

func WaveriderPriorityEnabled_Del(p *radius.Packet) {
	_Waverider_DelVendor(p, 2)
}

func WaveriderAuthenticationKey_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Waverider_AddVendor(p, 3, a)
}

func WaveriderAuthenticationKey_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Waverider_AddVendor(p, 3, a)
}

func WaveriderAuthenticationKey_Get(p *radius.Packet) (value []byte) {
	value, _ = WaveriderAuthenticationKey_Lookup(p)
	return
}

func WaveriderAuthenticationKey_GetString(p *radius.Packet) (value string) {
	value, _ = WaveriderAuthenticationKey_LookupString(p)
	return
}

func WaveriderAuthenticationKey_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Waverider_GetsVendor(p, 3) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WaveriderAuthenticationKey_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Waverider_GetsVendor(p, 3) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WaveriderAuthenticationKey_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Waverider_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func WaveriderAuthenticationKey_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Waverider_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func WaveriderAuthenticationKey_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Waverider_SetVendor(p, 3, a)
}

func WaveriderAuthenticationKey_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Waverider_SetVendor(p, 3, a)
}

func WaveriderAuthenticationKey_Del(p *radius.Packet) {
	_Waverider_DelVendor(p, 3)
}

func WaveriderCurrentPassword_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Waverider_AddVendor(p, 5, a)
}

func WaveriderCurrentPassword_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Waverider_AddVendor(p, 5, a)
}

func WaveriderCurrentPassword_Get(p *radius.Packet) (value []byte) {
	value, _ = WaveriderCurrentPassword_Lookup(p)
	return
}

func WaveriderCurrentPassword_GetString(p *radius.Packet) (value string) {
	value, _ = WaveriderCurrentPassword_LookupString(p)
	return
}

func WaveriderCurrentPassword_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Waverider_GetsVendor(p, 5) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WaveriderCurrentPassword_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Waverider_GetsVendor(p, 5) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WaveriderCurrentPassword_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Waverider_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func WaveriderCurrentPassword_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Waverider_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func WaveriderCurrentPassword_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Waverider_SetVendor(p, 5, a)
}

func WaveriderCurrentPassword_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Waverider_SetVendor(p, 5, a)
}

func WaveriderCurrentPassword_Del(p *radius.Packet) {
	_Waverider_DelVendor(p, 5)
}

func WaveriderNewPassword_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Waverider_AddVendor(p, 6, a)
}

func WaveriderNewPassword_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Waverider_AddVendor(p, 6, a)
}

func WaveriderNewPassword_Get(p *radius.Packet) (value []byte) {
	value, _ = WaveriderNewPassword_Lookup(p)
	return
}

func WaveriderNewPassword_GetString(p *radius.Packet) (value string) {
	value, _ = WaveriderNewPassword_LookupString(p)
	return
}

func WaveriderNewPassword_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Waverider_GetsVendor(p, 6) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WaveriderNewPassword_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Waverider_GetsVendor(p, 6) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WaveriderNewPassword_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Waverider_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func WaveriderNewPassword_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Waverider_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func WaveriderNewPassword_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Waverider_SetVendor(p, 6, a)
}

func WaveriderNewPassword_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Waverider_SetVendor(p, 6, a)
}

func WaveriderNewPassword_Del(p *radius.Packet) {
	_Waverider_DelVendor(p, 6)
}

type WaveriderRadioFrequency uint32

const (
	WaveriderRadioFrequency_Value_Auto    WaveriderRadioFrequency = 1
	WaveriderRadioFrequency_Value_Nomadic WaveriderRadioFrequency = 2
	WaveriderRadioFrequency_Value_F9050   WaveriderRadioFrequency = 3
	WaveriderRadioFrequency_Value_F9116   WaveriderRadioFrequency = 4
	WaveriderRadioFrequency_Value_F9184   WaveriderRadioFrequency = 5
	WaveriderRadioFrequency_Value_F9250   WaveriderRadioFrequency = 6
	WaveriderRadioFrequency_Value_F9084   WaveriderRadioFrequency = 7
	WaveriderRadioFrequency_Value_F9150   WaveriderRadioFrequency = 8
	WaveriderRadioFrequency_Value_F9216   WaveriderRadioFrequency = 9
)

var WaveriderRadioFrequency_Strings = map[WaveriderRadioFrequency]string{
	WaveriderRadioFrequency_Value_Auto:    "auto",
	WaveriderRadioFrequency_Value_Nomadic: "nomadic",
	WaveriderRadioFrequency_Value_F9050:   "f_9050",
	WaveriderRadioFrequency_Value_F9116:   "f_9116",
	WaveriderRadioFrequency_Value_F9184:   "f_9184",
	WaveriderRadioFrequency_Value_F9250:   "f_9250",
	WaveriderRadioFrequency_Value_F9084:   "f_9084",
	WaveriderRadioFrequency_Value_F9150:   "f_9150",
	WaveriderRadioFrequency_Value_F9216:   "f_9216",
}

func (a WaveriderRadioFrequency) String() string {
	if str, ok := WaveriderRadioFrequency_Strings[a]; ok {
		return str
	}
	return "WaveriderRadioFrequency(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func WaveriderRadioFrequency_Add(p *radius.Packet, value WaveriderRadioFrequency) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Waverider_AddVendor(p, 7, a)
}

func WaveriderRadioFrequency_Get(p *radius.Packet) (value WaveriderRadioFrequency) {
	value, _ = WaveriderRadioFrequency_Lookup(p)
	return
}

func WaveriderRadioFrequency_Gets(p *radius.Packet) (values []WaveriderRadioFrequency, err error) {
	var i uint32
	for _, attr := range _Waverider_GetsVendor(p, 7) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, WaveriderRadioFrequency(i))
	}
	return
}

func WaveriderRadioFrequency_Lookup(p *radius.Packet) (value WaveriderRadioFrequency, err error) {
	a, ok := _Waverider_LookupVendor(p, 7)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = WaveriderRadioFrequency(i)
	return
}

func WaveriderRadioFrequency_Set(p *radius.Packet, value WaveriderRadioFrequency) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Waverider_SetVendor(p, 7, a)
}

func WaveriderRadioFrequency_Del(p *radius.Packet) {
	_Waverider_DelVendor(p, 7)
}

func WaveriderSNMPReadCommunity_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Waverider_AddVendor(p, 8, a)
}

func WaveriderSNMPReadCommunity_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Waverider_AddVendor(p, 8, a)
}

func WaveriderSNMPReadCommunity_Get(p *radius.Packet) (value []byte) {
	value, _ = WaveriderSNMPReadCommunity_Lookup(p)
	return
}

func WaveriderSNMPReadCommunity_GetString(p *radius.Packet) (value string) {
	value, _ = WaveriderSNMPReadCommunity_LookupString(p)
	return
}

func WaveriderSNMPReadCommunity_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Waverider_GetsVendor(p, 8) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WaveriderSNMPReadCommunity_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Waverider_GetsVendor(p, 8) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WaveriderSNMPReadCommunity_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Waverider_LookupVendor(p, 8)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func WaveriderSNMPReadCommunity_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Waverider_LookupVendor(p, 8)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func WaveriderSNMPReadCommunity_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Waverider_SetVendor(p, 8, a)
}

func WaveriderSNMPReadCommunity_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Waverider_SetVendor(p, 8, a)
}

func WaveriderSNMPReadCommunity_Del(p *radius.Packet) {
	_Waverider_DelVendor(p, 8)
}

func WaveriderSNMPWriteCommunity_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Waverider_AddVendor(p, 9, a)
}

func WaveriderSNMPWriteCommunity_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Waverider_AddVendor(p, 9, a)
}

func WaveriderSNMPWriteCommunity_Get(p *radius.Packet) (value []byte) {
	value, _ = WaveriderSNMPWriteCommunity_Lookup(p)
	return
}

func WaveriderSNMPWriteCommunity_GetString(p *radius.Packet) (value string) {
	value, _ = WaveriderSNMPWriteCommunity_LookupString(p)
	return
}

func WaveriderSNMPWriteCommunity_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Waverider_GetsVendor(p, 9) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WaveriderSNMPWriteCommunity_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Waverider_GetsVendor(p, 9) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WaveriderSNMPWriteCommunity_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Waverider_LookupVendor(p, 9)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func WaveriderSNMPWriteCommunity_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Waverider_LookupVendor(p, 9)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func WaveriderSNMPWriteCommunity_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Waverider_SetVendor(p, 9, a)
}

func WaveriderSNMPWriteCommunity_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Waverider_SetVendor(p, 9, a)
}

func WaveriderSNMPWriteCommunity_Del(p *radius.Packet) {
	_Waverider_DelVendor(p, 9)
}

func WaveriderSNMPTrapServer_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Waverider_AddVendor(p, 10, a)
}

func WaveriderSNMPTrapServer_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Waverider_AddVendor(p, 10, a)
}

func WaveriderSNMPTrapServer_Get(p *radius.Packet) (value []byte) {
	value, _ = WaveriderSNMPTrapServer_Lookup(p)
	return
}

func WaveriderSNMPTrapServer_GetString(p *radius.Packet) (value string) {
	value, _ = WaveriderSNMPTrapServer_LookupString(p)
	return
}

func WaveriderSNMPTrapServer_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Waverider_GetsVendor(p, 10) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WaveriderSNMPTrapServer_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Waverider_GetsVendor(p, 10) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WaveriderSNMPTrapServer_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Waverider_LookupVendor(p, 10)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func WaveriderSNMPTrapServer_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Waverider_LookupVendor(p, 10)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func WaveriderSNMPTrapServer_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Waverider_SetVendor(p, 10, a)
}

func WaveriderSNMPTrapServer_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Waverider_SetVendor(p, 10, a)
}

func WaveriderSNMPTrapServer_Del(p *radius.Packet) {
	_Waverider_DelVendor(p, 10)
}

func WaveriderSNMPContact_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Waverider_AddVendor(p, 11, a)
}

func WaveriderSNMPContact_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Waverider_AddVendor(p, 11, a)
}

func WaveriderSNMPContact_Get(p *radius.Packet) (value []byte) {
	value, _ = WaveriderSNMPContact_Lookup(p)
	return
}

func WaveriderSNMPContact_GetString(p *radius.Packet) (value string) {
	value, _ = WaveriderSNMPContact_LookupString(p)
	return
}

func WaveriderSNMPContact_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Waverider_GetsVendor(p, 11) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WaveriderSNMPContact_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Waverider_GetsVendor(p, 11) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WaveriderSNMPContact_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Waverider_LookupVendor(p, 11)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func WaveriderSNMPContact_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Waverider_LookupVendor(p, 11)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func WaveriderSNMPContact_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Waverider_SetVendor(p, 11, a)
}

func WaveriderSNMPContact_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Waverider_SetVendor(p, 11, a)
}

func WaveriderSNMPContact_Del(p *radius.Packet) {
	_Waverider_DelVendor(p, 11)
}

func WaveriderSNMPLocation_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Waverider_AddVendor(p, 12, a)
}

func WaveriderSNMPLocation_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Waverider_AddVendor(p, 12, a)
}

func WaveriderSNMPLocation_Get(p *radius.Packet) (value []byte) {
	value, _ = WaveriderSNMPLocation_Lookup(p)
	return
}

func WaveriderSNMPLocation_GetString(p *radius.Packet) (value string) {
	value, _ = WaveriderSNMPLocation_LookupString(p)
	return
}

func WaveriderSNMPLocation_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Waverider_GetsVendor(p, 12) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WaveriderSNMPLocation_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Waverider_GetsVendor(p, 12) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WaveriderSNMPLocation_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Waverider_LookupVendor(p, 12)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func WaveriderSNMPLocation_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Waverider_LookupVendor(p, 12)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func WaveriderSNMPLocation_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Waverider_SetVendor(p, 12, a)
}

func WaveriderSNMPLocation_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Waverider_SetVendor(p, 12, a)
}

func WaveriderSNMPLocation_Del(p *radius.Packet) {
	_Waverider_DelVendor(p, 12)
}

func WaveriderSNMPName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Waverider_AddVendor(p, 13, a)
}

func WaveriderSNMPName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Waverider_AddVendor(p, 13, a)
}

func WaveriderSNMPName_Get(p *radius.Packet) (value []byte) {
	value, _ = WaveriderSNMPName_Lookup(p)
	return
}

func WaveriderSNMPName_GetString(p *radius.Packet) (value string) {
	value, _ = WaveriderSNMPName_LookupString(p)
	return
}

func WaveriderSNMPName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Waverider_GetsVendor(p, 13) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WaveriderSNMPName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Waverider_GetsVendor(p, 13) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WaveriderSNMPName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Waverider_LookupVendor(p, 13)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func WaveriderSNMPName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Waverider_LookupVendor(p, 13)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func WaveriderSNMPName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Waverider_SetVendor(p, 13, a)
}

func WaveriderSNMPName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Waverider_SetVendor(p, 13, a)
}

func WaveriderSNMPName_Del(p *radius.Packet) {
	_Waverider_DelVendor(p, 13)
}

type WaveriderMaxCustomers uint32

var WaveriderMaxCustomers_Strings = map[WaveriderMaxCustomers]string{}

func (a WaveriderMaxCustomers) String() string {
	if str, ok := WaveriderMaxCustomers_Strings[a]; ok {
		return str
	}
	return "WaveriderMaxCustomers(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func WaveriderMaxCustomers_Add(p *radius.Packet, value WaveriderMaxCustomers) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Waverider_AddVendor(p, 14, a)
}

func WaveriderMaxCustomers_Get(p *radius.Packet) (value WaveriderMaxCustomers) {
	value, _ = WaveriderMaxCustomers_Lookup(p)
	return
}

func WaveriderMaxCustomers_Gets(p *radius.Packet) (values []WaveriderMaxCustomers, err error) {
	var i uint32
	for _, attr := range _Waverider_GetsVendor(p, 14) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, WaveriderMaxCustomers(i))
	}
	return
}

func WaveriderMaxCustomers_Lookup(p *radius.Packet) (value WaveriderMaxCustomers, err error) {
	a, ok := _Waverider_LookupVendor(p, 14)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = WaveriderMaxCustomers(i)
	return
}

func WaveriderMaxCustomers_Set(p *radius.Packet, value WaveriderMaxCustomers) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Waverider_SetVendor(p, 14, a)
}

func WaveriderMaxCustomers_Del(p *radius.Packet) {
	_Waverider_DelVendor(p, 14)
}

type WaveriderRfPower uint32

const (
	WaveriderRfPower_Value_P15 WaveriderRfPower = 1
	WaveriderRfPower_Value_P16 WaveriderRfPower = 2
	WaveriderRfPower_Value_P17 WaveriderRfPower = 3
	WaveriderRfPower_Value_P18 WaveriderRfPower = 4
	WaveriderRfPower_Value_P19 WaveriderRfPower = 5
	WaveriderRfPower_Value_P20 WaveriderRfPower = 6
	WaveriderRfPower_Value_P21 WaveriderRfPower = 7
	WaveriderRfPower_Value_P22 WaveriderRfPower = 8
	WaveriderRfPower_Value_P23 WaveriderRfPower = 9
	WaveriderRfPower_Value_P24 WaveriderRfPower = 10
	WaveriderRfPower_Value_P25 WaveriderRfPower = 11
	WaveriderRfPower_Value_P26 WaveriderRfPower = 12
)

var WaveriderRfPower_Strings = map[WaveriderRfPower]string{
	WaveriderRfPower_Value_P15: "p_15",
	WaveriderRfPower_Value_P16: "p_16",
	WaveriderRfPower_Value_P17: "p_17",
	WaveriderRfPower_Value_P18: "p_18",
	WaveriderRfPower_Value_P19: "p_19",
	WaveriderRfPower_Value_P20: "p_20",
	WaveriderRfPower_Value_P21: "p_21",
	WaveriderRfPower_Value_P22: "p_22",
	WaveriderRfPower_Value_P23: "p_23",
	WaveriderRfPower_Value_P24: "p_24",
	WaveriderRfPower_Value_P25: "p_25",
	WaveriderRfPower_Value_P26: "p_26",
}

func (a WaveriderRfPower) String() string {
	if str, ok := WaveriderRfPower_Strings[a]; ok {
		return str
	}
	return "WaveriderRfPower(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func WaveriderRfPower_Add(p *radius.Packet, value WaveriderRfPower) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Waverider_AddVendor(p, 15, a)
}

func WaveriderRfPower_Get(p *radius.Packet) (value WaveriderRfPower) {
	value, _ = WaveriderRfPower_Lookup(p)
	return
}

func WaveriderRfPower_Gets(p *radius.Packet) (values []WaveriderRfPower, err error) {
	var i uint32
	for _, attr := range _Waverider_GetsVendor(p, 15) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, WaveriderRfPower(i))
	}
	return
}

func WaveriderRfPower_Lookup(p *radius.Packet) (value WaveriderRfPower, err error) {
	a, ok := _Waverider_LookupVendor(p, 15)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = WaveriderRfPower(i)
	return
}

func WaveriderRfPower_Set(p *radius.Packet, value WaveriderRfPower) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Waverider_SetVendor(p, 15, a)
}

func WaveriderRfPower_Del(p *radius.Packet) {
	_Waverider_DelVendor(p, 15)
}
