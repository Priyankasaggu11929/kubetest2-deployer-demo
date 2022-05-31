package azureaks

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/ppc64le-cloud/kubetest2-plugins/pkg/providers"
	"github.com/ppc64le-cloud/kubetest2-plugins/pkg/utils"
	"github.com/spf13/pflag"
)

const (
	Name = "azureaks"
)

var _ providers.Provider = &Provider{}

var AzureAKSProvider = &Provider{}

type Provider struct {
	ClusterName    string `json:"cluster_name"`
	BootstrapToken string `json:"bootstrap_token"`
	KubeconfigPath string `json:"kubeconfig_path"`
}

func (p *Provider) Initialize() error {
	randPostFix, err := utils.RandString(6)
	if err != nil {
		return fmt.Errorf("failed to generate a random string, error: %v", err)
	}
	p.ClusterName = "k8s-cluster-" + randPostFix

	bootstrapToken, err := utils.GenerateBootstrapToken()
	if err != nil {
		return fmt.Errorf("failed to generate a random string, error: %v", err)
	}
	p.BootstrapToken = bootstrapToken

	p.KubeconfigPath = path.Join(p.ClusterName, "kubeconfig")

	p.KubeconfigPath, err = filepath.Abs(p.KubeconfigPath)
	if err != nil {
		return fmt.Errorf("errored while getting absolute path for kubeconfig file")
	}
	return nil
}

func (p *Provider) BindFlags(flags *pflag.FlagSet) {
}

func (p *Provider) DumpConfig(dir string) error {
	return nil
}
