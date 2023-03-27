package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
	"math/rand"
	"os"
	"time"
)

func convertCharacter(value string) string {
	codonMapping := map[string][]string{
		"I": {"ATA", "ATC", "ATT"},
		"M": {"ATG"},
		"T": {"ACA", "ACC", "ACG", "ACT"},
		"N": {"AAC", "AAT"},
		"K": {"AAA", "AAG"},
		"S": {"AGC", "AGT", "TCA", "TCC", "TCG", "TCT"},
		"R": {"AGA", "AGG", "CGA", "CGC", "CGG", "CGT"},
		"L": {"CTA", "CTC", "CTG", "CTT", "TTA", "TTG"},
		"P": {"CCA", "CCC", "CCG", "CCT"},
		"H": {"CAC", "CAT"},
		"Q": {"CAA", "CAG"},
		"V": {"GTA", "GTC", "GTG", "GTT"},
		"A": {"GCA", "GCC", "GCG", "GCT"},
		"D": {"GAC", "GAT"},
		"E": {"GAA", "GAG"},
		"G": {"GGA", "GGC", "GGG", "GGT"},
		"F": {"TTC", "TTT"},
		"Y": {"TAC", "TAT"},
		"":  {"TAA", "TAG", "TGA"},
		"C": {"TGC", "TGT"},
		"W": {"TGG"},
	}

	rand.Seed(time.Now().UnixNano())
	list := codonMapping[value]
	if list == nil {
		fmt.Println("Invalid value")
		os.Exit(1)
	}
	randomIndex := rand.Intn(len(list))
	return list[randomIndex]
}

func convertString(value string) string {
	result := ""
	for i := 0; i < len(value); i++ {
		result += convertCharacter(string(value[i]))
	}
	return result
}

func getStringFromClipboard() string {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	value := clipboard.Read(clipboard.FmtText)
	return string(value)
}

func main() {
	rootCmd := &cobra.Command{
		Use:  "dna [AA sequence]",
		Args: cobra.RangeArgs(0, 1),
		Run: func(cmd *cobra.Command, args []string) {
			fromClipboard, err := cmd.PersistentFlags().GetBool("from-clipboard")
			if err != nil {
				panic(err)
			}

			value := ""
			if fromClipboard {
				value = getStringFromClipboard()
			} else if len(args) > 0 {
				value = args[0]
			} else {
				fmt.Println(cmd.UsageString())
				os.Exit(1)
			}

			fmt.Println(convertString(value))
		},
	}

	rootCmd.PersistentFlags().BoolP("from-clipboard", "c", false, "get input from clipboard")

	_ = rootCmd.Execute()
}
