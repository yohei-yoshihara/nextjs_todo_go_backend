package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var WaitValue time.Duration = 3

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "リバースプロキシとAPIサーバの起動",
	RunE: func(cmd *cobra.Command, args []string) error {
		connect, err := cmd.Flags().GetString("connect")
		if err != nil {
			return err
		}
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			return err
		}
		waitValue, err := cmd.Flags().GetInt("wait")
		if err != nil {
			return err
		}
		WaitValue = time.Duration(waitValue)

		run(connect, port)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().Int("port", 8000, "待ち受けるポート番号")
	serveCmd.Flags().String("connect", "http://localhost:3000", "接続先")
	serveCmd.Flags().Int("wait", 3, "待機時間(秒)")
}

type Folder struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Task struct {
	ID       int64  `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	FolderId int64  `json:"folder_id,omitempty"`
}

/*
  /folders
	  * GET - get all folders
	/folders/create
	  * POST - create a new folder
	/folders/delete
	  * POST - delete a folder
	/folders/:id
	  * GET - get a folder

  /tasks
	  * GET - get all tasks
	/tasks/create
	  * POST - create a new task
	/tasks/delete
	  * POST - delete a task
	/tasks/:id
	  * GET - get a task
*/

func run(connect string, port int) {
	remote, err := url.Parse(connect)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// ***** folder *****

	http.HandleFunc("/api/folders", foldersHandler(db))
	http.HandleFunc("/api/folders/{id}", folderHandler(db))
	http.HandleFunc("/api/folders/create", createFolderHandler(db))
	http.HandleFunc("/api/folders/delete", deleteFolderHandler(db))

	http.HandleFunc("/api/tasks", tasksHandler(db))
	http.HandleFunc("/api/tasks/{id}", taskHandler(db))
	http.HandleFunc("/api/tasks/create", createTaskHandler(db))
	http.HandleFunc("/api/tasks/delete", deleteTaskHandler(db))

	// ***** Reverse Proxy *****
	reverseProxyHandler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL)
			r.Host = remote.Host
			// w.Header().Set("xxx", "yyy")
			p.ServeHTTP(w, r)
		}
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", reverseProxyHandler(proxy))

	fmt.Println("listening...")
	err = http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		panic(err)
	}
}
