package handlers

import (
	"database/sql"
	"encoding/json"
	"example.com/m/v2/models"
	"example.com/m/v2/utils"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	Db *sql.DB
}

func parseTime(s string) time.Time {
	layout := "2006-01-02T15:04:05.000-07:00"

	// Parse the string into a time.Time object
	parsedTime, err := time.Parse(layout, s)
	if err != nil {
		log.Fatalf("Failed to parse datetime: %s", err)
	}
	return parsedTime
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
	rows, err := h.Db.Query("SELECT * FROM entries") //Iter
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var entry models.Entry
		if err := rows.Scan(
			&entry.UuId,
			&entry.AppId,
			&entry.AnswerTime,
			&entry.Created,
			&entry.Elaborate,
			&entry.ElaborateTimestampEnd,
			&entry.ElaborateTimestampStart,
			&entry.Greeting,
			&entry.GreetingTimestampEnd,
			&entry.GreetingTimestampStart,
			&entry.LanguageMedication,
			&entry.LanguageMedicationTimestampEnd,
			&entry.LanguageMedicationTimestampStart,
			&entry.Language1,
			&entry.Language1TimestampEnd,
			&entry.Language1TimestampStart,
			&entry.Language2,
			&entry.Language2TimestampEnd,
			&entry.Language2TimestampStart,
			&entry.Name,
			&entry.NameInk,
			&entry.NameInkTimestampEnd,
			&entry.NameInkTimestampStart,
			&entry.NameTimestampEnd,
			&entry.NameTimestampStart,
			&entry.Picture,
			&entry.PictureTimestampEnd,
			&entry.PictureTimestampStart,
			&entry.Radio1,
			&entry.Radio1TimestampStart,
			&entry.Radio2,
			&entry.Radio2TimestampStart,
			&entry.Retell,
			&entry.RetellTimestampEnd,
			&entry.RetellTimestampStart,
			&entry.Slider1,
			&entry.Slider1TimestampStart,
			&entry.Slider2,
			&entry.Slider2TimestampStart,
			&entry.Slider3,
			&entry.Slider3TimestampStart,
			&entry.SubmissionID,
			&entry.Suggestion,
			&entry.SuggestionTimestampEnd,
			&entry.SuggestionTimestampStart,
			&entry.TapRate,
			&entry.TapRateTimestampEnd,
			&entry.TapRateTimestampStart,
			&entry.DataVersion,
			&entry.NotificationDatetime,
			&entry.SessionId,
			&entry.StudyId,
			&entry.Annotations,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		entries = append(entries, entry)
	}
	Conv, _ := json.MarshalIndent(entries, "", " ")
	fmt.Fprintf(w, "%s\n", string(Conv))
}

func (h Handler) GetEntry(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	uuID := vars.Get("id")

	var entry models.Entry
	err := h.Db.QueryRow("SELECT * FROM entries WHERE app_id=?", uuID).Scan(
		&entry.UuId,
		&entry.AppId,
		&entry.AnswerTime,
		&entry.Created,
		&entry.Elaborate,
		&entry.ElaborateTimestampEnd,
		&entry.ElaborateTimestampStart,
		&entry.Greeting,
		&entry.GreetingTimestampEnd,
		&entry.GreetingTimestampStart,
		&entry.LanguageMedication,
		&entry.LanguageMedicationTimestampEnd,
		&entry.LanguageMedicationTimestampStart,
		&entry.Language1,
		&entry.Language1TimestampEnd,
		&entry.Language1TimestampStart,
		&entry.Language2,
		&entry.Language2TimestampEnd,
		&entry.Language2TimestampStart,
		&entry.Name,
		&entry.NameInk,
		&entry.NameInkTimestampEnd,
		&entry.NameInkTimestampStart,
		&entry.NameTimestampEnd,
		&entry.NameTimestampStart,
		&entry.Picture,
		&entry.PictureTimestampEnd,
		&entry.PictureTimestampStart,
		&entry.Radio1,
		&entry.Radio1TimestampStart,
		&entry.Radio2,
		&entry.Radio2TimestampStart,
		&entry.Retell,
		&entry.RetellTimestampEnd,
		&entry.RetellTimestampStart,
		&entry.Slider1,
		&entry.Slider1TimestampStart,
		&entry.Slider2,
		&entry.Slider2TimestampStart,
		&entry.Slider3,
		&entry.Slider3TimestampStart,
		&entry.SubmissionID,
		&entry.Suggestion,
		&entry.SuggestionTimestampEnd,
		&entry.SuggestionTimestampStart,
		&entry.TapRate,
		&entry.TapRateTimestampEnd,
		&entry.TapRateTimestampStart,
		&entry.DataVersion,
		&entry.NotificationDatetime,
		&entry.SessionId,
		&entry.StudyId,
		&entry.Annotations,
	)

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
	err := h.Db.QueryRow("SELECT count(*) FROM entries").Scan(&Count)
	utils.Panic(err, "Failed to count entries!")
	fmt.Fprintf(w, "%s\n", string(Count))
}

