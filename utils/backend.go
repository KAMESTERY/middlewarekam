package utils

import "os"

var (
	backendUrl = func() string {
		beUrl := os.Getenv("BACKEND")
		if len(beUrl) == 0 {
			beUrl = "http://localhost:8088"
		}
		return beUrl
	}

	BackendGQL = backendUrl() + "/graphql"

	rpcBackendUrl = func() string {
		rpcBeUrl := os.Getenv("RPC_BACKEND")
		if len(rpcBeUrl) == 0 {
			rpcBeUrl = "localhost:9099"
		}
		return rpcBeUrl
	}

	BackendRPC = rpcBackendUrl()
)
