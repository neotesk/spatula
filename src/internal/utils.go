/*
    Spatula, Serving SPA's made easy.
    Open-Source, WTFPL License.

    Copyright (C) 2025-20xx Neo <ntsk@airmail.cc>
*/

package Internal

import (
    "os"
)

func HandleError [ T any ] ( thing T, err error ) T {
    if err != nil {
        ErrPrintf( "Fatal Error: %s\n", err.Error() );
        os.Exit( 1 );
    }
    return thing;
}

func Make [ T any ] ( thing any ) T {
    output, ok := thing.( T );
    if !ok {
        ErrPrintf( "Fatal Error: Cannot convert object into desired type.\n" );
        os.Exit( 1 );
    }
    return output;
}

func MakeCoalesce [ T any ] ( thing any, def T ) T {
    if thing == nil {
        return def;
    }
    output, ok := thing.( T );
    if !ok {
        ErrPrintf( "Fatal Error: Cannot convert object into desired type.\n" );
        os.Exit( 1 );
    }
    return output;
}

func MakeArray [ T any ] ( thing any, def []T ) []T {
    if thing == nil {
        return def;
    }
    output, ok := thing.( []any );
    if !ok {
        ErrPrintf( "Fatal Error: Cannot convert object into desired type.\n" );
        os.Exit( 1 );
    }
    retval := []T {};
    for _, ret := range( output ) {
        typed, ok := ret.( T );
        if !ok {
            ErrPrintf( "Fatal Error: Cannot convert object into desired type.\n" );
            os.Exit( 1 );
        }
        retval = append( retval, typed );
    }
    return retval;
}

func PossibleItem [ T any ] ( arr []T, idx int ) any {
    if len( arr ) <= idx {
        return nil;
    }
    return arr[ idx ];
}