// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package client

// ClientOption is an option used to configure the Runtime Client
type ClientOption interface {
	apply(*clientOptions)
}
