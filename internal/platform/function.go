// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package platform

import (
	"time"
)

type CreateFunctionOptions struct {
	Name                  string
	Description           string
	MemorySize            int64
	Environment           map[string]string
	InitializationTimeout time.Duration
	RuntimeTimeout        time.Duration
	HTTPPort              int
	File                  string
}
