/*
 *  Copyright (c) 2022 NetEase Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

/*
 * Project: CurveCli
 * Created Date: 2022-08-25
 * Author: chengyi (Cyber-SiKu)
 */

package list

import (
	basecmd "github.com/opencurve/curve/tools-v2/pkg/cli/command"
	logicalpool "github.com/opencurve/curve/tools-v2/pkg/cli/command/curvebs/list/logicalPool"
	"github.com/spf13/cobra"
)

type ListCommand struct {
	basecmd.MidCurveCmd
}

var _ basecmd.MidCurveCmdFunc = (*ListCommand)(nil) // check interface

func (listCmd *ListCommand) AddSubCommands() {
	listCmd.Cmd.AddCommand(
		logicalpool.NewLogicalPoolCommand(),
	)
}

func NewListCommand() *cobra.Command {
	listCmd := &ListCommand{
		basecmd.MidCurveCmd{
			Use:   "list",
			Short: "list resources in the curvebs",
		},
	}
	return basecmd.NewMidCurveCli(&listCmd.MidCurveCmd, listCmd)
}
