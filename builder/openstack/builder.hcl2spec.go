// Code generated by "mapstructure-to-hcl2 -type Config,ImageFilter,ImageFilterOptions"; DO NOT EDIT.

package openstack

import (
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName               *string                 `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType             *string                 `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion             *string                 `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug                   *bool                   `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce                   *bool                   `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError                 *string                 `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars                map[string]string       `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars           []string                `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	Username                      *string                 `mapstructure:"username" required:"true" cty:"username" hcl:"username"`
	UserID                        *string                 `mapstructure:"user_id" cty:"user_id" hcl:"user_id"`
	Password                      *string                 `mapstructure:"password" required:"true" cty:"password" hcl:"password"`
	IdentityEndpoint              *string                 `mapstructure:"identity_endpoint" required:"true" cty:"identity_endpoint" hcl:"identity_endpoint"`
	TenantID                      *string                 `mapstructure:"tenant_id" required:"false" cty:"tenant_id" hcl:"tenant_id"`
	TenantName                    *string                 `mapstructure:"tenant_name" cty:"tenant_name" hcl:"tenant_name"`
	DomainID                      *string                 `mapstructure:"domain_id" cty:"domain_id" hcl:"domain_id"`
	DomainName                    *string                 `mapstructure:"domain_name" required:"false" cty:"domain_name" hcl:"domain_name"`
	Insecure                      *bool                   `mapstructure:"insecure" required:"false" cty:"insecure" hcl:"insecure"`
	Region                        *string                 `mapstructure:"region" required:"false" cty:"region" hcl:"region"`
	EndpointType                  *string                 `mapstructure:"endpoint_type" required:"false" cty:"endpoint_type" hcl:"endpoint_type"`
	CACertFile                    *string                 `mapstructure:"cacert" required:"false" cty:"cacert" hcl:"cacert"`
	ClientCertFile                *string                 `mapstructure:"cert" required:"false" cty:"cert" hcl:"cert"`
	ClientKeyFile                 *string                 `mapstructure:"key" required:"false" cty:"key" hcl:"key"`
	Token                         *string                 `mapstructure:"token" required:"false" cty:"token" hcl:"token"`
	ApplicationCredentialName     *string                 `mapstructure:"application_credential_name" required:"false" cty:"application_credential_name" hcl:"application_credential_name"`
	ApplicationCredentialID       *string                 `mapstructure:"application_credential_id" required:"false" cty:"application_credential_id" hcl:"application_credential_id"`
	ApplicationCredentialSecret   *string                 `mapstructure:"application_credential_secret" required:"false" cty:"application_credential_secret" hcl:"application_credential_secret"`
	Cloud                         *string                 `mapstructure:"cloud" required:"false" cty:"cloud" hcl:"cloud"`
	ImageName                     *string                 `mapstructure:"image_name" required:"true" cty:"image_name" hcl:"image_name"`
	ImageMetadata                 map[string]string       `mapstructure:"metadata" required:"false" cty:"metadata" hcl:"metadata"`
	ImageVisibility               *images.ImageVisibility `mapstructure:"image_visibility" required:"false" cty:"image_visibility" hcl:"image_visibility"`
	ImageMembers                  []string                `mapstructure:"image_members" required:"false" cty:"image_members" hcl:"image_members"`
	ImageAutoAcceptMembers        *bool                   `mapstructure:"image_auto_accept_members" required:"false" cty:"image_auto_accept_members" hcl:"image_auto_accept_members"`
	ImageDiskFormat               *string                 `mapstructure:"image_disk_format" required:"false" cty:"image_disk_format" hcl:"image_disk_format"`
	ImageTags                     []string                `mapstructure:"image_tags" required:"false" cty:"image_tags" hcl:"image_tags"`
	ImageMinDisk                  *int                    `mapstructure:"image_min_disk" required:"false" cty:"image_min_disk" hcl:"image_min_disk"`
	SkipCreateImage               *bool                   `mapstructure:"skip_create_image" required:"false" cty:"skip_create_image" hcl:"skip_create_image"`
	Type                          *string                 `mapstructure:"communicator" cty:"communicator" hcl:"communicator"`
	PauseBeforeConnect            *string                 `mapstructure:"pause_before_connecting" cty:"pause_before_connecting" hcl:"pause_before_connecting"`
	SSHHost                       *string                 `mapstructure:"ssh_host" cty:"ssh_host" hcl:"ssh_host"`
	SSHPort                       *int                    `mapstructure:"ssh_port" cty:"ssh_port" hcl:"ssh_port"`
	SSHUsername                   *string                 `mapstructure:"ssh_username" cty:"ssh_username" hcl:"ssh_username"`
	SSHPassword                   *string                 `mapstructure:"ssh_password" cty:"ssh_password" hcl:"ssh_password"`
	SSHKeyPairName                *string                 `mapstructure:"ssh_keypair_name" undocumented:"true" cty:"ssh_keypair_name" hcl:"ssh_keypair_name"`
	SSHTemporaryKeyPairName       *string                 `mapstructure:"temporary_key_pair_name" undocumented:"true" cty:"temporary_key_pair_name" hcl:"temporary_key_pair_name"`
	SSHTemporaryKeyPairType       *string                 `mapstructure:"temporary_key_pair_type" cty:"temporary_key_pair_type" hcl:"temporary_key_pair_type"`
	SSHTemporaryKeyPairBits       *int                    `mapstructure:"temporary_key_pair_bits" cty:"temporary_key_pair_bits" hcl:"temporary_key_pair_bits"`
	SSHCiphers                    []string                `mapstructure:"ssh_ciphers" cty:"ssh_ciphers" hcl:"ssh_ciphers"`
	SSHClearAuthorizedKeys        *bool                   `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys" hcl:"ssh_clear_authorized_keys"`
	SSHKEXAlgos                   []string                `mapstructure:"ssh_key_exchange_algorithms" cty:"ssh_key_exchange_algorithms" hcl:"ssh_key_exchange_algorithms"`
	SSHPrivateKeyFile             *string                 `mapstructure:"ssh_private_key_file" undocumented:"true" cty:"ssh_private_key_file" hcl:"ssh_private_key_file"`
	SSHCertificateFile            *string                 `mapstructure:"ssh_certificate_file" cty:"ssh_certificate_file" hcl:"ssh_certificate_file"`
	SSHPty                        *bool                   `mapstructure:"ssh_pty" cty:"ssh_pty" hcl:"ssh_pty"`
	SSHTimeout                    *string                 `mapstructure:"ssh_timeout" cty:"ssh_timeout" hcl:"ssh_timeout"`
	SSHWaitTimeout                *string                 `mapstructure:"ssh_wait_timeout" undocumented:"true" cty:"ssh_wait_timeout" hcl:"ssh_wait_timeout"`
	SSHAgentAuth                  *bool                   `mapstructure:"ssh_agent_auth" undocumented:"true" cty:"ssh_agent_auth" hcl:"ssh_agent_auth"`
	SSHDisableAgentForwarding     *bool                   `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding" hcl:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts          *int                    `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts" hcl:"ssh_handshake_attempts"`
	SSHBastionHost                *string                 `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host" hcl:"ssh_bastion_host"`
	SSHBastionPort                *int                    `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port" hcl:"ssh_bastion_port"`
	SSHBastionAgentAuth           *bool                   `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth" hcl:"ssh_bastion_agent_auth"`
	SSHBastionUsername            *string                 `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username" hcl:"ssh_bastion_username"`
	SSHBastionPassword            *string                 `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password" hcl:"ssh_bastion_password"`
	SSHBastionInteractive         *bool                   `mapstructure:"ssh_bastion_interactive" cty:"ssh_bastion_interactive" hcl:"ssh_bastion_interactive"`
	SSHBastionPrivateKeyFile      *string                 `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file" hcl:"ssh_bastion_private_key_file"`
	SSHBastionCertificateFile     *string                 `mapstructure:"ssh_bastion_certificate_file" cty:"ssh_bastion_certificate_file" hcl:"ssh_bastion_certificate_file"`
	SSHFileTransferMethod         *string                 `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method" hcl:"ssh_file_transfer_method"`
	SSHProxyHost                  *string                 `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host" hcl:"ssh_proxy_host"`
	SSHProxyPort                  *int                    `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port" hcl:"ssh_proxy_port"`
	SSHProxyUsername              *string                 `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username" hcl:"ssh_proxy_username"`
	SSHProxyPassword              *string                 `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password" hcl:"ssh_proxy_password"`
	SSHKeepAliveInterval          *string                 `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval" hcl:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout           *string                 `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout" hcl:"ssh_read_write_timeout"`
	SSHRemoteTunnels              []string                `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels" hcl:"ssh_remote_tunnels"`
	SSHLocalTunnels               []string                `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels" hcl:"ssh_local_tunnels"`
	SSHPublicKey                  []byte                  `mapstructure:"ssh_public_key" undocumented:"true" cty:"ssh_public_key" hcl:"ssh_public_key"`
	SSHPrivateKey                 []byte                  `mapstructure:"ssh_private_key" undocumented:"true" cty:"ssh_private_key" hcl:"ssh_private_key"`
	WinRMUser                     *string                 `mapstructure:"winrm_username" cty:"winrm_username" hcl:"winrm_username"`
	WinRMPassword                 *string                 `mapstructure:"winrm_password" cty:"winrm_password" hcl:"winrm_password"`
	WinRMHost                     *string                 `mapstructure:"winrm_host" cty:"winrm_host" hcl:"winrm_host"`
	WinRMNoProxy                  *bool                   `mapstructure:"winrm_no_proxy" cty:"winrm_no_proxy" hcl:"winrm_no_proxy"`
	WinRMPort                     *int                    `mapstructure:"winrm_port" cty:"winrm_port" hcl:"winrm_port"`
	WinRMTimeout                  *string                 `mapstructure:"winrm_timeout" cty:"winrm_timeout" hcl:"winrm_timeout"`
	WinRMUseSSL                   *bool                   `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl" hcl:"winrm_use_ssl"`
	WinRMInsecure                 *bool                   `mapstructure:"winrm_insecure" cty:"winrm_insecure" hcl:"winrm_insecure"`
	WinRMUseNTLM                  *bool                   `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm" hcl:"winrm_use_ntlm"`
	SSHInterface                  *string                 `mapstructure:"ssh_interface" required:"false" cty:"ssh_interface" hcl:"ssh_interface"`
	SSHIPVersion                  *string                 `mapstructure:"ssh_ip_version" required:"false" cty:"ssh_ip_version" hcl:"ssh_ip_version"`
	SourceImage                   *string                 `mapstructure:"source_image" required:"true" cty:"source_image" hcl:"source_image"`
	SourceImageName               *string                 `mapstructure:"source_image_name" required:"true" cty:"source_image_name" hcl:"source_image_name"`
	ExternalSourceImageURL        *string                 `mapstructure:"external_source_image_url" required:"true" cty:"external_source_image_url" hcl:"external_source_image_url"`
	ExternalSourceImageFormat     *string                 `mapstructure:"external_source_image_format" required:"false" cty:"external_source_image_format" hcl:"external_source_image_format"`
	ExternalSourceImageProperties map[string]string       `mapstructure:"external_source_image_properties" required:"false" cty:"external_source_image_properties" hcl:"external_source_image_properties"`
	SourceImageFilters            *FlatImageFilter        `mapstructure:"source_image_filter" required:"true" cty:"source_image_filter" hcl:"source_image_filter"`
	Flavor                        *string                 `mapstructure:"flavor" required:"true" cty:"flavor" hcl:"flavor"`
	AvailabilityZone              *string                 `mapstructure:"availability_zone" required:"false" cty:"availability_zone" hcl:"availability_zone"`
	RackconnectWait               *bool                   `mapstructure:"rackconnect_wait" required:"false" cty:"rackconnect_wait" hcl:"rackconnect_wait"`
	FloatingIPNetwork             *string                 `mapstructure:"floating_ip_network" required:"false" cty:"floating_ip_network" hcl:"floating_ip_network"`
	InstanceFloatingIPNet         *string                 `mapstructure:"instance_floating_ip_net" required:"false" cty:"instance_floating_ip_net" hcl:"instance_floating_ip_net"`
	FloatingIP                    *string                 `mapstructure:"floating_ip" required:"false" cty:"floating_ip" hcl:"floating_ip"`
	ReuseIPs                      *bool                   `mapstructure:"reuse_ips" required:"false" cty:"reuse_ips" hcl:"reuse_ips"`
	SecurityGroups                []string                `mapstructure:"security_groups" required:"false" cty:"security_groups" hcl:"security_groups"`
	Networks                      []string                `mapstructure:"networks" required:"false" cty:"networks" hcl:"networks"`
	Ports                         []string                `mapstructure:"ports" required:"false" cty:"ports" hcl:"ports"`
	NetworkDiscoveryCIDRs         []string                `mapstructure:"network_discovery_cidrs" required:"false" cty:"network_discovery_cidrs" hcl:"network_discovery_cidrs"`
	UserData                      *string                 `mapstructure:"user_data" required:"false" cty:"user_data" hcl:"user_data"`
	UserDataFile                  *string                 `mapstructure:"user_data_file" required:"false" cty:"user_data_file" hcl:"user_data_file"`
	InstanceName                  *string                 `mapstructure:"instance_name" required:"false" cty:"instance_name" hcl:"instance_name"`
	InstanceMetadata              map[string]string       `mapstructure:"instance_metadata" required:"false" cty:"instance_metadata" hcl:"instance_metadata"`
	ForceDelete                   *bool                   `mapstructure:"force_delete" required:"false" cty:"force_delete" hcl:"force_delete"`
	ConfigDrive                   *bool                   `mapstructure:"config_drive" required:"false" cty:"config_drive" hcl:"config_drive"`
	FloatingIPPool                *string                 `mapstructure:"floating_ip_pool" required:"false" cty:"floating_ip_pool" hcl:"floating_ip_pool"`
	UseBlockStorageVolume         *bool                   `mapstructure:"use_blockstorage_volume" required:"false" cty:"use_blockstorage_volume" hcl:"use_blockstorage_volume"`
	VolumeName                    *string                 `mapstructure:"volume_name" required:"false" cty:"volume_name" hcl:"volume_name"`
	VolumeType                    *string                 `mapstructure:"volume_type" required:"false" cty:"volume_type" hcl:"volume_type"`
	VolumeSize                    *int                    `mapstructure:"volume_size" required:"false" cty:"volume_size" hcl:"volume_size"`
	VolumeAvailabilityZone        *string                 `mapstructure:"volume_availability_zone" required:"false" cty:"volume_availability_zone" hcl:"volume_availability_zone"`
	OpenstackProvider             *string                 `mapstructure:"openstack_provider" cty:"openstack_provider" hcl:"openstack_provider"`
	UseFloatingIp                 *bool                   `mapstructure:"use_floating_ip" required:"false" cty:"use_floating_ip" hcl:"use_floating_ip"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":                &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":              &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":              &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":                     &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                     &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":                  &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":            &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables":       &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"username":                         &hcldec.AttrSpec{Name: "username", Type: cty.String, Required: false},
		"user_id":                          &hcldec.AttrSpec{Name: "user_id", Type: cty.String, Required: false},
		"password":                         &hcldec.AttrSpec{Name: "password", Type: cty.String, Required: false},
		"identity_endpoint":                &hcldec.AttrSpec{Name: "identity_endpoint", Type: cty.String, Required: false},
		"tenant_id":                        &hcldec.AttrSpec{Name: "tenant_id", Type: cty.String, Required: false},
		"tenant_name":                      &hcldec.AttrSpec{Name: "tenant_name", Type: cty.String, Required: false},
		"domain_id":                        &hcldec.AttrSpec{Name: "domain_id", Type: cty.String, Required: false},
		"domain_name":                      &hcldec.AttrSpec{Name: "domain_name", Type: cty.String, Required: false},
		"insecure":                         &hcldec.AttrSpec{Name: "insecure", Type: cty.Bool, Required: false},
		"region":                           &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"endpoint_type":                    &hcldec.AttrSpec{Name: "endpoint_type", Type: cty.String, Required: false},
		"cacert":                           &hcldec.AttrSpec{Name: "cacert", Type: cty.String, Required: false},
		"cert":                             &hcldec.AttrSpec{Name: "cert", Type: cty.String, Required: false},
		"key":                              &hcldec.AttrSpec{Name: "key", Type: cty.String, Required: false},
		"token":                            &hcldec.AttrSpec{Name: "token", Type: cty.String, Required: false},
		"application_credential_name":      &hcldec.AttrSpec{Name: "application_credential_name", Type: cty.String, Required: false},
		"application_credential_id":        &hcldec.AttrSpec{Name: "application_credential_id", Type: cty.String, Required: false},
		"application_credential_secret":    &hcldec.AttrSpec{Name: "application_credential_secret", Type: cty.String, Required: false},
		"cloud":                            &hcldec.AttrSpec{Name: "cloud", Type: cty.String, Required: false},
		"image_name":                       &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"metadata":                         &hcldec.AttrSpec{Name: "metadata", Type: cty.Map(cty.String), Required: false},
		"image_visibility":                 &hcldec.AttrSpec{Name: "image_visibility", Type: cty.String, Required: false},
		"image_members":                    &hcldec.AttrSpec{Name: "image_members", Type: cty.List(cty.String), Required: false},
		"image_auto_accept_members":        &hcldec.AttrSpec{Name: "image_auto_accept_members", Type: cty.Bool, Required: false},
		"image_disk_format":                &hcldec.AttrSpec{Name: "image_disk_format", Type: cty.String, Required: false},
		"image_tags":                       &hcldec.AttrSpec{Name: "image_tags", Type: cty.List(cty.String), Required: false},
		"image_min_disk":                   &hcldec.AttrSpec{Name: "image_min_disk", Type: cty.Number, Required: false},
		"skip_create_image":                &hcldec.AttrSpec{Name: "skip_create_image", Type: cty.Bool, Required: false},
		"communicator":                     &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":          &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                         &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                         &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                     &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                     &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":                 &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":          &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"temporary_key_pair_type":          &hcldec.AttrSpec{Name: "temporary_key_pair_type", Type: cty.String, Required: false},
		"temporary_key_pair_bits":          &hcldec.AttrSpec{Name: "temporary_key_pair_bits", Type: cty.Number, Required: false},
		"ssh_ciphers":                      &hcldec.AttrSpec{Name: "ssh_ciphers", Type: cty.List(cty.String), Required: false},
		"ssh_clear_authorized_keys":        &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_key_exchange_algorithms":      &hcldec.AttrSpec{Name: "ssh_key_exchange_algorithms", Type: cty.List(cty.String), Required: false},
		"ssh_private_key_file":             &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_certificate_file":             &hcldec.AttrSpec{Name: "ssh_certificate_file", Type: cty.String, Required: false},
		"ssh_pty":                          &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                      &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_wait_timeout":                 &hcldec.AttrSpec{Name: "ssh_wait_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":                   &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding":     &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":           &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":                 &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":                 &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":           &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":             &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":             &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_interactive":          &hcldec.AttrSpec{Name: "ssh_bastion_interactive", Type: cty.Bool, Required: false},
		"ssh_bastion_private_key_file":     &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_bastion_certificate_file":     &hcldec.AttrSpec{Name: "ssh_bastion_certificate_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":         &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":                   &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":                   &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":               &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":               &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":          &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":           &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":               &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":                &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":                   &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":                  &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":                   &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":                   &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                       &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_no_proxy":                   &hcldec.AttrSpec{Name: "winrm_no_proxy", Type: cty.Bool, Required: false},
		"winrm_port":                       &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                    &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                    &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":                   &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":                   &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"ssh_interface":                    &hcldec.AttrSpec{Name: "ssh_interface", Type: cty.String, Required: false},
		"ssh_ip_version":                   &hcldec.AttrSpec{Name: "ssh_ip_version", Type: cty.String, Required: false},
		"source_image":                     &hcldec.AttrSpec{Name: "source_image", Type: cty.String, Required: false},
		"source_image_name":                &hcldec.AttrSpec{Name: "source_image_name", Type: cty.String, Required: false},
		"external_source_image_url":        &hcldec.AttrSpec{Name: "external_source_image_url", Type: cty.String, Required: false},
		"external_source_image_format":     &hcldec.AttrSpec{Name: "external_source_image_format", Type: cty.String, Required: false},
		"external_source_image_properties": &hcldec.AttrSpec{Name: "external_source_image_properties", Type: cty.Map(cty.String), Required: false},
		"source_image_filter":              &hcldec.BlockSpec{TypeName: "source_image_filter", Nested: hcldec.ObjectSpec((*FlatImageFilter)(nil).HCL2Spec())},
		"flavor":                           &hcldec.AttrSpec{Name: "flavor", Type: cty.String, Required: false},
		"availability_zone":                &hcldec.AttrSpec{Name: "availability_zone", Type: cty.String, Required: false},
		"rackconnect_wait":                 &hcldec.AttrSpec{Name: "rackconnect_wait", Type: cty.Bool, Required: false},
		"floating_ip_network":              &hcldec.AttrSpec{Name: "floating_ip_network", Type: cty.String, Required: false},
		"instance_floating_ip_net":         &hcldec.AttrSpec{Name: "instance_floating_ip_net", Type: cty.String, Required: false},
		"floating_ip":                      &hcldec.AttrSpec{Name: "floating_ip", Type: cty.String, Required: false},
		"reuse_ips":                        &hcldec.AttrSpec{Name: "reuse_ips", Type: cty.Bool, Required: false},
		"security_groups":                  &hcldec.AttrSpec{Name: "security_groups", Type: cty.List(cty.String), Required: false},
		"networks":                         &hcldec.AttrSpec{Name: "networks", Type: cty.List(cty.String), Required: false},
		"ports":                            &hcldec.AttrSpec{Name: "ports", Type: cty.List(cty.String), Required: false},
		"network_discovery_cidrs":          &hcldec.AttrSpec{Name: "network_discovery_cidrs", Type: cty.List(cty.String), Required: false},
		"user_data":                        &hcldec.AttrSpec{Name: "user_data", Type: cty.String, Required: false},
		"user_data_file":                   &hcldec.AttrSpec{Name: "user_data_file", Type: cty.String, Required: false},
		"instance_name":                    &hcldec.AttrSpec{Name: "instance_name", Type: cty.String, Required: false},
		"instance_metadata":                &hcldec.AttrSpec{Name: "instance_metadata", Type: cty.Map(cty.String), Required: false},
		"force_delete":                     &hcldec.AttrSpec{Name: "force_delete", Type: cty.Bool, Required: false},
		"config_drive":                     &hcldec.AttrSpec{Name: "config_drive", Type: cty.Bool, Required: false},
		"floating_ip_pool":                 &hcldec.AttrSpec{Name: "floating_ip_pool", Type: cty.String, Required: false},
		"use_blockstorage_volume":          &hcldec.AttrSpec{Name: "use_blockstorage_volume", Type: cty.Bool, Required: false},
		"volume_name":                      &hcldec.AttrSpec{Name: "volume_name", Type: cty.String, Required: false},
		"volume_type":                      &hcldec.AttrSpec{Name: "volume_type", Type: cty.String, Required: false},
		"volume_size":                      &hcldec.AttrSpec{Name: "volume_size", Type: cty.Number, Required: false},
		"volume_availability_zone":         &hcldec.AttrSpec{Name: "volume_availability_zone", Type: cty.String, Required: false},
		"openstack_provider":               &hcldec.AttrSpec{Name: "openstack_provider", Type: cty.String, Required: false},
		"use_floating_ip":                  &hcldec.AttrSpec{Name: "use_floating_ip", Type: cty.Bool, Required: false},
	}
	return s
}

// FlatImageFilter is an auto-generated flat version of ImageFilter.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatImageFilter struct {
	Filters    *FlatImageFilterOptions `mapstructure:"filters" required:"false" cty:"filters" hcl:"filters"`
	MostRecent *bool                   `mapstructure:"most_recent" required:"false" cty:"most_recent" hcl:"most_recent"`
}

// FlatMapstructure returns a new FlatImageFilter.
// FlatImageFilter is an auto-generated flat version of ImageFilter.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*ImageFilter) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatImageFilter)
}

// HCL2Spec returns the hcl spec of a ImageFilter.
// This spec is used by HCL to read the fields of ImageFilter.
// The decoded values from this spec will then be applied to a FlatImageFilter.
func (*FlatImageFilter) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"filters":     &hcldec.BlockSpec{TypeName: "filters", Nested: hcldec.ObjectSpec((*FlatImageFilterOptions)(nil).HCL2Spec())},
		"most_recent": &hcldec.AttrSpec{Name: "most_recent", Type: cty.Bool, Required: false},
	}
	return s
}

// FlatImageFilterOptions is an auto-generated flat version of ImageFilterOptions.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatImageFilterOptions struct {
	Name       *string           `mapstructure:"name" cty:"name" hcl:"name"`
	Owner      *string           `mapstructure:"owner" cty:"owner" hcl:"owner"`
	Tags       []string          `mapstructure:"tags" cty:"tags" hcl:"tags"`
	Visibility *string           `mapstructure:"visibility" cty:"visibility" hcl:"visibility"`
	Properties map[string]string `mapstructure:"properties" cty:"properties" hcl:"properties"`
}

// FlatMapstructure returns a new FlatImageFilterOptions.
// FlatImageFilterOptions is an auto-generated flat version of ImageFilterOptions.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*ImageFilterOptions) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatImageFilterOptions)
}

// HCL2Spec returns the hcl spec of a ImageFilterOptions.
// This spec is used by HCL to read the fields of ImageFilterOptions.
// The decoded values from this spec will then be applied to a FlatImageFilterOptions.
func (*FlatImageFilterOptions) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"name":       &hcldec.AttrSpec{Name: "name", Type: cty.String, Required: false},
		"owner":      &hcldec.AttrSpec{Name: "owner", Type: cty.String, Required: false},
		"tags":       &hcldec.AttrSpec{Name: "tags", Type: cty.List(cty.String), Required: false},
		"visibility": &hcldec.AttrSpec{Name: "visibility", Type: cty.String, Required: false},
		"properties": &hcldec.AttrSpec{Name: "properties", Type: cty.Map(cty.String), Required: false},
	}
	return s
}
