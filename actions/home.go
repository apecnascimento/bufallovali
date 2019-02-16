package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	c.Set("task-title", "Dashboard")
	c.Set("task-icon", "fa fa-dashboard")
	c.Set("task-desc", "A free and open source Bootstrap 4 admin template")
	return c.Render(200, r.HTML("index.html"))
}

// ChartHandler is a handler to make charts works
func ChartHandler(c buffalo.Context) error {

	return c.Render(200, r.HTML("charts.html"))
}
