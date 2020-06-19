// Code generated by radius-dict-gen. DO NOT EDIT.

package hillstone

import (
	"net"
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Hillstone_VendorID = 28557
)

func _Hillstone_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Hillstone_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Hillstone_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Hillstone_VendorID {
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

func _Hillstone_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Hillstone_VendorID {
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

func _Hillstone_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Hillstone_VendorID {
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
	return _Hillstone_AddVendor(p, typ, attr)
}

func _Hillstone_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Hillstone_VendorID {
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

type HillstoneUserVsysID uint32

var HillstoneUserVsysID_Strings = map[HillstoneUserVsysID]string{}

func (a HillstoneUserVsysID) String() string {
	if str, ok := HillstoneUserVsysID_Strings[a]; ok {
		return str
	}
	return "HillstoneUserVsysID(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func HillstoneUserVsysID_Add(p *radius.Packet, value HillstoneUserVsysID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Hillstone_AddVendor(p, 1, a)
}

func HillstoneUserVsysID_Get(p *radius.Packet) (value HillstoneUserVsysID) {
	value, _ = HillstoneUserVsysID_Lookup(p)
	return
}

func HillstoneUserVsysID_Gets(p *radius.Packet) (values []HillstoneUserVsysID, err error) {
	var i uint32
	for _, attr := range _Hillstone_GetsVendor(p, 1) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, HillstoneUserVsysID(i))
	}
	return
}

func HillstoneUserVsysID_Lookup(p *radius.Packet) (value HillstoneUserVsysID, err error) {
	a, ok := _Hillstone_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = HillstoneUserVsysID(i)
	return
}

func HillstoneUserVsysID_Set(p *radius.Packet, value HillstoneUserVsysID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Hillstone_SetVendor(p, 1, a)
}

func HillstoneUserVsysID_Del(p *radius.Packet) {
	_Hillstone_DelVendor(p, 1)
}

type HillstoneUserType uint32

const (
	HillstoneUserType_Value_HSUserL2tp     HillstoneUserType = 1
	HillstoneUserType_Value_HSUser8021x    HillstoneUserType = 2
	HillstoneUserType_Value_HSUserSmartvpn HillstoneUserType = 4
	HillstoneUserType_Value_HSUserNormal   HillstoneUserType = 8
	HillstoneUserType_Value_HSUserAdmin    HillstoneUserType = 16
)

var HillstoneUserType_Strings = map[HillstoneUserType]string{
	HillstoneUserType_Value_HSUserL2tp:     "HS-User-l2tp",
	HillstoneUserType_Value_HSUser8021x:    "HS-User-8021x",
	HillstoneUserType_Value_HSUserSmartvpn: "HS-User-smartvpn",
	HillstoneUserType_Value_HSUserNormal:   "HS-User-normal",
	HillstoneUserType_Value_HSUserAdmin:    "HS-User-Admin",
}

func (a HillstoneUserType) String() string {
	if str, ok := HillstoneUserType_Strings[a]; ok {
		return str
	}
	return "HillstoneUserType(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func HillstoneUserType_Add(p *radius.Packet, value HillstoneUserType) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Hillstone_AddVendor(p, 2, a)
}

func HillstoneUserType_Get(p *radius.Packet) (value HillstoneUserType) {
	value, _ = HillstoneUserType_Lookup(p)
	return
}

func HillstoneUserType_Gets(p *radius.Packet) (values []HillstoneUserType, err error) {
	var i uint32
	for _, attr := range _Hillstone_GetsVendor(p, 2) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, HillstoneUserType(i))
	}
	return
}

func HillstoneUserType_Lookup(p *radius.Packet) (value HillstoneUserType, err error) {
	a, ok := _Hillstone_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = HillstoneUserType(i)
	return
}

func HillstoneUserType_Set(p *radius.Packet, value HillstoneUserType) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Hillstone_SetVendor(p, 2, a)
}

func HillstoneUserType_Del(p *radius.Packet) {
	_Hillstone_DelVendor(p, 2)
}

type HillstoneUserAdminPrivilege uint32

