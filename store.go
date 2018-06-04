package main

import(
	"database/sql"
)

type Store interface {
	CreateStorage() error
	GetRegions() ([]*Region, error)
	GetCategorys() ([]*Category, error)
	GetExpenses() ([]*Expense, error)
	GetRecipients() ([]*Recipient, error)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateStorage() error {
	_, err := store.db.Query("CREATE TABLE IF NOT EXISTS region (id int(5) PRIMARY KEY NOT NULL, description VARCHAR(256)) ENGINE=InnoDB;")
	_, err =store.db.Query("CREATE TABLE IF NOT EXISTS category (id int(5) PRIMARY KEY NOT NULL, description VARCHAR(256)) ENGINE=InnoDB;")
	_, err = store.db.Query("CREATE TABLE IF NOT EXISTS recipient (id int(5) PRIMARY KEY NOT NULL, name VARCHAR(256)) ENGINE=InnoDB;")
	_, err = store.db.Query("CREATE TABLE IF NOT EXISTS expense (id int(9) PRIMARY KEY NOT NULL, description VARCHAR(256), amount DECIMAL(10,2) NOT NULL, date DATE, category_id int(5), region_id int(5), recipient_id int(5), CONSTRAINT `fk_expense_region` FOREIGN KEY (region_id) REFERENCES region(id), CONSTRAINT `fk_expense_category` FOREIGN KEY (category_id) REFERENCES category(id), CONSTRAINT `fk_expense_recipient` FOREIGN KEY (recipient_id) REFERENCES recipient(id)) ENGINE=InnoDB;")

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

func (store *dbStore) GetCategorys() ([]*Category, error) {
        rows, err := store.db.Query("SELECT description from category")
        if err != nil {
                return nil, err
        }
        defer rows.Close()
        categorys := []*Category{}
        for rows.Next() {
                category := &Category{}
                if err := rows.Scan(&category.Description); err != nil {
                        return nil, err
                }
                categorys = append(categorys, category)
        }
        return categorys, nil
}

func (store *dbStore) GetRecipients() ([]*Recipient, error) {
        rows, err := store.db.Query("SELECT name from recipient")
        if err != nil {
                return nil, err
        }
        defer rows.Close()
        recipients := []*Recipient{}
        for rows.Next() {
                recipient := &Recipient{}
                if err := rows.Scan(&recipient.Name); err != nil {
                        return nil, err
                }
                recipients = append(recipients, recipient)
        }
        return recipients, nil
}

func (store *dbStore) GetExpenses() ([]*Expense, error) {
        rows, err := store.db.Query("SELECT description, amount, date, category_id, region_id, recipient_id from expense")
        if err != nil {
                return nil, err
        }
        defer rows.Close()
        expenses := []*Expense{}
        for rows.Next() {
                expense := &Expense{}
                if err := rows.Scan(&expense.Description, &expense.Amount, &expense.Date, &expense.CategoryId, &expense.RegionId, &expense.RecipientId); err != nil {
                        return nil, err
                }
                expenses = append(expenses, expense)
        }
        return expenses, nil
}

var store Store

func InitStore(s Store) {
	store = s
}
