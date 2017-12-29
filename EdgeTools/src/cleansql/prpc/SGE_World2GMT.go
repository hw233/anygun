package prpc

import (
	"bytes"
	"errors"
	"suzuki/prpc"
)

type SGE_World2GMTStub struct {
	Sender prpc.StubSender
}
type SGE_World2GMTProxy interface {
}

func SGE_World2GMTDispatch(buffer *bytes.Buffer, p SGE_World2GMTProxy) error {
	return errors.New(prpc.NoneMethodError)
}
