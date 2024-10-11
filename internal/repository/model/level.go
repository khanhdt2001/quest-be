package model

type Level struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	CategoryID  int `gorm:"index"`
	LevelNumber int
	Category    Category `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
