package jsonpb

import (
	"io"

	gogojsonpb "github.com/gogo/protobuf/jsonpb"
	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/utrack/clay/v2/transport/httpruntime"
)

type marshaller struct {
	httpruntime.Marshaler

	clayMarshaler       *httpruntime.MarshalerPbJSON
	gogoCustomMarshaler *GogoMarshaler
}

func (m *marshaller) ContentType() string {
	return "application/json"
}

func (m *marshaller) Unmarshal(r io.Reader, dst interface{}) error {
	return m.clayMarshaler.Unmarshal(r, dst)
}

func (m *marshaller) Marshal(w io.Writer, src interface{}) error {
	if pm, ok := src.(proto.Message); ok {
		if gogoproto.MessageName(pm) != "" {
			return m.gogoCustomMarshaler.Marshal(w, pm)
		}
	}
	return m.clayMarshaler.Marshal(w, src)
}

func NewCustomMarshaler(enumsAsInts bool, emitDefaults bool) httpruntime.Marshaler {
	return &marshaller{
		clayMarshaler: &httpruntime.MarshalerPbJSON{
			Marshaler:       &runtime.JSONPb{EnumsAsInts: enumsAsInts, EmitDefaults: emitDefaults},
			Unmarshaler:     &runtime.JSONPb{},
			GogoMarshaler:   &gogojsonpb.Marshaler{EnumsAsInts: enumsAsInts, EmitDefaults: emitDefaults},
			GogoUnmarshaler: &gogojsonpb.Unmarshaler{},
		},
		gogoCustomMarshaler: &GogoMarshaler{
			EnumsAsInts:  enumsAsInts,
			EmitDefaults: emitDefaults,
		},
	}
}
