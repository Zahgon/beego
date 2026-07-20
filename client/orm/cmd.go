package orm

type commander interface {
	Parse([]string)
	Run() error
}

var commands = make(map[string]commander)

func printHelp(errs ...string) { _ = "STUB: not implemented"; return }

func RunCommand() { _ = "STUB: not implemented"; return }

type commandSyncDb struct {
	al        *alias
	force     bool
	verbose   bool
	noInfo    bool
	rtOnError bool
}

func (d *commandSyncDb) Parse(args []string) { _ = "STUB: not implemented"; return }

func (d *commandSyncDb) Run() error { _ = "STUB: not implemented"; return nil }

type commandSQLAll struct {
	al *alias
}

func (d *commandSQLAll) Parse(args []string) { _ = "STUB: not implemented"; return }

func (d *commandSQLAll) Run() error { _ = "STUB: not implemented"; return nil }

func init() {
	commands["syncdb"] = new(commandSyncDb)
	commands["sqlall"] = new(commandSQLAll)
}

func RunSyncdb(name string, force bool, verbose bool) error { _ = "STUB: not implemented"; return nil }
