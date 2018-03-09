package cloudprovider

import (
	"github.com/pkg/errors"
)

type Operations interface {
	CreateOrUpdateSSL(appName, cert string, port int) error
}

type K8sOperations interface {
	CloudProviderName() (string, error)
	SetServiceAnnotations(namespace, service string, annotations map[string]string) error
}

func NewOperations(k8s K8sOperations) (Operations, error) {
	name, err := k8s.CloudProviderName()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get cloud provider name")
	}
	switch name {
	case "aws":
		return &awsOperations{k8s: k8s}, nil
	case "gce":
		return &gceOperations{k8s: k8s}, nil
	default:
		return nil, ErrInvalidCloudProvider
	}
}
