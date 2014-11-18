package commands

import (
	"flag"
	"log"
	"net"
	"os"
	"time"

	"github.com/ostrost/ostent"
	"github.com/ostrost/ostent/client"
	"github.com/ostrost/ostent/types"
	metrics "github.com/rcrowley/go-metrics"
)

type graphite struct {
	logger      *loggerWriter
	RefreshFlag types.PeriodValue
	ServerAddr  types.BindValue
}

func graphiteCommandLine(cli *flag.FlagSet) commandLineHandler {
	gr := &graphite{
		logger: &loggerWriter{
			log.New(os.Stderr, "[ostent sendto-graphite] ", log.LstdFlags),
		},
		RefreshFlag: types.PeriodValue{Duration: types.Duration(10 * time.Second)}, // 10s default
		ServerAddr:  types.NewBindValue(2003),
	}
	cli.Var(&gr.RefreshFlag, "graphite-refresh", "Graphite refresh interval")
	cli.Var(&gr.ServerAddr, "sendto-graphite", "Graphite server address")
	return func() (atexitHandler, bool, error) {
		if gr.ServerAddr.Host == "" {
			return nil, false, nil
		}
		ostent.AddBackground(func(defaultPeriod types.PeriodValue) {
			/* if gr.RefreshFlag.Duration == 0 { // if .RefreshFlag had no default
				gr.RefreshFlag = defaultPeriod
			} */
			gc := &carbond{
				logger:     gr.logger,
				serveraddr: gr.ServerAddr.String(),
				Client:     client.DefaultClient(types.Duration(gr.RefreshFlag.Duration)),
			}
			ostent.Register <- gc
		})
		return nil, false, nil
	}
}

type carbond struct {
	logger        *loggerWriter
	serveraddr    string
	conn          net.Conn
	client.Client // expires() bool method
	failing       bool
}

func (_ *carbond) CloseChans() {} // intentionally empty
func (_ *carbond) Reload()     {} // intentionally empty

func (_ *carbond) Push(*ostent.IndexUpdate) {} // TODO?

func (cd *carbond) Tack() {
	addr, err := net.ResolveTCPAddr("tcp", cd.serveraddr)
	if err != nil {
		cd.logger.Printf("Resolve Addr %s: %s\n", cd.serveraddr, err)
		return
	}
	// go metrics.Graphite(ostent.Reg1s.Registry, 1*time.Second, "ostent", addr)
	err = metrics.GraphiteOnce(metrics.GraphiteConfig{
		DurationUnit: time.Nanosecond, // default, used(divided by thus must not be 0) with Timer metrics
		Addr:         addr,
		Registry:     ostent.Reg1s.Registry,
		Prefix:       "ostent",
	})
	if err != nil {
		if !cd.failing {
			cd.failing = true
			cd.logger.Printf("Sending: %s\n", err)
		}
	} else if cd.failing {
		cd.failing = false
		cd.logger.Printf("Recovered\n")
	}
}

func init() {
	AddCommandLine(graphiteCommandLine)
}
