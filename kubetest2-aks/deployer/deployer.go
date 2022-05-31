package deployer

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/ppc64le-cloud/kubetest2-plugins/pkg/providers"
	"github.com/ppc64le-cloud/kubetest2-plugins/pkg/providers/azureaks"
	"github.com/ppc64le-cloud/kubetest2-plugins/pkg/terraform"
	"github.com/ppc64le-cloud/kubetest2-plugins/pkg/utils"
	"github.com/spf13/pflag"
	"k8s.io/klog/v2"
	"sigs.k8s.io/kubetest2/pkg/types"
)

const (
	//Name = "tf"
	Name = "aks"
)

type deployer struct {
	commonOptions types.Options
	logsDir       string
	doInit        sync.Once
	tmpDir        string
	provider      providers.Provider
}

func (d *deployer) init() error {
	var err error
	d.doInit.Do(func() { err = d.initialize() })
	return err
}

func (d *deployer) initialize() error {
	d.provider = azureaks.AzureAKSProvider
	randPostFix, err := utils.RandString(6)
	if err != nil {
		return fmt.Errorf("failed to generate a random string, error: %v", err)
	}
	d.tmpDir = "k8s-cluster-" + randPostFix
	return nil
}

var _ types.Deployer = &deployer{}

var (
	ignoreClusterDir      bool
	autoApprove           bool
	retryOnTfFailure      int
	breakKubetestOnUpFail bool
)

func New(opts types.Options) (types.Deployer, *pflag.FlagSet) {
	d := &deployer{
		commonOptions: opts,
		logsDir:       filepath.Join(opts.RunDir(), "logs"),
	}
	return d, bindFlags(d)
}

func bindFlags(d *deployer) *pflag.FlagSet {
	flags := pflag.NewFlagSet(Name, pflag.ContinueOnError)
	flags.BoolVar(
		&ignoreClusterDir, "ignore-cluster-dir", false, "Ignore the cluster folder if exists",
	)
	flags.BoolVar(
		&autoApprove, "auto-approve", false, "Terraform Auto Approve",
	)
	flags.IntVar(
		&retryOnTfFailure, "retry-on-tf-failure", 1, "Retry on Terraform Apply Failure",
	)
	flags.BoolVar(
		&breakKubetestOnUpFail, "break-kubetest-on-upfail", false, "Breaks kubetest2 when up fails",
	)
	flags.MarkHidden("ignore-cluster-dir")
	azureaks.AzureAKSProvider.BindFlags(flags)

	return flags
}

func (d *deployer) Up() error {
	if err := d.init(); err != nil {
		return fmt.Errorf("up failed to init: %s", err)
	}

	for i := 0; i <= retryOnTfFailure; i++ {
		path, err := terraform.Apply(d.tmpDir, "azureaks", autoApprove)
		op, oerr := terraform.Output(d.tmpDir, "azureaks")
		if err != nil {
			if i == retryOnTfFailure {
				fmt.Printf("terraform.Output: %s\nterraform.Output error: %v\n", op, oerr)
				if !breakKubetestOnUpFail {
					return fmt.Errorf("Terraform Apply failed. Error: %v\n", err)
				}
				klog.Infof("Terraform Apply failed. Look into it and delete the resources")
				klog.Infof("terraform.Apply error: %v", err)
				os.Exit(1)
			}
			continue
		} else {
			fmt.Printf("terraform.Output: %s\nterraform.Output error: %v\n", op, oerr)
			fmt.Printf("Terraform State at: %s\n", path)
			break
		}
	}

	return nil
}

func (d *deployer) Down() error {
	if err := d.init(); err != nil {
		return fmt.Errorf("down failed to init: %s", err)
	}
	err := terraform.Destroy(d.tmpDir, "azureaks", autoApprove)
	if err != nil {
		return fmt.Errorf("terraform.Destroy failed: %v", err)
	}
	return nil
}

func (d *deployer) IsUp() (up bool, err error) {
	panic("implement me")
}

func (d *deployer) DumpClusterLogs() error {
	panic("implement me")
}

func (d *deployer) Build() error {
	panic("implement me")
}
