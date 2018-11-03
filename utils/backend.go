package utils

import "os"

var (
	backendUrl = func() string {
		beUrl := os.Getenv("BACKEND")
		if len(beUrl) == 0 {
			beUrl = "localhost:9099"
		}
		return beUrl
	}

	BackendRPC = backendUrl()
)
