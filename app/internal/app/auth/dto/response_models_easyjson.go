// TEMPORARY AUTOGENERATED FILE: easyjson stub code to make the package
// compilable during generation.

package dto

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

func (AuthResponse) MarshalJSON() ([]byte, error)       { return nil, nil }
func (*AuthResponse) UnmarshalJSON([]byte) error        { return nil }
func (AuthResponse) MarshalEasyJSON(w *jwriter.Writer)  {}
func (*AuthResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_AuthResponse *AuthResponse

func (IDResponse) MarshalJSON() ([]byte, error)       { return nil, nil }
func (*IDResponse) UnmarshalJSON([]byte) error        { return nil }
func (IDResponse) MarshalEasyJSON(w *jwriter.Writer)  {}
func (*IDResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_IDResponse *IDResponse

func (UpdateResponse) MarshalJSON() ([]byte, error)       { return nil, nil }
func (*UpdateResponse) UnmarshalJSON([]byte) error        { return nil }
func (UpdateResponse) MarshalEasyJSON(w *jwriter.Writer)  {}
func (*UpdateResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_UpdateResponse *UpdateResponse
