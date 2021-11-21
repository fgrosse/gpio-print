<h1 align="center">GPIO Print for Raspberry Pi</h1>
<p align="center">A small command line tool to print the current status of GPIO pins.</p>
<p align="center">
   <a href="https://github.com/fgrosse/gpio-print/releases"><img src="https://img.shields.io/github/tag/fgrosse/gpio-print.svg?label=version&color=brightgreen"></a>
   <a href="https://github.com/fgrosse/gpio-print/actions/workflows/test.yml"><img src="https://github.com/fgrosse/gpio-print/actions/workflows/test.yml/badge.svg"></a>
   <a href="https://github.com/fgrosse/gpio-print/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-BSD--3--Clause-blue.svg"></a>
</p>

---

### Usage

You can run `gpio-print` without any arguments to print a colored overview of
the GPIO pins. The output is similar to the one at the bottom of the `pinout`
command.

For the GPIO pins, green coloring means the pin is set to HIGH, while yellow
coloring means it is set to LOW. 

![Screenshot](screenshot.png)

You can use the `-w` flag in order to continuously watch the pins.  

```bash
$ gpio-print -h
Usage of gpio-print:
  -w	watch mode
```
