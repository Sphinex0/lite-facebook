package service

import (
	"fmt"

	"social-network/internal/models"
	utils "social-network/pkg"
)

func (service *Service) CreateEvent(Events models.Event) (err error) {
	valid, err := service.Database.GetCreatorGroup(Events.GroupID, Events.UserID)
	if err != nil {
		return
	}
	if valid {
		err = service.Database.SaveEvent(&Events)
	}
	return
}

func (service *Service) AllEvents() ([]models.Event, error) {
	rows, err := service.Database.GetallEvents()
    if err != nil {
        return nil, err
    }
    defer rows.Close()


    var events []models.Event

    for rows.Next() {
        var event models.Event
        if err := rows.Scan(utils.GetScanFields(&event)); err != nil {
            fmt.Println(err)
            return nil, err
        }
        events = append(events, event)
    }
    return events, nil
}



func (S *Service) GetEventsById(Event *models.Event) (*models.Event, error) {
	row := S.Database.GetEventById(Event.ID)
	if row == nil {
		return nil, fmt.Errorf("no group found with ID: %s", Event.ID)
	}

	// Scan the row into the Event struct
	if err := row.Scan(utils.GetScanFields(&Event)); err != nil {
		return nil, fmt.Errorf("error scanning Event data: %v", err)
	}

	return Event, nil
}

func (S *Service) PostEventsOption(OptionEvent models.EventOption) (err error) {
	err = S.Database.SaveOptionEvent(&OptionEvent) 
	return
}


func (S *Service) GetEventsOption(OptionEvent models.EventOption) ([]models.EventOption, error) {
	rows ,err:= S.Database.OptionEvent(OptionEvent.EventID) 
	if err != nil {
        return nil, err
    }
    defer rows.Close()


    var events []models.EventOption

    for rows.Next() {
        var event models.EventOption
        if err := rows.Scan(utils.GetScanFields(&event)); err != nil {
            fmt.Println(err)
            return nil, err
        }
        events = append(events, event)
    }
    return events, nil
}
