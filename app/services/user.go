package services

import (
	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	models "github.com/go-dyn/app/models"
)

type IUser interface {
	IBase
	List() ([]models.User, error)
	Paginate(page int, limit int) ([]models.User, int64, error)
	Create() (models.User, error)
	Get(Id uint64) (models.User, error)
	GetAuthUser(ctx http.Context) (models.User, error)
}

type UserImpl struct {
	BaseImpl
	query orm.Query
}

func NewUserImpl() *UserImpl {
	return &UserImpl{
		query: facades.Orm().Query(),
	}
}

func (r *UserImpl) Create() (models.User, error) {
	user := models.User{
		Name: "name",
	}
	err := r.query.Create(&user)
	return user, err
}

func (r *UserImpl) Get(Id uint64) (models.User, error) {
	var item models.User
	err := r.query.FindOrFail(&item, Id)
	return item, err
}

func (r *UserImpl) List() ([]models.User, error) {
	var list []models.User
	err := r.query.Find(&list)
	return list, err
}

func (r *UserImpl) Paginate(page int, limit int) ([]models.User, int64, error) {
	var list []models.User
	var total int64
	err := r.query.Paginate(page, limit, &list, &total)
	return list, total, err
}

func (r *UserImpl) GetAuthUser(ctx http.Context) (models.User, error) {
	var user models.User
	err := facades.Auth().User(ctx, &user)
	return user, err
}
