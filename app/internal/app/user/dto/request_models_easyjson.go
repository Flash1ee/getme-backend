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

func easyjson7df0efccDecodeGetmeBackendInternalAppUserDto(in *jlexer.Lexer, out *RequestUserUpdate) {
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
		case "first_name":
			out.FirstName = string(in.String())
		case "last_name":
			out.LastName = string(in.String())
		case "about":
			out.About = string(in.String())
		case "tg_tag":
			out.TgTag = string(in.String())
		case "skills":
			if in.IsNull() {
				in.Skip()
				out.Skills = nil
			} else {
				in.Delim('[')
				if out.Skills == nil {
					if !in.IsDelim(']') {
						out.Skills = make([]string, 0, 4)
					} else {
						out.Skills = []string{}
					}
				} else {
					out.Skills = (out.Skills)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Skills = append(out.Skills, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson7df0efccEncodeGetmeBackendInternalAppUserDto(out *jwriter.Writer, in RequestUserUpdate) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FirstName != "" {
		const prefix string = ",\"first_name\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.FirstName))
	}
	if in.LastName != "" {
		const prefix string = ",\"last_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LastName))
	}
	if in.About != "" {
		const prefix string = ",\"about\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.About))
	}
	if in.TgTag != "" {
		const prefix string = ",\"tg_tag\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.TgTag))
	}
	if len(in.Skills) != 0 {
		const prefix string = ",\"skills\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Skills {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestUserUpdate) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7df0efccEncodeGetmeBackendInternalAppUserDto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestUserUpdate) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7df0efccEncodeGetmeBackendInternalAppUserDto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestUserUpdate) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7df0efccDecodeGetmeBackendInternalAppUserDto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestUserUpdate) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7df0efccDecodeGetmeBackendInternalAppUserDto(l, v)
}
func easyjson7df0efccDecodeGetmeBackendInternalAppUserDto1(in *jlexer.Lexer, out *RequestUpdateStatus) {
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
		case "IsMentor":
			out.IsMentor = bool(in.Bool())
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
func easyjson7df0efccEncodeGetmeBackendInternalAppUserDto1(out *jwriter.Writer, in RequestUpdateStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"IsMentor\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.IsMentor))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestUpdateStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7df0efccEncodeGetmeBackendInternalAppUserDto1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestUpdateStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7df0efccEncodeGetmeBackendInternalAppUserDto1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestUpdateStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7df0efccDecodeGetmeBackendInternalAppUserDto1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestUpdateStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7df0efccDecodeGetmeBackendInternalAppUserDto1(l, v)
}
