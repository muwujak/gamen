package utils

import (
	"encoding/json"

	"github.com/mujak27/gamen/src/core/internal/api/dto"
	model_catalogue "github.com/mujak27/gamen/src/core/internal/extensions/plugins/restart_deployment/dto"
)

func Transform(j interface{}, target interface{}) error {
	jsonData, err := json.Marshal(j)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonData, target); err != nil {
		return err
	}
	return nil
}

func TransformPayloadToModel(payload dto.PluginActionSchema) (model_catalogue.RestartDeploymentForm, error) {
	var data model_catalogue.RestartDeploymentForm
	if err := Transform(payload.RawPluginData, &data); err != nil {
		return model_catalogue.RestartDeploymentForm{}, err
	}
	return data, nil
}
