package _package

import (
	"fmt"
	"github.com/cbuschka/go-pkgdiff/internal/model"
	"github.com/cbuschka/go-pkgdiff/internal/transport"
	"regexp"
	"strings"
)

type GenericPackageManager struct {
	configFile        string
	listCommand       []string
	splitVersionRegex string
}

var rpmWithDnf = GenericPackageManager{configFile: "/etc/dnf/dnf.conf", listCommand: []string{"rpm", "-qa", "--queryformat", "%{NAME}-%{VERSION}\\t%{RELEASE}:%{ARCH}\\n"}, splitVersionRegex: "^\\s*(.+)\\t([^\\s]+)\\s*$"}

var pms = []GenericPackageManager{
	rpmWithDnf,
	// "dpkg-query", "--show", "--showformat", "${source:Package}\\t${Version}:${Architecture}\\n"
}

func GetPackageManager(transport transport.Transport) (PackageManager, error) {
	for _, pm := range pms {
		available, err := pm.IsAvailable(transport)
		if err == nil && available {
			return PackageManager(&pm), nil
		}
	}

	return nil, fmt.Errorf("no package manager")
}

func (pm *GenericPackageManager) IsAvailable(channel transport.Transport) (bool, error) {
	_, _, err := channel.ExecCommand("which", pm.listCommand[0])
	if err != nil {
		return false, nil
	}

	return channel.IsFile(pm.configFile)
}

func (pm *GenericPackageManager) ListPackages(channel transport.Transport) ([]model.Package, error) {

	out, _, err := channel.ExecCommand(pm.listCommand...)
	if err != nil {
		return nil, err
	}

	packages, err := pm.readPackages(out)
	if err != nil {
		return nil, err
	}

	return packages, nil
}

func (pm *GenericPackageManager) readPackages(out string) ([]model.Package, error) {

	lines := strings.Split(out, "\n")

	re, err := regexp.Compile(pm.splitVersionRegex)
	if err != nil {
		return nil, err
	}

	packages := []model.Package{}
	for _, line := range lines {
		parts := re.FindAllStringSubmatch(line, -1)
		for _, part := range parts {
			packages = append(packages, model.Package{Name: part[1], Version: part[2]})
		}
	}

	return packages, nil
}
