/*
   Spatula, Serving SPA's made easy.
   Open-Source, WTFPL License.

   Copyright (C) 2025-20xx Neo <ntsk@airmail.cc>
*/

package Cli

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	FSN "github.com/fsnotify/fsnotify"
	Internal "github.com/neotesk/spatula/src/internal"
	Types "github.com/neotesk/spatula/src/types"
);

var Serve = Types.Command {
    Name: "serve",
    ShortDesc: "Serves the SPA",
    LongDesc: "This command will serve the Single Page Application you've created.",
    Descriptor: "<html_path> [static_path]",
    Action: func ( args []string, flags map[ string ] bool, adjectives map[ string ] string ) {
        if len( args ) < 1 {
            Internal.ErrPrintf( "Fatal Error: 'html_path' parameter is missing.\n" );
            os.Exit( 1 );
        }
        htmlPath := args[ 0 ];
        servePath := Internal.MakeCoalesce( Internal.PossibleItem( args, 1 ), filepath.Join( htmlPath, "../" ) );
        port := adjectives[ "port" ];

        handleIndex := func ( w http.ResponseWriter, r *http.Request, path string ) {
            w.Header().Set( "Content-Type", "text/html; charset=utf-8" );
            w.Header().Set( "Cache-Control", "no-cache" );
            w.Header().Set( "Connection", "keep-alive" );
            w.Header().Set( "Access-Control-Allow-Origin", "*" );
            data, err := os.ReadFile( path );
            if err != nil {
                fmt.Fprint( w, err );
                return;
            }
            moddedData := string( data );
            modder := "</body>";
            splitted := strings.Split( moddedData, modder );
            if ( len( splitted ) > 1 ) {
                moddedData = splitted[ 0 ] + "<script>const _ev=new EventSource('http://127.0.0.1:" + port + "/events');_ev.onmessage=()=>window.location.reload()</script>" + modder + splitted[ 1 ];
            }
            http.ServeContent( w, r, ".html", time.Now(), bytes.NewReader( []byte( moddedData ) ) );
        }

        handler := func ( w http.ResponseWriter, r *http.Request ) {
            if Internal.IsDebug {
                fmt.Println( "Requested path from remote: " + r.URL.Path );
            }
            p := "." + r.URL.Path;
            x := filepath.Join( servePath, p );
            if b, _ := Internal.FileSystem.Exists( x ); b && r.URL.Path != "/" {
                http.ServeFile( w, r, x );
                return;
            }
            handleIndex( w, r, htmlPath );
        }

        if ( !flags[ "l" ] ) {
            watcher, err := FSN.NewWatcher();
            if err != nil {
                Internal.ErrPrintf( "Fatal Error: %s\n", err.Error() );
                os.Exit( 1 );
            }
            defer watcher.Close();

            go func() {
                for {
                    select {
                        case event, ok := <-watcher.Events:
                            if !ok {
                                return;
                            }
                            if event.Op != FSN.Chmod {
                                select {
                                    case Internal.SSEChannel <- "0":
                                    default:
                                }
                                if Internal.IsDebug {
                                    fmt.Println( "Change detected at", event.Name )
                                }
                            }
                        case err, ok := <-watcher.Errors:
                            if !ok {
                                return
                            }
                            fmt.Println( Internal.Colorify( "Error: " + err.Error(), "FF4040" ) )
                    }
                }
            }();

            err = watcher.Add( servePath );
            if err != nil {
                Internal.ErrPrintf( "Fatal Error: %s\n", err.Error() );
                os.Exit( 1 );
            }
            http.HandleFunc( "/events", Internal.SSEHandler );
        }

        http.HandleFunc( "/", handler );

        fmt.Printf( "Now serving at %s\n", Internal.Colorify( fmt.Sprintf( "http://127.0.0.1:%s/", port ), "ada440" ) );
        http.ListenAndServe( ":" + port, nil );
    },
};