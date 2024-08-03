package packet

type Column struct {
	name         string
	databaseType string
}

func NewColumn(name string, databaseType string) Column {
	return Column{
		name:         name,
		databaseType: databaseType,
	}
}

func (c Column) Name() string {
	return c.name
}

func (c Column) DatabaseTypeName() string {
	return c.databaseType
}