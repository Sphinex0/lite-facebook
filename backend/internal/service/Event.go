package service

import (
	"database/sql"
	"fmt"
	"log"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (service *Service) CreateEvent(event_pardse models.EventPardse) (err error) {
	// valid, err := service.Database.GetCreatorGroup(Events.GroupID, Events.UserID)
	Events := event_pardse.Event
	err2 := service.VerifyGroup(Events.GroupID, Events.UserID)
	if err2.Err != nil {
		log.Println("err => ", err)
		err = err2.Err
		return
	}
	err = service.Database.SaveEvent(&Events)
	if err != nil {
		return
	}
	//
	option := models.EventOption{
		Going:   event_pardse.Going,
		UserID:  Events.UserID,
		EventID: Events.ID,
	}
	err = service.Database.SaveOptionEvent(&option)
	if err != nil {
		return
	}
	//
	var ids []int
	ids, err = service.Database.GetGroupMembers(Events.GroupID)
	if err != nil {
		return
	}
	var notification models.Notification
	notification.Type = "event-created"
	notification.InvokerID = Events.UserID
	notification.EventID = Events.ID
	notification.GroupID = Events.GroupID
	for _, id := range ids {
		if id != Events.UserID {
			notification.UserID = id
			err = service.AddNotification(notification)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
	return
}

func (service *Service) AllEvents(Event models.Event) ([]models.Event, error) {
	rows, err := service.Database.GetallEvents(Event.GroupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event

	for rows.Next() {
		var event models.Event
		if err := rows.Scan(utils.GetScanFields(&event)...); err != nil {
			log.Println(err)
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (S *Service) GetEventsById(Event *models.Event) (*models.Event, error) {
	row := S.Database.GetEventById(Event.ID)
	if row == nil {
		return nil, fmt.Errorf("no group found with ID: %d", Event.ID)
	}

	// Scan the row into the Event struct
	if err := row.Scan(utils.GetScanFields(Event)...); err != nil {
		return nil, fmt.Errorf("error scanning Event data: %v", err)
	}

	return Event, nil
}

func (S *Service) PostEventsOption(OptionEvent models.EventOption) (err error) {
	booll, err := S.Database.CheckEvent(OptionEvent.EventID, OptionEvent.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = S.Database.SaveOptionEvent(&OptionEvent)
			return
		}
	}
	if booll == OptionEvent.Going {
		return
	} else if booll != OptionEvent.Going {
		err = S.Database.UpdateOptionEvent(OptionEvent)
		log.Println(err)
		return
	}
	err = S.Database.SaveOptionEvent(&OptionEvent)
	return
}

func (S *Service) GetEventsOption(OptionEvent models.EventOption) ([]models.EventOption, error) {
	rows, err := S.Database.OptionEvent(OptionEvent.EventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.EventOption

	for rows.Next() {
		var event models.EventOption
		if err := rows.Scan(utils.GetScanFields(&event)...); err != nil {
			log.Println(err)
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

// GetEventGoingInfo processes event attendance data
func (s *Service) GetEventGoingInfo(eventID int, userID int) (int, int, string, error) {
	goingCount, notGoingCount, userGoing, err := s.Database.GetEventCountsAndUserChoice(eventID, userID)
	if err != nil {
		return 0, 0, "", err
	}

	var action string
	if userGoing.Valid {
		if userGoing.Bool {
			action = "going"
		} else {
			action = "not going"
		}
	} else {
		action = "undecided"
	}

	return goingCount, notGoingCount, action, nil
}
