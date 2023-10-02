package services

import (
	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"

	models "github.com/go-dyn/app/models"
)

type IExample interface {
	IBase
	List() ([]models.Example, error)
	Paginate(page int, limit int) ([]models.Example, int64, error)
	Get(id uint64) (models.Example, error)
	Create(models.Example) (models.Example, error)
	Update(models.Example) (models.Example, error)
	Destroy(item models.Example) (models.Example, *orm.Result, error)
}

type ExampleImpl struct {
	BaseImpl
	query orm.Query
}

func NewExampleImpl() *ExampleImpl {
	return &ExampleImpl{
		query: facades.Orm().Query(),
	}
}

func (r *ExampleImpl) List() ([]models.Example, error) {
	var list []models.Example
	err := r.query.Find(&list)
	return list, err
}

func (r *ExampleImpl) Paginate(page int, limit int) (list []models.Example, total int64, err error) {
	err = r.query.Paginate(page, limit, &list, &total)
	return
}

func (r *ExampleImpl) Get(id uint64) (item models.Example, err error) {
	err = r.query.FindOrFail(id, &item)
	return
}

func (r *ExampleImpl) Create(item models.Example) (models.Example, error) {
	if err := r.query.Create(item); err != nil {
		return item, err
	}

	return item, nil
}

func (r *ExampleImpl) Update(item models.Example) (models.Example, error) {
	err := r.query.Save(item)
	return item, err
}

func (r *ExampleImpl) Destroy(item models.Example) (models.Example, *orm.Result, error) {
	result, err := r.query.Delete(item)
	return item, result, err
}
