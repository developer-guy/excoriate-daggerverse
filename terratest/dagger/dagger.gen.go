// Code generated by dagger. DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"Terratest/internal/dagger"
	"Terratest/internal/telemetry"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

var dag = dagger.Connect()

func Tracer() trace.Tracer {
	return otel.Tracer("dagger.io/sdk.go")
}

// used for local MarshalJSON implementations
var marshalCtx = context.Background()

// called by main()
func setMarshalContext(ctx context.Context) {
	marshalCtx = ctx
	dagger.SetMarshalContext(ctx)
}

type DaggerObject = dagger.DaggerObject

type ExecError = dagger.ExecError

// The `CacheVolumeID` scalar type represents an identifier for an object of type CacheVolume.
type CacheVolumeID = dagger.CacheVolumeID

// The `ContainerID` scalar type represents an identifier for an object of type Container.
type ContainerID = dagger.ContainerID

// The `CurrentModuleID` scalar type represents an identifier for an object of type CurrentModule.
type CurrentModuleID = dagger.CurrentModuleID

// The `DirectoryID` scalar type represents an identifier for an object of type Directory.
type DirectoryID = dagger.DirectoryID

// The `EnvVariableID` scalar type represents an identifier for an object of type EnvVariable.
type EnvVariableID = dagger.EnvVariableID

// The `FieldTypeDefID` scalar type represents an identifier for an object of type FieldTypeDef.
type FieldTypeDefID = dagger.FieldTypeDefID

// The `FileID` scalar type represents an identifier for an object of type File.
type FileID = dagger.FileID

// The `FunctionArgID` scalar type represents an identifier for an object of type FunctionArg.
type FunctionArgID = dagger.FunctionArgID

// The `FunctionCallArgValueID` scalar type represents an identifier for an object of type FunctionCallArgValue.
type FunctionCallArgValueID = dagger.FunctionCallArgValueID

// The `FunctionCallID` scalar type represents an identifier for an object of type FunctionCall.
type FunctionCallID = dagger.FunctionCallID

// The `FunctionID` scalar type represents an identifier for an object of type Function.
type FunctionID = dagger.FunctionID

// The `GeneratedCodeID` scalar type represents an identifier for an object of type GeneratedCode.
type GeneratedCodeID = dagger.GeneratedCodeID

// The `GitModuleSourceID` scalar type represents an identifier for an object of type GitModuleSource.
type GitModuleSourceID = dagger.GitModuleSourceID

// The `GitRefID` scalar type represents an identifier for an object of type GitRef.
type GitRefID = dagger.GitRefID

// The `GitRepositoryID` scalar type represents an identifier for an object of type GitRepository.
type GitRepositoryID = dagger.GitRepositoryID

// The `InputTypeDefID` scalar type represents an identifier for an object of type InputTypeDef.
type InputTypeDefID = dagger.InputTypeDefID

// The `InterfaceTypeDefID` scalar type represents an identifier for an object of type InterfaceTypeDef.
type InterfaceTypeDefID = dagger.InterfaceTypeDefID

// An arbitrary JSON-encoded value.
type JSON = dagger.JSON

// The `LabelID` scalar type represents an identifier for an object of type Label.
type LabelID = dagger.LabelID

// The `ListTypeDefID` scalar type represents an identifier for an object of type ListTypeDef.
type ListTypeDefID = dagger.ListTypeDefID

// The `LocalModuleSourceID` scalar type represents an identifier for an object of type LocalModuleSource.
type LocalModuleSourceID = dagger.LocalModuleSourceID

// The `ModuleDependencyID` scalar type represents an identifier for an object of type ModuleDependency.
type ModuleDependencyID = dagger.ModuleDependencyID

// The `ModuleID` scalar type represents an identifier for an object of type Module.
type ModuleID = dagger.ModuleID

// The `ModuleSourceID` scalar type represents an identifier for an object of type ModuleSource.
type ModuleSourceID = dagger.ModuleSourceID

