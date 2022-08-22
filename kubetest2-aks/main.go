package main

import (
	"sigs.k8s.io/kubetest2/pkg/app"

	"github.com/Priyankasaggu11929/kubetest2-deployer-demo/kubetest2-aks/deployer"
)

func main() {
	app.Main(deployer.Name, deployer.New)
}
