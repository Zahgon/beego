package alils

type MachineGroupAttribute struct {
	ExternalName string `json:"externalName"`
	TopicName    string `json:"groupTopic"`
}

type MachineGroup struct {
	Name          string   `json:"groupName"`
	Type          string   `json:"groupType"`
	MachineIDType string   `json:"machineIdentifyType"`
	MachineIDList []string `json:"machineList"`

	Attribute MachineGroupAttribute `json:"groupAttribute"`

	CreateTime     uint32
	LastModifyTime uint32

	project *LogProject
}

type Machine struct {
	IP            string
	UniqueID      string `json:"machine-uniqueid"`
	UserdefinedID string `json:"userdefined-id"`
}

type MachineList struct {
	Total    int
	Machines []*Machine
}

func (m *MachineGroup) ListMachines() (ms []*Machine, total int, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (m *MachineGroup) GetAppliedConfigs() (confNames []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