// The `ModuleSourceViewID` scalar type represents an identifier for an object of type ModuleSourceView.
type ModuleSourceViewID = dagger.ModuleSourceViewID

// The `ObjectTypeDefID` scalar type represents an identifier for an object of type ObjectTypeDef.
type ObjectTypeDefID = dagger.ObjectTypeDefID

// The platform config OS and architecture in a Container.
//
// The format is [os]/[platform]/[version] (e.g., "darwin/arm64/v7", "windows/amd64", "linux/arm64").
type Platform = dagger.Platform

// The `PortID` scalar type represents an identifier for an object of type Port.
type PortID = dagger.PortID

// The `SecretID` scalar type represents an identifier for an object of type Secret.
type SecretID = dagger.SecretID

// The `ServiceID` scalar type represents an identifier for an object of type Service.
type ServiceID = dagger.ServiceID

// The `SocketID` scalar type represents an identifier for an object of type Socket.
type SocketID = dagger.SocketID

// The `TerminalID` scalar type represents an identifier for an object of type Terminal.
type TerminalID = dagger.TerminalID

// The `TypeDefID` scalar type represents an identifier for an object of type TypeDef.
type TypeDefID = dagger.TypeDefID

// The absence of a value.
//
// A Null Void is used as a placeholder for resolvers that do not return anything.
type Void = dagger.Void

// Key value object that represents a build argument.
type BuildArg = dagger.BuildArg

// Key value object that represents a pipeline label.
type PipelineLabel = dagger.PipelineLabel

// Port forwarding rules for tunneling network traffic.
type PortForward = dagger.PortForward

// A directory whose contents persist across runs.
type CacheVolume = dagger.CacheVolume

// An OCI-compatible container, also known as a Docker container.
type Container = dagger.Container

type WithContainerFunc = dagger.WithContainerFunc

// ContainerAsTarballOpts contains options for Container.AsTarball
type ContainerAsTarballOpts = dagger.ContainerAsTarballOpts

// ContainerBuildOpts contains options for Container.Build
type ContainerBuildOpts = dagger.ContainerBuildOpts

// ContainerExportOpts contains options for Container.Export
type ContainerExportOpts = dagger.ContainerExportOpts

// ContainerImportOpts contains options for Container.Import
type ContainerImportOpts = dagger.ContainerImportOpts

// ContainerPipelineOpts contains options for Container.Pipeline
type ContainerPipelineOpts = dagger.ContainerPipelineOpts

// ContainerPublishOpts contains options for Container.Publish
type ContainerPublishOpts = dagger.ContainerPublishOpts

// ContainerTerminalOpts contains options for Container.Terminal
type ContainerTerminalOpts = dagger.ContainerTerminalOpts

// ContainerWithDefaultTerminalCmdOpts contains options for Container.WithDefaultTerminalCmd
type ContainerWithDefaultTerminalCmdOpts = dagger.ContainerWithDefaultTerminalCmdOpts

// ContainerWithDirectoryOpts contains options for Container.WithDirectory
type ContainerWithDirectoryOpts = dagger.ContainerWithDirectoryOpts

// ContainerWithEntrypointOpts contains options for Container.WithEntrypoint
type ContainerWithEntrypointOpts = dagger.ContainerWithEntrypointOpts

// ContainerWithEnvVariableOpts contains options for Container.WithEnvVariable
type ContainerWithEnvVariableOpts = dagger.ContainerWithEnvVariableOpts

// ContainerWithExecOpts contains options for Container.WithExec
type ContainerWithExecOpts = dagger.ContainerWithExecOpts

// ContainerWithExposedPortOpts contains options for Container.WithExposedPort
type ContainerWithExposedPortOpts = dagger.ContainerWithExposedPortOpts

// ContainerWithFileOpts contains options for Container.WithFile
type ContainerWithFileOpts = dagger.ContainerWithFileOpts

