package dicionario

const (
	ErrNaoEncontrado    = ErrDicionario("não foi possível encontrar a palavra que você procura")
	ErrPalavraExistente = ErrDicionario("não é possível adicionar a palavra pois ela já existe")
)

type ErrDicionario string

func (e ErrDicionario) Error() string {
	return string(e)
}
