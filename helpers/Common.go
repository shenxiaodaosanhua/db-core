package helpers

import (
	"context"
	"google.golang.org/grpc/codes"
	"log"
)

func Error(err error, msg string) {
	if err != nil {
		panic("msg")
	}
}

func ContextIsBroken(ctx context.Context) (codes.Code, bool) {
	if ctx.Err() != nil {
		log.Println("broken:", ctx.Err().Error())
		switch ctx.Err() {
		case context.Canceled:
			return codes.Canceled, true
		case context.DeadlineExceeded:
			return codes.DeadlineExceeded, true
		default:
			return codes.Unavailable, true
		}
	}
	return codes.OK, false
}
