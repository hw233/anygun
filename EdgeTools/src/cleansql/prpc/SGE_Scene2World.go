package prpc

import (
	"bytes"
	"errors"
	"suzuki/prpc"
)

type SGE_Scene2WorldStub struct {
	Sender prpc.StubSender
}
type SGE_Scene2WorldProxy interface {
}

func SGE_Scene2WorldDispatch(buffer *bytes.Buffer, p SGE_Scene2WorldProxy) error {
	return errors.New(prpc.NoneMethodError)
}
