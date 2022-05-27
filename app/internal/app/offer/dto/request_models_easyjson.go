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

func easyjson7df0efccDecodeGetmeBackendInternalAppOfferDto(in *jlexer.Lexer, out *RequestCreateOffer) {
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
		case "skill_name":
			out.SkillName = string(in.String())
		case "mentor_id":
			out.MentorID = int64(in.Int64())
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
func easyjson7df0efccEncodeGetmeBackendInternalAppOfferDto(out *jwriter.Writer, in RequestCreateOffer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"skill_name\":"
		out.RawString(prefix[1:])
		out.String(string(in.SkillName))
	}
	{
		const prefix string = ",\"mentor_id\":"
		out.RawString(prefix)
		out.Int64(int64(in.MentorID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestCreateOffer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7df0efccEncodeGetmeBackendInternalAppOfferDto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestCreateOffer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7df0efccEncodeGetmeBackendInternalAppOfferDto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestCreateOffer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7df0efccDecodeGetmeBackendInternalAppOfferDto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestCreateOffer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7df0efccDecodeGetmeBackendInternalAppOfferDto(l, v)
}
