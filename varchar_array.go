package pgtype

import (
	"io"
)

type VarcharArray TextArray

func (dst *VarcharArray) Set(src interface{}) error {
	return (*TextArray)(dst).Set(src)
}

func (dst *VarcharArray) Get() interface{} {
	return (*TextArray)(dst).Get()
}

func (src *VarcharArray) AssignTo(dst interface{}) error {
	return (*TextArray)(src).AssignTo(dst)
}

func (dst *VarcharArray) DecodeText(src []byte) error {
	return (*TextArray)(dst).DecodeText(src)
}

func (dst *VarcharArray) DecodeBinary(src []byte) error {
	return (*TextArray)(dst).DecodeBinary(src)
}

func (src *VarcharArray) EncodeText(w io.Writer) (bool, error) {
	return (*TextArray)(src).EncodeText(w)
}

func (src *VarcharArray) EncodeBinary(w io.Writer) (bool, error) {
	return (*TextArray)(src).encodeBinary(w, VarcharOid)
}