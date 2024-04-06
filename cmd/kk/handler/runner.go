package handler

import (
	"keykeeper/cmd/kk/file"
	"keykeeper/cmd/kk/model"
	"strings"
)

type Runner struct {
	env        *HandlerEnv
	dispatcher *Dispatcher
}

func (r *Runner) Run(args []string) error {
	arguments, flags := r.separateArgumentsAndFlags(args)

	r.env.Flags = flags
	r.env.Args = arguments

	return registeredHandlers["dispatch"].Handle(r.env)
}

func (r *Runner) RegisterRepo(repo *model.Repo) {
	r.env.Repo = repo
}

func (r *Runner) separateArgumentsAndFlags(args []string) ([]string, []string) {
	flags := make([]string, 0)
	arguments := make([]string, 0)

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			flags = append(flags, arg)
		} else {
			arguments = append(arguments, arg)
		}
	}

	return arguments, flags
}

func (r *Runner) RegisterFileHandler(handler file.Handler) {
	r.env.FileHandler[handler.GetFileExtension()] = handler
}

func NewRunner() *Runner {
	return &Runner{
		env: &HandlerEnv{
			FileHandler: make(map[string]file.Handler),
		},
	}
}
