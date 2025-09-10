/*
    Spatula, Serving SPA's made easy.
    Open-Source, WTFPL License.

    Copyright (C) 2025-20xx Neo <ntsk@airmail.cc>
*/

package main

import (
    Cli "github.com/neotesk/spatula/src/cli"
	Internal "github.com/neotesk/spatula/src/internal"
	Types "github.com/neotesk/spatula/src/types"
)

func main () {
    // Setup default argument parameters
    defaultArgs := Types.DefaultArgs {
        Flags: []Types.Flag {
            {
                Name: "d",
                ShortDesc: "Enables Debug Mode",
                DefaultValue: false,
            },
            {
                Name: "l",
                ShortDesc: "Disables Live Reload Mode",
                DefaultValue: false,
            },
            {
                Name: "v",
                ShortDesc: "Prints version",
                DefaultValue: false,
            },
        },
        Arguments: []Types.Argument {
            {
                Name: "port",
                ShortDesc: "Override port number",
                DefaultValue: "8080",
            },
        },
    };

    // After setting the default arguments, let's
    // feed them inside the Arguments function
    // so we can get a good output of what we have
    // in our hands.
    argsList := Arguments( defaultArgs );

    // After gathering the output, we will
    // create a program
    program := Types.Program {
        Name: "spatula",
        Desc: "Spatula is an SPA tool to serve your Single Page Applications for testing/debugging.",
        Footer: "",
        DefaultArgs: defaultArgs,
        Commands: Cli.Commands,
    };

    Internal.IsDebug = argsList.Flags[ "d" ];

    // Now we run the program.
    RunProgram( program, argsList );
}