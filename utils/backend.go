package utils

import "os"

var (
	rpcBackendUrl = func() string {
		rpcBeUrl := os.Getenv("RPC_BACKEND")
		if len(rpcBeUrl) == 0 {
			rpcBeUrl = "localhost:9099"
		}
		return rpcBeUrl
	}

	BackendRPC = rpcBackendUrl()
)
