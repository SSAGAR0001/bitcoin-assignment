package db

type TransactionStruct struct {
	FromAddr        string `gorm:"type:varchar(100);NOT NULL"`
	ToAddr          string `gorm:"type:varchar(100);NOT NULL"`
	BlockNumber     string `gorm:"type:varchar(100);NOT NULL"`
	TransactionHash string `gorm:"type:varchar(100);PRIMARY_KEY"`
}
