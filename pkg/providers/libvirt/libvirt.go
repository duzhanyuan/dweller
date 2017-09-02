package libvirt

import (
	"github.com/libvirt/libvirt-go"

	"github.com/cloudflavor/dweller/pkg/providers/config"
)

var (
	defaultSystemURI = "qemu:///system"
)

// LibvirtProvider contains the libvirt connection instance information.
type LibvirtProvider struct {
	client *libvirt.Connect
}

// NewInfra creates a new cloudflavor infrastructure on top of qemu formed out
// of 1xMaster and 2xWorker Nodes with an attached storage.
func (lb *LibvirtProvider) NewInfra(config *config.Infra) error {
	return nil
}

// RegisterInstances adds a new instance to the infrastructure.
func (lb *LibvirtProvider) RegisterInstances(config *config.Infra) error {
	return nil
}

// DestroyInstances destroys a specific instance from the infrastructure.
func (lb *LibvirtProvider) DestroyInstances(config *config.Infra) error {
	return nil
}

// HaltInfra halts the current running infrastructure. It can also pause the
// current running infra and even destroy it if --destroy flag was passed.
func (lb *LibvirtProvider) HaltInfra(config *config.Infra) error {
	return nil
}

// NewLibvirtProvider creates a new libvirt provider with an active connection.
func NewLibvirtProvider(config *config.Infra) (*LibvirtProvider, error) {
	if config.LibvirtURI == nil {
		config.LibvirtURI = &defaultSystemURI
	}

	conn, err := libvirt.NewConnect(*config.LibvirtURI)
	if err != nil {
		return nil, err
	}

	return &LibvirtProvider{
		client: conn,
	}, nil
}
