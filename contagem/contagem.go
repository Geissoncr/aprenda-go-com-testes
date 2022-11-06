package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const ultimaPalavra = "Vai!"
const inicioContagem = 3

type Sleeper interface {
	Sleep()
}

type SleeperPadrao struct{}

func (d *SleeperPadrao) Sleep() {
	time.Sleep(1 * time.Second)
}

func Contagem(escritor io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(escritor, i)
	}

	for i := inicioContagem; i > 0; i-- {
		fmt.Fprintln(escritor, i)
	}

	sleeper.Sleep()
	fmt.Fprint(escritor, ultimaPalavra)
}

func main() {
	sleeper := &SleeperPadrao{}
	Contagem(os.Stdout, sleeper)
}
