package handlers

import (
	"database/sql"
	"encoding/json"
	"example.com/m/v2/models"
	"example.com/m/v2/utils"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type Handler struct {
	Db *sql.DB
}

func (h Handler) HomeLink(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "My Restful API")
	utils.Panic(err, "Failed to write a message!")
}

//TODO unittests
/*
GET Requests
*/
func (h Handler) GetEntries(w http.ResponseWriter, r *http.Request) {
	var entries []models.Entry
	rows, err := h.Db.Query("SELECT * FROM db_entries") //Iter
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var entry models.Entry
		if err := rows.Scan(&entry.ID, &entry.EntryVal); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		entries = append(entries, entry)
	}
	Conv, _ := json.MarshalIndent(entries, "", " ")
	fmt.Fprintf(w, "%s\n", string(Conv))
}

func (h Handler) GetEntry(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	var entry models.Entry
	err := h.Db.QueryRow("SELECT * FROM db_entries WHERE id=?", ID).Scan(&entry.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Entry not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Conv, _ := json.MarshalIndent(entry, "", " ")
	fmt.Fprintf(w, "%s\n", string(Conv))
}

func (h Handler) CountEntries(w http.ResponseWriter, r *http.Request) {
	var Count string
	err := h.Db.QueryRow("SELECT count(*) FROM db_entries").Scan(&Count)
	utils.Panic(err, "Failed to count entries!")
	fmt.Fprintf(w, "%s\n", string(Count))
}

/**/
/*
POST Requests
*/
func (h Handler) CreateEntry(w http.ResponseWriter, r *http.Request) {
	var entry models.Entry
	entry.ID = uuid.New().String()
	requestBody, err := io.ReadAll(r.Body)
	utils.Panic(err, "Failed to read request body!")

	unmarshallErr := json.Unmarshal(requestBody, &entry)
	utils.Panic(unmarshallErr, "Failed to process request body!")

	if _, insertErr := h.Db.Exec("INSERT INTO db_entries(id,entry_val) VALUES (?, ?)", entry.ID, entry.EntryVal); insertErr != nil {
		utils.Panic(insertErr, "Failed to insert data")
	}
	w.WriteHeader(http.StatusCreated)
	Conv, _ := json.MarshalIndent(entry, "", " ")
	fmt.Fprintf(w, "%s\n", string(Conv))

}

func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	var updated models.Entry
	reqBody, err := io.ReadAll(r.Body)
	utils.Panic(err, "Failed to read request body!")
	json.Unmarshal(reqBody, &updated)
	if _, postErr := h.Db.Exec("UPDATE db_entries SET entry_val = ? WHERE id = ?", updated.EntryVal, ID); postErr != nil {
		utils.Panic(postErr, "Failed to update entry!")
	}
	fmt.Fprintf(w, "Entry %s updated!\n", ID)
}

func (h Handler) CreateEntries(w http.ResponseWriter, r *http.Request) {
	var entries []models.Entry
	requestBody, err := io.ReadAll(r.Body)
	utils.Panic(err, "Failed to read request body!")
	unmarshallErr := json.Unmarshal(requestBody, &entries)
	utils.Panic(unmarshallErr, "Failed to process request body!")

	for i := range entries {
		entries[i].ID = uuid.New().String()
	}

	tx, err := h.Db.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stmt, err := tx.Prepare("INSERT INTO db_entries(id,entry_val) VALUES (?, ?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	for _, entry := range entries {
		_, err = stmt.Exec(entry.ID, entry.EntryVal)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	Conv, _ := json.MarshalIndent(entries, "", " ")
	fmt.Fprintf(w, "%s\n", string(Conv))

}

/*
DELETE Requests
*/
func (h Handler) DeleteOne(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	if _, err := h.Db.Exec("DELETE FROM db_entries WHERE id=?", ID); err != nil {
		utils.Panic(err, "Failed to delete entry!")
	}
	fmt.Fprintf(w, "Entry %s deleted!\n", ID)
}

func (h Handler) DeleteAll(w http.ResponseWriter, r *http.Request) {
	if _, err := h.Db.Exec("TRUNCATE db_entries"); err != nil {
		utils.Panic(err, "Failed to delete entries!")
	}
	fmt.Fprintf(w, "All entries deleted!\n")
}
