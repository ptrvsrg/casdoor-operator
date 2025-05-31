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

	"go.uber.org/zap"
)

type logger struct {
	l *zap.Logger
}

func newLogger() *logger {
	return &logger{
		l: zap.L().Named("resty"),
	}
}

func (l *logger) Errorf(format string, v ...any) {
	l.l.Error(fmt.Sprintf(format, v...))
}

func (l *logger) Warnf(format string, v ...any) {
	l.l.Warn(fmt.Sprintf(format, v...))
}

func (l *logger) Debugf(format string, v ...any) {
	l.l.Debug(fmt.Sprintf(format, v...))
}
