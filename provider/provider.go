package provider

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/cloudbase/garm-provider-azure/config"

	"github.com/cloudbase/garm/params"
	"github.com/cloudbase/garm/runner/providers/external/execution"
)

var _ execution.ExternalProvider = &azureProvider{}

const (
	controllerIDTagName = "garm-controller-id"
	poolIDTagName       = "garm-pool-id"
)

func NewAzureProvider(execEnv execution.Environment) (execution.ExternalProvider, error) {
	if err := execEnv.Validate(); err != nil {
		return nil, fmt.Errorf("error validating execution environment: %w", err)
	}

	conf, err := config.NewConfig(execEnv.ProviderConfigFile)
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}
	creds, err := conf.DefaultCredentials()
	if err != nil {
		return nil, fmt.Errorf("failed to get credentials: %w", err)
	}
	return &azureProvider{
		cfg:          conf,
		defaultCreds: creds,
		execEnv:      execEnv,
	}, nil
}

type azureProvider struct {
	cfg          *config.Config
	defaultCreds azcore.TokenCredential
	execEnv      execution.Environment
}

// CreateInstance creates a new compute instance in the provider.
func (a *azureProvider) CreateInstance(ctx context.Context, bootstrapParams params.BootstrapInstance) (params.Instance, error) {
	return params.Instance{}, nil
}

// Delete instance will delete the instance in a provider.
func (a *azureProvider) DeleteInstance(ctx context.Context, instance string) error {
	return nil
}

// GetInstance will return details about one instance.
func (a *azureProvider) GetInstance(ctx context.Context, instance string) (params.Instance, error) {
	return params.Instance{}, nil
}

// ListInstances will list all instances for a provider.
func (a *azureProvider) ListInstances(ctx context.Context, poolID string) ([]params.Instance, error) {
	return nil, nil
}

// RemoveAllInstances will remove all instances created by this provider.
func (a *azureProvider) RemoveAllInstances(ctx context.Context) error {
	return nil
}

// Stop shuts down the instance.
func (a *azureProvider) Stop(ctx context.Context, instance string, force bool) error {
	return nil
}

// Start boots up an instance.
func (a *azureProvider) Start(ctx context.Context, instance string) error {
	return nil
}