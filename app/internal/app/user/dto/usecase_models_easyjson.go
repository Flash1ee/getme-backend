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

func easyjsonB1c6a43eDecodeGetmeBackendInternalAppUserDto(in *jlexer.Lexer, out *UserWithSkillsUsecaseSlice) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(UserWithSkillsUsecaseSlice, 0, 0)
			} else {
				*out = UserWithSkillsUsecaseSlice{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 UserWithSkillsUsecase
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB1c6a43eEncodeGetmeBackendInternalAppUserDto(out *jwriter.Writer, in UserWithSkillsUsecaseSlice) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v UserWithSkillsUsecaseSlice) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB1c6a43eEncodeGetmeBackendInternalAppUserDto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserWithSkillsUsecaseSlice) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB1c6a43eEncodeGetmeBackendInternalAppUserDto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserWithSkillsUsecaseSlice) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB1c6a43eDecodeGetmeBackendInternalAppUserDto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserWithSkillsUsecaseSlice) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB1c6a43eDecodeGetmeBackendInternalAppUserDto(l, v)
}
func easyjsonB1c6a43eDecodeGetmeBackendInternalAppUserDto1(in *jlexer.Lexer, out *UserWithSkillsUsecase) {
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
		case "Skills":
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
					var v4 string
					v4 = string(in.String())
					out.Skills = append(out.Skills, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ID":
			out.ID = int64(in.Int64())
		case "FirstName":
			out.FirstName = string(in.String())
		case "LastName":
			out.LastName = string(in.String())
		case "Nickname":
			out.Nickname = string(in.String())
		case "About":
			out.About = string(in.String())
		case "Avatar":
			out.Avatar = string(in.String())
		case "TgTag":
			out.TgTag = string(in.String())
		case "Email":
			out.Email = string(in.String())
		case "IsSearchable":
			out.IsSearchable = bool(in.Bool())
		case "CreatedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "UpdatedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
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
func easyjsonB1c6a43eEncodeGetmeBackendInternalAppUserDto1(out *jwriter.Writer, in UserWithSkillsUsecase) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Skills\":"
		out.RawString(prefix[1:])
		if in.Skills == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Skills {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix)
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"FirstName\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"LastName\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"Nickname\":"
		out.RawString(prefix)
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"About\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	{
		const prefix string = ",\"Avatar\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	{
		const prefix string = ",\"TgTag\":"
		out.RawString(prefix)
		out.String(string(in.TgTag))
	}
	{
		const prefix string = ",\"Email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"IsSearchable\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsSearchable))
	}
	{
		const prefix string = ",\"CreatedAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"UpdatedAt\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserWithSkillsUsecase) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB1c6a43eEncodeGetmeBackendInternalAppUserDto1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserWithSkillsUsecase) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB1c6a43eEncodeGetmeBackendInternalAppUserDto1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserWithSkillsUsecase) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB1c6a43eDecodeGetmeBackendInternalAppUserDto1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserWithSkillsUsecase) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB1c6a43eDecodeGetmeBackendInternalAppUserDto1(l, v)
}
func easyjsonB1c6a43eDecodeGetmeBackendInternalAppUserDto2(in *jlexer.Lexer, out *UserWithOfferIDUsecase) {
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
		case "OfferID":
			out.OfferID = int64(in.Int64())
		case "ID":
			out.ID = int64(in.Int64())
		case "FirstName":
			out.FirstName = string(in.String())
		case "LastName":
			out.LastName = string(in.String())
		case "Nickname":
			out.Nickname = string(in.String())
		case "About":
			out.About = string(in.String())
		case "Avatar":
			out.Avatar = string(in.String())
		case "TgTag":
			out.TgTag = string(in.String())
		case "Email":
			out.Email = string(in.String())
		case "IsSearchable":
			out.IsSearchable = bool(in.Bool())
		case "CreatedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "UpdatedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
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
func easyjsonB1c6a43eEncodeGetmeBackendInternalAppUserDto2(out *jwriter.Writer, in UserWithOfferIDUsecase) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"OfferID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.OfferID))
	}
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix)
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"FirstName\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"LastName\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"Nickname\":"
		out.RawString(prefix)
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"About\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	{
		const prefix string = ",\"Avatar\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	{
		const prefix string = ",\"TgTag\":"
		out.RawString(prefix)
		out.String(string(in.TgTag))
	}
	{
		const prefix string = ",\"Email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"IsSearchable\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsSearchable))
	}
	{
		const prefix string = ",\"CreatedAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"UpdatedAt\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserWithOfferIDUsecase) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB1c6a43eEncodeGetmeBackendInternalAppUserDto2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserWithOfferIDUsecase) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB1c6a43eEncodeGetmeBackendInternalAppUserDto2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserWithOfferIDUsecase) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB1c6a43eDecodeGetmeBackendInternalAppUserDto2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserWithOfferIDUsecase) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB1c6a43eDecodeGetmeBackendInternalAppUserDto2(l, v)
}
func easyjsonB1c6a43eDecodeGetmeBackendInternalAppUserDto3(in *jlexer.Lexer, out *UserUsecase) {
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
		case "ID":
			out.ID = int64(in.Int64())
		case "FirstName":
			out.FirstName = string(in.String())
		case "LastName":
			out.LastName = string(in.String())
		case "Nickname":
			out.Nickname = string(in.String())
		case "About":
			out.About = string(in.String())
		case "Avatar":
			out.Avatar = string(in.String())
		case "TgTag":
			out.TgTag = string(in.String())
		case "Email":
			out.Email = string(in.String())
		case "IsSearchable":
			out.IsSearchable = bool(in.Bool())
		case "CreatedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "UpdatedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
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
func easyjsonB1c6a43eEncodeGetmeBackendInternalAppUserDto3(out *jwriter.Writer, in UserUsecase) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"FirstName\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"LastName\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"Nickname\":"
		out.RawString(prefix)
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"About\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	{
		const prefix string = ",\"Avatar\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	{
		const prefix string = ",\"TgTag\":"
		out.RawString(prefix)
		out.String(string(in.TgTag))
	}
	{
		const prefix string = ",\"Email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"IsSearchable\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsSearchable))
	}
	{
		const prefix string = ",\"CreatedAt\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"UpdatedAt\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserUsecase) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB1c6a43eEncodeGetmeBackendInternalAppUserDto3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserUsecase) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB1c6a43eEncodeGetmeBackendInternalAppUserDto3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserUsecase) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB1c6a43eDecodeGetmeBackendInternalAppUserDto3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserUsecase) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB1c6a43eDecodeGetmeBackendInternalAppUserDto3(l, v)
}
func easyjsonB1c6a43eDecodeGetmeBackendInternalAppUserDto4(in *jlexer.Lexer, out *UserStatusUsecase) {
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
		case "UserID":
			out.UserID = int64(in.Int64())
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
func easyjsonB1c6a43eEncodeGetmeBackendInternalAppUserDto4(out *jwriter.Writer, in UserStatusUsecase) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"UserID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.UserID))
	}
	{
		const prefix string = ",\"IsMentor\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsMentor))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserStatusUsecase) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB1c6a43eEncodeGetmeBackendInternalAppUserDto4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserStatusUsecase) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB1c6a43eEncodeGetmeBackendInternalAppUserDto4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserStatusUsecase) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB1c6a43eDecodeGetmeBackendInternalAppUserDto4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserStatusUsecase) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB1c6a43eDecodeGetmeBackendInternalAppUserDto4(l, v)
}