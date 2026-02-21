package domain

import (
	"time"
	"github.com/sashabaranov/go-openai"
	"github.com/shopspring/decimal"
)

// Rebranding all types implicitly by keeping their names but putting them in the sophon-core module.

type Org struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	IsTrial            bool   `json:"isTrial"`
	AutoAddDomainUsers bool   `json:"autoAddDomainUsers"`
	IntegratedModelsMode bool                `json:"integratedModelsMode,omitempty"`
}

type User struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
}

type Project struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Plan struct {
	Id              string      `json:"id"`
	OwnerId         string      `json:"ownerId"`
	ProjectId       string      `json:"projectId"`
	Name            string      `json:"name"`
	CreatedAt       time.Time   `json:"createdAt"`
	UpdatedAt       time.Time   `json:"updatedAt"`
}

type ContextType string

const (
	ContextFileType          ContextType = "file"
	ContextURLType           ContextType = "url"
	ContextNoteType          ContextType = "note"
	ContextDirectoryTreeType ContextType = "directory tree"
	ContextPipedDataType     ContextType = "piped data"
	ContextImageType         ContextType = "image"
	ContextMapType           ContextType = "map"
)

type Context struct {
	Id              string                `json:"id"`
	ContextType     ContextType           `json:"contextType"`
	Name            string                `json:"name"`
	FilePath        string                `json:"file_path"`
	Body            string                `json:"body,omitempty"`
}

type ConvoMessage struct {
	Id               string            `json:"id"`
	UserId           string            `json:"userId"`
	Role             string            `json:"role"`
	Message          string            `json:"message"`
	CreatedAt        time.Time         `json:"createdAt"`
}
