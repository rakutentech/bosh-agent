package fakes

import (
	boshsettings "bosh/settings"
	boshsys "bosh/system"
)

type FakeSettingsServiceProvider struct {
	NewServiceFs              boshsys.FileSystem
	NewServiceDir             string
	NewServiceFetcher         boshsettings.SettingsFetcher
	NewServiceSettingsService *FakeSettingsService
}

func NewServiceProvider() *FakeSettingsServiceProvider {
	return &FakeSettingsServiceProvider{
		NewServiceSettingsService: &FakeSettingsService{},
	}
}

func (provider *FakeSettingsServiceProvider) NewService(
	fs boshsys.FileSystem,
	dir string,
	fetcher boshsettings.SettingsFetcher,
) boshsettings.Service {
	provider.NewServiceFs = fs
	provider.NewServiceDir = dir
	provider.NewServiceFetcher = fetcher
	return provider.NewServiceSettingsService
}

type FakeSettingsService struct {
	FetchSettingsError  error
	SettingsWereFetched bool

	ForceNextLoadToFetchSettingsError   error
	SettingsWereForcedToFetchOnNextLoad bool

	Settings boshsettings.Settings

	Blobstore boshsettings.Blobstore
	AgentId   string
	Vm        boshsettings.Vm
	MbusUrl   string
	Disks     boshsettings.Disks
	DefaultIp string
	Ips       []string
}

func (service *FakeSettingsService) ForceNextLoadToFetchSettings() error {
	service.SettingsWereForcedToFetchOnNextLoad = true
	return service.ForceNextLoadToFetchSettingsError
}

func (service *FakeSettingsService) FetchSettings() error {
	service.SettingsWereFetched = true
	return service.FetchSettingsError
}

func (service FakeSettingsService) GetSettings() boshsettings.Settings {
	return service.Settings
}

func (service FakeSettingsService) GetBlobstore() boshsettings.Blobstore {
	return service.Blobstore
}

func (service FakeSettingsService) GetAgentId() string {
	return service.AgentId
}

func (service FakeSettingsService) GetVm() boshsettings.Vm {
	return service.Vm
}

func (service FakeSettingsService) GetMbusUrl() string {
	return service.MbusUrl
}

func (service FakeSettingsService) GetDisks() boshsettings.Disks {
	return service.Disks
}

func (service FakeSettingsService) GetDefaultIp() (string, bool) {
	return service.DefaultIp, service.DefaultIp != ""
}

func (service FakeSettingsService) GetIps() []string {
	return service.Ips
}
