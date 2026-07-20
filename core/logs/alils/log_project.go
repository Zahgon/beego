package alils

import (
	"net/http"
)

type errorMessage struct {
	Code    string `json:"errorCode"`
	Message string `json:"errorMessage"`
}

type LogProject struct {
	Name            string
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
}

func NewLogProject(name, endpoint, AccessKeyID, accessKeySecret string) (p *LogProject, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func handleResponse(r *http.Response, actionDesc string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *LogProject) sendRequest(method, uri string, headers map[string]string, body []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createStandardHeaders(bodyLen int) map[string]string { _ = "STUB: not implemented"; return nil }

func (p *LogProject) ListLogStore() (storeNames []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *LogProject) GetLogStore(name string) (s *LogStore, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *LogProject) CreateLogStore(name string, ttl, shardCnt int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *LogProject) DeleteLogStore(name string) (err error) { _ = "STUB: not implemented"; return nil }

func (p *LogProject) UpdateLogStore(name string, ttl, shardCnt int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *LogProject) ListMachineGroup(offset, size int) (m []string, total int, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p *LogProject) GetMachineGroup(name string) (m *MachineGroup, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *LogProject) CreateMachineGroup(m *MachineGroup) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *LogProject) UpdateMachineGroup(m *MachineGroup) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *LogProject) DeleteMachineGroup(name string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *LogProject) ListConfig(offset, size int) (cfgNames []string, total int, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p *LogProject) GetConfig(name string) (c *LogConfig, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *LogProject) UpdateConfig(c *LogConfig) (err error) { _ = "STUB: not implemented"; return nil }

func (p *LogProject) CreateConfig(c *LogConfig) (err error) { _ = "STUB: not implemented"; return nil }

func (p *LogProject) DeleteConfig(name string) (err error) { _ = "STUB: not implemented"; return nil }

func (p *LogProject) GetAppliedMachineGroups(confName string) (groupNames []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *LogProject) GetAppliedConfigs(groupName string) (confNames []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *LogProject) ApplyConfigToMachineGroup(confName, groupName string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *LogProject) RemoveConfigFromMachineGroup(confName, groupName string) (err error) {
	_ = "STUB: not implemented"
	return nil
}
