package veiculo

//struct que representa a entidade do neg?cio (business object)
//representacao da tabela
type Veiculo struct {
	ID     int    `json:"id"`
	Nome   string `json:"nome"`
	Marca  string `json:"marca"`
	Ano    int    `json:"ano"`
	Modelo int    `json:"modelo"`
}
