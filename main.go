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
    product := Product{"Table ",200,true}
    pk := InsertProduct(db,product)
    fmt.Println("ID : %d",pk)
	var name  string
	var price float64
	var available bool
	query := "SELECT name ,available, price FROM product WHERE id = $1"
    errr := db.QueryRow(query,pk).Scan(&name, &available, &price)
	if errr != nil{
		log.Fatal(errr)
	}
	fmt.Println("name : %v",name)
	fmt.Println("price : %v",price)
	fmt.Println("available : %v",available)
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
// This function insert  data from Database 

func InsertProduct(db *sql.DB ,product Product ) int{
   query := ` INSERT INTO product (name,price,available)
   VALUES ($1, $2,$3 ) RETURNING id`
   var pk int   
   err := db.QueryRow(query,product.Name,product.Price,product.Available).Scan(&pk)
   if err != nil{
	log.Fatal(err)
   }
   return pk
}