var HillstoneUserAdminPrivilege_Strings = map[HillstoneUserAdminPrivilege]string{}

func (a HillstoneUserAdminPrivilege) String() string {
	if str, ok := HillstoneUserAdminPrivilege_Strings[a]; ok {
		return str
	}
	return "HillstoneUserAdminPrivilege(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func HillstoneUserAdminPrivilege_Add(p *radius.Packet, value HillstoneUserAdminPrivilege) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Hillstone_AddVendor(p, 3, a)
}

func HillstoneUserAdminPrivilege_Get(p *radius.Packet) (value HillstoneUserAdminPrivilege) {
	value, _ = HillstoneUserAdminPrivilege_Lookup(p)
	return
}

func HillstoneUserAdminPrivilege_Gets(p *radius.Packet) (values []HillstoneUserAdminPrivilege, err error) {
	var i uint32
	for _, attr := range _Hillstone_GetsVendor(p, 3) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, HillstoneUserAdminPrivilege(i))
	}
	return
}

func HillstoneUserAdminPrivilege_Lookup(p *radius.Packet) (value HillstoneUserAdminPrivilege, err error) {
	a, ok := _Hillstone_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = HillstoneUserAdminPrivilege(i)
	return
}

func HillstoneUserAdminPrivilege_Set(p *radius.Packet, value HillstoneUserAdminPrivilege) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Hillstone_SetVendor(p, 3, a)
}

func HillstoneUserAdminPrivilege_Del(p *radius.Packet) {
	_Hillstone_DelVendor(p, 3)
}

type HillstoneUserLoginType uint32

const (
	HillstoneUserLoginType_Value_HSAdminConsole HillstoneUserLoginType = 1
	HillstoneUserLoginType_Value_HSAdminTelnet  HillstoneUserLoginType = 2
	HillstoneUserLoginType_Value_HSAdminSSH     HillstoneUserLoginType = 4
	HillstoneUserLoginType_Value_HSAdminHTTP    HillstoneUserLoginType = 8
	HillstoneUserLoginType_Value_HSAdminHTTPS   HillstoneUserLoginType = 16
)

var HillstoneUserLoginType_Strings = map[HillstoneUserLoginType]string{
	HillstoneUserLoginType_Value_HSAdminConsole: "HS-Admin-Console",
	HillstoneUserLoginType_Value_HSAdminTelnet:  "HS-Admin-Telnet",
	HillstoneUserLoginType_Value_HSAdminSSH:     "HS-Admin-SSH",
	HillstoneUserLoginType_Value_HSAdminHTTP:    "HS-Admin-HTTP",
	HillstoneUserLoginType_Value_HSAdminHTTPS:   "HS-Admin-HTTPS",
}

func (a HillstoneUserLoginType) String() string {
	if str, ok := HillstoneUserLoginType_Strings[a]; ok {
		return str
	}
	return "HillstoneUserLoginType(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func HillstoneUserLoginType_Add(p *radius.Packet, value HillstoneUserLoginType) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Hillstone_AddVendor(p, 4, a)
}

func HillstoneUserLoginType_Get(p *radius.Packet) (value HillstoneUserLoginType) {
	value, _ = HillstoneUserLoginType_Lookup(p)
	return
}

func HillstoneUserLoginType_Gets(p *radius.Packet) (values []HillstoneUserLoginType, err error) {
	var i uint32
	for _, attr := range _Hillstone_GetsVendor(p, 4) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, HillstoneUserLoginType(i))
	}
	return
}

func HillstoneUserLoginType_Lookup(p *radius.Packet) (value HillstoneUserLoginType, err error) {
	a, ok := _Hillstone_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = HillstoneUserLoginType(i)
	return
}

func HillstoneUserLoginType_Set(p *radius.Packet, value HillstoneUserLoginType) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Hillstone_SetVendor(p, 4, a)
}

func HillstoneUserLoginType_Del(p *radius.Packet) {
	_Hillstone_DelVendor(p, 4)
}

func HillstoneUserMobileNumber_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Hillstone_AddVendor(p, 5, a)
}

func HillstoneUserMobileNumber_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Hillstone_AddVendor(p, 5, a)
}

