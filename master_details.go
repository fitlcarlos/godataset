package godata

type MasterSource struct {
	DataSource   *DataSet
	DetailFields []string
	MasterFields []string
}

func NewMasterSource() *MasterSource {
	ms := &MasterSource{}
	return ms
}
func (ms *MasterSource) AddMasterSource(dataSet *DataSet) *MasterSource {
	ms.DataSource = dataSet
	return ms
}

func (ms *MasterSource) AddDetailFields(fields ...string) *MasterSource {
	ms.DetailFields = fields
	return ms
}

func (ms *MasterSource) AddMasterFields(fields ...string) *MasterSource {
	ms.MasterFields = fields
	return ms
}

func (ms *MasterSource) ClearMasterFields() *MasterSource {
	ms.MasterFields = nil
	return ms
}

func (ms *MasterSource) ClearDetailFields() *MasterSource {
	ms.DetailFields = nil
	return ms
}

func (ms *MasterSource) And() *DataSet {
	return ms.DataSource
}
