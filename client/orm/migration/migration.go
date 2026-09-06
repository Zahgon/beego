package migration

const (
	DateFormat   = "20060102_150405"
	DBDateFormat = "2006-01-02 15:04:05"
)

type Migrationer interface {
	Up()
	Down()
	Reset()
	Exec(name, status string) error
	GetCreated() int64
}

type Migration struct {
	sqls           []string
	Created        string
	TableName      string
	Engine         string
	Charset        string
	ModifyType     string
	Columns        []*Column
	Indexes        []*Index
	Primary        []*Column
	Uniques        []*Unique
	Foreigns       []*Foreign
	Renames        []*RenameColumn
	RemoveColumns  []*Column
	RemoveIndexes  []*Index
	RemoveUniques  []*Unique
	RemoveForeigns []*Foreign
}

var migrationMap map[string]Migrationer

func init() {
	migrationMap = make(map[string]Migrationer)
}

func (m *Migration) Up() { _ = "STUB: not implemented"; return }

func (m *Migration) Down() { _ = "STUB: not implemented"; return }

func (m *Migration) Migrate(migrationType string) { _ = "STUB: not implemented"; return }

func (m *Migration) SQL(sql string) { _ = "STUB: not implemented"; return }

func (m *Migration) Reset() { _ = "STUB: not implemented"; return }

func (m *Migration) Exec(name, status string) error { _ = "STUB: not implemented"; return nil }

func (m *Migration) addOrUpdateRecord(name, status string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Migration) GetCreated() int64 { _ = "STUB: not implemented"; return 0 }

func Register(name string, m Migrationer) error { _ = "STUB: not implemented"; return nil }

func Upgrade(lasttime int64) error { _ = "STUB: not implemented"; return nil }

func Rollback(name string) error { _ = "STUB: not implemented"; return nil }

func Reset() error { _ = "STUB: not implemented"; return nil }

func Refresh() error { _ = "STUB: not implemented"; return nil }

type dataSlice []data

type data struct {
	created int64
	name    string
	m       Migrationer
}

func (d dataSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (d dataSlice) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (d dataSlice) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func sortMap(m map[string]Migrationer) dataSlice { _ = "STUB: not implemented"; return *new(dataSlice) }

func isRollBack(name string) bool { _ = "STUB: not implemented"; return false }

func getAllMigrations() (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }
