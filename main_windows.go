package main

import "github.com/gonutz/w32/v2"

func main() {
	w32.SendMessage(w32.GetConsoleWindow(), w32.WM_CLOSE, 0, 0)
}
