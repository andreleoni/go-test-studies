package dicionario

import "errors"

type Dictionary map[string]string

func (d Dictionary) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]

	if !existe {
		return "", errors.New("Não foi possível encontrar a palavra.")
	}

	return definicao, nil
}
