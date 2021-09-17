// Go byte units and localized formatting
// 17-Sep-2021
// https://devtidbits.com/2021/09/17/go-byte-units-and-localized-formatting/
package main

import (
	"fmt"
	"math"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type (
	Decimal float64
	Binary  float64
)

const (
	kb Decimal = 1e+03
	mb Decimal = 1e+06
	gb Decimal = 1e+09
	tb Decimal = 1e+12
	pb Decimal = 1e+15
	eb Decimal = 1e+18
)

const (
	kib Binary = 1 << 10
	mib Binary = 1 << 20
	gib Binary = 1 << 30
	tib Binary = 1 << 40
	pib Binary = 1 << 50
	eib Binary = 1 << 60
)

const (
	precision0 = "%.0f"
	precision1 = "%.1f\u00A0%s"
	precision2 = "%.2f\u00A0%s"
)

func main() {
	const n = 10000000000
	en := language.English
	de := language.German
	fr := language.French
	fmt.Println("🇬🇧", Decimal(n).String(en))
	fmt.Println("🇬🇧", Binary(n).String(en))
	fmt.Println("🇩🇪", Decimal(n).String(de))
	fmt.Println("🇩🇪", Binary(n).String(de))
	fmt.Println("🇫🇷", Decimal(n).String(fr))
	fmt.Println("🇫🇷", Binary(n).String(fr))
}

func (n Decimal) String(t language.Tag) string {
	p := message.NewPrinter(t)
	f := n
	x := Decimal(math.Abs(float64(n)))
	switch {
	case x >= eb:
		f /= eb
		return p.Sprintf(precision2, f, "EB")
	case x >= pb:
		f /= pb
		return p.Sprintf(precision2, f, "PB")
	case x >= tb:
		f /= tb
		return p.Sprintf(precision2, f, "TB")
	case x >= gb:
		f /= gb
		return p.Sprintf(precision2, f, "GB")
	case x >= mb:
		f /= mb
		return p.Sprintf(precision2, f, "MB")
	case x >= kb:
		f /= kb
		return p.Sprintf(precision1, f, "kB")
	default:
		return p.Sprintf(precision0, f)
	}
}

func (n Binary) String(t language.Tag) string {
	p := message.NewPrinter(t)
	f := n
	x := Binary(math.Abs(float64(n)))
	switch {
	case x >= eib:
		f /= eib
		return p.Sprintf(precision2, f, "EiB")
	case x >= pib:
		f /= pib
		return p.Sprintf(precision2, f, "PiB")
	case x >= tib:
		f /= tib
		return p.Sprintf(precision2, f, "TiB")
	case x >= gib:
		f /= gib
		return p.Sprintf(precision2, f, "GiB")
	case x >= mib:
		f /= mib
		return p.Sprintf(precision2, f, "MiB")
	case x >= kib:
		f /= kib
		return p.Sprintf(precision1, f, "KiB")
	default:
		return p.Sprintf(precision0, f)
	}
}
