package commands

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/prashantwosti/ticker_cli/ticker"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type market string

const (
	US market = ""
	AU market = "AX"
)

var rootCmd = &cobra.Command{
	Use:   "ticker_cli",
	Short: "commandline ticker",
	Long: "\nThis is a commandline ticker.\n" +
		"Search using [Exchange] followed by the [Ticker] symbol." +
		"\n\ne.g $ ticker_cli us tsla   - \"us\" is the exchange and \"tsla\" is the symbol." +
		"\n\nhttps://github.com/prashantwosti",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows version name.",
	Long:  "Prints the current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1.0.0-WTF!")
	},
}

var asxCmd = &cobra.Command{
	Use:     "asx",
	Aliases: []string{"ASX"},
	Short:   "Shows ASX listed companies.",
	Run: func(cmd *cobra.Command, args []string) {
		getTickerResult(args, AU)
	},
	Args: cobra.ExactArgs(1),
}

var usCmd = &cobra.Command{
	Use:     "us",
	Aliases: []string{"US"},
	Short:   "Shows NASDAQ and NYSE listed companies.",
	Run: func(cmd *cobra.Command, args []string) {
		getTickerResult(args, US)
	},
	Args: cobra.ExactArgs(1),
}

func InitCommandline() {
	var commands = rootCmd
	commands.AddCommand(versionCmd, asxCmd, usCmd)
	commands.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
	if err := commands.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func getTickerResult(args []string, market market) {
	if len(args) > 0 {
		symbol := args[0]
		if market == AU {
			symbol += "." + string(AU)
		}

		drawLine := strings.Repeat("$€£¥", 6)

		tickerInfo, err := ticker.NewTicker().Get(symbol)

		if len(tickerInfo) > 0 {
			color.Green.Println("\n\n" + drawLine)
			fmt.Println(tickerInfo)
			color.Green.Println(drawLine + "\n\n")
		} else {
			color.Red.Println(err)
		}
	}
}
