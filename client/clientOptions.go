// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package client

import (
	"go.uber.org/zap"
)

type clientOptions struct {
	logger        *zap.Logger
	loggerOptions []zap.Option

	grpcRuntimeTarget string
}
