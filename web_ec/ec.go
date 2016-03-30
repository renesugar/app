package web_ec

import (
	"github.com/qor/app"
	"github.com/qor/app/modules/products"
)

type EC struct {
	Path string
	app.Theme
}

func (ec *EC) GetTemplatesPath() string {
	if ec.TemplatesPath == "" {
		ec.TemplatesPath = "github.com/qor/app/web_ec/templates"
	}
	return ec.Theme.GetTemplatesPath()
}

func (ec *EC) ConfigureQorApplication(theme app.ThemeInterface) {
	ec.Theme.Path = ec.Path

	// Add Product Plugin
	ec.UsePlugin(&products.Product{})
	return
}

func (ec *EC) Build(theme app.ThemeInterface) error {
	return nil
}
