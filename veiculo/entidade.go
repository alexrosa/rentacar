package veiculo

type Veiculo struct {
	ID     int    `json:"id"`
	Nome   string `json:"nome"`
	Marca  string `json:"marca"`
	Ano    int    `json:"ano"`
	Modelo int    `json:"modelo"`
}
