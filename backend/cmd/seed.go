package cmd

import (
	"database/sql"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	_ "github.com/mattn/go-sqlite3"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "データベースへのデータ登録",
	Run: func(cmd *cobra.Command, args []string) {
		RunSeed()
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}

func RunSeed() {
	os.Remove("./database.db")

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(`
	create table folders (
	  id integer primary key autoincrement, 
		name text not null);
	`)
	if err != nil {
		panic(err)
	}

	folderNames := []string{"プライベート", "仕事", "その他"}
	stmt, err := db.Prepare(`
	insert into folders(name) values(?)
	`)
	if err != nil {
		panic(err)
	}
	folderIds := []int64{}
	for _, name := range folderNames {
		result, err := stmt.Exec(name)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		folderIds = append(folderIds, id)
	}

	_, err = db.Exec(`
	create table tasks (
	  id integer primary key autoincrement,
		title text not null,
		folder_id integer not null
	)
	`)
	if err != nil {
		panic(err)
	}

	stmt, err = db.Prepare("insert into tasks(title, folder_id) values(?, ?)")
	if err != nil {
		panic(err)
	}
	taskNames := []string{}
	for i := 0; i < 20; i++ {
		taskNames = append(taskNames, "Task "+strconv.Itoa(i))
	}
	for i, name := range taskNames {
		_, err = stmt.Exec(name, folderIds[i%len(folderIds)])
		if err != nil {
			panic(err)
		}
	}

}
