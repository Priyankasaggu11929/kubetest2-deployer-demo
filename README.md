# kubetest2-deployer-demo

The project contains source code for a custom kubetest2 deployer that can be used to install and test Kubernetes on a variety of cloud platforms. 

In this specific use-case, we're attempting to demonstrate how to write a custom Azure cloud platform deployer.

## kubetest2-aks

kubetest2-aks is a deployer created for deploying kubernetes using [Azure Kubernetes Service (AKS) Engine](https://azure.microsoft.com/en-in/services/kubernetes-service/#overview)

## Steps to build locally
```shell
$ make aks_deployer
```

## Resources

- [Project - Kubetest2](https://github.com/kubernetes-sigs/kubetest2)
- [Provision an AKS Cluster (Azure)](https://learn.hashicorp.com/tutorials/terraform/aks)
- [Extend kubetest2 for Your Own Cloud](https://developer.ibm.com/conferences/kubesummit/extend-kubetest-for-deployer/)
