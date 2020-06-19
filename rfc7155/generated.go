// Code generated by radius-dict-gen. DO NOT EDIT.

package rfc7155

import (
	"errors"

	"layeh.com/radius"
)

const (
	OriginatingLineInfo_Type radius.Type = 94
)

func OriginatingLineInfo_Add(p *radius.Packet, value []byte) (err error) {
	if len(value) != 2 {
		err = errors.New("invalid value length")
		return
	}
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(OriginatingLineInfo_Type, a)
	return
}

func OriginatingLineInfo_AddString(p *radius.Packet, value string) (err error) {
	if len(value) != 2 {
		err = errors.New("invalid value length")
		return
	}
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(OriginatingLineInfo_Type, a)
	return
}

func OriginatingLineInfo_Get(p *radius.Packet) (value []byte) {
	value, _ = OriginatingLineInfo_Lookup(p)
	return
}

func OriginatingLineInfo_GetString(p *radius.Packet) (value string) {
	value, _ = OriginatingLineInfo_LookupString(p)
	return
}

func OriginatingLineInfo_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, avp := range p.Attributes {
		if avp.Type != OriginatingLineInfo_Type {
			continue
		}
		attr := avp.Attribute
		i = radius.Bytes(attr)
		if err == nil && len(i) != 2 {
			err = errors.New("invalid value length")
		}
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func OriginatingLineInfo_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, avp := range p.Attributes {
		if avp.Type != OriginatingLineInfo_Type {
			continue
		}
		attr := avp.Attribute
		i = radius.String(attr)
		if err == nil && len(i) != 2 {
			err = errors.New("invalid value length")
		}
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func OriginatingLineInfo_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(OriginatingLineInfo_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	if err == nil && len(value) != 2 {
		err = errors.New("invalid value length")
	}
	return
}

func OriginatingLineInfo_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(OriginatingLineInfo_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	if err == nil && len(value) != 2 {
		err = errors.New("invalid value length")
	}
	return
}

func OriginatingLineInfo_Set(p *radius.Packet, value []byte) (err error) {
	if len(value) != 2 {
		err = errors.New("invalid value length")
		return
	}
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(OriginatingLineInfo_Type, a)
	return
}

func OriginatingLineInfo_SetString(p *radius.Packet, value string) (err error) {
	if len(value) != 2 {
		err = errors.New("invalid value length")
		return
	}
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(OriginatingLineInfo_Type, a)
	return
}

func OriginatingLineInfo_Del(p *radius.Packet) {
	p.Attributes.Del(OriginatingLineInfo_Type)
}
