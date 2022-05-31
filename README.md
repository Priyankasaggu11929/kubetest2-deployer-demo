# kubetest2-deployer-demo

This project contains an example [kubetest2](https://github.com/kubernetes-sigs/kubetest2) custom deployer implementation code for deploying kubernetes on different cloud and run tests against/on them. In this specific case, we're trying to demonstrate writing a custom deployer for Azure cloud.

## kubetest2-aks

kubetest2-aks is a deployer created for deploying kubernetes using [Azure Kubernetes Service (AKS) Engine](https://azure.microsoft.com/en-in/services/kubernetes-service/#overview)

## Steps to build
```shell
$ make aks_deployer
```
