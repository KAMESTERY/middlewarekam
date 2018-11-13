package utils

import "os"

var (
	namespace = func() string {
		ns := os.Getenv("NAMESPACE")
		if len(ns) == 0 {
			panic("Namespace Required!")
		}
		return ns
	}

	Namespace = namespace()
)
