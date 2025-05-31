/*
Copyright 2025 ptrvsrg.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
	"resty.dev/v3"
)

type Option func(*resty.Client) error

func New(opts ...Option) (*resty.Client, error) {
	client := resty.New().SetLogger(newLogger())

	transport, ok := client.Transport().(*http.Transport)
	if ok {
		if err := http2.ConfigureTransport(transport); err == nil {
			client.SetTransport(transport)
		}
	}

	for _, opt := range opts {
		if err := opt(client); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}

	return client, nil
}