// ContainerWithFilesOpts contains options for Container.WithFiles
type ContainerWithFilesOpts = dagger.ContainerWithFilesOpts

// ContainerWithMountedCacheOpts contains options for Container.WithMountedCache
type ContainerWithMountedCacheOpts = dagger.ContainerWithMountedCacheOpts

// ContainerWithMountedDirectoryOpts contains options for Container.WithMountedDirectory
type ContainerWithMountedDirectoryOpts = dagger.ContainerWithMountedDirectoryOpts

// ContainerWithMountedFileOpts contains options for Container.WithMountedFile
type ContainerWithMountedFileOpts = dagger.ContainerWithMountedFileOpts

// ContainerWithMountedSecretOpts contains options for Container.WithMountedSecret
type ContainerWithMountedSecretOpts = dagger.ContainerWithMountedSecretOpts

// ContainerWithNewFileOpts contains options for Container.WithNewFile
type ContainerWithNewFileOpts = dagger.ContainerWithNewFileOpts

// ContainerWithUnixSocketOpts contains options for Container.WithUnixSocket
type ContainerWithUnixSocketOpts = dagger.ContainerWithUnixSocketOpts

// ContainerWithoutEntrypointOpts contains options for Container.WithoutEntrypoint
type ContainerWithoutEntrypointOpts = dagger.ContainerWithoutEntrypointOpts

// ContainerWithoutExposedPortOpts contains options for Container.WithoutExposedPort
type ContainerWithoutExposedPortOpts = dagger.ContainerWithoutExposedPortOpts

// Reflective module API provided to functions at runtime.
type CurrentModule = dagger.CurrentModule

// CurrentModuleWorkdirOpts contains options for CurrentModule.Workdir
type CurrentModuleWorkdirOpts = dagger.CurrentModuleWorkdirOpts

// A directory.
type Directory = dagger.Directory

type WithDirectoryFunc = dagger.WithDirectoryFunc

// DirectoryAsModuleOpts contains options for Directory.AsModule
type DirectoryAsModuleOpts = dagger.DirectoryAsModuleOpts

// DirectoryDockerBuildOpts contains options for Directory.DockerBuild
type DirectoryDockerBuildOpts = dagger.DirectoryDockerBuildOpts

// DirectoryEntriesOpts contains options for Directory.Entries
type DirectoryEntriesOpts = dagger.DirectoryEntriesOpts

// DirectoryExportOpts contains options for Directory.Export
type DirectoryExportOpts = dagger.DirectoryExportOpts

// DirectoryPipelineOpts contains options for Directory.Pipeline
type DirectoryPipelineOpts = dagger.DirectoryPipelineOpts

// DirectoryWithDirectoryOpts contains options for Directory.WithDirectory
type DirectoryWithDirectoryOpts = dagger.DirectoryWithDirectoryOpts

// DirectoryWithFileOpts contains options for Directory.WithFile
type DirectoryWithFileOpts = dagger.DirectoryWithFileOpts

// DirectoryWithFilesOpts contains options for Directory.WithFiles
type DirectoryWithFilesOpts = dagger.DirectoryWithFilesOpts

// DirectoryWithNewDirectoryOpts contains options for Directory.WithNewDirectory
type DirectoryWithNewDirectoryOpts = dagger.DirectoryWithNewDirectoryOpts

// DirectoryWithNewFileOpts contains options for Directory.WithNewFile
type DirectoryWithNewFileOpts = dagger.DirectoryWithNewFileOpts

// An environment variable name and value.
type EnvVariable = dagger.EnvVariable

// A definition of a field on a custom object defined in a Module.
//
// A field on an object has a static value, as opposed to a function on an object whose value is computed by invoking code (and can accept arguments).
type FieldTypeDef = dagger.FieldTypeDef

// A file.
type File = dagger.File

type WithFileFunc = dagger.WithFileFunc

// FileExportOpts contains options for File.Export
type FileExportOpts = dagger.FileExportOpts

