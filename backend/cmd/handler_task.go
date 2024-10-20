package cmd

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func tasksHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(WaitValue * time.Second)

		folderId := r.URL.Query().Get("folderId")

		var rows *sql.Rows
		var err error

		if folderId == "" {
			rows, err = db.Query("select id, title, folder_id from tasks")
		} else {
			rows, err = db.Query("select id, title, folder_id from tasks where folder_id = ?", folderId)
		}
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		defer rows.Close()

		tasks := []Task{}
		for rows.Next() {
			var id int64
			var name string
			var folderId int64
			err := rows.Scan(&id, &name, &folderId)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			tasks = append(tasks, Task{ID: id, Title: name, FolderId: folderId})
		}

		data, err := json.Marshal(tasks)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

// /folders/{id}
func taskHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(WaitValue * time.Second)
		idString := r.PathValue("id")
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			http.Error(w, "illegal id", http.StatusBadRequest)
			return
		}
		row := db.QueryRow("select title, folder_id from tasks where id = ?", id)
		var title string
		var folderId int64
		err = row.Scan(&title, &folderId)
		if err != nil {
			http.Error(w, "database error", http.StatusBadRequest)
			return
		}
		folder := Task{ID: id, Title: title, FolderId: folderId}
		data, err := json.Marshal(folder)
		if err != nil {
			http.Error(w, "failed to marshal JSON", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func createTaskHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(WaitValue * time.Second)
		var task Task

		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := db.Exec("insert into tasks(title, folder_id) values(?, ?)", task.Title, task.FolderId)
		if err != nil {
			http.Error(w, "failed to insert a folder", http.StatusBadRequest)
			return
		}
		id, err := result.LastInsertId()
		if err != nil {
			http.Error(w, "failed to get id", http.StatusBadRequest)
			return
		}
		task.ID = id
		data, err := json.Marshal(task)
		if err != nil {
			http.Error(w, "failed to marshal JSON", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func deleteTaskHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(WaitValue * time.Second)
		var task Task

		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = db.Exec("delete from tasks where id=?", task.ID)
		if err != nil {
			http.Error(w, "failed to insert", http.StatusBadRequest)
			return
		}
		data, err := json.Marshal(task)
		if err != nil {
			http.Error(w, "failed to marshal JSON", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
