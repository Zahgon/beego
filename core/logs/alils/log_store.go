package alils

type LogStore struct {
	Name       string `json:"logstoreName"`
	TTL        int
	ShardCount int

	CreateTime     uint32
	LastModifyTime uint32

	project *LogProject
}

type Shard struct {
	ShardID int `json:"shardID"`
}

func (s *LogStore) ListShards() (shardIDs []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *LogStore) PutLogs(lg *LogGroup) (err error) { _ = "STUB: not implemented"; return nil }

func (s *LogStore) GetCursor(shardID int, from string) (cursor string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *LogStore) GetLogsBytes(shardID int, cursor string,
	logGroupMaxCount int) (out []byte, nextCursor string, err error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func LogsBytesDecode(data []byte) (gl *LogGroupList, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *LogStore) GetLogs(shardID int, cursor string,
	logGroupMaxCount int) (gl *LogGroupList, nextCursor string, err error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}
