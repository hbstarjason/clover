// Copyright (c) Authors of Clover
//
// All rights reserved. This program and the accompanying materials
// are made available under the terms of the Apache License, Version 2.0
// which accompanies this distribution, and is available at
// http://www.apache.org/licenses/LICENSE-2.0

package cmd

import (
    "github.com/spf13/cobra"
    "cloverkube"
)

var servicesCmd = &cobra.Command{
    Use:   "services",
    Short: "Get listing of Kubernetes services",
    Long: ``,
    Run: func(cmd *cobra.Command, args []string) {
        cloverkube.GetServices()
    },
}

func init() {
    getCmd.AddCommand(servicesCmd)
}
