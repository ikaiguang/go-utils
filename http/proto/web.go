package pb

import (
	"fmt"
)

// WebJSON : web json
type WebJSON struct {
	Code    int32       `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message string      `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Data    interface{} `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (w *WebJSON) Reset()                      { *w = WebJSON{} }
func (w *WebJSON) String() string              { return fmt.Sprintf("%#v", w) }
func (w *WebJSON) ProtoMessage()               {}
func (w *WebJSON) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }
