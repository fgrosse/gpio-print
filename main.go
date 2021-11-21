package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ahmetb/go-cursor"
	"github.com/fatih/color"
	"github.com/stianeikeland/go-rpio"
)

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Printf("Initialization failed: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		_ = rpio.Close()
	}()

	watch := flag.Bool("w", false, "watch mode")
	flag.Parse()

	for {
		states := readStates()
		printTable(states)
		if !*watch {
			return
		}

		time.Sleep(2*time.Second)
		fmt.Print(cursor.MoveUp(20))
	}
}

func readStates() [28]rpio.State {
	states := [28]rpio.State{}
	for i := range states {
		if i <= 1 {
			continue // first two GPIO pins are reserved
		}

		pin := rpio.Pin(i)
		states[i] = pin.Read()
	}

	return states
}

func printTable(states [28]rpio.State) {
	gpio := func(i int) string {
		label := fmt.Sprintf("GPIO%d", i)
		if i < 10 {
			label = " " + label
		}

		var c *color.Color
		if states[i] == rpio.High {
			c = color.New(color.FgGreen)
		} else {
			c = color.New(color.FgYellow)
		}
		return c.Sprint(label)
	}

	number := func(n int) string {
		s := fmt.Sprintf("(%d)", n)
		if n < 10 {
			s = " " + s
		}
		return s
	}

	rows := [][2]string{
		{blue("   3V3"), red("5V")},
		{gpio(2), red("5V")},
		{gpio(3), gray("GND")},
		{gpio(4), gpio(14)},
		{gray("   GND"), gpio(15)},
		{gpio(17), gpio(18)},
		{gpio(27), gray("GND")},
		{gpio(22), gpio(23)},
		{blue("   3V3"), gpio(24)},
		{gpio(10), gray("GND")},
		{gpio(9), gpio(25)},
		{gpio(11), gpio(8)},
		{gray("   GND"), gpio(7)},
		{gray(" ID_SD"), gray(" ID_SC")},
		{gpio(5), gray("GND")},
		{gpio(6), gpio(12)},
		{gpio(13), gray("GND")},
		{gpio(19), gpio(16)},
		{gpio(26), gpio(20)},
		{gray("   GND"), gpio(21)},
	}

	buf := new(bytes.Buffer)
	for i, row := range rows {
		left := row[0]
		right := row[1]
		numberLeft := number(i*2+1)
		numberRight := number(i*2+2)
		fmt.Fprintf(buf, "%s %s %s %s\n", left, numberLeft, numberRight, right)
	}

	fmt.Print(buf.String())
}

func blue(s string) string {
	c := color.New(color.FgBlue)
	return c.Sprint(s)
}

func red(s string) string {
	c := color.New(color.FgRed)
	return c.Sprint(s)
}

func gray(s string) string {
	c := color.New(color.FgHiBlack)
	return c.Sprint(s)
}
