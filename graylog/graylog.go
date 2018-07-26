package graylog

import (
	"github.com/aphistic/golf"
	"github.com/rs/zerolog"
)

//GraylogHook hook for graylog
type GraylogHook struct {
	logger *golf.Logger
	client *golf.Client
}

//NewGraylogHook return new hook
func NewGraylogHook(url string) (GraylogHook, error) {
	hook := GraylogHook{}
	c, err := golf.NewClient()
	if err != nil {
		return hook, err
	}

	err = c.Dial(url)
	if err != nil {
		return hook, err
	}

	l, err := c.NewLogger()
	if err != nil {
		return hook, err
	}

	return GraylogHook{
		logger: l,
		client: c,
	}, nil
}

//Run run hook when having a log event
func (h GraylogHook) Run(e *zerolog.Event, level zerolog.Level, message string) {
	var fields map[string]interface{}
	switch level {
	case zerolog.DebugLevel:
		h.logger.Dbgm(fields, message)
	case zerolog.InfoLevel:
		h.logger.Infom(fields, message)
	case zerolog.WarnLevel:
		h.logger.Warnm(fields, message)
	case zerolog.ErrorLevel:
		h.logger.Errm(fields, message)
	case zerolog.FatalLevel:
		h.logger.Critm(fields, message)
	}
}

//Close close graylog client
func (h GraylogHook) Close() {
	h.client.Close()
}
