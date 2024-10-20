package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestFolders(t *testing.T) {
	resp, err := http.Get("http://localhost:8000/api/folders")
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	folders := []Folder{}
	err = json.Unmarshal(body, &folders)
	if err != nil {
		t.Error(err)
		return
	}
	if len(folders) != 3 {
		t.Errorf("folders must have 3 elements, %v", folders)
	}
	if folders[0].ID != 1 || folders[0].Name != "プライベート" {
		t.Errorf("folder 1 is invalid")
	}
	if folders[1].ID != 2 || folders[1].Name != "仕事" {
		t.Errorf("folder 2 is invalid")
	}
}

func TestFolder(t *testing.T) {
	resp, err := http.Get("http://localhost:8000/api/folders/1")
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var folder Folder
	err = json.Unmarshal(body, &folder)
	if err != nil {
		t.Error(err)
		return
	}
	if folder.ID != 1 || folder.Name != "プライベート" {
		t.Errorf("folder 1 is invalid")
	}
}

func TestCreateFolder(t *testing.T) {
	var folderId int64
	{ // create
		folder := Folder{
			Name: "テスト",
		}
		folderJson, err := json.Marshal(folder)
		if err != nil {
			t.Error(err)
			return
		}
		resp, err := http.Post("http://localhost:8000/api/folders/create", "application/json", bytes.NewBuffer(folderJson))
		if err != nil {
			t.Error(err)
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		var retFolder Folder
		err = json.Unmarshal(body, &retFolder)
		if err != nil {
			t.Error(err)
			return
		}
		if retFolder.Name != "テスト" {
			t.Errorf("failed to create a new folder, %v", retFolder)
		}
		folderId = retFolder.ID
	}

	{ // delete
		folder := Folder{
			ID: folderId,
		}
		folderJson, err := json.Marshal(folder)
		if err != nil {
			t.Error(err)
			return
		}
		resp, err := http.Post("http://localhost:8000/api/folders/delete", "application/json", bytes.NewBuffer(folderJson))
		if err != nil {
			t.Error(err)
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		var retFolder Folder
		err = json.Unmarshal(body, &retFolder)
		if err != nil {
			t.Error(err)
			return
		}
		if retFolder.ID != folderId {
			t.Errorf("failed to delete a new folder, %v", retFolder)
		}
	}
}
