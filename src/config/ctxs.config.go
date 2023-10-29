package config

import (
	"context"
	"net"
	"net/http"
	"strconv"
)

type contextKey int

// CONTEXT_CANCELLED_ERROR by user
const CONTEXT_CANCELLED_ERROR = "context canceled"

const (
	IP_ADDRESS_KEY contextKey = iota
	DEVICE_KEY
	URL_KEY
	MSG_KEY
	HASH_KEY
)

// Extract some value from *http.Request and append it into context
// If the value is not valid or there's any error, return the parent context as result
// add another exportable func to fetch the desired value from context
// make sure that req http.Request pointer is not nil
func (ck contextKey) getString(ctx context.Context) (string, bool) {
	str, ok := ctx.Value(ck).(string)
	return str, ok
}
func (ck contextKey) getInt(ctx context.Context) (int, bool) {
	str, _ := ck.getString(ctx)
	atoi, err := strconv.Atoi(str)
	if err != nil {
		return 0, false
	}
	return atoi, true
}
func (ck contextKey) getInt64(ctx context.Context) (int64, bool) {
	str, _ := ck.getString(ctx)
	atoi, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, false
	}
	return atoi, true
}

func IPAddressToContext(ctx context.Context, req *http.Request) context.Context {
	ip, _, _ := net.SplitHostPort(req.RemoteAddr)
	return context.WithValue(ctx, IP_ADDRESS_KEY, net.ParseIP(ip))
}

func IPAddressFromContext(ctx context.Context) (net.IP, bool) {
	ip, ok := ctx.Value(IP_ADDRESS_KEY).(net.IP)
	return ip, ok
}

func DeviceToContext(ctx context.Context, req *http.Request) context.Context {
	return context.WithValue(ctx, DEVICE_KEY, req.UserAgent())
}

func DeviceFromContext(ctx context.Context) (string, bool) {
	ua, ok := ctx.Value(DEVICE_KEY).(string)
	return ua, ok
}

func UrlToContext(ctx context.Context, req *http.Request) context.Context {
	return context.WithValue(ctx, URL_KEY, req.URL)
}

func UrlFromContext(ctx context.Context) (string, bool) {
	url, ok := ctx.Value(URL_KEY).(string)
	return url, ok
}

func MsgToContext(ctx context.Context, errMsg string) context.Context {
	return context.WithValue(ctx, MSG_KEY, errMsg)
}

func MsgFromContext(ctx context.Context) (string, bool) {
	url, ok := ctx.Value(MSG_KEY).(string)
	return url, ok
}

func HashKeyToContext(ctx context.Context, hashKey string) context.Context {
	return context.WithValue(ctx, HASH_KEY, hashKey)
}

func HashKeyFromContext(ctx context.Context) (string, bool) {
	url, ok := ctx.Value(HASH_KEY).(string)
	return url, ok
}

// Function to fetch context from *http.Request
// set timeout if QS timeout is parseable into time.Duration
// implement the save value to context funcs as optional parameter
func GetContextFromRequest(req *http.Request, opt ...func(context.Context, *http.Request) context.Context) context.Context {
	ctx := req.Context()

	for _, f := range opt {
		ctx = f(ctx, req)
	}

	return ctx
}

func GetAllContextFromRequest(req *http.Request) context.Context {
	return GetContextFromRequest(req, IPAddressToContext, DeviceToContext, UrlToContext)
}
