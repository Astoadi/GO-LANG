package models

import (
	"cars/config"
	"fmt"
)

var Cars = make(map[int64]Car)

type Car struct {
	ID    int64   `json:"id" db:"id"` //`json:"id" gorm:"primaryKey"` in case of gorm
	Name  string  `json:"name" db:"name"`
	Model string  `json:"model" db:"model"`
	Brand string  `json:"brand" db:"brand"`
	Year  int64   `json:"year" db:"year"`
	Price float64 `json:"price" db:"price"`
}

func (c *Car) Insert() {
	//sqlx
	query := `INSERT INTO cars (name, model, brand, year, price ) VALUES (:name, :model, :brand, :year, :price) RETURNING id`
	row, err := config.DB.NamedQuery(query, c)
	if err != nil {
		fmt.Printf("error inserting car:%v", err)
	}
	defer row.Close()
	if row.Next() {
		if err := row.Scan(&c.ID); err != nil {
			fmt.Printf("error inseriting car:%v", err)
		}
	}

	//gorm
	// if err := config.DB.Create(&c).Error; err != nil {
	// 	fmt.Printf("error inserting car %v", err)
	// }
	//native psql code
	// query := `INSERT INTO cars (name, model, brand, year, price ) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	// if err := config.DB.QueryRow(query, c.Name, c.Model, c.Brand, c.Year, c.Price).Scan(&c.ID); err != nil {
	// 	error := fmt.Errorf("error inserting car %v", err)
	// 	fmt.Println(error)
	// }
}

func (c *Car) Get() error {

	query := `SELECT name, model, brand, year, price FROM cars WHERE id=$1`
	if err := config.DB.Get(c, query, c.ID); err != nil {
		fmt.Printf("error getting car %v", err)
		return err
	}
	//gorm
	// if err := config.DB.First(c, c.ID).Error; err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		fmt.Printf("error getting car %v", err)
	// 		return err
	// 	}

	// }
	//native psql code
	// query := `SELECT name, model, brand, year, price FROM cars WHERE id=$1`
	// if err := config.DB.QueryRow(query, c.ID).Scan(&c.Name, &c.Model, &c.Brand, &c.Year, &c.Price); err != nil {
	// 	if err == sql.ErrNoRows {
	// 		error := fmt.Errorf("error getting car %v", err)
	// 		fmt.Println(error)
	// 		return err
	// 	}
	// }
	return nil
}

func (c *Car) Update() {
	//sqlx
	query := `UPDATE cars SET name = :name, model = :model, brand = :brand, year = :year , price = :price WHERE id = :id`
	if _, err := config.DB.NamedQuery(query, c); err != nil {
		fmt.Printf("error updating car: %v\n", err)
	}
	//gorm
	// if err := config.DB.Save(c).Error; err != nil {
	// 	fmt.Printf("error updating car %v", err)
	// }
	//native psql code
	// query := `UPDATE cars SET name = $1, model = $2, brand = $3, year = $4 , price = $5 WHERE id = $6`
	// _, err := config.DB.Exec(query, c.Name, c.Model, c.Brand, c.Year, c.Price, c.ID)
	// if err != nil {
	// 	error := fmt.Errorf("error updating car %v", err)
	// 	fmt.Println(error)

	// }
}
func (c *Car) Delete() {
	//sqlx
	query := `DELETE FROM cars WHERE id=$1`
	if _, err := config.DB.Exec(query, c.ID); err != nil {
		fmt.Printf("Error deleting car with id %v ,error:%v\n", c.ID, err)
	}
	//gorm
	// if err := config.DB.Delete(c).Error; err != nil {
	// 	fmt.Printf("error updating car %v", err)
	// }
	//native psql code
	// query := `DELETE FROM cars WHERE id=$1`
	// _, err := config.DB.Exec(query, c.ID)
	// if err != nil {
	// 	fmt.Printf("Error deleting car with id %v ,error:%v\n", c.ID, err)
	// }
}
