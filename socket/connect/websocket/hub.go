package main

type hub struct {
	c map[*connection]bool
}
