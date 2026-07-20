package bean

type BeanMetadata struct {
	Fields map[string]*FieldMetadata
}

type FieldMetadata struct {
	DftValue string
}
