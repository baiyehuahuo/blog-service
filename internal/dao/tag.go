package dao

import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
)

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	return model.Tag{Name: name, State: state}.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{ID: id, ModifiedBy: modifiedBy},
	}
	values := map[string]interface{}{
		"name":       name,
		"state":      state,
		"modifiedBy": modifiedBy,
	}
	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{
		Model: &model.Model{ID: id},
	}
	return tag.Delete(d.engine)
}
