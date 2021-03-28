// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package client

import (
	"go.uber.org/zap"
)

// WithLoggingOptions adds options to use when creating a zap.Logger for logging in the Client
func WithLoggingOptions(options ...zap.Option) ClientOption {
	return newFuncClientOption(func(o *clientOptions) {
		o.loggerOptions = append(o.loggerOptions, options...)
	})
}
