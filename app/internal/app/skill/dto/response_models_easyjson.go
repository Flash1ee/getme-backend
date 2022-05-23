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

func easyjson316682a0DecodeGetmeBackendInternalAppSkillDto(in *jlexer.Lexer, out *SkillsResponse) {
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
		case "skills":
			if in.IsNull() {
				in.Skip()
				out.Skills = nil
			} else {
				in.Delim('[')
				if out.Skills == nil {
					if !in.IsDelim(']') {
						out.Skills = make([]SkillResponse, 0, 2)
					} else {
						out.Skills = []SkillResponse{}
					}
				} else {
					out.Skills = (out.Skills)[:0]
				}
				for !in.IsDelim(']') {
					var v1 SkillResponse
					(v1).UnmarshalEasyJSON(in)
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
func easyjson316682a0EncodeGetmeBackendInternalAppSkillDto(out *jwriter.Writer, in SkillsResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"skills\":"
		out.RawString(prefix[1:])
		if in.Skills == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Skills {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SkillsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppSkillDto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SkillsResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppSkillDto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SkillsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppSkillDto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SkillsResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppSkillDto(l, v)
}
func easyjson316682a0DecodeGetmeBackendInternalAppSkillDto1(in *jlexer.Lexer, out *SkillResponse) {
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
		case "name":
			out.Name = string(in.String())
		case "color":
			out.Color = string(in.String())
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
func easyjson316682a0EncodeGetmeBackendInternalAppSkillDto1(out *jwriter.Writer, in SkillResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if in.Color != "" {
		const prefix string = ",\"color\":"
		out.RawString(prefix)
		out.String(string(in.Color))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SkillResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppSkillDto1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SkillResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppSkillDto1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SkillResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppSkillDto1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SkillResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppSkillDto1(l, v)
}
