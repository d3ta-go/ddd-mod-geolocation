package indexer

import (
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/indexer"
)

// BaseIndexerSvc represent BaseIndexerSvc
type BaseIndexerSvc struct {
	handler *handler.Handler
	index   string
	indexer *indexer.Indexer

	dbConnectionName string
}

// SetHandler set Handler
func (s *BaseIndexerSvc) SetHandler(h *handler.Handler) {
	s.handler = h
}

// GetHandler set Handler
func (s *BaseIndexerSvc) GetHandler() *handler.Handler {
	return s.handler
}

// SetDBConnectionName set DBConnectionName
func (s *BaseIndexerSvc) SetDBConnectionName(v string) {
	s.dbConnectionName = v
}

// SetIndexer set Indexer
func (s *BaseIndexerSvc) SetIndexer(context, container, component string) (*indexer.Indexer, error) {
	cfg, err := s.handler.GetDefaultConfig()
	if err != nil {
		return nil, err
	}

	ce, err := s.handler.GetIndexer(cfg.Indexers.DataIndexer.ConnectionName)
	if err != nil {
		return nil, err
	}
	ce.Context = context
	ce.Container = container
	ce.Component = component

	s.indexer = ce

	return ce, nil
}
