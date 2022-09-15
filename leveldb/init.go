package leveldb

import leveldb1 "github.com/jeffcail/leveldb"

var (
	LevelDb *leveldb1.LevelDB
	err     error
)

func init() {
	LevelDb, err = leveldb1.CreateLevelDB("./level_view_data")
	if err != nil {
		panic(err)
	}
}
