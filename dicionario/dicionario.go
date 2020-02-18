package dicionario

import "errors"

type Dictionary map[string]string

var ErrNaoEncontrado = errors.New("Não foi possível encontrar a palavra.")
var ErrPalavraExistente = errors.New("Palavra já existe.")

func (d Dictionary) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]

	if !existe {
		return "", ErrNaoEncontrado
	}

	return definicao, nil
}

func (d Dictionary) Add(palavra, definicao string) error {
	_, err := d.Busca(palavra)
	switch err {
	case ErrNaoEncontrado:
		d[palavra] = definicao
	case nil:
		return ErrPalavraExistente
	default:
		return err
	}

	return nil
}
