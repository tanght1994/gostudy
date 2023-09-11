package sqlite

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	g_sensitiveWordDbName    = "./sensitive_word.db"
	g_sensitiveWordTableName = "stat"
)

var (
	g_sensitiveWordStore *sensitiveWordStore
)

type sensitiveData struct {
	t1 string
	t2 int64
	s  string
	d  string
}

func init() {
	g_sensitiveWordStore = newSensitiveWordStore()
}

func create_table(db *sql.DB, name string) error {
	sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS [%s]([t1] VARCHAR(50), [t2] INTEGER, [s] TEXT, [d] TEXT)", name)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	exist, err := table_exist(db, name)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("create table error")
	}
	return nil
}

func table_exist(db *sql.DB, name string) (bool, error) {
	sql := fmt.Sprintf(`SELECT [name] FROM sqlite_master WHERE [type]="table" AND [name]="%s"`, name)
	rows, err := db.Query(sql)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		return true, nil
	}
	return false, nil
}

type sensitiveWordStore struct {
	db        *sql.DB
	ch        chan sensitiveData
	buf       []sensitiveData
	dbname    string
	tablename string
}

func newSensitiveWordStore() *sensitiveWordStore {
	db, err := sql.Open("sqlite3", "./sensitive_word.db")
	if err != nil {
		return nil
	}
	err = create_table(db, "stat")
	if err != nil {
		return nil
	}
	store := &sensitiveWordStore{
		db:        db,
		ch:        make(chan sensitiveData, 10000),
		buf:       make([]sensitiveData, 0),
		dbname:    g_sensitiveWordDbName,
		tablename: g_sensitiveWordTableName,
	}
	go store.ch2db()
	return store
}

func (sws *sensitiveWordStore) unblockSave(data sensitiveData) error {
	select {
	case sws.ch <- data:
		return nil
	default:
		return fmt.Errorf("channel full")
	}
}

func (sws *sensitiveWordStore) blockSave(data sensitiveData) error {
	sws.ch <- data
	return nil
}

func (sws *sensitiveWordStore) ch2db() {
	tkInterval := 1 * time.Second
	tk := time.NewTicker(tkInterval)
	for {
		select {
		case x, ok := <-sws.ch:
			if !ok {
				break
			}
			sws.buf = append(sws.buf, x)
			if len(sws.buf) > 5000 {
				sws.buf2db()
			}
		case <-tk.C:
			sws.buf2db()
			tk.Reset(tkInterval)
		}
	}
}

func (sws *sensitiveWordStore) buf2db() {
	defer func() { sws.buf = sws.buf[:0] }()
	if len(sws.buf) == 0 {
		return
	}
	fmt.Println("save ", len(sws.buf))
	tx, err := sws.db.Begin()
	if err != nil {
		return
	}
	s := fmt.Sprintf(`INSERT INTO %s ([t1], [t2], [s], [d]) VALUES (?, ?, ?, ?)`, sws.tablename)
	stmt, err := tx.Prepare(s)
	if err != nil {
		return
	}
	defer stmt.Close()
	for _, data := range sws.buf {
		stmt.Exec(data.t1, data.t2, data.s, data.d)
	}
	tx.Commit()
}

func gett1andt2() (t1 string, t2 int64) {
	t0 := time.Now()
	t1 = t0.Format("2006-01-02 15:04:05 -07:00")
	t2 = t0.Unix()
	return
}

func SaveSensitiveWordUnblock(s string, d string) error {
	t1, t2 := gett1andt2()
	return g_sensitiveWordStore.unblockSave(sensitiveData{t1: t1, t2: t2, s: s, d: d})
}

func SaveSensitiveWordBlock(s string, d string) error {
	t1, t2 := gett1andt2()
	return g_sensitiveWordStore.blockSave(sensitiveData{t1: t1, t2: t2, s: s, d: d})
}
