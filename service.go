package KubernetesUtil

import (
	"log"
	"os"
	"strconv"
	"strings"
)

//Kubernetes sets the env var xxx_SERVICE_HOST and xxx_SERVICE_PORT for each service where xxx is the name of your service
//see https://kubernetes.io/docs/concepts/services-networking/service/#environment-variables

// For this to work, you must set the service name to the same value as the deployment name

func GetInternalServiceName() string {
	deploymentName := GetDeploymentName()

	if deploymentName == "" {
		//We can try to fallback to the hostname
		hostname := FullHostname()
		//splice the hostname with "-" and take the first part
		deploymentSplit := strings.Split(hostname, "-")
		if len(deploymentSplit) > 0 {
			deploymentName = deploymentSplit[0]
		} else {
			deploymentName = ""
		}
	}

	return GetDeploymentName() + "-svc"
}

func GetInternalServiceIP() string {
	envName := transformStringToEnvVarName(GetInternalServiceName() + "_HOST")

	serviceHost, ok := os.LookupEnv(envName + "_SERVICE_HOST")
	if ok {
		return serviceHost
	} else {
		return ""
	}
}

func GetInternalServicePort() int {
	envName := transformStringToEnvVarName(GetInternalServiceName() + "_PORT")

	// Debug
	log.Println("KU: envName = ", envName)
	//

	servicePort, ok := os.LookupEnv(envName + "_SERVICE_PORT")
	//Debug
	log.Println("KU: SERVICE_PORT = ", servicePort)
	//
	if ok {
		port, err := strconv.Atoi(servicePort)
		if err == nil {
			return port
		}

		return 0
	} else {
		return 0
	}
}