func HillstoneUserMobileNumber_Get(p *radius.Packet) (value []byte) {
	value, _ = HillstoneUserMobileNumber_Lookup(p)
	return
}

func HillstoneUserMobileNumber_GetString(p *radius.Packet) (value string) {
	value, _ = HillstoneUserMobileNumber_LookupString(p)
	return
}

func HillstoneUserMobileNumber_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Hillstone_GetsVendor(p, 5) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func HillstoneUserMobileNumber_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Hillstone_GetsVendor(p, 5) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func HillstoneUserMobileNumber_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Hillstone_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func HillstoneUserMobileNumber_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Hillstone_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func HillstoneUserMobileNumber_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Hillstone_SetVendor(p, 5, a)
}

func HillstoneUserMobileNumber_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Hillstone_SetVendor(p, 5, a)
}

func HillstoneUserMobileNumber_Del(p *radius.Packet) {
	_Hillstone_DelVendor(p, 5)
}

type HillstoneUserMobileOperator uint32

const (
	HillstoneUserMobileOperator_Value_HSMobileChinaMobile  HillstoneUserMobileOperator = 1
	HillstoneUserMobileOperator_Value_HSMobileChinaUnicom  HillstoneUserMobileOperator = 2
	HillstoneUserMobileOperator_Value_HSMobileChinaTelecom HillstoneUserMobileOperator = 3
)

var HillstoneUserMobileOperator_Strings = map[HillstoneUserMobileOperator]string{
	HillstoneUserMobileOperator_Value_HSMobileChinaMobile:  "HS-Mobile-ChinaMobile",
	HillstoneUserMobileOperator_Value_HSMobileChinaUnicom:  "HS-Mobile-ChinaUnicom",
	HillstoneUserMobileOperator_Value_HSMobileChinaTelecom: "HS-Mobile-ChinaTelecom",
}

func (a HillstoneUserMobileOperator) String() string {
	if str, ok := HillstoneUserMobileOperator_Strings[a]; ok {
		return str
	}
	return "HillstoneUserMobileOperator(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func HillstoneUserMobileOperator_Add(p *radius.Packet, value HillstoneUserMobileOperator) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Hillstone_AddVendor(p, 6, a)
}

func HillstoneUserMobileOperator_Get(p *radius.Packet) (value HillstoneUserMobileOperator) {
	value, _ = HillstoneUserMobileOperator_Lookup(p)
	return
}

func HillstoneUserMobileOperator_Gets(p *radius.Packet) (values []HillstoneUserMobileOperator, err error) {
	var i uint32
	for _, attr := range _Hillstone_GetsVendor(p, 6) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, HillstoneUserMobileOperator(i))
	}
	return
}

func HillstoneUserMobileOperator_Lookup(p *radius.Packet) (value HillstoneUserMobileOperator, err error) {
	a, ok := _Hillstone_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = HillstoneUserMobileOperator(i)
	return
}

func HillstoneUserMobileOperator_Set(p *radius.Packet, value HillstoneUserMobileOperator) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Hillstone_SetVendor(p, 6, a)
}

func HillstoneUserMobileOperator_Del(p *radius.Packet) {
	_Hillstone_DelVendor(p, 6)
}

func HillstoneUserPolicyDstIPBegin_Add(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Hillstone_AddVendor(p, 7, a)
}

func HillstoneUserPolicyDstIPBegin_Get(p *radius.Packet) (value net.IP) {
	value, _ = HillstoneUserPolicyDstIPBegin_Lookup(p)
	return
}