// Function represents a resolver provided by a Module.
//
// A function always evaluates against a parent object and is given a set of named arguments.
type Function = dagger.Function

type WithFunctionFunc = dagger.WithFunctionFunc

// FunctionWithArgOpts contains options for Function.WithArg
type FunctionWithArgOpts = dagger.FunctionWithArgOpts

// An argument accepted by a function.
//
// This is a specification for an argument at function definition time, not an argument passed at function call time.
type FunctionArg = dagger.FunctionArg

// An active function call.
type FunctionCall = dagger.FunctionCall

// A value passed as a named argument to a function call.
type FunctionCallArgValue = dagger.FunctionCallArgValue

// The result of running an SDK's codegen.
type GeneratedCode = dagger.GeneratedCode

type WithGeneratedCodeFunc = dagger.WithGeneratedCodeFunc

// Module source originating from a git repo.
type GitModuleSource = dagger.GitModuleSource

// A git ref (tag, branch, or commit).
type GitRef = dagger.GitRef

// GitRefTreeOpts contains options for GitRef.Tree
type GitRefTreeOpts = dagger.GitRefTreeOpts

// A git repository.
type GitRepository = dagger.GitRepository

type WithGitRepositoryFunc = dagger.WithGitRepositoryFunc

// A graphql input type, which is essentially just a group of named args.
// This is currently only used to represent pre-existing usage of graphql input types
// in the core API. It is not used by user modules and shouldn't ever be as user
// module accept input objects via their id rather than graphql input types.
type InputTypeDef = dagger.InputTypeDef

// A definition of a custom interface defined in a Module.
type InterfaceTypeDef = dagger.InterfaceTypeDef

// A simple key value object that represents a label.
type Label = dagger.Label

// A definition of a list type in a Module.
type ListTypeDef = dagger.ListTypeDef

// Module source that that originates from a path locally relative to an arbitrary directory.
type LocalModuleSource = dagger.LocalModuleSource

// A Dagger module.
type Module = dagger.Module

type WithModuleFunc = dagger.WithModuleFunc

// The configuration of dependency of a module.
type ModuleDependency = dagger.ModuleDependency

// The source needed to load and run a module, along with any metadata about the source such as versions/urls/etc.
type ModuleSource = dagger.ModuleSource

type WithModuleSourceFunc = dagger.WithModuleSourceFunc

// ModuleSourceResolveDirectoryFromCallerOpts contains options for ModuleSource.ResolveDirectoryFromCaller
type ModuleSourceResolveDirectoryFromCallerOpts = dagger.ModuleSourceResolveDirectoryFromCallerOpts

// A named set of path filters that can be applied to directory arguments provided to functions.
type ModuleSourceView = dagger.ModuleSourceView

// A definition of a custom object defined in a Module.
type ObjectTypeDef = dagger.ObjectTypeDef

// A port exposed by a container.
type Port = dagger.Port

// The root of the DAG.
type Client = dagger.Client

type WithClientFunc = dagger.WithClientFunc

// ContainerOpts contains options for Client.Container
type ContainerOpts = dagger.ContainerOpts

// DirectoryOpts contains options for Client.Directory
type DirectoryOpts = dagger.DirectoryOpts

// GitOpts contains options for Client.Git
type GitOpts = dagger.GitOpts

// HTTPOpts contains options for Client.HTTP
type HTTPOpts = dagger.HTTPOpts

// ModuleDependencyOpts contains options for Client.ModuleDependency
type ModuleDependencyOpts = dagger.ModuleDependencyOpts

// ModuleSourceOpts contains options for Client.ModuleSource
type ModuleSourceOpts = dagger.ModuleSourceOpts

// PipelineOpts contains options for Client.Pipeline
type PipelineOpts = dagger.PipelineOpts

// SecretOpts contains options for Client.Secret
type SecretOpts = dagger.SecretOpts

