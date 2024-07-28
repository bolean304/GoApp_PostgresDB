 package main
import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/lib/pq"
)
type Product struct {
	Name string
	Price float64
	Available bool
}
 func main(){
	connstr := "postgres://postgres:mysecretpassword@localhost:5432/gopgtest?sslmode=disable"
	db ,err := sql.Open("postgres",connstr)
	defer db.Close()
	if err != nil{
		log.Fatal(err)
	}
err = db.Ping();
if err != nil{
	log.Fatal(err)
}
CreateProductTable(db)
product := Product{"Book",15.55,true}
pk := insertProduct(db,product)
fmt.Println("ID : %d",pk)
 }

func CreateProductTable(db *sql.DB ){
	query := `CREATE TABLE IF NOT EXISTS product(
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	price NUMERIC(6,2) NOT NULL,
	available BOOLEAN,
	created timestamp DEFAULT NOW()
	)`
 _, err :=  db.Exec(query)
 if err != nil{
	log.Fatal(err)
}
} 

func insertProduct(db *sql.DB ,product Product ) int{
   query := ` INSERT INTO product (name,price,available)
   VALUES ($1, $2,$3 ) RETURNING id`
 var pk int   
   err := db.QueryRow(query,product.Name,product.Price,product.Available).Scan(&pk)
   if err != nil{
	log.Fatal(err)
}
return pk
}