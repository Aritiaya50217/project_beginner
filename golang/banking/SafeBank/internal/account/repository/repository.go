package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/Aritiaya50217/project_beginner/internal/account/model"
)

// AccountRepository interface กำหนด method สำหรัยจัดการข้อมูลบัญชีและธุรกรรม
type AccountRepository interface {
	Create(account model.Account) (*model.Account, error)
	FindByID(id int) (*model.Account, error)
	Update(account model.Account) error

	CreateTransaction(tx model.Transaction) error
	FindTransactionByAccountID(accountID int) ([]model.Transaction, error)
}

// mysqlAccountRepository struct เก็บ *sql.DB
type mysqlAccountRepository struct {
	db *sql.DB
}

// NewMySQLAccountRepository สร้าง repository ใหม่
func NewMySQLAccountRepository(db *sql.DB) AccountRepository {
	return &mysqlAccountRepository{db: db}
}

// Create สร้างบัญชีใหม่นฐานข้อมูล
func (r *mysqlAccountRepository) Create(account model.Account) (*model.Account, error) {
	query := "INSERT INTO accounts (name, balance) VALUES (?, ?)"
	result, err := r.db.Exec(query, account.Name, account.Balance)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	account.ID = int(id)
	return &account, nil
}

// FindByID ดึงบัญชีตาม ID
func (r *mysqlAccountRepository) FindByID(id int) (*model.Account, error) {
	query := "SELECT id, name, balance FROM accounts WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var account model.Account
	err := row.Scan(&account.ID, &account.Name, &account.Balance)
	if err == sql.ErrNoRows {
		return nil, errors.New("account not found")
	} else if err != nil {
		return nil, err
	}

	return &account, nil
}

// Update อัพเดตยอดเงินบัญชี
func (r *mysqlAccountRepository) Update(account model.Account) error {
	query := "UPDATE accounts SET balance = ? WHERE id = ?"
	_, err := r.db.Exec(query, account.Balance, account.ID)
	return err
}

// CreateTransaction บันทึกธุรกรรมฝาก-ถอน
func (r *mysqlAccountRepository) CreateTransaction(tx model.Transaction) error {
	query := `INSERT INTO transactions (account_id, amount, type, created_at) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(query, tx.AccountID, tx.Amount, tx.Type, time.Now())
	return err
}

// FindTransactionsByAccountID ดึงประวัติธุรกรรมตามบัญชี
func (r *mysqlAccountRepository) FindTransactionByAccountID(accountID int) ([]model.Transaction, error) {
	query := `SELECT id, account_id, amount, type, created_at FROM transactions WHERE account_id = ? ORDER BY created_at DESC`
	rows, err := r.db.Query(query, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.Transaction
	for rows.Next() {
		var tx model.Transaction
		err := rows.Scan(&tx.ID, &tx.AccountID, &tx.Amount, &tx.Type, &tx.CreatedAt)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, tx)
	}
	return transactions, nil
}
