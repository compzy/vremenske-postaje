package main

import (
	"context"
	"fmt"
	"os"

	"github.com/compzy/vremenske-postaje/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "vreme",
		Usage: "izpis meritev vremenskih postaj",
		Flags: []cli.Flag{
			&cli.FloatFlag{Name: "minTemp", Value: -50, Usage: "najnižja dovoljena temperatura meritve"},
			&cli.FloatFlag{Name: "maxTemp", Value: 50, Usage: "najvišja dovoljena temperatura meritve"},
			&cli.IntFlag{Name: "stMeritev", Value: 3, Usage: "najmanjše število meritev za statistiko"},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float("minTemp")
			maxTemp := cmd.Float("maxTemp")
			stMeritev := int(cmd.Int("stMeritev"))

			var postaje map[string]vreme.Postaja
			postaje = make(map[string]vreme.Postaja)
			vreme.DodajPostajo(postaje, "KP01", "Koper Luka", "(Koper)")
			vreme.DodajPostajo(postaje, "KR01", "Kredarica", "(Triglav)")
			vreme.DodajPostajo(postaje, "LJ01", "Ljubljana Bežigrad", "(Ljubljana)")
			vreme.DodajPostajo(postaje, "MB01", "Maribor Tabor", "(Maribor)")

			for _, t := range []float64{24.5, 27, 26.1, 28.4} {
				vreme.DodajMeritev(postaje, "KP01", t, minTemp, maxTemp)
			}
			for _, t := range []float64{2, -1.5, 3.5} {
				vreme.DodajMeritev(postaje, "KR01", t, minTemp, maxTemp)
			}
			for _, t := range []float64{18.2, 21.5, 19.8, 16.5} {
				vreme.DodajMeritev(postaje, "LJ01", t, minTemp, maxTemp)
			}
			vreme.DodajMeritev(postaje, "MB01", 12.4, minTemp, maxTemp)

			vreme.DodajMeritev(postaje, "IZ01", 34.1, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "MB01", -66, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "MB01", 10.2, minTemp, maxTemp)

			vreme.IzpisMeritev(postaje)
			vreme.IzpisPorocila(postaje, stMeritev)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
