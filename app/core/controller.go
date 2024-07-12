package core

type Controller struct {
    Commands map[string]ICommand
}

func NewController() *Controller {
    return &Controller{make(map[string]ICommand)}
}

func (controller *Controller) AddCommand(key string, command ICommand) {
    controller.Commands[key] = command
}

func (controller *Controller) Execute(key string) {
    if controller.Commands[key] == nil {
        return
    }

    controller.Commands[key].Execute()
}
