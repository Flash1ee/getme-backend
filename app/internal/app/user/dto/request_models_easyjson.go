// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package dto

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson7df0efccDecodeGetmeBackendInternalAppUserDto(in *jlexer.Lexer, out *UserAuthCheckRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int64(in.Int64())
		case "auth_date":
			out.AuthDate = int64(in.Int64())
		case "first_name":
			out.FirstName = string(in.String())
		case "last_name":
			out.LastName = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "photo_url":
			out.Avatar = string(in.String())
		case "hash":
			out.Hash = string(in.String())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7df0efccEncodeGetmeBackendInternalAppUserDto(out *jwriter.Writer, in UserAuthCheckRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"auth_date\":"
		out.RawString(prefix)
		out.Int64(int64(in.AuthDate))
	}
	{
		const prefix string = ",\"first_name\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"last_name\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"photo_url\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	{
		const prefix string = ",\"hash\":"
		out.RawString(prefix)
		out.String(string(in.Hash))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserAuthCheckRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7df0efccEncodeGetmeBackendInternalAppUserDto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserAuthCheckRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7df0efccEncodeGetmeBackendInternalAppUserDto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserAuthCheckRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7df0efccDecodeGetmeBackendInternalAppUserDto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserAuthCheckRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7df0efccDecodeGetmeBackendInternalAppUserDto(l, v)
}