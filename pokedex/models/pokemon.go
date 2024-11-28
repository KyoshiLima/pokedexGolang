package models

type Pokemon struct {
	ID	int	`json:"id"`
	Nome	string	`json:"nome"`
	Tipo	string	`json:"tipos"`
	Tamanho	float64	`json:"tamanho"`
	Peso	float64	`json:"peso"`
	Descricao	string	`json:"descrição"`
	Rotas	int	`json:"rotas"`
}