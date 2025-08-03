package dto

import "github.com/mujak27/gamen/src/core/internal/models"

type PluginFetchSchema struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Values []string `json:"values"`
}

type PluginActionSchema struct {
	Configuration models.Configuration   `json:"configuration"`
	RawPluginData map[string]interface{} `json:"raw_plugin_data"`
}

type FormFields struct {
	Label      string `json:"label"`
	IsRequired bool   `json:"isRequired"`
	Type       string `json:"type"`
	Value      string `json:"value"`
}

type PluginUISchemaForm struct {
	Title  string       `json:"title"`
	Fields []FormFields `json:"fields"`
}

type PluginUISchemaGraph struct {
	Title string `json:"title"`
	// Fields []PluginItemGraphField `json:"fields"`
}
