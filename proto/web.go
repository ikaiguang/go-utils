package gopb

import (
	"fmt"
)

// response
const (
	SuccessCode   = 0         // code
	SuccessMsg    = "success" // msg
	SuccessSecret = "success" // secret
)

// WebJSON : web json
type WebJSON struct {
	Code    int32       `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message string      `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Secret  string      `protobuf:"bytes,2,opt,name=secret" json:"secret,omitempty"`
	Data    interface{} `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (w *WebJSON) Reset()                      { *w = WebJSON{} }
func (w *WebJSON) String() string              { return fmt.Sprintf("%#v", w) }
func (w *WebJSON) ProtoMessage()               {}
func (w *WebJSON) Descriptor() ([]byte, []int) { return nil, []int{0} }
