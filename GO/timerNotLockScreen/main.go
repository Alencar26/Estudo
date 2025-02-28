package main

import (
	"time"

	"github.com/micmonay/keybd_event"
)

func main() {
	timerHoras(9)
}

func timerHoras(t int) {
	t = t * (60 * 60) //convertendo para segundos
	tempo := 0
	for tempo <= t {
		if tempo%60 == 0 {
			pressNumLock(keybd_event.VK_NUMLOCK)
			println("Press", time.Now().Local().String())
		}
		time.Sleep(1000 * time.Millisecond)
		tempo++
	}
}

// Para funcionar precisa dar permissão para acesso ao input do teclado.
// Comando temporário: (vão ao padrão após reiniciar)
// sudo chmod 0666 /dev/uinput
func pressNumLock(tecla int) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		println("Erro ao criar evendo no teclado: ", err.Error())
	}

	//precionar tecla
	kb.SetKeys(tecla)

	//precionar a tecla
	err = kb.Press()
	if err != nil {
		println("Erro ao precionar a tecla: ", err)
	}

	//esperando
	time.Sleep(100 * time.Millisecond)

	//liberar a tecla
	err = kb.Release()
	if err != nil {
		println("Erro ao liberar a tecla: ", err)
	}
}
