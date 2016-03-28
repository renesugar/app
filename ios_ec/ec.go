package ios_ec

import "github.com/qor/app"

type EC struct {
	app.Theme
}

func (ec *EC) GetTemplatesPath() string {
	if ec.TemplatesPath == "" {
		ec.TemplatesPath = "github.com/qor/app/ios_ec/templates"
	}
	return ec.Theme.GetTemplatesPath()
}

func (ec *EC) ConfigureQorApplication(theme app.ThemeInterface) {
	return
}

func (ec *EC) Build(theme app.ThemeInterface) error {
	return nil
}