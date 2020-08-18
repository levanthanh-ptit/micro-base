package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	auth "auth/proto"
)

type Auth struct {
	accountId uint
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Auth) Call(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	log.Info("Received Auth.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
