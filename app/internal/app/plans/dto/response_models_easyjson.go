// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package dto

import (
	json "encoding/json"
	dto "getme-backend/internal/app/task/dto"
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

func easyjson316682a0DecodeGetmeBackendInternalAppPlansDto(in *jlexer.Lexer, out *PlansWithSkillsResponseMentor) {
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
		case "plans":
			if in.IsNull() {
				in.Skip()
				out.Plans = nil
			} else {
				in.Delim('[')
				if out.Plans == nil {
					if !in.IsDelim(']') {
						out.Plans = make([]PlanWithSkillsResponseMentor, 0, 0)
					} else {
						out.Plans = []PlanWithSkillsResponseMentor{}
					}
				} else {
					out.Plans = (out.Plans)[:0]
				}
				for !in.IsDelim(']') {
					var v1 PlanWithSkillsResponseMentor
					(v1).UnmarshalEasyJSON(in)
					out.Plans = append(out.Plans, v1)
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
func easyjson316682a0EncodeGetmeBackendInternalAppPlansDto(out *jwriter.Writer, in PlansWithSkillsResponseMentor) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"plans\":"
		out.RawString(prefix[1:])
		if in.Plans == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Plans {
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
func (v PlansWithSkillsResponseMentor) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppPlansDto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PlansWithSkillsResponseMentor) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppPlansDto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PlansWithSkillsResponseMentor) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppPlansDto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PlansWithSkillsResponseMentor) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppPlansDto(l, v)
}
func easyjson316682a0DecodeGetmeBackendInternalAppPlansDto1(in *jlexer.Lexer, out *PlansWithSkillsResponseMentee) {
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
		case "plans":
			if in.IsNull() {
				in.Skip()
				out.Plans = nil
			} else {
				in.Delim('[')
				if out.Plans == nil {
					if !in.IsDelim(']') {
						out.Plans = make([]PlanWithSkillsResponseMentee, 0, 0)
					} else {
						out.Plans = []PlanWithSkillsResponseMentee{}
					}
				} else {
					out.Plans = (out.Plans)[:0]
				}
				for !in.IsDelim(']') {
					var v4 PlanWithSkillsResponseMentee
					(v4).UnmarshalEasyJSON(in)
					out.Plans = append(out.Plans, v4)
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
func easyjson316682a0EncodeGetmeBackendInternalAppPlansDto1(out *jwriter.Writer, in PlansWithSkillsResponseMentee) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"plans\":"
		out.RawString(prefix[1:])
		if in.Plans == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Plans {
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
func (v PlansWithSkillsResponseMentee) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppPlansDto1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PlansWithSkillsResponseMentee) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppPlansDto1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PlansWithSkillsResponseMentee) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppPlansDto1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PlansWithSkillsResponseMentee) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppPlansDto1(l, v)
}
func easyjson316682a0DecodeGetmeBackendInternalAppPlansDto2(in *jlexer.Lexer, out *PlanWithTaskResponseMentor) {
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
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "progress":
			out.Progress = float64(in.Float64())
		case "user":
			(out.UserResponse).UnmarshalEasyJSON(in)
		case "tasks":
			if in.IsNull() {
				in.Skip()
				out.Tasks = nil
			} else {
				in.Delim('[')
				if out.Tasks == nil {
					if !in.IsDelim(']') {
						out.Tasks = make([]dto.ResponseTask, 0, 0)
					} else {
						out.Tasks = []dto.ResponseTask{}
					}
				} else {
					out.Tasks = (out.Tasks)[:0]
				}
				for !in.IsDelim(']') {
					var v7 dto.ResponseTask
					(v7).UnmarshalEasyJSON(in)
					out.Tasks = append(out.Tasks, v7)
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
func easyjson316682a0EncodeGetmeBackendInternalAppPlansDto2(out *jwriter.Writer, in PlanWithTaskResponseMentor) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"progress\":"
		out.RawString(prefix)
		out.Float64(float64(in.Progress))
	}
	{
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		(in.UserResponse).MarshalEasyJSON(out)
	}
	if len(in.Tasks) != 0 {
		const prefix string = ",\"tasks\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v8, v9 := range in.Tasks {
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
func (v PlanWithTaskResponseMentor) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppPlansDto2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PlanWithTaskResponseMentor) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppPlansDto2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PlanWithTaskResponseMentor) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppPlansDto2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PlanWithTaskResponseMentor) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppPlansDto2(l, v)
}
func easyjson316682a0DecodeGetmeBackendInternalAppPlansDto3(in *jlexer.Lexer, out *PlanWithTaskResponseMentee) {
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
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "progress":
			out.Progress = float64(in.Float64())
		case "user":
			(out.UserResponse).UnmarshalEasyJSON(in)
		case "tasks":
			if in.IsNull() {
				in.Skip()
				out.Tasks = nil
			} else {
				in.Delim('[')
				if out.Tasks == nil {
					if !in.IsDelim(']') {
						out.Tasks = make([]dto.ResponseTask, 0, 0)
					} else {
						out.Tasks = []dto.ResponseTask{}
					}
				} else {
					out.Tasks = (out.Tasks)[:0]
				}
				for !in.IsDelim(']') {
					var v10 dto.ResponseTask
					(v10).UnmarshalEasyJSON(in)
					out.Tasks = append(out.Tasks, v10)
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
func easyjson316682a0EncodeGetmeBackendInternalAppPlansDto3(out *jwriter.Writer, in PlanWithTaskResponseMentee) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"progress\":"
		out.RawString(prefix)
		out.Float64(float64(in.Progress))
	}
	{
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		(in.UserResponse).MarshalEasyJSON(out)
	}
	if len(in.Tasks) != 0 {
		const prefix string = ",\"tasks\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v11, v12 := range in.Tasks {
				if v11 > 0 {
					out.RawByte(',')
				}
				(v12).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PlanWithTaskResponseMentee) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppPlansDto3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PlanWithTaskResponseMentee) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppPlansDto3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PlanWithTaskResponseMentee) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppPlansDto3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PlanWithTaskResponseMentee) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppPlansDto3(l, v)
}
func easyjson316682a0DecodeGetmeBackendInternalAppPlansDto4(in *jlexer.Lexer, out *PlanWithSkillsResponseMentor) {
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
					var v13 string
					v13 = string(in.String())
					out.Skills = append(out.Skills, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "id":
			out.ID = int64(in.Int64())
		case "name":
			out.Name = string(in.String())
		case "about":
			out.About = string(in.String())
		case "progress":
			out.Progress = float64(in.Float64())
		case "mentee_id":
			if in.IsNull() {
				in.Skip()
				out.MenteeID = nil
			} else {
				if out.MenteeID == nil {
					out.MenteeID = new(int64)
				}
				*out.MenteeID = int64(in.Int64())
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
func easyjson316682a0EncodeGetmeBackendInternalAppPlansDto4(out *jwriter.Writer, in PlanWithSkillsResponseMentor) {
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
			for v14, v15 := range in.Skills {
				if v14 > 0 {
					out.RawByte(',')
				}
				out.String(string(v15))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"about\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	{
		const prefix string = ",\"progress\":"
		out.RawString(prefix)
		out.Float64(float64(in.Progress))
	}
	if in.MenteeID != nil {
		const prefix string = ",\"mentee_id\":"
		out.RawString(prefix)
		out.Int64(int64(*in.MenteeID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PlanWithSkillsResponseMentor) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppPlansDto4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PlanWithSkillsResponseMentor) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppPlansDto4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PlanWithSkillsResponseMentor) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppPlansDto4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PlanWithSkillsResponseMentor) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppPlansDto4(l, v)
}
func easyjson316682a0DecodeGetmeBackendInternalAppPlansDto5(in *jlexer.Lexer, out *PlanWithSkillsResponseMentee) {
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
					var v16 string
					v16 = string(in.String())
					out.Skills = append(out.Skills, v16)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "id":
			out.ID = int64(in.Int64())
		case "name":
			out.Name = string(in.String())
		case "about":
			out.About = string(in.String())
		case "progress":
			out.Progress = float64(in.Float64())
		case "mentor_id":
			if in.IsNull() {
				in.Skip()
				out.MentorID = nil
			} else {
				if out.MentorID == nil {
					out.MentorID = new(int64)
				}
				*out.MentorID = int64(in.Int64())
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
func easyjson316682a0EncodeGetmeBackendInternalAppPlansDto5(out *jwriter.Writer, in PlanWithSkillsResponseMentee) {
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
			for v17, v18 := range in.Skills {
				if v17 > 0 {
					out.RawByte(',')
				}
				out.String(string(v18))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"about\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	{
		const prefix string = ",\"progress\":"
		out.RawString(prefix)
		out.Float64(float64(in.Progress))
	}
	if in.MentorID != nil {
		const prefix string = ",\"mentor_id\":"
		out.RawString(prefix)
		out.Int64(int64(*in.MentorID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PlanWithSkillsResponseMentee) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppPlansDto5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PlanWithSkillsResponseMentee) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppPlansDto5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PlanWithSkillsResponseMentee) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppPlansDto5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PlanWithSkillsResponseMentee) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppPlansDto5(l, v)
}
func easyjson316682a0DecodeGetmeBackendInternalAppPlansDto6(in *jlexer.Lexer, out *PlanResponseMentor) {
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
		case "name":
			out.Name = string(in.String())
		case "about":
			out.About = string(in.String())
		case "progress":
			out.Progress = float64(in.Float64())
		case "mentee_id":
			if in.IsNull() {
				in.Skip()
				out.MenteeID = nil
			} else {
				if out.MenteeID == nil {
					out.MenteeID = new(int64)
				}
				*out.MenteeID = int64(in.Int64())
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
func easyjson316682a0EncodeGetmeBackendInternalAppPlansDto6(out *jwriter.Writer, in PlanResponseMentor) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"about\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	{
		const prefix string = ",\"progress\":"
		out.RawString(prefix)
		out.Float64(float64(in.Progress))
	}
	if in.MenteeID != nil {
		const prefix string = ",\"mentee_id\":"
		out.RawString(prefix)
		out.Int64(int64(*in.MenteeID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PlanResponseMentor) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppPlansDto6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PlanResponseMentor) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppPlansDto6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PlanResponseMentor) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppPlansDto6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PlanResponseMentor) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppPlansDto6(l, v)
}
func easyjson316682a0DecodeGetmeBackendInternalAppPlansDto7(in *jlexer.Lexer, out *PlanResponseMentee) {
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
		case "name":
			out.Name = string(in.String())
		case "about":
			out.About = string(in.String())
		case "progress":
			out.Progress = float64(in.Float64())
		case "mentor_id":
			if in.IsNull() {
				in.Skip()
				out.MentorID = nil
			} else {
				if out.MentorID == nil {
					out.MentorID = new(int64)
				}
				*out.MentorID = int64(in.Int64())
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
func easyjson316682a0EncodeGetmeBackendInternalAppPlansDto7(out *jwriter.Writer, in PlanResponseMentee) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"about\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	{
		const prefix string = ",\"progress\":"
		out.RawString(prefix)
		out.Float64(float64(in.Progress))
	}
	if in.MentorID != nil {
		const prefix string = ",\"mentor_id\":"
		out.RawString(prefix)
		out.Int64(int64(*in.MentorID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PlanResponseMentee) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppPlansDto7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PlanResponseMentee) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppPlansDto7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PlanResponseMentee) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppPlansDto7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PlanResponseMentee) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppPlansDto7(l, v)
}
func easyjson316682a0DecodeGetmeBackendInternalAppPlansDto8(in *jlexer.Lexer, out *PlanResponse) {
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
		case "name":
			out.Name = string(in.String())
		case "about":
			out.About = string(in.String())
		case "mentee_id":
			if in.IsNull() {
				in.Skip()
				out.MenteeID = nil
			} else {
				if out.MenteeID == nil {
					out.MenteeID = new(int64)
				}
				*out.MenteeID = int64(in.Int64())
			}
		case "mentor_id":
			if in.IsNull() {
				in.Skip()
				out.MentorID = nil
			} else {
				if out.MentorID == nil {
					out.MentorID = new(int64)
				}
				*out.MentorID = int64(in.Int64())
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
func easyjson316682a0EncodeGetmeBackendInternalAppPlansDto8(out *jwriter.Writer, in PlanResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"about\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	if in.MenteeID != nil {
		const prefix string = ",\"mentee_id\":"
		out.RawString(prefix)
		out.Int64(int64(*in.MenteeID))
	}
	if in.MentorID != nil {
		const prefix string = ",\"mentor_id\":"
		out.RawString(prefix)
		out.Int64(int64(*in.MentorID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PlanResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppPlansDto8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PlanResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppPlansDto8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PlanResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppPlansDto8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PlanResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppPlansDto8(l, v)
}