func HillstoneUserPolicyDstIPBegin_Gets(p *radius.Packet) (values []net.IP, err error) {
	var i net.IP
	for _, attr := range _Hillstone_GetsVendor(p, 7) {
		i, err = radius.IPAddr(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func HillstoneUserPolicyDstIPBegin_Lookup(p *radius.Packet) (value net.IP, err error) {
	a, ok := _Hillstone_LookupVendor(p, 7)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value, err = radius.IPAddr(a)
	return
}

func HillstoneUserPolicyDstIPBegin_Set(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Hillstone_SetVendor(p, 7, a)
}

func HillstoneUserPolicyDstIPBegin_Del(p *radius.Packet) {
	_Hillstone_DelVendor(p, 7)
}

func HillstoneUserPolicyDstIPEnd_Add(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Hillstone_AddVendor(p, 8, a)
}

func HillstoneUserPolicyDstIPEnd_Get(p *radius.Packet) (value net.IP) {
	value, _ = HillstoneUserPolicyDstIPEnd_Lookup(p)
	return
}

func HillstoneUserPolicyDstIPEnd_Gets(p *radius.Packet) (values []net.IP, err error) {
	var i net.IP
	for _, attr := range _Hillstone_GetsVendor(p, 8) {
		i, err = radius.IPAddr(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func HillstoneUserPolicyDstIPEnd_Lookup(p *radius.Packet) (value net.IP, err error) {
	a, ok := _Hillstone_LookupVendor(p, 8)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value, err = radius.IPAddr(a)
	return
}

func HillstoneUserPolicyDstIPEnd_Set(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Hillstone_SetVendor(p, 8, a)
}

func HillstoneUserPolicyDstIPEnd_Del(p *radius.Packet) {
	_Hillstone_DelVendor(p, 8)
}

func HillstoneUserRoleBame_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Hillstone_AddVendor(p, 9, a)
}

func HillstoneUserRoleBame_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Hillstone_AddVendor(p, 9, a)
}

func HillstoneUserRoleBame_Get(p *radius.Packet) (value []byte) {
	value, _ = HillstoneUserRoleBame_Lookup(p)
	return
}

func HillstoneUserRoleBame_GetString(p *radius.Packet) (value string) {
	value, _ = HillstoneUserRoleBame_LookupString(p)
	return
}

func HillstoneUserRoleBame_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Hillstone_GetsVendor(p, 9) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func HillstoneUserRoleBame_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Hillstone_GetsVendor(p, 9) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func HillstoneUserRoleBame_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Hillstone_LookupVendor(p, 9)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func HillstoneUserRoleBame_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Hillstone_LookupVendor(p, 9)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func HillstoneUserRoleBame_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Hillstone_SetVendor(p, 9, a)
}

func HillstoneUserRoleBame_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Hillstone_SetVendor(p, 9, a)
}

func HillstoneUserRoleBame_Del(p *radius.Packet) {
	_Hillstone_DelVendor(p, 9)
}

func HillstoneVPNDHCPGateway_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Hillstone_AddVendor(p, 100, a)
}

func HillstoneVPNDHCPGateway_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Hillstone_AddVendor(p, 100, a)
}

func HillstoneVPNDHCPGateway_Get(p *radius.Packet) (value []byte) {
	value, _ = HillstoneVPNDHCPGateway_Lookup(p)
	return
}

func HillstoneVPNDHCPGateway_GetString(p *radius.Packet) (value string) {
	value, _ = HillstoneVPNDHCPGateway_LookupString(p)
	return
}

func HillstoneVPNDHCPGateway_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Hillstone_GetsVendor(p, 100) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func HillstoneVPNDHCPGateway_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Hillstone_GetsVendor(p, 100) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func HillstoneVPNDHCPGateway_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Hillstone_LookupVendor(p, 100)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func HillstoneVPNDHCPGateway_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Hillstone_LookupVendor(p, 100)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func HillstoneVPNDHCPGateway_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Hillstone_SetVendor(p, 100, a)
}

func HillstoneVPNDHCPGateway_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Hillstone_SetVendor(p, 100, a)
}

func HillstoneVPNDHCPGateway_Del(p *radius.Packet) {
	_Hillstone_DelVendor(p, 100)
}

func HillstoneVPNDHCPMask_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Hillstone_AddVendor(p, 101, a)
}

func HillstoneVPNDHCPMask_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Hillstone_AddVendor(p, 101, a)
}

func HillstoneVPNDHCPMask_Get(p *radius.Packet) (value []byte) {
	value, _ = HillstoneVPNDHCPMask_Lookup(p)
	return
}

func HillstoneVPNDHCPMask_GetString(p *radius.Packet) (value string) {
	value, _ = HillstoneVPNDHCPMask_LookupString(p)
	return
}

