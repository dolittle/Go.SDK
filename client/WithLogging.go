// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package client

import (
	"go.uber.org/zap"
)

// WithLogging sets the zap.Logger to use for logging in the Client
func WithLogging(logger *zap.Logger) ClientOption {
	return newFuncClientOption(func(o *clientOptions) {
		o.logger = logger
	})
}
