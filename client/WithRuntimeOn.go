// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package client

import (
	"fmt"
)

// WithRuntimeOn sets the host address and port to use for connecting to a Runtime
func WithRuntimeOn(host string, port int) ClientOption {
	return newFuncClientOption(func(o *clientOptions) {
		o.grpcRuntimeTarget = fmt.Sprintf("%s:%d", host, port)
	})
}
