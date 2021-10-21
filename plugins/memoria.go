// arquivo: plugins/cpu.go
package plugins

import (
	"fmt"
	"errors"

	"github.com/raffs/go-labs/sistema"
)

type Memoria struct {
	So string
}

func (m *Memoria) Iniciar() {
	fmt.Println("Incializando o plugin de Memória", m.So)
}

func (m *Memoria) Coletar() (err error) {
	usoCpu := sistema.MemUso()

	if usoCpu < 0 {
		err = errors.New("Falha ao coletar a memória usada")
		return err
	}

	fmt.Printf("Memória: %f \n", usoCpu)
	return nil
}
func (m *Memoria) Describe() {
	fmt.Println(" Plugin usado para monitorar o uso da Memória!")
}