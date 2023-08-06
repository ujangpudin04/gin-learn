package storage

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Storage struct{
	db *sqlx.DB
}

func NewStorage() *Storage {
	dsn := "root:@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	return &Storage{
		db: db,
	}
}

func (s *Storage) Store(title string,description string,price string,image string,qty string,idPerson int) {
	tx := s.db.MustBegin()
    tx.MustExec("INSERT INTO books (title, description, price, image,qty,id_person) VALUES (?,?,?,?,?,?)", title, description, price,image,qty,idPerson)
    tx.Commit()
}

func (s *Storage) GetData() []Book {
	
	books := []Book{}

	if err := s.db.Select(&books, "select * from books"); err != nil {
		log.Println("error", err.Error())
	}
	

	return books
}


func (s *Storage)GetOneData(id int) (Book, error) {
	
	row := s.db.QueryRow("SELECT * FROM books WHERE id=?", id)
	book:=Book{}
	if err := row.Scan(
		&book.ID,
		&book.Title,
		&book.Description,
		&book.Price,
		&book.Image,
		&book.Qty,
		&book.IdPerson,
		&book.UpdatedAt,
	);err!=nil {
		log.Println("error",err)
	}
	return book,nil
}

func (s *Storage) UpdateBook(id int,title string,description string,price string,image string,qty string,idPerson int)(Book, error)  {
	book:=Book{}
	
	tx := s.db.MustBegin()
	tx.MustExec(`UPDATE books SET title = ?, description = (?), price=(?) , image=(?), qty=(?), id_person=(?) WHERE id = ?`, title, description, price,image,qty,idPerson, id)
	
	err:=tx.Commit()
	if err!=nil {
		log.Println(err)
	}
	return book,nil	
}

func (s *Storage)DeleteBook(id int)(Book,error)  {
	book:=Book{}

	tx := s.db.MustBegin()
	tx.MustExec(`DELETE FROM books WHERE id=?`,id)
	
	err:=tx.Commit()
	if err!=nil {
		log.Println(err)
	}
	return book,nil	
}


// PERSON QUERY

func (s *Storage)StorePerson(name string, email string, alamat string)()  {
	tx:=s.db.MustBegin()
	tx.MustExec(`INSERT INTO person (name,email,alamat) VALUES (?,?,?)`,name,email,alamat)
	tx.Commit()
}



func (s *Storage) GetPerson() []Person {
	
	persons := []Person{}

	if err := s.db.Select(&persons, "SELECT * FROM person"); err != nil {
		log.Println("error", err.Error())
	}
	

	return persons
}

func (s *Storage) PostPersonBook(idPerson int,idBook int) {
	tx := s.db.MustBegin()
    tx.MustExec("INSERT INTO person_book (id_person,id_book) VALUES (?,?)",idPerson,idBook)
    tx.Commit()
}

func (s *Storage)UploadImage(filename string)()  {
	tx:=s.db.MustBegin()
	tx.MustExec(`INSERT INTO files (filename) VALUES (?)`,filename)
	tx.Commit()
}


type Book struct {
	ID          int `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Price       string `db:"price" json:"price"`
	Image       string `db:"image" json:"image"`
	Qty         string `db:"qty" json:"qty"`
	IdPerson          int `db:"id_person" json:"id_person"`
	UpdatedAt 	string `db:"updated_at" json:"updated_at"`
}

type Person struct {
	ID     int `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Alamat string `json:"alamat"`
}

type PersonBook struct{
	ID     int `json:"id"`
	IdPerson     int `json:"id_person"`
	IdBook     int `json:"id_book"`

}