// A reference to a secret value, which can be handled more safely than the value itself.
type Secret = dagger.Secret

// A content-addressed service providing TCP connectivity.
type Service = dagger.Service

// ServiceEndpointOpts contains options for Service.Endpoint
type ServiceEndpointOpts = dagger.ServiceEndpointOpts

// ServiceStopOpts contains options for Service.Stop
type ServiceStopOpts = dagger.ServiceStopOpts

// ServiceUpOpts contains options for Service.Up
type ServiceUpOpts = dagger.ServiceUpOpts

// A Unix or TCP/IP socket that can be mounted into a container.
type Socket = dagger.Socket

// An interactive terminal that clients can connect to.
type Terminal = dagger.Terminal

// A definition of a parameter or return type in a Module.
type TypeDef = dagger.TypeDef

type WithTypeDefFunc = dagger.WithTypeDefFunc

// TypeDefWithFieldOpts contains options for TypeDef.WithField
type TypeDefWithFieldOpts = dagger.TypeDefWithFieldOpts

// TypeDefWithInterfaceOpts contains options for TypeDef.WithInterface
type TypeDefWithInterfaceOpts = dagger.TypeDefWithInterfaceOpts

// TypeDefWithObjectOpts contains options for TypeDef.WithObject
type TypeDefWithObjectOpts = dagger.TypeDefWithObjectOpts

// Sharing mode of the cache volume.
type CacheSharingMode = dagger.CacheSharingMode

const (
	// Shares the cache volume amongst many build pipelines, but will serialize the writes
	Locked CacheSharingMode = dagger.Locked

	// Keeps a cache volume for a single build pipeline
	Private CacheSharingMode = dagger.Private

	// Shares the cache volume amongst many build pipelines
	Shared CacheSharingMode = dagger.Shared
)

// Compression algorithm to use for image layers.
type ImageLayerCompression = dagger.ImageLayerCompression

const (
	Estargz ImageLayerCompression = dagger.Estargz

	Gzip ImageLayerCompression = dagger.Gzip

	Uncompressed ImageLayerCompression = dagger.Uncompressed

	Zstd ImageLayerCompression = dagger.Zstd
)

// Mediatypes to use in published or exported image metadata.
type ImageMediaTypes = dagger.ImageMediaTypes

const (
	Dockermediatypes ImageMediaTypes = dagger.Dockermediatypes

	Ocimediatypes ImageMediaTypes = dagger.Ocimediatypes
)

// The kind of module source.
type ModuleSourceKind = dagger.ModuleSourceKind

const (
	GitSource ModuleSourceKind = dagger.GitSource

	LocalSource ModuleSourceKind = dagger.LocalSource
)

// Transport layer network protocol associated to a port.
type NetworkProtocol = dagger.NetworkProtocol

const (
	Tcp NetworkProtocol = dagger.Tcp

	Udp NetworkProtocol = dagger.Udp
)

// Distinguishes the different kinds of TypeDefs.
type TypeDefKind = dagger.TypeDefKind

const (
	// A boolean value.
	BooleanKind TypeDefKind = dagger.BooleanKind

	// A graphql input type, used only when representing the core API via TypeDefs.
	InputKind TypeDefKind = dagger.InputKind

	// An integer value.
	IntegerKind TypeDefKind = dagger.IntegerKind

	// A named type of functions that can be matched+implemented by other objects+interfaces.
	//
	// Always paired with an InterfaceTypeDef.
	InterfaceKind TypeDefKind = dagger.InterfaceKind

	// A list of values all having the same type.
	//
	// Always paired with a ListTypeDef.
	ListKind TypeDefKind = dagger.ListKind

	// A named type defined in the GraphQL schema, with fields and functions.
	//
	// Always paired with an ObjectTypeDef.
	ObjectKind TypeDefKind = dagger.ObjectKind

	// A string value.
	StringKind TypeDefKind = dagger.StringKind

	// A special kind used to signify that no value is returned.
	//
	// This is used for functions that have no return value. The outer TypeDef specifying this Kind is always Optional, as the Void is never actually represented.
	VoidKind TypeDefKind = dagger.VoidKind
)

