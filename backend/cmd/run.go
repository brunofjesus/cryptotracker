package cmd

import (
	"cryptotracker/repository/filesystem"
	"cryptotracker/repository/filesystem/persistence"
	"cryptotracker/rest/router"
	"cryptotracker/service"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
	"runtime"
	"time"
)

var FilePath string
var Host string
var Port int
var OpenBrowser bool

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the Crypto Tracker web app",
	Run: func(cmd *cobra.Command, args []string) {
		trackerRepository, err := filesystem.NewFsRepository(
			persistence.NewXmlPersistence(FilePath),
		)

		if err != nil {
			panic(err)
		}

		svc := service.NewWalletService(trackerRepository)

		time.AfterFunc(1*time.Second, func() {
			openBrowser(fmt.Sprintf("http://%s:%d", Host, Port))
		})

		err = router.Start(svc, Host, Port)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	runCmd.Flags().StringVarP(&FilePath, "file", "f", "cryptotracker.xml", "Transaction file to use")
	runCmd.Flags().StringVarP(&Host, "host", "i", "127.0.0.1", "Host to serve the app")
	runCmd.Flags().IntVarP(&Port, "port", "p", 40042, "Port to use")
	runCmd.Flags().BoolVarP(&OpenBrowser, "open-browser", "b", true, "Open your web browser")
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("cannot open browser, unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
