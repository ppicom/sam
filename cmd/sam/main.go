/*
Copyright © 2021 Pere Jover

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
	cfg2 "github.com/pjover/sam/internal"
	"github.com/pjover/sam/internal/adapters/cfg"
	"github.com/pjover/sam/internal/adapters/cli"
)

func main() {
	configService := cfg.NewConfigService()
	cmdManager := cli.NewCmdManager(configService)
	cfg2.MainDI(cmdManager, configService)
	cmdManager.Execute()
}
