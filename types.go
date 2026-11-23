package common

import (
	"time"

	"github.com/mergezone/scope"
)

type Timestamps struct {
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type User struct {
	Name string `json:"name"`
}

type Comment struct {
	Content string `json:"content"`

	Timestamps
	Author User `json:"author"`
}

type Review struct {
	Comment
	IsApproved bool `json:"is_approved"`
}

type WorkItem struct {
	ID int `json:"id"`
	URL string `json:"url"`
	Title string `json:"title"`
	Description string `json:"description"`

	Timestamps
	ClosedAt *time.Time `json:"closed_at"`
	CompletedAt *time.Time `json:"completed_at"`

	Owner User `json:"owner"`
	Assignees []User `json:"assignees"`
	Reviewers []User `json:"reviewers"`

	Discussion []Comment `json:"discussion"`
	Reviews []Review `json:"reviews"`
}

func (w WorkItem) Paths() []*scope.Path {
	return scope.ExtractPaths(w.Title)
}
