package schema

type HelloWorld struct {
	Lang    string `gorm:"type:varchar(2);not null;primaryKey"`
	Message string `gorm:"type:VARCHAR(40);not null"`
}

