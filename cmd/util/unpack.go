// Copyright © 2021 Ettore Di Giacinto <mudler@mocaccino.org>
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program; if not, see <http://www.gnu.org/licenses/>.

package util

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/docker/go-units"
	"github.com/mudler/luet/pkg/config"
	"github.com/mudler/luet/pkg/helpers"
	. "github.com/mudler/luet/pkg/logger"

	"github.com/spf13/cobra"
)

func NewUnpackCommand() *cobra.Command {

	return &cobra.Command{
		Use:   "unpack image path",
		Short: "Unpack a docker image natively",
		Long: `unpack doesn't need the docker daemon to run, and unpacks a docker image in the specified directory:
		
	luet util unpack golang:alpine /alpine
`,
		PreRun: func(cmd *cobra.Command, args []string) {

			if len(args) != 2 {
				Fatal("Expects an image and a path")
			}

		},
		Run: func(cmd *cobra.Command, args []string) {

			image := args[0]
			destination, err := filepath.Abs(args[1])
			if err != nil {
				Error("Invalid path %s", destination)
				os.Exit(1)
			}

			temp, err := config.LuetCfg.GetSystem().TempDir("contentstore")
			if err != nil {
				Fatal("Cannot create a tempdir", err.Error())
			}

			Info("Downloading", image, "to", destination)
			info, err := helpers.DownloadAndExtractDockerImage(temp, image, destination)
			if err != nil {
				Error(err.Error())
				os.Exit(1)
			}
			Info(fmt.Sprintf("Pulled: %s", info.Target.Digest))
			Info(fmt.Sprintf("Size: %s", units.BytesSize(float64(info.ContentSize))))
		},
	}

}
