package task

import (
	"errors"
	"net/http"
	"time"
)

type Request struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Status      string    `json:"status" binding:"required"`
	Priority    string    `json:"priority"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date" binding:"required"`
}

func (s *Request) Bind(r *http.Request) error {
	if s.Title == "" {
		return errors.New("title: cannot be blank")
	}
	if s.Description == "" {
		return errors.New("description: cannot be blank")
	}
	if s.Status == "" {
		return errors.New("status: cannot be blank")
	}
	if s.Priority == "" {
		return errors.New("priority: cannot be blank")
	}
	if s.StartDate.IsZero() {
		return errors.New("start_date: cannot be blank")
	}
	if s.EndDate.IsZero() {
		return errors.New("end_date: cannot be blank")
	}
	if s.EndDate.Before(s.StartDate) {
		return errors.New("end_date: must be after start_date")
	}
	return nil
}

type Response struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

func ParseFromEntity(data Entity) (res Response) {
	res = Response{
		ID:          data.ID,
		Title:       *data.Title,
		Description: *data.Description,
		Status:      *data.Status,
		Priority:    *data.Priority,
		StartDate:   data.StartDate,
		EndDate:     data.EndDate,
	}
	return
}

func ParseFromEntities(data []Entity) (res []Response) {
	res = make([]Response, 0)
	for _, object := range data {
		res = append(res, ParseFromEntity(object))
	}
	return
}