// ptr returns a pointer to the given value.
func ptr[T any](v T) *T {
	return &v
}

// convertSlice converts a slice of one type to a slice of another type using a
// converter function
func convertSlice[I any, O any](in []I, f func(I) O) []O {
	out := make([]O, len(in))
	for i, v := range in {
		out[i] = f(v)
	}
	return out
}

func (m Terratest) MarshalJSON() ([]byte, error) {
	var concrete struct {
		Version   string
		TfVersion string
		Image     string
		Src       *Directory
		Ctr       *Container
	}
	concrete.Version = m.Version
	concrete.TfVersion = m.TfVersion
	concrete.Image = m.Image
	concrete.Src = m.Src
	concrete.Ctr = m.Ctr
	return json.Marshal(&concrete)
}

func (m *Terratest) UnmarshalJSON(bs []byte) error {
	var concrete struct {
		Version   string
		TfVersion string
		Image     string
		Src       *Directory
		Ctr       *Container
	}
	err := json.Unmarshal(bs, &concrete)
	if err != nil {
		return err
	}
	m.Version = concrete.Version
	m.TfVersion = concrete.TfVersion
	m.Image = concrete.Image
	m.Src = concrete.Src
	m.Ctr = concrete.Ctr
	return nil
}

func main() {
	ctx := context.Background()

	// Direct slog to the new stderr. This is only for dev time debugging, and
	// runtime errors/warnings.
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	})))

	if err := dispatch(ctx); err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}

func dispatch(ctx context.Context) error {
	ctx = telemetry.InitEmbedded(ctx, resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("dagger-go-sdk"),
		// TODO version?
	))
	defer telemetry.Close()

	ctx, span := Tracer().Start(ctx, "Go runtime",
		trace.WithAttributes(
			// In effect, the following two attributes hide the exec /runtime span.
			//
			// Replace the parent span,
			attribute.Bool("dagger.io/ui.mask", true),
			// and only show our children.
			attribute.Bool("dagger.io/ui.passthrough", true),
		))
	defer span.End()

	// A lot of the "work" actually happens when we're marshalling the return
	// value, which entails getting object IDs, which happens in MarshalJSON,
	// which has no ctx argument, so we use this lovely global variable.
	setMarshalContext(ctx)

	fnCall := dag.CurrentFunctionCall()
	parentName, err := fnCall.ParentName(ctx)
	if err != nil {
		return fmt.Errorf("get parent name: %w", err)
	}
	fnName, err := fnCall.Name(ctx)
	if err != nil {
		return fmt.Errorf("get fn name: %w", err)
	}
	parentJson, err := fnCall.Parent(ctx)
	if err != nil {
		return fmt.Errorf("get fn parent: %w", err)
	}
	fnArgs, err := fnCall.InputArgs(ctx)
	if err != nil {
		return fmt.Errorf("get fn args: %w", err)
	}

	inputArgs := map[string][]byte{}
	for _, fnArg := range fnArgs {
		argName, err := fnArg.Name(ctx)
		if err != nil {
			return fmt.Errorf("get fn arg name: %w", err)
		}
		argValue, err := fnArg.Value(ctx)
		if err != nil {
			return fmt.Errorf("get fn arg value: %w", err)
		}
		inputArgs[argName] = []byte(argValue)
	}

	result, err := invoke(ctx, []byte(parentJson), parentName, fnName, inputArgs)
	if err != nil {
		return fmt.Errorf("invoke: %w", err)
	}
	resultBytes, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	_, err = fnCall.ReturnValue(ctx, JSON(resultBytes))
	if err != nil {
		return fmt.Errorf("store return value: %w", err)
	}
	return nil
}

