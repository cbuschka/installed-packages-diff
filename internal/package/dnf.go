package _package

import (
	"regexp"
	"strings"
)

type Dnf struct{}

func (dnf *Dnf) IsAvailable(channel Channel) (bool, error) {
	_, _, err := channel.Run("which", "dnf")
	if err != nil {
		return false, nil
	}

	return channel.IsFile("/etc/dnf/dnf.conf")
}

func (dnf *Dnf) ListPackages(channel Channel) ([]Package, error) {

	out, _, err := channel.Run("dnf", "list", "--installed")
	if err != nil {
		return nil, err
	}

	packages, err := ReadPackages(out)
	if err != nil {
		return nil, err
	}

	return packages, nil
}

func ReadPackages(out string) ([]Package, error) {

	lines := strings.Split(out, "\n")

	re, err := regexp.Compile("\\s*(\\S+)\\s+(\\S+)\\s+(\\S+)\\s*")
	if err != nil {
		return nil, err
	}

	packages := []Package{}
	for _, line := range lines {
		parts := re.FindAllStringSubmatch(line, -1)
		for _, part := range parts {
			packages = append(packages, Package{Name: part[1], Version: part[2]})
		}
	}

	return packages, nil
}
