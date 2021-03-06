package client

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	buildapi "github.com/openshift/origin/pkg/build/api"
	osclient "github.com/openshift/origin/pkg/client"
)

// BuildConfigGetter provides methods for getting BuildConfigs
type BuildConfigGetter interface {
	Get(namespace, name string, options metav1.GetOptions) (*buildapi.BuildConfig, error)
}

// BuildConfigUpdater provides methods for updating BuildConfigs
type BuildConfigUpdater interface {
	Update(buildConfig *buildapi.BuildConfig) error
}

// OSClientBuildConfigClient delegates get and update operations to the OpenShift client interface
type OSClientBuildConfigClient struct {
	Client osclient.Interface
}

// NewOSClientBuildConfigClient creates a new build config client that uses an openshift client to create and get BuildConfigs
func NewOSClientBuildConfigClient(client osclient.Interface) *OSClientBuildConfigClient {
	return &OSClientBuildConfigClient{Client: client}
}

// Get returns a BuildConfig using the OpenShift client.
func (c OSClientBuildConfigClient) Get(namespace, name string, options metav1.GetOptions) (*buildapi.BuildConfig, error) {
	return c.Client.BuildConfigs(namespace).Get(name, options)
}

// Update updates a BuildConfig using the OpenShift client.
func (c OSClientBuildConfigClient) Update(buildConfig *buildapi.BuildConfig) error {
	_, err := c.Client.BuildConfigs(buildConfig.Namespace).Update(buildConfig)
	return err
}

// BuildUpdater provides methods for updating existing Builds.
type BuildUpdater interface {
	Update(namespace string, build *buildapi.Build) error
}

type BuildPatcher interface {
	Patch(namespace, name string, patch []byte) (*buildapi.Build, error)
}

// BuildLister provides methods for listing the Builds.
type BuildLister interface {
	List(namespace string, opts metav1.ListOptions) (*buildapi.BuildList, error)
}

// BuildDeleter knows how to delete builds from OpenShift.
type BuildDeleter interface {
	// DeleteBuild removes the build from OpenShift's storage.
	DeleteBuild(build *buildapi.Build) error
}

// OSClientBuildClient delegates build create and update operations to the OpenShift client interface
type OSClientBuildClient struct {
	Client osclient.Interface
}

// NewOSClientBuildClient creates a new build client that uses an openshift client to update builds
func NewOSClientBuildClient(client osclient.Interface) *OSClientBuildClient {
	return &OSClientBuildClient{Client: client}
}

// Update updates builds using the OpenShift client.
func (c OSClientBuildClient) Update(namespace string, build *buildapi.Build) error {
	_, e := c.Client.Builds(namespace).Update(build)
	return e
}

// Patch patches builds using the OpenShift client.
func (c OSClientBuildClient) Patch(namespace, name string, patch []byte) (*buildapi.Build, error) {
	return c.Client.Builds(namespace).Patch(name, types.StrategicMergePatchType, patch)
}

// List lists the builds using the OpenShift client.
func (c OSClientBuildClient) List(namespace string, opts metav1.ListOptions) (*buildapi.BuildList, error) {
	return c.Client.Builds(namespace).List(opts)
}

// DeleteBuild deletes a build from OpenShift.
func (c OSClientBuildClient) DeleteBuild(build *buildapi.Build) error {
	return c.Client.Builds(build.Namespace).Delete(build.Name)
}

// BuildCloner provides methods for cloning builds
type BuildCloner interface {
	Clone(namespace string, request *buildapi.BuildRequest) (*buildapi.Build, error)
}

// OSClientBuildClonerClient creates a new build client that uses an openshift client to clone builds
type OSClientBuildClonerClient struct {
	Client osclient.Interface
}

// NewOSClientBuildClonerClient creates a new build client that uses an openshift client to clone builds
func NewOSClientBuildClonerClient(client osclient.Interface) *OSClientBuildClonerClient {
	return &OSClientBuildClonerClient{Client: client}
}

// Clone generates new build for given build name
func (c OSClientBuildClonerClient) Clone(namespace string, request *buildapi.BuildRequest) (*buildapi.Build, error) {
	return c.Client.Builds(namespace).Clone(request)
}

// BuildConfigInstantiator provides methods for instantiating builds from build configs
type BuildConfigInstantiator interface {
	Instantiate(namespace string, request *buildapi.BuildRequest) (*buildapi.Build, error)
}

// OSClientBuildConfigInstantiatorClient creates a new build client that uses an openshift client to create builds
type OSClientBuildConfigInstantiatorClient struct {
	Client osclient.Interface
}

// NewOSClientBuildConfigInstantiatorClient creates a new build client that uses an openshift client to create builds
func NewOSClientBuildConfigInstantiatorClient(client osclient.Interface) *OSClientBuildConfigInstantiatorClient {
	return &OSClientBuildConfigInstantiatorClient{Client: client}
}

// Instantiate generates new build for given buildConfig
func (c OSClientBuildConfigInstantiatorClient) Instantiate(namespace string, request *buildapi.BuildRequest) (*buildapi.Build, error) {
	return c.Client.BuildConfigs(namespace).Instantiate(request)
}
