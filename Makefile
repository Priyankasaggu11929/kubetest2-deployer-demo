aks_deployer:
	@MODE=dev ./hack/build.sh
	@export TF_DATA=${pwd}/data
	@echo "\n[INFO] kubetest2-aks local build successful! Start using it as: \"./bin/kubetest2-aks --help\""
