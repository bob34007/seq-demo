package util

import (
	"github.com/spf13/pflag"
)

const MAX_INT int32 = 1<<31 - 1

type Config struct {
	Threads  int32
	DSN      string
	CacheNum int32
	MaxVal   int32
	IsCycle  bool
	Counters int32
}

func (cfg *Config) ParseFlag(flags *pflag.FlagSet) {
	var isCycle string
	cfg.Threads = *(flags.Int32P("threads", "t", 10, "run threads "))
	flags.StringVarP(&cfg.DSN, "dsn", "d", "", "database server dsn")
	cfg.CacheNum = *(flags.Int32P("cache", "n", 10, "cache seq num"))
	cfg.MaxVal = *(flags.Int32P("maxval", "m", MAX_INT, "cache seq num"))
	flags.StringVarP(&isCycle, "iscycle", "c", "cycle", "seq can cycle")
	if isCycle == "cycle" {
		cfg.IsCycle = true
	} else {
		cfg.IsCycle = false
	}
	cfg.Threads = *(flags.Int32P("counter", "ct", 100000, " get count "))
}
