package almactx

import (
	"context"
	"time"
)

type ctxKey string

const almadata ctxKey = "almadata"

// CommonData .
type CommonData struct {
	TxTime time.Time
	Mid    string
	Locale string
	Email  string
}

// WithData .
func WithData(ctx context.Context, commonData *CommonData) context.Context {
	return context.WithValue(ctx, almadata, commonData)
}

// GetTxTime .
func GetTxTime(ctx context.Context) time.Time {
	return ctx.Value(almadata).(*CommonData).TxTime
}

// GetMid .
func GetMid(ctx context.Context) string {
	return ctx.Value(almadata).(*CommonData).Mid
}

// GetEmail .
func GetEmail(ctx context.Context) string {
	return ctx.Value(almadata).(*CommonData).Email
}

// GetLocale .
func GetLocale(ctx context.Context) string {
	return ctx.Value(almadata).(*CommonData).Locale
}
