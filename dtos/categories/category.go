package categories

// CategoryDTO : Estrutura para receber novas categorias
type CategoryDTO struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty" valid:"required"`
}
