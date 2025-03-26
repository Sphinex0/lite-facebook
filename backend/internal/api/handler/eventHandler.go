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
	Event.UserID = user.ID
	if err != nil || Event.GroupID == 0  {
		utils.WriteJson(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if Event.Title == "" || Event.Description == "" || Event.Day == "" {
		utils.WriteJson(w, http.StatusBadRequest, "Bad Request")
		return
	}

	if err := Handler.Service.CreateEvent(Event); err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

func (Handler *Handler) GetEvents(w http.ResponseWriter, r *http.Request) {
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
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	groupErr := Handler.Service.VerifyGroup(Event.GroupID, user.ID)
	err = groupErr.Err
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	Event.UserID=user.ID
	Events, err := Handler.Service.AllEvents(Event)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Events)
}

func (Handler *Handler) GetEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	user, ok := r.Context().Value(utils.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var Events models.Event
	err := utils.ParseBody(r, &Events)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	groupErr := Handler.Service.VerifyGroup(Events.GroupID, user.ID)
	err = groupErr.Err
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	
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

	// get event
	var event models.Event
	event.ID = OptionEvent.EventID
	ev  ,err := Handler.Service.GetEventsById(&event)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	groupErr := Handler.Service.VerifyGroup(ev.GroupID, user.ID)
	err = groupErr.Err
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
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

	user, ok := r.Context().Value(utils.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var EventOption models.EventOption
	err := utils.ParseBody(r, &EventOption)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	//
	var event models.Event
	event.ID = EventOption.EventID
	ev  ,err := Handler.Service.GetEventsById(&event)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	groupErr := Handler.Service.VerifyGroup(ev.GroupID, user.ID)
	err = groupErr.Err
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
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
	if r.Method != http.MethodPost {
		utils.WriteJson(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	user, ok := r.Context().Value(utils.UserIDKey).(models.UserInfo)
	if !ok {
		utils.WriteJson(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	//

	var EventOption models.EventOption
	err := utils.ParseBody(r, &EventOption)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	//
	//
	var event models.Event
	event.ID = EventOption.EventID
	ev  ,err := Handler.Service.GetEventsById(&event)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	groupErr := Handler.Service.VerifyGroup(ev.GroupID, user.ID)
	err = groupErr.Err
	if err != nil {
		utils.WriteJson(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	//
	fmt.Println("EventOption",EventOption)
	Event, action,err := Handler.Service.GetEventgoing(EventOption ,user.ID)	
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	
	}
	fmt.Println("Event",Event)
	fmt.Println("action",action)
	response := struct {
		Event  interface{} `json:"event"`
		Action interface{} `json:"action"`
	}{
		Event:  Event,
		Action: action,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