/**/
/*
POST Requests
*/
func (h Handler) Upload(w http.ResponseWriter, r *http.Request) {
	// Parse the uploaded file
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	parseUUID := func(s string) (uuid.UUID, error) {
		return uuid.Parse(s)
	}

	var entriesTable []models.Entry

	err, data := utils.ConvertToJson(file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to convert to json\n%s", err), http.StatusBadRequest)
	}

	if err := json.Unmarshal([]byte(data), &entriesTable); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse json\n%s", err), http.StatusBadRequest)
	}

	tx, err := h.Db.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stmt, err := tx.Prepare("INSERT INTO entries(uu_id,app_id,answer_time,created,elaborate,elaborate_timestamp_end,elaborate_timestamp_start,greeting,greeting_timestamp_end,greeting_timestamp_start,language_medication,language_medication_timestamp_end,language_medication_timestamp_start,language1,language1_timestamp_end,language1_timestamp_start,language2,language2_timestamp_end,language2_timestamp_start,name,name_ink,name_ink_timestamp_end,name_ink_timestamp_start,name_timestamp_end,name_timestamp_start,picture,picture_timestamp_end,picture_timestamp_start,radio1,radio1_timestamp_start,radio2,radio2_timestamp_start,retell,retell_timestamp_end,retell_timestamp_start,slider1,slider1_timestamp_start,slider2,slider2_timestamp_start,slider3,slider3_timestamp_start,submission_id,suggestion,suggestion_timestamp_end,suggestion_timestamp_start,tap_rate,tap_rate_timestamp_end,tap_rate_timestamp_start,data_version,notification_datetime,session_id,study_id,annotations) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	for _, entry := range entriesTable {

		uniqueID := uuid.New()

		appID, _ := parseUUID(entry.AppId)

		_, err = stmt.Exec(
			uniqueID,
			appID,
			entry.AnswerTime,
			entry.Created,
			entry.Elaborate,
			entry.ElaborateTimestampEnd,
			entry.ElaborateTimestampStart,
			entry.Greeting,
			entry.GreetingTimestampEnd,
			entry.GreetingTimestampStart,
			entry.LanguageMedication,
			entry.LanguageMedicationTimestampEnd,
			entry.LanguageMedicationTimestampStart,
			entry.Language1,
			entry.Language1TimestampEnd,
			entry.Language1TimestampStart,
			entry.Language2,
			entry.Language2TimestampEnd,
			entry.Language2TimestampStart,
			entry.Name,
			entry.NameInk,
			entry.NameInkTimestampEnd,
			entry.NameInkTimestampStart,
			entry.NameTimestampEnd,
			entry.NameTimestampStart,
			entry.Picture,
			entry.PictureTimestampEnd,
			entry.PictureTimestampStart,
			entry.Radio1,
			entry.Radio1TimestampStart,
			entry.Radio2,
			entry.Radio2TimestampStart,
			entry.Retell,
			entry.RetellTimestampEnd,
			entry.RetellTimestampStart,
			entry.Slider1,
			entry.Slider1TimestampStart,
			entry.Slider2,
			entry.Slider2TimestampStart,
			entry.Slider3,
			entry.Slider3TimestampStart,
			entry.SubmissionID,
			entry.Suggestion,
			entry.SuggestionTimestampEnd,
			entry.SuggestionTimestampStart,
			entry.TapRate,
			entry.TapRateTimestampEnd,
			entry.TapRateTimestampStart,
			entry.DataVersion,
			entry.NotificationDatetime,
			entry.SessionId,
			entry.StudyId,
			"", //empty by default
		)
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
	Conv, _ := json.MarshalIndent(entriesTable, "", " ")
	fmt.Fprintf(w, "%s\n", string(Conv))
}

func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
	//ID := mux.Vars(r)["id"]
	//var updated models.Entry
	//reqBody, err := io.ReadAll(r.Body)
	//utils.Panic(err, "Failed to read request body!")
	//json.Unmarshal(reqBody, &updated)
	//if _, postErr := h.Db.Exec("UPDATE db_entries SET entry_val = ? WHERE id = ?", updated.EntryVal, ID); postErr != nil {
	//	utils.Panic(postErr, "Failed to update entry!")
	//}
	//fmt.Fprintf(w, "Entry %s updated!\n", ID)
}

/*
DELETE Requests
*/
func (h Handler) DeleteOne(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	if _, err := h.Db.Exec("DELETE FROM entries WHERE id=?", ID); err != nil {
		utils.Panic(err, "Failed to delete entry!")
	}
	fmt.Fprintf(w, "Entry %s deleted!\n", ID)
}

func (h Handler) DeleteAll(w http.ResponseWriter, r *http.Request) {
	if _, err := h.Db.Exec("TRUNCATE entries"); err != nil {
		utils.Panic(err, "Failed to delete entries!")
	}
	fmt.Fprintf(w, "All entries deleted!\n")
}
