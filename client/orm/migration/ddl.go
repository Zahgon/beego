package migration

type Index struct {
	Name string
}

type Unique struct {
	Definition string
	Columns    []*Column
}

type Column struct {
	Name     string
	Inc      string
	Null     string
	Default  string
	Unsign   string
	DataType string
	remove   bool
	Modify   bool
}

type Foreign struct {
	ForeignTable  string
	ForeignColumn string
	OnDelete      string
	OnUpdate      string
	Column
}

type RenameColumn struct {
	OldName     string
	OldNull     string
	OldDefault  string
	OldUnsign   string
	OldDataType string
	NewName     string
	Column
}

func (m *Migration) CreateTable(tablename, engine, charset string, p ...func()) {
	_ = "STUB: not implemented"
	return
}

func (m *Migration) AlterTable(tablename string) { _ = "STUB: not implemented"; return }

func (m *Migration) NewCol(name string) *Column { _ = "STUB: not implemented"; return nil }

func (m *Migration) PriCol(name string) *Column { _ = "STUB: not implemented"; return nil }

func (m *Migration) UniCol(uni, name string) *Column { _ = "STUB: not implemented"; return nil }

func (m *Migration) ForeignCol(colname, foreigncol, foreigntable string) (foreign *Foreign) {
	_ = "STUB: not implemented"
	return nil
}

func (foreign *Foreign) SetOnDelete(del string) *Foreign { _ = "STUB: not implemented"; return nil }

func (foreign *Foreign) SetOnUpdate(update string) *Foreign { _ = "STUB: not implemented"; return nil }

func (c *Column) Remove() { _ = "STUB: not implemented"; return }

func (c *Column) SetAuto(inc bool) *Column { _ = "STUB: not implemented"; return nil }

func (c *Column) SetNullable(null bool) *Column { _ = "STUB: not implemented"; return nil }

func (c *Column) SetDefault(def string) *Column { _ = "STUB: not implemented"; return nil }

func (c *Column) SetUnsigned(unsign bool) *Column { _ = "STUB: not implemented"; return nil }

func (c *Column) SetDataType(dataType string) *Column { _ = "STUB: not implemented"; return nil }

func (c *RenameColumn) SetOldNullable(null bool) *RenameColumn {
	_ = "STUB: not implemented"
	return nil
}

func (c *RenameColumn) SetOldDefault(def string) *RenameColumn {
	_ = "STUB: not implemented"
	return nil
}

func (c *RenameColumn) SetOldUnsigned(unsign bool) *RenameColumn {
	_ = "STUB: not implemented"
	return nil
}

func (c *RenameColumn) SetOldDataType(dataType string) *RenameColumn {
	_ = "STUB: not implemented"
	return nil
}

func (c *Column) SetPrimary(m *Migration) *Column { _ = "STUB: not implemented"; return nil }

func (unique *Unique) AddColumnsToUnique(columns ...*Column) *Unique {
	_ = "STUB: not implemented"
	return nil
}

func (m *Migration) AddColumns(columns ...*Column) *Migration {
	_ = "STUB: not implemented"
	return nil
}

func (m *Migration) AddPrimary(primary *Column) *Migration { _ = "STUB: not implemented"; return nil }

func (m *Migration) AddUnique(unique *Unique) *Migration { _ = "STUB: not implemented"; return nil }

func (m *Migration) AddForeign(foreign *Foreign) *Migration { _ = "STUB: not implemented"; return nil }

func (m *Migration) AddIndex(index *Index) *Migration { _ = "STUB: not implemented"; return nil }

func (m *Migration) RenameColumn(from, to string) *RenameColumn {
	_ = "STUB: not implemented"
	return nil
}

func (m *Migration) GetSQL() (sql string) { _ = "STUB: not implemented"; return "" }