func HillstoneVPNDHCPMask_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Hillstone_GetsVendor(p, 101) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func HillstoneVPNDHCPMask_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Hillstone_GetsVendor(p, 101) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func HillstoneVPNDHCPMask_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Hillstone_LookupVendor(p, 101)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func HillstoneVPNDHCPMask_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Hillstone_LookupVendor(p, 101)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func HillstoneVPNDHCPMask_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Hillstone_SetVendor(p, 101, a)
}

func HillstoneVPNDHCPMask_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Hillstone_SetVendor(p, 101, a)
}

func HillstoneVPNDHCPMask_Del(p *radius.Packet) {
	_Hillstone_DelVendor(p, 101)
}

func HillstoneVPNDHCPPool_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Hillstone_AddVendor(p, 102, a)
}

func HillstoneVPNDHCPPool_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Hillstone_AddVendor(p, 102, a)
}

func HillstoneVPNDHCPPool_Get(p *radius.Packet) (value []byte) {
	value, _ = HillstoneVPNDHCPPool_Lookup(p)
	return
}

func HillstoneVPNDHCPPool_GetString(p *radius.Packet) (value string) {
	value, _ = HillstoneVPNDHCPPool_LookupString(p)
	return
}

func HillstoneVPNDHCPPool_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Hillstone_GetsVendor(p, 102) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func HillstoneVPNDHCPPool_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Hillstone_GetsVendor(p, 102) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func HillstoneVPNDHCPPool_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Hillstone_LookupVendor(p, 102)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func HillstoneVPNDHCPPool_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Hillstone_LookupVendor(p, 102)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func HillstoneVPNDHCPPool_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Hillstone_SetVendor(p, 102, a)
}

func HillstoneVPNDHCPPool_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Hillstone_SetVendor(p, 102, a)
}

func HillstoneVPNDHCPPool_Del(p *radius.Packet) {
	_Hillstone_DelVendor(p, 102)
}

func HillstoneVPNWINS_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Hillstone_AddVendor(p, 103, a)
}

func HillstoneVPNWINS_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Hillstone_AddVendor(p, 103, a)
}

func HillstoneVPNWINS_Get(p *radius.Packet) (value []byte) {
	value, _ = HillstoneVPNWINS_Lookup(p)
	return
}

func HillstoneVPNWINS_GetString(p *radius.Packet) (value string) {
	value, _ = HillstoneVPNWINS_LookupString(p)
	return
}

func HillstoneVPNWINS_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Hillstone_GetsVendor(p, 103) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func HillstoneVPNWINS_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Hillstone_GetsVendor(p, 103) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func HillstoneVPNWINS_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Hillstone_LookupVendor(p, 103)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func HillstoneVPNWINS_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Hillstone_LookupVendor(p, 103)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func HillstoneVPNWINS_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Hillstone_SetVendor(p, 103, a)
}

func HillstoneVPNWINS_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Hillstone_SetVendor(p, 103, a)
}

func HillstoneVPNWINS_Del(p *radius.Packet) {
	_Hillstone_DelVendor(p, 103)
}

func HillstoneVPNDNS_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Hillstone_AddVendor(p, 104, a)
}

func HillstoneVPNDNS_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Hillstone_AddVendor(p, 104, a)
}

func HillstoneVPNDNS_Get(p *radius.Packet) (value []byte) {
	value, _ = HillstoneVPNDNS_Lookup(p)
	return
}

func HillstoneVPNDNS_GetString(p *radius.Packet) (value string) {
	value, _ = HillstoneVPNDNS_LookupString(p)
	return
}

func HillstoneVPNDNS_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Hillstone_GetsVendor(p, 104) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func HillstoneVPNDNS_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Hillstone_GetsVendor(p, 104) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func HillstoneVPNDNS_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Hillstone_LookupVendor(p, 104)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func HillstoneVPNDNS_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Hillstone_LookupVendor(p, 104)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func HillstoneVPNDNS_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Hillstone_SetVendor(p, 104, a)
}

func HillstoneVPNDNS_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Hillstone_SetVendor(p, 104, a)
}

