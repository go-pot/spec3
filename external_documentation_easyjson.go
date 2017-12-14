// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package spec3

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

func easyjson1a518086DecodeGithubComGoOpenapiSpec3(in *jlexer.Lexer, out *ExternalDocumentation) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "description":
			out.Description = string(in.String())
		case "url":
			out.URL = string(in.String())
		case "extensions":
			(out.Extensions).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1a518086EncodeGithubComGoOpenapiSpec3(out *jwriter.Writer, in ExternalDocumentation) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Description != "" {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if in.URL != "" {
		const prefix string = ",\"url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.URL))
	}
	if true {
		const prefix string = ",\"extensions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Extensions).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ExternalDocumentation) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1a518086EncodeGithubComGoOpenapiSpec3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ExternalDocumentation) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1a518086EncodeGithubComGoOpenapiSpec3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ExternalDocumentation) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1a518086DecodeGithubComGoOpenapiSpec3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ExternalDocumentation) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1a518086DecodeGithubComGoOpenapiSpec3(l, v)
}
