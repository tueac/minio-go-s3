package initsystem

import (
	"flag"
	"fmt"
	"os"
)

func Help() {
	helpFlag := flag.Bool("h", false, "帮助信息")
	versionFlag := flag.Bool("v", false, "版本信息")

	flag.Parse()

	if *helpFlag {
		printHelp()
		os.Exit(0)
	}

	if *versionFlag {
		showVersion()
		os.Exit(0)
	}

	if flag.NArg() == 0 {
		printHelp()
		os.Exit(0)
	}
}

func printHelp() {
	fmt.Println("https://missyou.love\n")
	fmt.Println("usage: [-v] [bucket] [object]")
	fmt.Println("\nAvailable Commands:")
	fmt.Println(" <subcommand>")
	fmt.Println("  bucket create     Create bucket")
	fmt.Println("  bucket delete     Delete bucket,but bucket is empty")
	fmt.Println("  bucker exist      Check if bucket exists")
	fmt.Println("  bucker list       Print bucket name")
	fmt.Println("  object delete     Delete object files")
	fmt.Println("  object delall     Purge bucket,delete all objects")
	fmt.Println("  object download   Download object files")
	fmt.Println("  object get        Get(download) bucket files")
	fmt.Println("  object put        Put object files")
	fmt.Println("  object putdit     Put object directory")
	fmt.Println("\nFlags:")
	flag.PrintDefaults()
}

func showVersion() {
	fmt.Println("version: 0.1")
}
