package dicionario

type Dictionary map[string]string

func (d Dictionary) Busca(palavra string) string {
	return d[palavra]
}
