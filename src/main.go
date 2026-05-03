package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB;

type Album struct{
	ID int64
	Title string
	Artist string
	Price float32
}

func main() {
	errr := godotenv.Load("cfg.env")
	if errr!=nil{
		log.Fatal("Error loading env file")
	}
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "learn_go_db"

	var err error
	db, err = sql.Open("mysql",cfg.FormatDSN())

	if err!=nil{
		fmt.Println("error connecting to mysql db")
		log.Fatal(err)
	}
	pingError := db.Ping()
	if pingError!=nil{
		fmt.Println("ping error")
		log.Fatal(pingError)
	}
	fmt.Println("Connection established")

	// albumsByArtist,err:=albumByArtist("John Coltrane")	
	// if err!=nil{
	// log.Fatal("Eror returned")
	// }

	// fmt.Println(albumsByArtist)

	

	// albumToInsert := []Album{
	// 	{
	// 		Title: "Dhurandhar",
	// 		Artist: "Shashwat Sachdev",
	// 		Price: 72.49,
	// 	},
	// }
	// insertAlbumIntoAlumbs(albumToInsert)


	// albumsByPrice, err := filterAlbumByPrice()
	// if err!=nil{
	// 	log.Fatal(err)
	// }

	// fmt.Println(albumsByPrice)

	// deleteAlbumFromAlbums("Dhurandhar")

	updateTitleOfArtist("After Hours","The Weeknd")

	db.Close()

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

func deleteAlbumFromAlbums(title string){
	rowsUpdated, err := db.Exec("DELETE from album where title=?",title)
	if err!=nil{
		fmt.Println("Error executing delete query")
		log.Fatal(err)
		os.Exit(-1)
	}
	numRowsUpdated,err := rowsUpdated.RowsAffected()
	if err!=nil{
		fmt.Println("Error printing rows updated")
		log.Fatal(err)
		os.Exit(-1)
	}
	fmt.Println(numRowsUpdated)
}

func updateTitleOfArtist(title string, artist string){

	//check a title of the artist already exists beforehand
	var countRows int64
	err:= db.QueryRow("SELECT count(*) from album where artist=?",artist).Scan(&countRows)
	if err!=nil{
		log.Fatal("Error fetching count of albumn from db  for artist ",artist)
		os.Exit(-1)
	}
	if countRows==0{
		log.Fatal("No entries in db for artist ",artist)
	}
	rowsUpdated, err := db.Exec("UPDATE album SET title=? where artist=?",title,artist)
	if err!=nil{
		fmt.Println("Error executing update title of artist query")
		log.Fatal(err)
		os.Exit(-1)
	}

	numRowsUpdated,err := rowsUpdated.RowsAffected()
	if err!=nil{
		fmt.Println("error fetching number of rows updated query")
		log.Fatal(err)
		os.Exit(-1)
	}

	fmt.Println("Number of rows updated: ",numRowsUpdated)

	if countRows!=numRowsUpdated{
		fmt.Println("Mismatch in number of rows updated!!!")
		log.Fatal("Error!!")
		os.Exit(-1)
	}
}