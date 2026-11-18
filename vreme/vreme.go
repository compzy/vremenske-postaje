// Package vreme omogoča delo z vremenskimi postajami in njihovimi meritvami.
package vreme

import (
	"fmt"
	"math"
)

// Postaja predstavlja vremensko postajo z imenom, krajem in meritvami.
// Novo postajo ustvarite s funkcijo DodajPostajo.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo v slovar postaje doda novo postajo z dano oznako, imenom in krajem.
func DodajPostajo(postaje map[string]Postaja, oznaka string, ime string, kraj string) {
	postaje[oznaka] = Postaja{ime, kraj, []float64{}}
}

// DodajMeritev postaji z dano oznako doda meritev temperature.
// Meritve izven intervala [minTemp, maxTemp] in meritve neobstoječih postaj zavrne.
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64, minTemp float64, maxTemp float64) {
	p, ok := postaje[oznaka]
	if temperatura > maxTemp || temperatura < minTemp {
		fmt.Println("Temperatura", temperatura, "°C je izven dovoljenega območja")
	} else if !ok {
		fmt.Println("Postaja", oznaka, "ne obstaja")
	} else {
		p.meritve = append(p.meritve, temperatura)
		postaje[oznaka] = p
	}

	return
}

func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	p, ok := postaje[oznaka]
	if !ok {
		return
	}
	najnizja = math.Inf(1)
	najvisja = math.Inf(-1)
	d := len(p.meritve)
	if d < stMeritev {
		ok = false
	}
	var sum float64
	for _, x := range p.meritve {
		sum += x
		if x > najvisja {
			najvisja = x
		}
		if x < najnizja {
			najnizja = x
		}
	}
	povprecna = sum / float64(d)
	return
}

// IzpisMeritev izpiše vse postaje in njihove meritve.
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	for i, x := range postaje {
		fmt.Println(i, "-", x.ime, x.kraj+":", x.meritve)
	}
	return
}

// IzpisPorocila za vsako postajo izpiše najnižjo, najvišjo in povprečno temperaturo.
// Za postaje z manj kot stMeritev meritvami izpiše "Premalo meritev".
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	for i, x := range postaje {
		najnizja, najvisja, povprecna, ok := statistika(postaje, i, stMeritev)
		switch {
		case povprecna >= float64(25) && ok:
			fmt.Println(x.ime+":", "min", najnizja, "°C,", "max", najvisja, "°C,", "povprečje", povprecna, "°C -> Vroče")
		case povprecna >= float64(15) && ok:
			fmt.Println(x.ime+":", "min", najnizja, "°C,", "max", najvisja, "°C,", "povprečje", povprecna, "°C -> Toplo")
		case povprecna >= float64(5) && ok:
			fmt.Println(x.ime+":", "min", najnizja, "°C,", "max", najvisja, "°C,", "povprečje", povprecna, "°C -> Hladno")
		case povprecna < float64(5) && ok:
			fmt.Println(x.ime+":", "min", najnizja, "°C,", "max", najvisja, "°C,", "povprečje", povprecna, "°C -> Mrzlo")
		case !ok:
			fmt.Println(x.ime+":", "Premalo meritev")
		default:
			fmt.Println("BUG")
		}
	}
}
