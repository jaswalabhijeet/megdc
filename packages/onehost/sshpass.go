/*
** Copyright [2013-2015] [Megam Systems]
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
 */
package onehost

import (
	"github.com/megamsys/libgo/cmd"
	"github.com/megamsys/megdc/handler"
//		"github.com/megamsys/megdc/packages"
	"launchpad.net/gnuflag"

)
var SSH_PASS = []string{"SshPass"}

type Sshpass struct {
	Fs       *gnuflag.FlagSet
	HOST string
}


func (g *Sshpass) Info() *cmd.Info {
	return &cmd.Info{
		Name:    "sshpass",
		Usage:   `sshpass [--help/-h] ...`,
		Desc:    `Pass opennebula frontend sshpub key to opennebula host.
    The [[--host]] parameter defines the opennebulahost ip address
		Default  is localhost
	`,
		MinArgs: 0,
	}
}

func (c *Sshpass) Run(context *cmd.Context) error {
	handler.FunSpin(cmd.Colorfy(handler.Logo, "green", "", "bold"), "", "auth")
	w := handler.NewWrap(c)
	w.IfNoneAddPackages(SSH_PASS)
	if h, err := handler.NewHandler(w); err != nil {
		return err
	} else if err := h.Run(); err != nil {
		return err
	}
	return nil
}

func (c *Sshpass) Flags() *gnuflag.FlagSet {
	if c.Fs == nil {
		c.Fs = gnuflag.NewFlagSet("sshpass", gnuflag.ExitOnError)
		hostMsg := "The hostname or ip address of the remote opennebula host"
		c.Fs.StringVar(&c.HOST, "host", "localhost", hostMsg)
	}
	return c.Fs
}
