package tree

import "fmt"

type No struct {
	chave int
	NoEsq *No
	NoDir *No
}

var arvore No

func MyTree() No {
	arvore.chave = 15
	arvore.NoDir = nil
	arvore.NoEsq = nil
	fmt.Printf("%+v\n", arvore)
	return arvore
}
func Insert(arvore No) int {
	var inserirArvore = &arvore
	if inserirArvore.chave > 0 {
		fmt.Println("A arvore tem raiz maior do que 0")
		return arvore.chave
	}
	return 0
}
