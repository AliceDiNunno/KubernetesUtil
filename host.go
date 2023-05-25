package KubernetesUtil

import "os"

func FullHostname() string {
	//Kubernetes sets the env var HOSTNAME
	hostname, ok := os.LookupEnv("HOSTNAME")
	if ok {
		return hostname
	} else {
		return ""
	}
}
