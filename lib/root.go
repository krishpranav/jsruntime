package lib

import (
	_ "embed"
	"fmt"
	"os"
	"runtime"

	esbuild "github.com/evanw/esbuild/pkg/api"
	"github.com/lithdew/quickjs"
)

var codeGlobals string

func RunFile(fileToRun string) {
	code, errorReadingFile := os.ReadFile(fileToRun)

	JSError(errorReadingFile, true)

	parsedCode := esbuild.Transform(string(code), esbuild.TransformOptions{
		Loader: esbuild.LoaderTS,
	})

	if len(parsedCode.Errors) > 1 {
		fmt.Println("Errors:")
		for _, error := range parsedCode.Errors {
			fmt.Println(error)
		}
		os.Exit(1)
	}

	runtime.LockOSThread()

	JSRuntime := quickjs.NewRuntime()
	defer JSRuntime.Free()
	ctx := JSRuntime.NewContext()
	defer ctx.Free()

	globals := ctx.Globals()
	globals.Set("__dispatch", ctx.Function(Globals))

	k, errorInjectingGlobals := ctx.Eval(codeGlobals)
	JSError(errorInjectingGlobals, true)
	defer k.Free()

	result, error := ctx.EvalFile(string(parsedCode.Code), "s")

	JSError(error, true)

	result.Free()
}
