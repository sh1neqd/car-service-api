package repositories

import (
	"car-service1/app/internal/domain/car"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type CarRepo struct {
	db *sqlx.DB
}

func (r CarRepo) Update(id int, dto car.UpdateCarDTO) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if dto.RegNum != nil {
		setValues = append(setValues, fmt.Sprintf("reg_num=$%d", argId))
		args = append(args, *dto.RegNum)
		argId++
	}

	if dto.Mark != nil {
		setValues = append(setValues, fmt.Sprintf("mark=$%d", argId))
		args = append(args, *dto.Mark)
		argId++
	}

	if dto.Model != nil {
		setValues = append(setValues, fmt.Sprintf("model=$%d", argId))
		args = append(args, *dto.Model)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	var err error
	q := fmt.Sprintf("UPDATE public.car SET %s WHERE id=$%d", setQuery, argId)
	logrus.Println(q)
	args = append(args, id)
	_, err = r.db.Exec(q, args...)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (r CarRepo) GetAll() ([]car.Car, error) {
	var cars []car.Car
	q := `SELECT * FROM public.car`
	logrus.Println(q)
	err := r.db.Select(&cars, q)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	return cars, err
}

func (r CarRepo) GetById(id int) (car.Car, error) {
	var c car.Car
	q := `SELECT id, reg_num, mark, model FROM public.car WHERE id = $1`
	logrus.Println(q)
	row := r.db.QueryRow(q, id)
	err := row.Scan(&c.ID, &c.RegNum, &c.Mark, &c.Model)
	if err != nil {
		return car.Car{}, err
	}
	return c, err
}

func (r CarRepo) Create(car car.CarDTO) error {
	var id int
	q := `INSERT INTO public.car (reg_num, mark, model) VALUES ($1, $2, $3) RETURNING id`
	logrus.Println(q)
	row := r.db.QueryRow(q, car.RegNum, car.Mark, car.Model)
	if err := row.Scan(&id); err != nil {
		return err
	} else {
		logrus.Infof("car added with id: %v", id)
	}
	return nil
}

func (r CarRepo) Delete(id int) error {
	q := `DELETE FROM public.car WHERE id=$1`
	logrus.Println(q)
	_, err := r.db.Exec(q, id)
	if err != nil {
		return err
	}
	return nil
}

func NewCarRepo(db *sqlx.DB) *CarRepo {
	return &CarRepo{db: db}
}
