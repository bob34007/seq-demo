package cmd

import (
	"errors"
	"fmt"
	"sync"

	"github.com/seqDemo/sqll"
	"github.com/seqDemo/util"
	"github.com/spf13/cobra"
)

func getMaxSeq(cfg *util.Config) (int64, error) {
	db, err := sqll.ConnectDB(cfg.DSN)
	if err != nil {
		return 0, err
	}
	defer sqll.CloseConn(db)
	rows, err := sqll.QueryWithResult(db, util.GetMaxSQL)
	rows.Close()
	m, err := sqll.ParseResult(rows)
	if err != nil {
		return 0, err
	}
	return m, nil
}

func getSeq(cfg *util.Config, wg *sync.WaitGroup, ch1 chan error, ch2 chan int32) {
	//conn DB
	var count int32 = 0
	defer wg.Done()
	db, err := sqll.ConnectDB(cfg.DSN)
	if err != nil {
		ch1 <- err
		ch2 <- count
		return
	}
	defer sqll.CloseConn(db)
	for count < cfg.Counters {
		if len(ch1) > 0 {
			ch2 <- count
			return
		}
		//start transaction
		err = sqll.QueryWithNoResult(db, "begin;")
		if err != nil {
			ch1 <- err
			ch2 <- count
			sqll.QueryWithNoResult(db, "rollback;")
			return
		}
		//select for update
		rows, err := sqll.QueryWithResult(db, util.GetMaxSQL)
		if err != nil {
			ch1 <- err
			ch2 <- count
			sqll.QueryWithNoResult(db, "rollback;")
			return
		}
		rows.Close()
		sql := util.GeneralUpdateMax(cfg.CacheNum)
		//update
		err = sqll.QueryWithNoResult(db, sql)
		if err != nil {
			ch1 <- err
			ch2 <- count
			sqll.QueryWithNoResult(db, "rollback;")
			return
		}
		//commit
		err = sqll.QueryWithNoResult(db, "commit;")
		if err != nil {
			ch1 <- err
			ch2 <- count
			sqll.QueryWithNoResult(db, "rollback;")
			return
		}
		count += 1
	}
	ch2 <- count
	return
}

func getruncount(ch chan int32) int64 {
	var res int64
	for a := range ch {
		res += int64(a)
	}
	return res
}

func NewSeqCommand() *cobra.Command {
	//Replay sql from pcap files，and compare reslut from pcap file and
	//replay server
	var cfg *util.Config
	cmd := &cobra.Command{
		Use:   "seq",
		Short: "general run ",
		RunE: func(cmd *cobra.Command, args []string) error {
			maxseqbegin, err := getMaxSeq(cfg)
			if err != nil {
				return err
			}
			var wg sync.WaitGroup
			wg.Add(int(cfg.Threads))
			ch1 := make(chan error, cfg.Threads)
			ch2 := make(chan int32, cfg.Threads)
			defer close(ch1)
			defer close(ch2)
			var i int32 = 0
			for ; i < cfg.Threads; i++ {
				go getSeq(cfg, &wg, ch1, ch2)
			}
			wg.Wait()
			if len(ch1) > 0 {
				return <-ch1
			}
			a := getruncount(ch2)
			maxseqend, err := getMaxSeq(cfg)
			if err != nil {
				return err
			}
			if maxseqend-maxseqbegin != a*int64(cfg.CacheNum) {
				return errors.New(fmt.Sprintf("check error , runcount:%v,"+
					"maxseqbegin:%v,maxseqend:%v,cachenum:%v",
					a, maxseqbegin, maxseqend, cfg.CacheNum))
			}
			fmt.Println("check success ")
			return nil

		},
	}

	cfg.ParseFlag(cmd.Flags())
	return cmd
}

func NewGeneralCommand() *cobra.Command {
	//add sub command general
	cmd := &cobra.Command{
		Use:   "general",
		Short: "general run ",
	}
	fmt.Println("in general")
	cmd.AddCommand(NewSeqCommand())
	return cmd
}
