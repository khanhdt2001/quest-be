package model

type Option struct {
	ID          int      `gorm:"primaryKey;autoIncrement"`
	QuestionID  int      `gorm:"index"`
	OptionText  string   `gorm:"type:varchar(255)"`
	OptionLabel string   `gorm:"type:varchar(255)"`
	Question    Question `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
