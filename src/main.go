package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB;

type Album struct{
	ID int64
	Title string
	Artist string
	Price float32
}

func main() {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "learn_go_db"

	var err error
	db, err = sql.Open("mysql",cfg.FormatDSN())

	if err!=nil{
		log.Fatal(err)
	}
	pingError := db.Ping()
	if pingError!=nil{
		log.Fatal(pingError)
	}
	fmt.Println("Connection established")

	albumsByArtist,err:=albumByArtist("John Coltrane")	
	if err!=nil{
	log.Fatal("Eror returned")
	}

	fmt.Println(albumsByArtist)

	

	albumToInsert := []Album{
		{
			Title: "Dhurandhar",
			Artist: "Shashwat Sachdev",
			Price: 72.49,
		},
	}
	insertAlbumIntoAlumbs(albumToInsert)


	albumsByPrice, err := filterAlbumByPrice()
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println(albumsByPrice)
}


func albumByArtist(name string)([]Album,error){
	var albums []Album
	rows,err := db.Query("SELECT * FROM album WHERE artist=?",name)
	if err!=nil{
		log.Fatal("Select query returned nil\n")
		return nil,fmt.Errorf("error returned %q: %v",name,err)
	}

	defer rows.Close()

	for rows.Next(){
		var alb Album
		if err := rows.Scan(&alb.ID,&alb.Title,&alb.Artist,&alb.Price);err!=nil{
			log.Fatal("Error in scan block\n")
			return nil,fmt.Errorf("row %q: %v",name,err)
		}

		albums = append(albums, alb)

	}

	if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    return albums, nil

}

func filterAlbumByPrice()([]Album,error){
	var albums []Album
	rows,err := db.Query("SELECT * FROM album order by price desc")
	if err!=nil{
		log.Fatal(err)
	}

	for rows.Next(){
		var alb Album
		rows.Scan(&alb.ID,&alb.Title,&alb.Artist,&alb.Price)
		albums = append(albums, alb)
	}

	return albums,nil
}


func insertAlbumIntoAlumbs(album []Album){
	fmt.Printf("Album: %v",album)
	title := album[0].Title
	artist := album[0].Artist
	price := album[0].Price
	result,err :=db.Exec("INSERT INTO album (title,artist,price) values (?,?,?)",title,artist,price)
	if err!=nil{
		log.Fatal(err)
	}

	rowsAffected,err := result.RowsAffected()
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("Executed insert quuery. Rows affected: %d\n",rowsAffected)
}


