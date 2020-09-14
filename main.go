package main

import (
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/netlify/netlifyctl/commands"
)

func main() {
	
	if false {
		cpuProfile := "/tmp/netlifyctl-cpu.prof"
		cpuFile, err := os.Create(cpuProfile) ; if err != nil { panic (err) }
		defer cpuFile.Close ()
		if err := pprof.StartCPUProfile (cpuFile); err != nil { panic (err) }
		defer pprof.StopCPUProfile ()
	}
	
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:    true,
		DisableTimestamp: false,
		TimestampFormat:  time.RFC822Z,
	})
	commands.Execute()
	
	if false {
		memProfile := "/tmp/netlifyctl-mem.prof"
		memFile, err := os.Create(memProfile) ; if err != nil { panic (err) }
		defer memFile.Close ()
		runtime.GC ()
		if err := pprof.WriteHeapProfile (memFile); err != nil { panic (err) }
	}
}
