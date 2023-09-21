package main

import (
	"fmt"

	capiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
	kubeadmbootstrapv1beta1 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1beta1"
	clusterctlv1alpha3 "sigs.k8s.io/cluster-api/cmd/clusterctl/api/v1alpha3"
	kubeadmcontrolplanev1beta1 "sigs.k8s.io/cluster-api/controlplane/kubeadm/api/v1beta1"
	expv1beta1 "sigs.k8s.io/cluster-api/exp/api/v1beta1"
	ipamv1alpha1 "sigs.k8s.io/cluster-api/exp/ipam/api/v1alpha1"
	runtimev1alpha1 "sigs.k8s.io/cluster-api/exp/runtime/api/v1alpha1"
	dockerv1beta1 "sigs.k8s.io/cluster-api/test/infrastructure/docker/api/v1beta1"
	dockerexpv1beta1 "sigs.k8s.io/cluster-api/test/infrastructure/docker/exp/api/v1beta1"
	inmemoryv1alpha1 "sigs.k8s.io/cluster-api/test/infrastructure/inmemory/api/v1alpha1"
)

func main() {
	myCluster := &capiv1beta1.Cluster{}

	fmt.Printf("My Cluster: %+v\n", myCluster)

	myKubeControlPlane := &kubeadmcontrolplanev1beta1.KubeadmControlPlane{}

	fmt.Printf("My KubeControlPlane: %+v\n", myKubeControlPlane)

	myKubeADMConfig := &kubeadmbootstrapv1beta1.KubeadmConfig{}

	fmt.Printf("My KubeADMConfig: %+v\n", myKubeADMConfig)

	myMachinePool := &expv1beta1.MachinePool{}

	fmt.Printf("My MachinePool: %+v\n", myMachinePool)

	myIPAddressClaim := &ipamv1alpha1.IPAddressClaim{}

	fmt.Printf("My IPAddressClaim: %+v\n", myIPAddressClaim)

	myExtensionConfig := &runtimev1alpha1.ExtensionConfig{}

	fmt.Printf("My ExtensionConfig: %+v\n", myExtensionConfig)

	myProvider := &clusterctlv1alpha3.Provider{}

	fmt.Printf("My Provider: %+v\n", myProvider)

	myDockerCluster := &dockerv1beta1.DockerCluster{}

	fmt.Printf("My DockerCluster: %+v\n", myDockerCluster)

	myDockerMachinePool := &dockerexpv1beta1.DockerMachinePool{}

	fmt.Printf("My DockerMachinePool: %+v\n", myDockerMachinePool)

	myInMemoryCluster := &inmemoryv1alpha1.InMemoryCluster{}

	fmt.Printf("My InMemoryCluster: %+v\n", myInMemoryCluster)
}
