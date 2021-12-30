package command

import (
	configPkg "github.com/cbuschka/go-pkgdiff/internal/config"
	"github.com/cbuschka/go-pkgdiff/internal/model"
	_package "github.com/cbuschka/go-pkgdiff/internal/package"
	transportPkg "github.com/cbuschka/go-pkgdiff/internal/transport"
)

func Run() error {

	config, err := configPkg.LoadConfigFromFile("./config.yaml")
	if err != nil {
		return err
	}

	for _, group := range config.Groups {

		packageLists := []model.PackageList{}
		for _, serverConfig := range group.Servers {
			transport, err := transportPkg.Connect(&serverConfig)
			if err != nil {
				return err
			}

			manager, err := _package.GetPackageManager(transport)
			if err != nil {
				return err
			}

			packageList, err := manager.ListPackages(transport)
			if err != nil {
				return err
			}

			packageLists = append(packageLists, model.PackageList{Url: serverConfig.Url, Packages: packageList})
		}

		// diff(packageLists)
	}

	return nil
}
