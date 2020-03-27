package main

import (
	"context"
	"mms/auth"
)

type (
	myDefaultCallCredentials struct{}
)

func (dcc *myDefaultCallCredentials) GetRequestMetadata(ctx context.Context,
	uri ...string) (map[string]string, error) {
	m := make(map[string]string)
	m[auth.MethodKey2] = auth.MethodValue2

	return m, nil
}

func (dcc *myDefaultCallCredentials) RequireTransportSecurity() bool {
	return true
}