func invoke(ctx context.Context, parentJSON []byte, parentName string, fnName string, inputArgs map[string][]byte) (_ any, err error) {
	switch parentName {
	case "Terratest":
		switch fnName {
		case "Base":
			var parent Terratest
			err = json.Unmarshal(parentJSON, &parent)
			if err != nil {
				panic(fmt.Errorf("%s: %w", "failed to unmarshal parent object", err))
			}
			var goVersion string
			if inputArgs["goVersion"] != nil {
				err = json.Unmarshal([]byte(inputArgs["goVersion"]), &goVersion)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg goVersion", err))
				}
			}
			var tfVersion string
			if inputArgs["tfVersion"] != nil {
				err = json.Unmarshal([]byte(inputArgs["tfVersion"]), &tfVersion)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg tfVersion", err))
				}
			}
			return (*Terratest).Base(&parent, goVersion, tfVersion), nil
		case "WithModule":
			var parent Terratest
			err = json.Unmarshal(parentJSON, &parent)
			if err != nil {
				panic(fmt.Errorf("%s: %w", "failed to unmarshal parent object", err))
			}
			var src *Directory
			if inputArgs["src"] != nil {
				err = json.Unmarshal([]byte(inputArgs["src"]), &src)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg src", err))
				}
			}
			return (*Terratest).WithModule(&parent, src), nil
		case "WithContainer":
			var parent Terratest
			err = json.Unmarshal(parentJSON, &parent)
			if err != nil {
				panic(fmt.Errorf("%s: %w", "failed to unmarshal parent object", err))
			}
			var ctr *Container
			if inputArgs["ctr"] != nil {
				err = json.Unmarshal([]byte(inputArgs["ctr"]), &ctr)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg ctr", err))
				}
			}
			return (*Terratest).WithContainer(&parent, ctr), nil
		case "Run":
			var parent Terratest
			err = json.Unmarshal(parentJSON, &parent)
			if err != nil {
				panic(fmt.Errorf("%s: %w", "failed to unmarshal parent object", err))
			}
			var testDir string
			if inputArgs["testDir"] != nil {
				err = json.Unmarshal([]byte(inputArgs["testDir"]), &testDir)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg testDir", err))
				}
			}
			var args string
			if inputArgs["args"] != nil {
				err = json.Unmarshal([]byte(inputArgs["args"]), &args)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg args", err))
				}
			}
			return (*Terratest).Run(&parent, testDir, args)
		case "":
			var parent Terratest
			err = json.Unmarshal(parentJSON, &parent)
			if err != nil {
				panic(fmt.Errorf("%s: %w", "failed to unmarshal parent object", err))
			}
			var version string
			if inputArgs["version"] != nil {
				err = json.Unmarshal([]byte(inputArgs["version"]), &version)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg version", err))
				}
			}
			var tfVersion string
			if inputArgs["tfVersion"] != nil {
				err = json.Unmarshal([]byte(inputArgs["tfVersion"]), &tfVersion)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg tfVersion", err))
				}
			}
			var image string
			if inputArgs["image"] != nil {
				err = json.Unmarshal([]byte(inputArgs["image"]), &image)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg image", err))
				}
			}
			var src *Directory
			if inputArgs["src"] != nil {
				err = json.Unmarshal([]byte(inputArgs["src"]), &src)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg src", err))
				}
			}
			var ctr *Container
			if inputArgs["ctr"] != nil {
				err = json.Unmarshal([]byte(inputArgs["ctr"]), &ctr)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg ctr", err))
				}
			}
			var envVars string
			if inputArgs["envVars"] != nil {
				err = json.Unmarshal([]byte(inputArgs["envVars"]), &envVars)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg envVars", err))
				}
			}
			return New(version, tfVersion, image, src, ctr, envVars), nil
		default:
			return nil, fmt.Errorf("unknown function %s", fnName)
		}
	case "":
		return dag.Module().
			WithObject(
				dag.TypeDef().WithObject("Terratest").
					WithFunction(
						dag.Function("Base",
							dag.TypeDef().WithObject("Terratest")).
							WithDescription("Base sets up the Container with a golang image and cache volumes\nversion string").
							WithArg("goVersion", dag.TypeDef().WithKind(StringKind)).
							WithArg("tfVersion", dag.TypeDef().WithKind(StringKind))).
					WithFunction(
						dag.Function("WithModule",
							dag.TypeDef().WithObject("Terratest")).
							WithDescription("WithModule specifies the module to use in the Terraform module by the 'Src' directory.").
							WithArg("src", dag.TypeDef().WithObject("Directory"))).
					WithFunction(
						dag.Function("WithContainer",
							dag.TypeDef().WithObject("Terratest")).
							WithDescription("WithContainer specifies the container to use in the Terraform module.").
							WithArg("ctr", dag.TypeDef().WithObject("Container"))).
					WithFunction(
						dag.Function("Run",
							dag.TypeDef().WithObject("Container")).
							WithArg("testDir", dag.TypeDef().WithKind(StringKind), FunctionWithArgOpts{Description: "testDir is the directory that contains all the test code."}).
							WithArg("args", dag.TypeDef().WithKind(StringKind).WithOptional(true), FunctionWithArgOpts{Description: "args is the arguments to pass to the 'go test' command."})).
					WithField("Version", dag.TypeDef().WithKind(StringKind), TypeDefWithFieldOpts{Description: "The Version of the Golang image that'll host the 'terratest' test"}).
					WithField("TfVersion", dag.TypeDef().WithKind(StringKind), TypeDefWithFieldOpts{Description: "TfVersion is the Version of the Terraform to use, e.g., \"0.12.24\".\nby default, it uses the latest Version."}).
					WithField("Image", dag.TypeDef().WithKind(StringKind), TypeDefWithFieldOpts{Description: "Image of the container to use."}).
					WithField("Src", dag.TypeDef().WithObject("Directory"), TypeDefWithFieldOpts{Description: "Src is the directory that contains all the source code, including the module directory."}).
					WithField("Ctr", dag.TypeDef().WithObject("Container"), TypeDefWithFieldOpts{Description: "Ctr is the container to use as a base container."}).
					WithConstructor(
						dag.Function("New",
							dag.TypeDef().WithObject("Terratest")).
							WithArg("version", dag.TypeDef().WithKind(StringKind).WithOptional(true), FunctionWithArgOpts{Description: "the Version of the Terraform to use, e.g., \"0.12.24\".\nby default, it uses the latest Version.", DefaultValue: JSON("\"1.22.0-alpine3.19\"")}).
							WithArg("tfVersion", dag.TypeDef().WithKind(StringKind).WithOptional(true), FunctionWithArgOpts{Description: "the Version of the Terraform to use, e.g., \"0.12.24\".\nby default, it uses the latest Version.", DefaultValue: JSON("\"1.6.0\"")}).
							WithArg("image", dag.TypeDef().WithKind(StringKind).WithOptional(true), FunctionWithArgOpts{Description: "Image of the container to use.\nby default, it uses the official HashiCorp Terraform Image hashicorp/terraform.", DefaultValue: JSON("\"gcr.io/distroless/static-debian11\"")}).
							WithArg("src", dag.TypeDef().WithObject("Directory"), FunctionWithArgOpts{Description: "Src is the directory that contains all the source code,\nincluding the module directory."}).
							WithArg("ctr", dag.TypeDef().WithObject("Container").WithOptional(true), FunctionWithArgOpts{Description: "ctr is the container to use as a base container.\nIt's an optional parameter. If it's not set, it's going to create a new container."}).
							WithArg("envVars", dag.TypeDef().WithKind(StringKind).WithOptional(true), FunctionWithArgOpts{Description: "envVars is a string of environment variables in the form of \"key1=value1,key2=value2\""}))), nil
	default:
		return nil, fmt.Errorf("unknown object %s", parentName)
	}
}
