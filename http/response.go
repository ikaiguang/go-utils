package gohttp

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	gopb "github.com/ikaiguang/go-utils/proto"
	"github.com/pkg/errors"
)

// GetParameters get parameters
func GetParameters(ctx *gin.Context, parameters proto.Message) error {
	// buffer
	var buf = new(bytes.Buffer)

	// read body
	if _, err := buf.ReadFrom(ctx.Request.Body); err != nil {
		return errors.WithStack(err)
	}
	if buf.Len() == 0 {
		return nil
	}

	// unmarshal
	if err := jsonpb.Unmarshal(buf, parameters); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// PbJSON pb json
func PbJSON(data proto.Message) (*gopb.PbJson, error) {
	anyData, err := ptypes.MarshalAny(data)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &gopb.PbJson{
		Code:    gopb.SuccessCode,
		Message: gopb.SuccessMsg,
		Secret:  gopb.SuccessSecret,
		Data:    anyData,
	}, nil
}
