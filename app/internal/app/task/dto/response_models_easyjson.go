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

func easyjson316682a0DecodeGetmeBackendInternalAppTaskDto(in *jlexer.Lexer, out *TaskIDResponse) {
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
		case "task_id":
			out.ID = int64(in.Int64())
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
func easyjson316682a0EncodeGetmeBackendInternalAppTaskDto(out *jwriter.Writer, in TaskIDResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"task_id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TaskIDResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppTaskDto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TaskIDResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppTaskDto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TaskIDResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppTaskDto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TaskIDResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppTaskDto(l, v)
}
func easyjson316682a0DecodeGetmeBackendInternalAppTaskDto1(in *jlexer.Lexer, out *ResponseTask) {
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
		case "title":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "deadline":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Deadline).UnmarshalJSON(data))
			}
		case "status":
			out.Status = string(in.String())
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
func easyjson316682a0EncodeGetmeBackendInternalAppTaskDto1(out *jwriter.Writer, in ResponseTask) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"deadline\":"
		out.RawString(prefix)
		out.Raw((in.Deadline).MarshalJSON())
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResponseTask) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppTaskDto1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResponseTask) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppTaskDto1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResponseTask) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppTaskDto1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResponseTask) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppTaskDto1(l, v)
}
