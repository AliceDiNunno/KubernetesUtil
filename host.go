package KubernetesUtil

import (
	"log"
	"os"
)

func FullHostname() string {
	//Kubernetes sets the env var HOSTNAME
	hostname, ok := os.LookupEnv("HOSTNAME")
	if ok {
		//Debug
		log.Println("KU: HOSTNAME = ", hostname)
		//

		return hostname
	} else {
		//Debug
		log.Println("KU: HOSTNAME NOT FOUND")
		//

		return ""
	}
}
