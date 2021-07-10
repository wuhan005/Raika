// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	log "unknwon.dev/clog/v2"

	"github.com/wuhan005/Raika/internal/config"
	"github.com/wuhan005/Raika/internal/platform"
	"github.com/wuhan005/Raika/internal/platform/aliyun"
	"github.com/wuhan005/Raika/internal/types"
)

var Platform = &cli.Command{
	Name:  "platform",
	Usage: "Manage the cloud service",
	Subcommands: []*cli.Command{
		{
			Name:   "login",
			Usage:  "Login in to a new cloud service",
			Action: loginPlatform,
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "platform", Usage: "Cloud platform name", Required: true},
				&cli.StringFlag{Name: "region-id", Usage: "Cloud platform region ID"},
				&cli.StringFlag{Name: "account-id", Usage: "Cloud platform account ID"},
				&cli.StringFlag{Name: "access-key-id", Usage: "Cloud platform access key ID"},
				&cli.StringFlag{Name: "access-key-secret", Usage: "Cloud platform access key secret"},
				&cli.StringFlag{Name: "name", Usage: "Name of this account"},
			},
		},
		{
			Name:   "list",
			Usage:  "List the current cloud service",
			Action: listPlatform,
		},
	},
	Flags: []cli.Flag{
		stringFlag("config-file", config.DefaultConfigPath, "Config file path"),
	},
}

func loginPlatform(c *cli.Context) error {
	name := c.String("name")
	configFilePath := c.String("config-file")
	regionID := c.String("region-id")
	accountID := c.String("account-id")
	accessKeyID := c.String("access-key-id")
	accessKeySecret := c.String("access-key-secret")

	var client platform.Cloud
	p := types.Platform(c.String("platform"))
	switch p {
	case types.Aliyun:
		client = aliyun.New(platform.AuthenticateOptions{
			aliyun.RegionIDField:        regionID,
			aliyun.AccountIDField:       accountID,
			aliyun.AccessKeyIDField:     accessKeyID,
			aliyun.AccessKeySecretField: accessKeySecret,
		})
	case types.TencentCloud:

	default:
		return errors.Errorf("unsupported platform: %q", p)
	}

	if client == nil {
		return errors.New("unexpected error, client is nil")
	}

	if err := client.Authenticate(); err != nil {
		return errors.Wrapf(err, "authenticate %q", p)
	}

	log.Info("Authenticate to %q succeed.", p)

	// Save the authenticate config to file.
	configFile := config.New(configFilePath)
	if err := configFile.Load(); err != nil {
		return errors.Wrap(err, "load config file")
	}

	if name == "" {
		name = string(p)
	}

	configFile.AuthConfigs[name] = types.AuthConfig{
		Platform:        p,
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
	}
	return configFile.Save()
}

func listPlatform(c *cli.Context) error {
	configFilePath := c.String("config-file")
	configFile := config.New(configFilePath)
	if err := configFile.Load(); err != nil {
		return errors.Wrap(err, "load config file")
	}

	if len(configFile.AuthConfigs) == 0 {
		log.Warn("The platform is empty. Run `Raika platform login` to create one.")
		return nil
	}

	i := 0
	for _, p := range configFile.AuthConfigs {
		i++
		log.Trace("%02d - [ %s ] %s", i, p.Platform, p.AccessKeyID)
	}
	return nil
}