package model

type UserAnswer struct {
	ID               int      `gorm:"primaryKey;autoIncrement"`
	UserID           int      `gorm:"index"`
	QuestionID       int      `gorm:"index"`
	SelectedOptionID int      `gorm:"index"`
	User             User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Question         Question `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	SelectedOption   Option   `gorm:"foreignKey:SelectedOptionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
