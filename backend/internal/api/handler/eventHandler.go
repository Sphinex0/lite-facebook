package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (Handler *Handler) AddEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	user, ok := r.Context().Value(utils.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	var Event models.Event
	err := utils.ParseBody(r, &Event)
	fmt.Println(Event)
	Event.UserID = user.ID
	if err != nil || Event.UserID == 0 || Event.GroupID == 0 {
		utils.WriteJson(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if Event.Title == "" || Event.Description == "" {
		utils.WriteJson(w, http.StatusBadRequest, "Bad Request")
		return
	}

	if err := Handler.Service.CreateEvent(Event); err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

func (Handler *Handler) GetEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("fffffffffffffffffffffffffffffff")
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	user, ok := r.Context().Value(utils.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	var Event models.Event
	err := utils.ParseBody(r, &Event)
	Event.UserID=user.ID
	fmt.Println(Event)
	Events, err := Handler.Service.AllEvents(Event)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	fmt.Println("eeeeeee", Events)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Events)
}

func (Handler *Handler) GetEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	var Events models.Event
	err := utils.ParseBody(r, &Events)
	Event, err := Handler.Service.GetEventsById(&Events)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Event)
}

// skip

func (Handler *Handler) OptionEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	user, ok := r.Context().Value(utils.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var OptionEvent models.EventOption

	OptionEvent.UserID = user.ID



	err := utils.ParseBody(r, &OptionEvent)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	
	err = Handler.Service.PostEventsOption(OptionEvent)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

func (Handler *Handler) GetEventOption(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	var EventOption models.EventOption
	err := utils.ParseBody(r, &EventOption)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	EventOptions, err := Handler.Service.GetEventsOption(EventOption)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(EventOptions)
}


func (Handler *Handler) GetEventchoise(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetEventchoise")
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}




	var EventOption models.EventOption
	err := utils.ParseBody(r, &EventOption)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
	fmt.Println("EventOption",EventOption)
	Event, err := Handler.Service.GetEventgoing(EventOption)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Event)
}
