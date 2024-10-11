package model

type Question struct {
	ID           int    `gorm:"primaryKey;autoIncrement"`
	LevelID      int    `gorm:"index"`
	QuestionText string `gorm:"type:varchar(255)"`
	Level        Level  `gorm:"foreignKey:LevelID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
