package postgresql

import (
	"context"
	"database/sql"
	"shopping_backend/domain"

	"github.com/sirupsen/logrus"
)

// 定義好需要哪些依賴注入(DI): 這是 Clean Architecture 的核心之一，將依賴的事物由外部注入，而不是寫死在裡頭
type postgresqlAccountRepository struct {
	db *sql.DB
}

// 設計一個DI注入的function: 我們還需要一個注入的 function，將 db 確實注入。為什麼要這樣呢？因為你是可以這樣產生實例的postgresqlAccountRepository{}，可以發現其實沒有帶 db，但實際上還是可以運行。我們可以透過此 function 來避免這個情形。
// 透過domain裡定義的interface來約束回傳值: domain.AccountRepository這個 interface 定義了有哪些方法postgresqlAccountRepository必須要實作，有了此 interface 才能讓呼叫的程式在還沒run起來就知道哪些呼叫方法存在。
func NewpostgresqlAccountRepository(db *sql.DB) domain.AccountRepository {
	return &postgresqlAccountRepository{db}
}

// 實作domain.AccountRepository interface定義的方法: GetByID要確實符合 interface 定義的方法，不然 Golang 在運行前就會報錯。
func (p *postgresqlAccountRepository) GetByID(ctx context.Context, id string) (*domain.Account, error) {
	row := p.db.QueryRow("SELECT id, name, status FROM digimons WHERE id =$1", id)
	d := &domain.Account{}
	if err := row.Scan(&d.ID, &d.Name, &d.Status); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return d, nil
}

func (p *postgresqlAccountRepository) Store(ctx context.Context, d *domain.Account) error {
	_, err := p.db.Exec(
		"INSERT INTO digimons (id, name, status) VALUES ($1, $2, $3)",
		d.ID, d.Name, d.Status,
	)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (p *postgresqlAccountRepository) UpdateStatus(ctx context.Context, d *domain.Account) error {
	_, err := p.db.Exec(
		"UPDATE digimons SET status=$1 WHERE id=$2",
		d.Status, d.ID,
	)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
