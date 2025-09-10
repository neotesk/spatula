/*
   Spatula, Serving SPA's made easy.
   Open-Source, WTFPL License.

   Copyright (C) 2025-20xx Neo <ntsk@airmail.cc>
*/

package Internal

import (
    "fmt"
    "net/http"
)

var SSEChannel = make( chan string );

func SSEHandler ( w http.ResponseWriter, r *http.Request ) {
    w.Header().Set( "Content-Type", "text/event-stream" );
    w.Header().Set( "Cache-Control", "no-cache" );
    w.Header().Set( "Connection", "keep-alive" );
    w.Header().Set( "Access-Control-Allow-Origin", "*" );

    gone := r.Context().Done();
    respc := http.NewResponseController( w );

    for {
        select {
            case <-gone:
                return
            case <-SSEChannel:
                _, err := fmt.Fprintf( w, "data: refresh\n\n" )
                if err != nil {
                    return
                }
                err = respc.Flush()
                if err != nil {
                    return
                }
        }
    }
}