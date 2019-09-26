package veiculo

//define a struct que representa a entidade do negócio (business object)
type Veiculo struct {
	ID     int    `json:"id"`
	Nome   string `json:"nome"`
	Marca  string `json:"marca"`
	Ano    int    `json:"ano"`
	Modelo int    `json:"modelo"`
}
