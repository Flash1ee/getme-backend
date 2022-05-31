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

func easyjson316682a0DecodeGetmeBackendInternalAppUserDto(in *jlexer.Lexer, out *UsersWithSkillResponse) {
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
		case "users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]UserWithSkillsResponse, 0, 0)
					} else {
						out.Users = []UserWithSkillsResponse{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v1 UserWithSkillsResponse
					(v1).UnmarshalEasyJSON(in)
					out.Users = append(out.Users, v1)
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
func easyjson316682a0EncodeGetmeBackendInternalAppUserDto(out *jwriter.Writer, in UsersWithSkillResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"users\":"
		out.RawString(prefix[1:])
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Users {
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
func (v UsersWithSkillResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppUserDto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UsersWithSkillResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppUserDto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UsersWithSkillResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppUserDto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UsersWithSkillResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppUserDto(l, v)
}
func easyjson316682a0DecodeGetmeBackendInternalAppUserDto1(in *jlexer.Lexer, out *UsersWithOfferIDResponse) {
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
		case "users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]UserWithOfferIDResponse, 0, 0)
					} else {
						out.Users = []UserWithOfferIDResponse{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v4 UserWithOfferIDResponse
					(v4).UnmarshalEasyJSON(in)
					out.Users = append(out.Users, v4)
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
func easyjson316682a0EncodeGetmeBackendInternalAppUserDto1(out *jwriter.Writer, in UsersWithOfferIDResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"users\":"
		out.RawString(prefix[1:])
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Users {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UsersWithOfferIDResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppUserDto1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UsersWithOfferIDResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppUserDto1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UsersWithOfferIDResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppUserDto1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UsersWithOfferIDResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppUserDto1(l, v)
}
func easyjson316682a0DecodeGetmeBackendInternalAppUserDto2(in *jlexer.Lexer, out *UsersResponse) {
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
		case "users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]UserResponse, 0, 0)
					} else {
						out.Users = []UserResponse{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v7 UserResponse
					(v7).UnmarshalEasyJSON(in)
					out.Users = append(out.Users, v7)
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
func easyjson316682a0EncodeGetmeBackendInternalAppUserDto2(out *jwriter.Writer, in UsersResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"users\":"
		out.RawString(prefix[1:])
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Users {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UsersResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppUserDto2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UsersResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppUserDto2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UsersResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppUserDto2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UsersResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppUserDto2(l, v)
}
func easyjson316682a0DecodeGetmeBackendInternalAppUserDto3(in *jlexer.Lexer, out *UserWithSkillsResponse) {
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
						out.Skills = make([]string, 0, 4)
					} else {
						out.Skills = []string{}
					}
				} else {
					out.Skills = (out.Skills)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.Skills = append(out.Skills, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "id":
			out.ID = int64(in.Int64())
		case "first_name":
			out.FirstName = string(in.String())
		case "last_name":
			out.LastName = string(in.String())
		case "about":
			out.About = string(in.String())
		case "avatar":
			out.Avatar = string(in.String())
		case "tg_tag":
			out.TgTag = string(in.String())
		case "is_mentor":
			out.IsSearchable = bool(in.Bool())
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
func easyjson316682a0EncodeGetmeBackendInternalAppUserDto3(out *jwriter.Writer, in UserWithSkillsResponse) {
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
			for v11, v12 := range in.Skills {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Int64(int64(in.ID))
	}
	if in.FirstName != "" {
		const prefix string = ",\"first_name\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	if in.LastName != "" {
		const prefix string = ",\"last_name\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	if in.About != "" {
		const prefix string = ",\"about\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	if in.Avatar != "" {
		const prefix string = ",\"avatar\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	{
		const prefix string = ",\"tg_tag\":"
		out.RawString(prefix)
		out.String(string(in.TgTag))
	}
	{
		const prefix string = ",\"is_mentor\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsSearchable))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserWithSkillsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppUserDto3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserWithSkillsResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppUserDto3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserWithSkillsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppUserDto3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserWithSkillsResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppUserDto3(l, v)
}
func easyjson316682a0DecodeGetmeBackendInternalAppUserDto4(in *jlexer.Lexer, out *UserWithOfferIDResponse) {
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
		case "offer_id":
			out.OfferID = int64(in.Int64())
		case "id":
			out.ID = int64(in.Int64())
		case "first_name":
			out.FirstName = string(in.String())
		case "last_name":
			out.LastName = string(in.String())
		case "about":
			out.About = string(in.String())
		case "avatar":
			out.Avatar = string(in.String())
		case "tg_tag":
			out.TgTag = string(in.String())
		case "is_mentor":
			out.IsSearchable = bool(in.Bool())
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
func easyjson316682a0EncodeGetmeBackendInternalAppUserDto4(out *jwriter.Writer, in UserWithOfferIDResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"offer_id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.OfferID))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Int64(int64(in.ID))
	}
	if in.FirstName != "" {
		const prefix string = ",\"first_name\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	if in.LastName != "" {
		const prefix string = ",\"last_name\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	if in.About != "" {
		const prefix string = ",\"about\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	if in.Avatar != "" {
		const prefix string = ",\"avatar\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	{
		const prefix string = ",\"tg_tag\":"
		out.RawString(prefix)
		out.String(string(in.TgTag))
	}
	{
		const prefix string = ",\"is_mentor\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsSearchable))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserWithOfferIDResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppUserDto4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserWithOfferIDResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppUserDto4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserWithOfferIDResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppUserDto4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserWithOfferIDResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppUserDto4(l, v)
}
func easyjson316682a0DecodeGetmeBackendInternalAppUserDto5(in *jlexer.Lexer, out *UserStatusResponse) {
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
		case "is_mentor":
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
func easyjson316682a0EncodeGetmeBackendInternalAppUserDto5(out *jwriter.Writer, in UserStatusResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"is_mentor\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.IsMentor))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserStatusResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppUserDto5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserStatusResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppUserDto5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserStatusResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppUserDto5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserStatusResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppUserDto5(l, v)
}
func easyjson316682a0DecodeGetmeBackendInternalAppUserDto6(in *jlexer.Lexer, out *UserResponse) {
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
		case "first_name":
			out.FirstName = string(in.String())
		case "last_name":
			out.LastName = string(in.String())
		case "about":
			out.About = string(in.String())
		case "avatar":
			out.Avatar = string(in.String())
		case "tg_tag":
			out.TgTag = string(in.String())
		case "is_mentor":
			out.IsSearchable = bool(in.Bool())
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
func easyjson316682a0EncodeGetmeBackendInternalAppUserDto6(out *jwriter.Writer, in UserResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	if in.FirstName != "" {
		const prefix string = ",\"first_name\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	if in.LastName != "" {
		const prefix string = ",\"last_name\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	if in.About != "" {
		const prefix string = ",\"about\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	if in.Avatar != "" {
		const prefix string = ",\"avatar\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	{
		const prefix string = ",\"tg_tag\":"
		out.RawString(prefix)
		out.String(string(in.TgTag))
	}
	{
		const prefix string = ",\"is_mentor\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsSearchable))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppUserDto6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppUserDto6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppUserDto6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppUserDto6(l, v)
}
