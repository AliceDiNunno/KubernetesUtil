package KubernetesUtil

import (
	"log"
	"os"
	"strings"
)

func IsRunningInKubernetes() bool {
	//Kubernetes sets the env var KUBERNETES_SERVICE_HOST
	_, ok := os.LookupEnv("KUBERNETES_SERVICE_HOST")

	//Debug
	log.Println("KU: KUBERNETES_SERVICE_HOST = ", ok)
	//

	return ok
}

func transformStringToEnvVarName(s string) string {
	s = strings.Replace(s, "-", "_", -1)
	s = strings.Replace(s, ".", "_", -1)
	s = strings.Replace(s, "__", "_", -1)
	s = strings.ToUpper(s)

	return s
}
