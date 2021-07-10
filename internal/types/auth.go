// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package types

// AuthConfig contains authorization information for connecting to a cloud service.
type AuthConfig struct {
	Platform        Platform `json:"platform"`
	AccessKeyID     string   `json:"access_key_id,omitempty"`
	AccessKeySecret string   `json:"access_key_secret,omitempty"`
}