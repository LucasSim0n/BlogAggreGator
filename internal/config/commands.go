package config

import "fmt"

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Cmds map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	f, ok := c.Cmds[cmd.Name]
	if !ok {
		return fmt.Errorf("Command '%s' not found\n", cmd.Name)
	}

	err := f(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	if c.Cmds == nil {
		c.Cmds = make(map[string]func(*State, Command) error)
	}
	c.Cmds[name] = f
}
