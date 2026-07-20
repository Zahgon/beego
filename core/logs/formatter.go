package logs

var formatterMap = make(map[string]LogFormatter, 4)

type LogFormatter interface {
	Format(lm *LogMsg) string
}

type PatternLogFormatter struct {
	Pattern    string
	WhenFormat string
}

func (p *PatternLogFormatter) getWhenFormatter() string { _ = "STUB: not implemented"; return "" }

func (p *PatternLogFormatter) Format(lm *LogMsg) string { _ = "STUB: not implemented"; return "" }

func RegisterFormatter(name string, fmtr LogFormatter) { _ = "STUB: not implemented"; return }

func GetFormatter(name string) (LogFormatter, bool) {
	_ = "STUB: not implemented"
	return *new(LogFormatter), false
}

func (p *PatternLogFormatter) ToString(lm *LogMsg) string { _ = "STUB: not implemented"; return "" }
