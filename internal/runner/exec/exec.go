package exec

import (
	"fmt"

	"github.com/jahvon/flow/config"
	"github.com/jahvon/flow/internal/context"
	"github.com/jahvon/flow/internal/runner"
	"github.com/jahvon/flow/internal/services/run"
)

type execRunner struct{}

func NewRunner() runner.Runner {
	return &execRunner{}
}

func (r *execRunner) Name() string {
	return "exec"
}

func (r *execRunner) IsCompatible(executable *config.Executable) bool {
	if executable == nil || executable.Type == nil || executable.Type.Exec == nil {
		return false
	}
	return true
}

func (r *execRunner) Exec(ctx *context.Context, executable *config.Executable, promptedEnv map[string]string) error {
	execSpec := executable.Type.Exec
	envMap, err := runner.ParametersToEnvMap(&execSpec.ParameterizedExecutable, promptedEnv)
	if err != nil {
		return fmt.Errorf("env setup failed - %w", err)
	}
	envList, err := runner.ParametersToEnvList(&execSpec.ParameterizedExecutable, promptedEnv)
	if err != nil {
		return fmt.Errorf("env setup failed  - %w", err)
	}

	targetDir, isTmp, err := execSpec.ExpandDirectory(
		executable.WorkspacePath(),
		executable.DefinitionPath(),
		ctx.ProcessTmpDir,
		envMap,
	)
	if err != nil {
		return fmt.Errorf("unable to expand directory - %w", err)
	} else if isTmp {
		ctx.ProcessTmpDir = targetDir
	}

	logMode := execSpec.LogMode
	var logFields map[string]interface{}

	if logMode == config.StructuredLogMode {
		logFields = execSpec.GetLogFields()
	}

	switch {
	case execSpec.Command == "" && execSpec.File == "":
		return fmt.Errorf("either cmd or file must be specified")
	case execSpec.Command != "" && execSpec.File != "":
		return fmt.Errorf("cannot set both cmd and file")
	case execSpec.Command != "":
		return run.RunCmd(execSpec.Command, targetDir, envList, logMode, ctx.Logger, logFields)
	case execSpec.File != "":
		return run.RunFile(execSpec.File, targetDir, envList, logMode, ctx.Logger, logFields)
	default:
		return fmt.Errorf("unable to determine how executable should be run")
	}
}
