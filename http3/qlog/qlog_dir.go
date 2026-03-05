package qlog

import (
	"context"

	"github.com/MahamShakir/quic-go-patch"
	"github.com/MahamShakir/quic-go-patch/qlog"
	"github.com/MahamShakir/quic-go-patch/qlogwriter"
)

const EventSchema = "urn:ietf:params:qlog:events:http3-12"

func DefaultConnectionTracer(ctx context.Context, isClient bool, connID quic.ConnectionID) qlogwriter.Trace {
	return qlog.DefaultConnectionTracerWithSchemas(ctx, isClient, connID, []string{qlog.EventSchema, EventSchema})
}
