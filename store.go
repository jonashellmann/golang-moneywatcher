package main

import(
	"database/sql"
	"crypto/sha256"
	"encoding/hex"
)

type Store interface {
	CreateStorage() error

	GetRegions(userId int) ([]*Region, error)
	GetRegion(userId int, regionId int) (Region, error)
	GetCategorys(userId int) ([]*Category, error)
	GetCategory(userId int, categoryId int) (Category, error)
	GetExpenses(userId int) ([]*Expense, error)
	GetExpense(userId int, expenseId int) (Expense, error)
	GetRecipients(userId int) ([]*Recipient, error)
	GetRecipient(userId int, recipientId int) (Recipient, error)

	CreateRegion(region *Region) error
	CreateCategory(category *Category) error
	CreateRecipient(recipient *Recipient) error
	CreateExpense(expense *Expense) error

	DeleteExpense(userId int, expenseId int) error

	CheckCredentials(username string, password string) error
	GetUserId(username string) (int, error)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateStorage() error {
	_, err := store.db.Query("CREATE TABLE IF NOT EXISTS user (id int(5) PRIMARY KEY NOT NULL AUTO_INCREMENT, username VARCHAR(256) NOT NULL, email VARCHAR(256) NOT NULL, password VARCHAR(256) NOT NULL) ENGINE=InnoDB CHARACTER SET 'utf8' COLLATE 'utf8_bin';")
	_, err = store.db.Query("CREATE TABLE IF NOT EXISTS region (id int(5) PRIMARY KEY NOT NULL AUTO_INCREMENT, description VARCHAR(256) NOT NULL, user_id int(5) NOT NULL, CONSTRAINT `fk_region_user` FOREIGN KEY (user_id) REFERENCES user(id)) ENGINE=InnoDB CHARACTER SET 'utf8' COLLATE 'utf8_bin';")
	_, err = store.db.Query("CREATE TABLE IF NOT EXISTS category (id int(5) PRIMARY KEY NOT NULL AUTO_INCREMENT, description VARCHAR(256) NOT NULL, user_id int(5) NOT NULL, CONSTRAINT `fk_category_user` FOREIGN KEY (user_id) REFERENCES user(id)) ENGINE=InnoDB CHARACTER SET 'utf8' COLLATE 'utf8_bin';")
	_, err = store.db.Query("CREATE TABLE IF NOT EXISTS recipient (id int(5) PRIMARY KEY NOT NULL AUTO_INCREMENT, name VARCHAR(256) NOT NULL, user_id int(5) NOT NULL, CONSTRAINT `fk_recipient_user` FOREIGN KEY (user_id) REFERENCES user(id)) ENGINE=InnoDB CHARACTER SET 'utf8' COLLATE 'utf8_bin';")
	_, err = store.db.Query("CREATE TABLE IF NOT EXISTS expense (id int(9) PRIMARY KEY NOT NULL AUTO_INCREMENT, description VARCHAR(256), amount DECIMAL(10,2) NOT NULL, date DATE, category_id int(5), region_id int(5), source_id int(5), destination_id int(5), user_id int(5) NOT NULL, CONSTRAINT `fk_expense_region` FOREIGN KEY (region_id) REFERENCES region(id), CONSTRAINT `fk_expense_category` FOREIGN KEY (category_id) REFERENCES category(id), CONSTRAINT `fk_expense_source` FOREIGN KEY (source_id) REFERENCES recipient(id), CONSTRAINT `fk_expense_destination` FOREIGN KEY (destination_id) REFERENCES recipient(id), CONSTRAINT `fk_expense_user` FOREIGN KEY (user_id) REFERENCES user(id)) ENGINE=InnoDB CHARACTER SET 'utf8' COLLATE 'utf8_bin';")

	return err
}

func (store *dbStore) GetRegions(userId int) ([]*Region, error) {
	rows, err := store.db.Query("SELECT id, description FROM region WHERE user_id = ? ORDER BY description ASC, id DESC", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	regions := []*Region{}
	for rows.Next() {
		region := &Region{}
		if err := rows.Scan(&region.Id, &region.Description); err != nil {
			return nil, err
		}
		regions = append(regions, region)
	}
	return regions, nil
}

func (store *dbStore) GetRegion(userId int, regionId int) (Region, error) {
	region := Region{}
	err := store.db.QueryRow("SELECT id, description FROM region WHERE user_id = ? AND id = ?", userId, regionId).Scan(&region.Id, &region.Description)

	if err != nil {
		return region, err
	}

	return region, nil
}

func (store *dbStore) GetCategorys(userId int) ([]*Category, error) {
	rows, err := store.db.Query("SELECT id, description FROM category WHERE user_id = ? ORDER BY description ASC, id DESC", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categorys := []*Category{}
	for rows.Next() {
		category := &Category{}
		if err := rows.Scan(&category.Id, &category.Description); err != nil {
				return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}

func (store *dbStore) GetCategory(userId int, categoryId int) (Category, error) {
	category := Category{}
	err := store.db.QueryRow("SELECT id, description FROM category WHERE user_id = ? AND id = ?", userId, categoryId).Scan(&category.Id, &category.Description)

	if err != nil {
		return category, err
	}

	return category, nil
}

func (store *dbStore) GetRecipients(userId int) ([]*Recipient, error) {
	rows, err := store.db.Query("SELECT id, name FROM recipient WHERE user_id = ? ORDER BY name ASC, id DESC", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	recipients := []*Recipient{}
	for rows.Next() {
		recipient := &Recipient{}
		if err := rows.Scan(&recipient.Id, &recipient.Name); err != nil {
				return nil, err
		}
		recipients = append(recipients, recipient)
	}
	return recipients, nil
}

func (store *dbStore) GetRecipient(userId int, recipientId int) (Recipient, error) {
	recipient := Recipient{}
	err := store.db.QueryRow("SELECT id, name FROM recipient WHERE user_id = ? AND id = ?", userId, recipientId).Scan(&recipient.Id, &recipient.Name)

	if err != nil {
		return recipient, err
	}

	return recipient, nil
}

func (store *dbStore) GetExpenses(userId int) ([]*Expense, error) {
	rows, err := store.db.Query("SELECT id, description, amount, date, category_id, region_id, source_id, destination_id FROM expense WHERE user_id = ? ORDER BY date DESC, id DESC", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	expenses := []*Expense{}
	for rows.Next() {
		expense := &Expense{}
		regionId := sql.NullInt64{}
		categoryId := sql.NullInt64{}
		sourceId := sql.NullInt64{}
		destinationId := sql.NullInt64{}
		if err := rows.Scan(&expense.Id, &expense.Description, &expense.Amount, &expense.Date, &categoryId, &regionId, &sourceId, &destinationId); err != nil {
				return nil, err
		}

		if categoryId.Valid {
			category, err := store.GetCategory(userId, int(categoryId.Int64))
			if err != nil {
				return expenses, err
			}
			expense.Category = category
		}
		if regionId.Valid {
                        region, err := store.GetRegion(userId, int(regionId.Int64))
                        if err != nil {
                                return expenses, err
                        }
                        expense.Region = region
                }
		if sourceId.Valid {
                        source, err := store.GetRecipient(userId, int(sourceId.Int64))
                        if err != nil {
                                return expenses, err
                        }
                        expense.Source = source
                }
		if destinationId.Valid {
                        destination, err := store.GetRecipient(userId, int(destinationId.Int64))
                        if err != nil {
                                return expenses, err
                        }
                        expense.Destination = destination
                }

		expenses = append(expenses, expense)
	}
	return expenses, nil
}

func (store *dbStore) GetExpense(userId int, expenseId int) (Expense, error) {
	expense := Expense{}
	regionId := sql.NullInt64{}
        categoryId := sql.NullInt64{}
	sourceId := sql.NullInt64{}
        destinationId := sql.NullInt64{}

	err := store.db.QueryRow("SELECT id, description, amount, date, category_id, region_id, source_id, destination_id FROM expense WHERE user_id = ? AND id = ?", userId, expenseId).Scan(&expense.Id, &expense.Description, &expense.Amount, &expense.Date, &categoryId, &regionId, &sourceId, &destinationId)

	if categoryId.Valid {
		category, err := store.GetCategory(userId, int(categoryId.Int64))
                if err != nil {
			return expense, err
                }
                expense.Category = category
        }
        if regionId.Valid {
                region, err := store.GetRegion(userId, int(regionId.Int64))
                if err != nil {
                        return expense, err
                }
                expense.Region = region
        }
	if sourceId.Valid {
		source, err := store.GetRecipient(userId, int(sourceId.Int64))
                if err != nil {
                        return expense, err
                }
                expense.Source = source
        }
        if destinationId.Valid {
                destination, err := store.GetRecipient(userId, int(destinationId.Int64))
                if err != nil {
                        return expense, err
                }
                expense.Destination = destination
        }

	if err != nil {
		return expense, err
	}

	return expense, nil
}

func (store *dbStore) CreateRegion(region *Region) error {
	_, err := store.db.Query("INSERT INTO region (description, user_id) VALUES (?, ?)", region.Description, region.UserId)
	return err
}

func (store *dbStore) CreateCategory(category *Category) error {
	_, err := store.db.Query("INSERT INTO category (description, user_id) VALUES (?, ?)", category.Description, category.UserId)
	return err
}

func (store *dbStore) CreateRecipient(recipient *Recipient) error {
	_, err := store.db.Query("INSERT INTO recipient (name, user_id) VALUES (?, ?)", recipient.Name, recipient.UserId)
	return err
}

func (store *dbStore) CreateExpense(expense *Expense) error {
	var categoryId sql.NullInt64
	var regionId sql.NullInt64
	var sourceId sql.NullInt64
	var destinationId sql.NullInt64

	if expense.Category.Id == 0 {
		categoryId = sql.NullInt64{Valid: false}
	} else {
		categoryId = sql.NullInt64{Int64: int64(expense.Category.Id), Valid: true}
	}
	if expense.Source.Id == 0 {
                sourceId = sql.NullInt64{Valid: false}
        } else {
                sourceId = sql.NullInt64{Int64: int64(expense.Source.Id), Valid: true}
        }
	if expense.Destination.Id == 0 {
		destinationId = sql.NullInt64{Valid: false}
	} else {
		destinationId = sql.NullInt64{Int64: int64(expense.Destination.Id), Valid: true}
	}
	if expense.Region.Id == 0 {
                regionId = sql.NullInt64{Valid: false}
        } else {
                regionId = sql.NullInt64{Int64: int64(expense.Region.Id), Valid: true}
        }

	_, err := store.db.Query("INSERT INTO expense (description, amount, date, category_id, region_id, source_id, destination_id, user_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", expense.Description, expense.Amount, expense.Date, categoryId, regionId, sourceId, destinationId, expense.UserId)
	return err
}

func (store *dbStore) DeleteExpense(userId int, expenseId int) error {
	_, err := store.db.Query("DELETE FROM expense WHERE id = ? AND user_id = ?", expenseId, userId)

	if err != nil {
		return err
	}

	return nil
}

func (store *dbStore) CheckCredentials(username string, password string) error {
	hash := sha256.New()
	hash.Write([]byte(password))
	passwordHash := hex.EncodeToString(hash.Sum(nil))
	user := User{}
	err := store.db.QueryRow("SELECT username FROM user WHERE username = ? AND password = ?", username, passwordHash).Scan(&user.Username)

	if err != nil {
		return err
	}

	return nil
}

func (store *dbStore) GetUserId(username string) (int, error) {
	user := User{}
	err := store.db.QueryRow("SELECT id FROM user WHERE username = ?", username).Scan(&user.Id)

	if err != nil {
		return 0, err
	}

	return user.Id, nil
}

var store Store

func InitStore(s Store) {
	store = s
}
