// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package client

type funcClientOption struct {
	callback func(*clientOptions)
}

func newFuncClientOption(callback func(*clientOptions)) ClientOption {
	return &funcClientOption{
		callback: callback,
	}
}

func (f *funcClientOption) apply(options *clientOptions) {
	f.callback(options)
}
