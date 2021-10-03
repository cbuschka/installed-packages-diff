package internal

import (
	configPkg "github.com/cbuschka/pkgdiff/internal/config"
	_package "github.com/cbuschka/pkgdiff/internal/package"
	serverPkg "github.com/cbuschka/pkgdiff/internal/server"
)

func Run() error {

	config, err := configPkg.LoadConfig("./config.yaml")
	if err != nil {
		return err
	}

	for _, group := range config.Groups {

		packageLists := []*serverPkg.PackageList{}
		for _, serverConfig := range group.Servers {
			server, err := serverPkg.Connect(&serverConfig)
			if err != nil {
				return err
			}

			packageList, err := server.ListPackages()
			if err != nil {
				return err
			}

			packageLists = append(packageLists, packageList)
		}

    diff(packageLists)
	}

	return nil
}
