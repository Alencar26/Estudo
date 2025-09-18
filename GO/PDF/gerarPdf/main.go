package main

import (
	"fmt"
	"os"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	buildHeading(m)

	m.RegisterFooter(func() {
		m.Row(10, func() {
			m.Col(5, func() {
				m.Text("TESTE TESTE TESTE")
			})
		})
	})

	// buildHeading(m)

	if err := m.OutputFileAndClose("pdfs/novo-pdf.pdf"); err != nil {
		fmt.Println("Arquivo não pode ser salvo:", err)
		os.Exit(1)
	}

	fmt.Println("Arquivo salvo com sucesso!")

}

func buildHeading(m pdf.Maroto) {
	//registerHeader define um conjunto de linha, tabelas, lista.
	//cada chamada dessa função cria uma nova página para cada conjunto.
	m.RegisterHeader(func() {
		m.SetPageMargins(10, 20, 10)
		m.Row(50, func() {
			m.Col(12, func() {
				if err := m.FileImage("images/logo.png", props.Rect{
					Center:  true,
					Percent: 75,
				}); err != nil {
					fmt.Println("Image file was not loaded:", err)
				}
			})
		})

		m.Row(10, func() {
			m.Col(12, func() {
				m.Text("Texto escrito por mim", props.Text{
					Top:   3,
					Style: consts.Bold,
					Align: consts.Center,
					Color: getDarkPurpleColor(),
				})
			})
		})
	})
}

func getDarkPurpleColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}
