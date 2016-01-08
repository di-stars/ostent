// +build !bin
// http://blog.golang.org/profiling-go-programs

package cmd

import (
	"os"
	"runtime/pprof"
)

var (
	// ProfileHeapOutput is the flag value.
	ProfileHeapOutput string
	// ProfileCPUOutput is the flag value.
	ProfileCPUOutput string
)

// ProfileHeapRun is a prerun func for starting heap profile.
func ProfileHeapRun() error {
	if ProfileHeapOutput == "" {
		return nil
	}
	file, err := os.Create(ProfileHeapOutput)
	if err != nil {
		return err
	}
	PostRuns.Add(func() error {
		logger := NewLog(nil, "[ostent profile-heap] ")
		if err := pprof.Lookup("heap").WriteTo(file, 1); err != nil {
			logger.Print(err) // just print
		}
		if err := file.Close(); err != nil {
			logger.Print(err) // just print
		}
		return nil
	})
	return nil
}

// ProfileCPURun is a prerun func for starting CPU profile.
func ProfileCPURun() error {
	if ProfileCPUOutput == "" {
		return nil
	}
	file, err := os.Create(ProfileCPUOutput)
	if err != nil {
		return err
	}
	if err := pprof.StartCPUProfile(file); err != nil {
		return err
	}
	PostRuns.Add(func() error {
		logger := NewLog(nil, "[ostent profile-cpu] ")
		logger.Print("Writing CPU profile")
		pprof.StopCPUProfile()
		if err := file.Close(); err != nil {
			logger.Print(err) // just print
		}
		return nil
	})
	return nil
}

func init() {
	OstentCmd.PersistentFlags().StringVar(&ProfileHeapOutput, "profile-heap", "", "Profiling heap output `filename`")
	OstentCmd.PersistentFlags().StringVar(&ProfileCPUOutput, "profile-cpu", "", "Profiling CPU output `filename`")
	PreRuns.Adds(ProfileHeapRun, ProfileCPURun)
}
