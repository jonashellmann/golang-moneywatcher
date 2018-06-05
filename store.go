package main

import(
	"database/sql"
	"crypto/sha256"
)

type Store interface {
	CreateStorage() error

	GetRegions(userId int) ([]*Region, error)
	GetCategorys(userId int) ([]*Category, error)
	GetExpenses(userId int) ([]*Expense, error)
	GetRecipients(userId int) ([]*Recipient, error)

	CreateRegion(region *Region) error
	CreateCategory(category *Category) error
	CreateRecipient(recipient *Recipient) error
	CreateExpense(expense *Expense) error

	CheckCredentials(username string, password string) bool
	GetUserId(username string) (int, error)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateStorage() error {
	_, err := store.db.Query("CREATE TABLE IF NOT EXISTS user (id int(5) PRIMARY KEY NOT NULL AUTO_INCREMENT, username VARCHAR(256) NOT NULL, email VARCHAR(256) NOT NULL, password VARCHAR(256) NOT NULL) ENGINE=InnoDB")
	_, err = store.db.Query("CREATE TABLE IF NOT EXISTS region (id int(5) PRIMARY KEY NOT NULL AUTO_INCREMENT, description VARCHAR(256) NOT NULL, user_id int(5) NOT NULL, CONSTRAINT `fk_region_user` FOREIGN KEY (user_id) REFERENCES user(id)) ENGINE=InnoDB;")
	_, err = store.db.Query("CREATE TABLE IF NOT EXISTS category (id int(5) PRIMARY KEY NOT NULL AUTO_INCREMENT, description VARCHAR(256) NOT NULL, user_id int(5) NOT NULL, CONSTRAINT `fk_category_user` FOREIGN KEY (user_id) REFERENCES user(id)) ENGINE=InnoDB;")
	_, err = store.db.Query("CREATE TABLE IF NOT EXISTS recipient (id int(5) PRIMARY KEY NOT NULL AUTO_INCREMENT, name VARCHAR(256) NOT NULL, user_id int(5) NOT NULL, CONSTRAINT `fk_recipient_user` FOREIGN KEY (user_id) REFERENCES user(id)) ENGINE=InnoDB;")
	_, err = store.db.Query("CREATE TABLE IF NOT EXISTS expense (id int(9) PRIMARY KEY NOT NULL AUTO_INCREMENT, description VARCHAR(256), amount DECIMAL(10,2) NOT NULL, date DATE, category_id int(5), region_id int(5), recipient_id int(5), user_id int(5), CONSTRAINT `fk_expense_region` FOREIGN KEY (region_id) REFERENCES region(id), CONSTRAINT `fk_expense_category` FOREIGN KEY (category_id) REFERENCES category(id), CONSTRAINT `fk_expense_recipient` FOREIGN KEY (recipient_id) REFERENCES recipient(id), CONSTRAINT `fk_expense_user` FOREIGN KEY (user_id) REFERENCES user(id)) ENGINE=InnoDB;")

	return err
}

func (store *dbStore) GetRegions(userId int) ([]*Region, error) {
	rows, err := store.db.Query("SELECT description FROM region WHERE user_id = ?", userId)
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

func (store *dbStore) GetCategorys(userId int) ([]*Category, error) {
        rows, err := store.db.Query("SELECT description FROM category WHERE user_id = ?", userId)
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

func (store *dbStore) GetRecipients(userId int) ([]*Recipient, error) {
        rows, err := store.db.Query("SELECT name FROM recipient WHERE user_id = ?", userId)
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

func (store *dbStore) GetExpenses(userId int) ([]*Expense, error) {
        rows, err := store.db.Query("SELECT description, amount, date, category_id, region_id, recipient_id FROM expense WHERE user_id = ?", userId)
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

func (store *dbStore) CreateRegion(region *Region) error {
	_, err := store.db.Query("INSERT INTO region (description) VALUES (?)", region.Description)
	return err
}

func (store *dbStore) CreateCategory(category *Category) error {
	_, err := store.db.Query("INSERT INTO category (description) VALUES (?)", category.Description)
	return err
}

func (store *dbStore) CreateRecipient(recipient *Recipient) error {
	_, err := store.db.Query("INSERT INTO recipient (name) VALUES (?)", recipient.Name)
	return err
}

func (store *dbStore) CreateExpense(expense *Expense) error {
	_, err := store.db.Query("INSERT INTO expense (description, amount, date, category_id, region_id, recipient_id) VALUES (?, ?, ?, ?, ?, ?)", expense.Description, expense.Amount, expense.Date, expense.CategoryId, expense.RegionId, expense.RecipientId)
	return err
}

func (store *dbStore) CheckCredentials(username string, password string) bool {
	hash := sha256.New()
	hash.Write([]byte(password))
	user := User{}
	err := store.db.QueryRow("SELECT username FROM users WHERE username = ? AND password = ?", username, hash.Sum(nil)).Scan(user.Username)

	if err != nil {
		return false
	}

	return true
}

func (store *dbStore) GetUserId(username string) (int, error) {
	user := User{}
	err := store.db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(user.Id)

	if err != nil {
		return 0, err
	}

	return user.Id, nil
}

var store Store

func InitStore(s Store) {
	store = s
}
