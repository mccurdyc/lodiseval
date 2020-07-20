package command

type Helper interface {
	ShortHelp() string
	Help() string
}
