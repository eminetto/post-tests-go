package mysql

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/eminetto/post-tests-go/person"
)

// MySQL mysql repo
type MySQL struct {
	db *sql.DB
}

// NewMySQL create new repository
func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{
		db: db,
	}
}

// Create a person
func (r *MySQL) Create(p *person.Person) (person.ID, error) {
	stmt, err := r.db.Prepare(`
		insert into person (first_name, last_name, created_at)
		values(?,?,?)`)
	defer stmt.Close()
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(
		p.Name,
		p.LastName,
		time.Now().Format("2006-01-02"),
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return person.ID(id), nil
}

// Get a person
func (r *MySQL) Get(id person.ID) (*person.Person, error) {
	stmt, err := r.db.Prepare(`select id, first_name, last_name from person where id = ?`)
	if err != nil {
		return nil, err
	}
	var p person.Person
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, fmt.Errorf("not found")
	}
	err = rows.Scan(&p.ID, &p.Name, &p.LastName)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// Update a person
func (r *MySQL) Update(p *person.Person) error {
	_, err := r.db.Exec("update person set first_name = ?, last_name = ?, updated_at = ? where id = ?", p.Name, p.LastName, time.Now().Format("2006-01-02"), p.ID)
	if err != nil {
		return err
	}
	return nil
}

// Search person
func (r *MySQL) Search(query string) ([]*person.Person, error) {
	stmt, err := r.db.Prepare(`select id, first_name, last_name from person where first_name like ? or last_name like ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var people []*person.Person
	query = "%" + strings.ToLower(query) + "%"
	rows, err := stmt.Query(query, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p person.Person
		err = rows.Scan(&p.ID, &p.Name, &p.LastName)
		if err != nil {
			return nil, err
		}
		people = append(people, &p)
	}
	if len(people) == 0 {
		return nil, fmt.Errorf("not found")
	}

	return people, nil
}

// List person
func (r *MySQL) List() ([]*person.Person, error) {
	stmt, err := r.db.Prepare(`select id, first_name, last_name from person`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var people []*person.Person
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p person.Person
		err = rows.Scan(&p.ID, &p.Name, &p.LastName)
		if err != nil {
			return nil, err
		}
		people = append(people, &p)
	}
	if len(people) == 0 {
		return nil, fmt.Errorf("not found")
	}

	return people, nil
}

// Delete a person
func (r *MySQL) Delete(id person.ID) error {
	p, _ := r.Get(id)
	if p == nil {
		return fmt.Errorf("not found")
	}
	_, err := r.db.Exec("delete from person where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
