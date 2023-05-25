package KubernetesUtil

import (
	"log"
	"os"
)

func GetDeploymentName() string {
	//you must set the env var DEPLOYMENT_NAME in your deployment yaml using the value of metadata.name to use this api
	//see https://kubernetes.io/docs/tasks/inject-data-application/environment-variable-expose-pod-information/

	deploymentName, ok := os.LookupEnv("DEPLOYMENT_NAME")
	if ok {

		//Debug
		log.Println("KU: DEPLOYMENT_NAME = ", deploymentName)
		//

		return deploymentName
	} else {
		//Debug
		log.Println("KU: DEPLOYMENT_NAME NOT FOUND")
		//
		return ""
	}
}
