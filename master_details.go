package godataset

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

func (ms *MasterSource) Clear() {
	ms.DataSource = nil
	ms.ClearMasterFields()
	ms.ClearDetailFields()
}

func (ms *MasterSource) CountMasterFields() int {
	return len(ms.MasterFields)
}

func (ms *MasterSource) CountDetailFields() int {
	return len(ms.DetailFields)
}

func (ms *MasterSource) ClearMasterFields() {
	ClearSlice(ms.MasterFields)
	ms.MasterFields = nil
}

func (ms *MasterSource) ClearDetailFields() {
	ClearSlice(ms.DetailFields)
	ms.DetailFields = nil
}

func (ms *MasterSource) And() *DataSet {
	return ms.DataSource
}
