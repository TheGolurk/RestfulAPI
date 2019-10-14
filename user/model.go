package user

// Model es la estructura para el usuario
type Model struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int32  `json:"Age"`
}

// Storage es un mapa del modelo
type Storage map[string]*Model

var storage Storage

// Create crea un nuevo usuario
func (s Storage) Create(m *Model) *Model {
	return m
}
