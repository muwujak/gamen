package services

import (
	"github.com/mujak27/gamen/src/core/internal/api/dto"
	repositoryInterfaces "github.com/mujak27/gamen/src/core/internal/interfaces/repository"
	serviceInterfaces "github.com/mujak27/gamen/src/core/internal/interfaces/services"
)

type WidgetService struct {
	catalogueService     serviceInterfaces.CatalogueService
	configurationService serviceInterfaces.ConfigurationService
	widgetRepository     repositoryInterfaces.WidgetRepository
}

func NewWidgetService(catalogueService serviceInterfaces.CatalogueService, configurationService serviceInterfaces.ConfigurationService, widgetRepository repositoryInterfaces.WidgetRepository) *WidgetService {
	return &WidgetService{catalogueService: catalogueService, configurationService: configurationService, widgetRepository: widgetRepository}
}

func (s *WidgetService) Action(payload dto.WidgetActionPayload) error {

	widget, err := s.widgetRepository.GetWidgetById(payload.WidgetID)
	if err != nil {
		return err
	}
	plugin, err := s.catalogueService.GetPluginFunctionById(widget.PluginID)
	if err != nil {
		return err
	}
	err = plugin.Action(dto.PluginActionSchema{
		Configuration: widget.Configuration,
		RawPluginData: payload.Data,
	})
	if err != nil {
		return err
	}
	return nil
}

// func (s *WidgetService) Fetch(payload dto.WidgetActionPayload) error {

// 	widget, err := s.widgetRepository.GetWidgetById(payload.WidgetID)
// 	if err != nil {
// 		return err
// 	}
// 	plugin, err := s.catalogueService.GetPluginFunctionById(widget.PluginID.String())
// 	if err != nil {
// 		return err
// 	}
// 	plugin.Fetch(payload.Data)
// 	return nil
// }
