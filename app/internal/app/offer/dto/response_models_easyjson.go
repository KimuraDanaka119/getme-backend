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

func easyjson316682a0DecodeGetmeBackendInternalAppOfferDto(in *jlexer.Lexer, out *ResponseOffer) {
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
func easyjson316682a0EncodeGetmeBackendInternalAppOfferDto(out *jwriter.Writer, in ResponseOffer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"offer_id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResponseOffer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppOfferDto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResponseOffer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppOfferDto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResponseOffer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppOfferDto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResponseOffer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppOfferDto(l, v)
}
func easyjson316682a0DecodeGetmeBackendInternalAppOfferDto1(in *jlexer.Lexer, out *RespondOffersWithUser) {
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
			out.ID = int64(in.Int64())
		case "skill_name":
			out.SkillName = string(in.String())
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
func easyjson316682a0EncodeGetmeBackendInternalAppOfferDto1(out *jwriter.Writer, in RespondOffersWithUser) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"offer_id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"skill_name\":"
		out.RawString(prefix)
		out.String(string(in.SkillName))
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
func (v RespondOffersWithUser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson316682a0EncodeGetmeBackendInternalAppOfferDto1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RespondOffersWithUser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson316682a0EncodeGetmeBackendInternalAppOfferDto1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RespondOffersWithUser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson316682a0DecodeGetmeBackendInternalAppOfferDto1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RespondOffersWithUser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson316682a0DecodeGetmeBackendInternalAppOfferDto1(l, v)
}