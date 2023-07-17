package main

import (
	"fmt"

	capiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

func main() {
	myCluster := &capiv1beta1.Cluster{}

	fmt.Printf("My Cluster: %+v\n", myCluster)
}
