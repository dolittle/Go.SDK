// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package client

import (
	"go.uber.org/zap"
)

// Client is the main entrypoint for interacting with the Dolittle Runtime
type Client interface {
	SayHello()
}

// NewClient builds a new Client with the provided ClientOptions.
// Most options have a sensible default value.
func NewClient(options ...ClientOption) (Client, error) {
	opts := clientOptions{
		logger:            nil,
		loggerOptions:     nil,
		grpcRuntimeTarget: "localhost:50053",
	}

	for _, option := range options {
		option.apply(&opts)
	}

	c := &client{}

	if err := c.buildLogger(opts.logger, opts.loggerOptions); err != nil {
		return nil, err
	}

	c.grpcRuntimeTarget = opts.grpcRuntimeTarget

	return c, nil
}

type client struct {
	logger            *zap.Logger
	grpcRuntimeTarget string
}

func (c *client) buildLogger(logger *zap.Logger, options []zap.Option) error {
	if logger != nil {
		if len(options) > 0 {
			logger.Warn("Logger provided through 'client.WithLogging', logging options provided through 'client.WithLoggingOptions' will be ignored")
		}
		c.logger = logger
		return nil
	}
	logger, err := zap.NewProduction(options...)
	c.logger = logger
	return err
}

// SayHello logs the gRPC target address currently configured
func (c *client) SayHello() {
	c.logger.Info("Hello I will connect to the runtime using " + c.grpcRuntimeTarget)
}
