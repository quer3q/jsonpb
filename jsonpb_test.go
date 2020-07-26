package jsonpb

import (
	"bytes"
	"strings"
	"testing"
	"time"

	jsonpbtestpb "github.com/quer3q/jsonpb/test_pb"

	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJsonPb_EmittedTaggedOnly(t *testing.T) {
	msg := &jsonpbtestpb.Test{
		StrEmitAlways:  "",
		StrEmitDefault: "",
		TsEmitAlways:   nil,
		TsEmitDefault:  nil,
	}

	marshaler := GogoMarshaler{
		EmitDefaults: false,
		EnumsAsInts:  false,
	}

	buf := bytes.NewBuffer(make([]byte, 0, 1024))

	err := marshaler.Marshal(buf, msg)
	assert.NoError(t, err)

	strJson := buf.String()

	require.True(t, strings.Contains(strJson, "strEmitAlways"))
	require.True(t, strings.Contains(strJson, "tsEmitAlways"))
	require.False(t, strings.Contains(strJson, "strEmitDefault"))
	require.False(t, strings.Contains(strJson, "tsEmitDefault"))
}

func TestJsonPb_EmittedAll(t *testing.T) {
	ts, err := types.TimestampProto(time.Now())
	assert.NoError(t, err)

	msg := &jsonpbtestpb.Test{
		StrEmitAlways:  "a",
		StrEmitDefault: "b",
		TsEmitAlways:   ts,
		TsEmitDefault:  ts,
	}

	marshaler := GogoMarshaler{
		EmitDefaults: false,
		EnumsAsInts:  false,
	}

	buf := bytes.NewBuffer(make([]byte, 0, 1024))

	err = marshaler.Marshal(buf, msg)
	assert.NoError(t, err)

	strJson := buf.String()

	require.True(t, strings.Contains(strJson, "strEmitAlways"))
	require.True(t, strings.Contains(strJson, "tsEmitAlways"))
	require.True(t, strings.Contains(strJson, "strEmitDefault"))
	require.True(t, strings.Contains(strJson, "tsEmitDefault"))
}
