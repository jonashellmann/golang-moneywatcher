package main

import(
	"database/sql"
)

type Store interface {
	CreateStorage() error
	GetRegions() ([]*Region, error)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateStorage() error {
	_, err := store.db.Query("CREATE TABLE IF NOT EXISTS region (id int(5) PRIMARY KEY NOT NULL, description VARCHAR(256));")
	_, err = store.db.Query("CREATE TABLE IF NOT EXISTS category (id int(5) PRIMARY KEY NOT NULL, description VARCHAR(256);)")
	_, err = store.db.Query("CREATE TABLE IF NOT EXISTS recipient (id int(5) PRIMARY KEY NOT NULL, name VARCHAR(256);)")
	_, err = store.db.Query("CREATE TABLE IF NOT EXISTS expense (id int(9) PRIMARY KEY NOT NULL, description VARCHR(256), amount DECIMAL(10,2), date DATE, category_id int(5), region_id int(5), recipient_id int(5), FOREIGN KEY (region_id) REFERENCES region(id), FOREIGN KEY (category_id) REFERENCES category(id), FOREIGN KEY (recipient_id) REFERENCES recipient(id));")

	return err
}

func (store *dbStore) GetRegions() ([]*Region, error) {
	rows, err := store.db.Query("SELECT description from region")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	regions := []*Region{}
	for rows.Next() {
		region := &Region{}

		if err := rows.Scan(&region.Description); err != nil {
			return nil, err
		}

		regions = append(regions, region)
	}

	return regions, nil
}

var store Store

func InitStore(s Store) {
	store = s
}