func HillstoneVPNDNS_Del(p *radius.Packet) {
	_Hillstone_DelVendor(p, 104)
}

func HillstoneVPNSplitRoute_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Hillstone_AddVendor(p, 105, a)
}

func HillstoneVPNSplitRoute_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Hillstone_AddVendor(p, 105, a)
}

func HillstoneVPNSplitRoute_Get(p *radius.Packet) (value []byte) {
	value, _ = HillstoneVPNSplitRoute_Lookup(p)
	return
}

func HillstoneVPNSplitRoute_GetString(p *radius.Packet) (value string) {
	value, _ = HillstoneVPNSplitRoute_LookupString(p)
	return
}

func HillstoneVPNSplitRoute_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Hillstone_GetsVendor(p, 105) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func HillstoneVPNSplitRoute_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Hillstone_GetsVendor(p, 105) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func HillstoneVPNSplitRoute_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Hillstone_LookupVendor(p, 105)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func HillstoneVPNSplitRoute_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Hillstone_LookupVendor(p, 105)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func HillstoneVPNSplitRoute_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Hillstone_SetVendor(p, 105, a)
}

func HillstoneVPNSplitRoute_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Hillstone_SetVendor(p, 105, a)
}

func HillstoneVPNSplitRoute_Del(p *radius.Packet) {
	_Hillstone_DelVendor(p, 105)
}

func HillstoneVPNTunnelIP_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Hillstone_AddVendor(p, 106, a)
}

func HillstoneVPNTunnelIP_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Hillstone_AddVendor(p, 106, a)
}

func HillstoneVPNTunnelIP_Get(p *radius.Packet) (value []byte) {
	value, _ = HillstoneVPNTunnelIP_Lookup(p)
	return
}

func HillstoneVPNTunnelIP_GetString(p *radius.Packet) (value string) {
	value, _ = HillstoneVPNTunnelIP_LookupString(p)
	return
}

func HillstoneVPNTunnelIP_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Hillstone_GetsVendor(p, 106) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func HillstoneVPNTunnelIP_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Hillstone_GetsVendor(p, 106) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func HillstoneVPNTunnelIP_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Hillstone_LookupVendor(p, 106)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func HillstoneVPNTunnelIP_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Hillstone_LookupVendor(p, 106)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func HillstoneVPNTunnelIP_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Hillstone_SetVendor(p, 106, a)
}

func HillstoneVPNTunnelIP_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Hillstone_SetVendor(p, 106, a)
}

func HillstoneVPNTunnelIP_Del(p *radius.Packet) {
	_Hillstone_DelVendor(p, 106)
}

type HillstoneVPNSNAT uint32

var HillstoneVPNSNAT_Strings = map[HillstoneVPNSNAT]string{}

func (a HillstoneVPNSNAT) String() string {
	if str, ok := HillstoneVPNSNAT_Strings[a]; ok {
		return str
	}
	return "HillstoneVPNSNAT(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func HillstoneVPNSNAT_Add(p *radius.Packet, value HillstoneVPNSNAT) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Hillstone_AddVendor(p, 107, a)
}

func HillstoneVPNSNAT_Get(p *radius.Packet) (value HillstoneVPNSNAT) {
	value, _ = HillstoneVPNSNAT_Lookup(p)
	return
}

func HillstoneVPNSNAT_Gets(p *radius.Packet) (values []HillstoneVPNSNAT, err error) {
	var i uint32
	for _, attr := range _Hillstone_GetsVendor(p, 107) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, HillstoneVPNSNAT(i))
	}
	return
}

func HillstoneVPNSNAT_Lookup(p *radius.Packet) (value HillstoneVPNSNAT, err error) {
	a, ok := _Hillstone_LookupVendor(p, 107)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = HillstoneVPNSNAT(i)
	return
}

func HillstoneVPNSNAT_Set(p *radius.Packet, value HillstoneVPNSNAT) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Hillstone_SetVendor(p, 107, a)
}

func HillstoneVPNSNAT_Del(p *radius.Packet) {
	_Hillstone_DelVendor(p, 107)
}
