package models

type Tag struct {
	Model
	Name string `json:"name"`
	CreateBy string `json:"create_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface {}) (tags []Tag)  {
	db.Where(maps).Order("id desc").Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int)  {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}
