package tree

import "fmt"

type No struct {
	chave int
	NoEsq *No
	NoDir *No
}

func MyTree() No {
	var arvore No
	arvore.chave = 15
	arvore.NoDir = nil
	arvore.NoEsq = nil
	fmt.Printf("%+v\n", arvore)
	return arvore
}
func Search(arvore No, valor int) bool {
	var inserirArvore = &arvore
	if inserirArvore == nil {
		return false
	}
	if inserirArvore.chave == valor {
		fmt.Print("Valor encontrado!: ", inserirArvore.chave)
		return true
	} else {
		fmt.Printf("O valor não foi encontrado")
		return false
	}
}
