// Package templates deals with creating template assets that will be used by other assets
package templates

import (
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content/bootkube"
	"github.com/openshift/installer/pkg/asset/templates/content/tectonic"
)

var _ asset.WritableAsset = (*Templates)(nil)

// Templates generates the dependent unrendered template files
type Templates struct {
	FileList []*asset.File
}

// Name returns a human friendly name for the templates asset
func (m *Templates) Name() string {
	return "Common Templates"
}

// Dependencies returns all of the dependencies directly needed by a
// Templates asset.
func (m *Templates) Dependencies() []asset.Asset {
	return []asset.Asset{
		&bootkube.KubeCloudConfig{},
		&bootkube.MachineConfigServerTLSSecret{},
		&bootkube.OpenshiftServiceCertSignerSecret{},
		&bootkube.Pull{},
		&bootkube.CVOOverrides{},
		&bootkube.LegacyCVOOverrides{},
		&bootkube.EtcdServiceEndpointsKubeSystem{},
		&bootkube.HostEtcdServiceEndpointsKubeSystem{},
		&bootkube.KubeSystemConfigmapEtcdServingCA{},
		&bootkube.KubeSystemConfigmapRootCA{},
		&bootkube.KubeSystemSecretEtcdClient{},
		&bootkube.OpenshiftWebConsoleNamespace{},
		&bootkube.OpenshiftMachineConfigOperator{},
		&bootkube.OpenshiftClusterAPINamespace{},
		&bootkube.OpenshiftServiceCertSignerNamespace{},
		&bootkube.EtcdServiceKubeSystem{},
		&bootkube.HostEtcdServiceKubeSystem{},
		&tectonic.BindingDiscovery{},
		&tectonic.CloudCredsSecret{},
		&tectonic.RoleCloudCredsSecretReader{},
	}
}

// Generate generates the respective operator config.yml files
func (m *Templates) Generate(dependencies asset.Parents) error {
	kubeCloudConfig := &bootkube.KubeCloudConfig{}
	machineConfigServerTLSSecret := &bootkube.MachineConfigServerTLSSecret{}
	openshiftServiceCertSignerSecret := &bootkube.OpenshiftServiceCertSignerSecret{}
	pull := &bootkube.Pull{}
	cVOOverrides := &bootkube.CVOOverrides{}
	legacyCVOOverrides := &bootkube.LegacyCVOOverrides{}
	etcdServiceEndpointsKubeSystem := &bootkube.EtcdServiceEndpointsKubeSystem{}
	hostEtcdServiceEndpointsKubeSystem := &bootkube.HostEtcdServiceEndpointsKubeSystem{}
	kubeSystemConfigmapEtcdServingCA := &bootkube.KubeSystemConfigmapEtcdServingCA{}
	kubeSystemConfigmapRootCA := &bootkube.KubeSystemConfigmapRootCA{}
	kubeSystemSecretEtcdClient := &bootkube.KubeSystemSecretEtcdClient{}
	openshiftWebConsoleNamespace := &bootkube.OpenshiftWebConsoleNamespace{}
	openshiftMachineConfigOperator := &bootkube.OpenshiftMachineConfigOperator{}
	openshiftClusterAPINamespace := &bootkube.OpenshiftClusterAPINamespace{}
	openshiftServiceCertSignerNamespace := &bootkube.OpenshiftServiceCertSignerNamespace{}
	etcdServiceKubeSystem := &bootkube.EtcdServiceKubeSystem{}
	hostEtcdServiceKubeSystem := &bootkube.HostEtcdServiceKubeSystem{}

	bindingDiscovery := &tectonic.BindingDiscovery{}
	cloudCredsSecret := &tectonic.CloudCredsSecret{}
	roleCloudCredsSecretReader := &tectonic.RoleCloudCredsSecretReader{}

	dependencies.Get(
		kubeCloudConfig,
		machineConfigServerTLSSecret,
		openshiftServiceCertSignerSecret,
		pull,
		cVOOverrides,
		legacyCVOOverrides,
		etcdServiceEndpointsKubeSystem,
		hostEtcdServiceEndpointsKubeSystem,
		kubeSystemConfigmapEtcdServingCA,
		kubeSystemConfigmapRootCA,
		kubeSystemSecretEtcdClient,
		openshiftWebConsoleNamespace,
		openshiftMachineConfigOperator,
		openshiftClusterAPINamespace,
		openshiftServiceCertSignerNamespace,
		etcdServiceKubeSystem,
		hostEtcdServiceKubeSystem,
		bindingDiscovery,
		cloudCredsSecret,
		roleCloudCredsSecretReader)

	m.FileList = []*asset.File{}
	m.FileList = append(m.FileList, kubeCloudConfig.Files()...)
	m.FileList = append(m.FileList, machineConfigServerTLSSecret.Files()...)
	m.FileList = append(m.FileList, openshiftServiceCertSignerSecret.Files()...)
	m.FileList = append(m.FileList, pull.Files()...)
	m.FileList = append(m.FileList, cVOOverrides.Files()...)
	m.FileList = append(m.FileList, legacyCVOOverrides.Files()...)
	m.FileList = append(m.FileList, etcdServiceEndpointsKubeSystem.Files()...)
	m.FileList = append(m.FileList, hostEtcdServiceEndpointsKubeSystem.Files()...)
	m.FileList = append(m.FileList, kubeSystemConfigmapEtcdServingCA.Files()...)
	m.FileList = append(m.FileList, kubeSystemConfigmapRootCA.Files()...)
	m.FileList = append(m.FileList, kubeSystemSecretEtcdClient.Files()...)
	m.FileList = append(m.FileList, openshiftWebConsoleNamespace.Files()...)
	m.FileList = append(m.FileList, openshiftMachineConfigOperator.Files()...)
	m.FileList = append(m.FileList, openshiftClusterAPINamespace.Files()...)
	m.FileList = append(m.FileList, openshiftServiceCertSignerNamespace.Files()...)
	m.FileList = append(m.FileList, etcdServiceKubeSystem.Files()...)
	m.FileList = append(m.FileList, hostEtcdServiceKubeSystem.Files()...)

	m.FileList = append(m.FileList, bindingDiscovery.Files()...)
	m.FileList = append(m.FileList, cloudCredsSecret.Files()...)
	m.FileList = append(m.FileList, roleCloudCredsSecretReader.Files()...)

	return nil
}

// Files returns the files generated by the asset.
func (m *Templates) Files() []*asset.File {
	return m.FileList
}

// Load returns the manifests asset from disk.
func (m *Templates) Load(f asset.FileFetcher) (bool, error) {
	return false, nil
}
