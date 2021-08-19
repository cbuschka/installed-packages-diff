import re


class Package(object):
  PKG_NAME_REGEX = re.compile("^\s*(.*)-([^\-\s]+-[^\-\s]+)\s*$")

  @classmethod
  def parse(cls, pkg_name: str):
    result = Package.PKG_NAME_REGEX.fullmatch(pkg_name)
    if not result:
      raise ValueError(f"Invalid package name '{pkg_name}'.")

    return Package(result.group(1), result.group(2))

  def __init__(self, name, version):
    self.name = name
    self.version = version

  def __repr__(self):
    return f"Package(name={self.name},version={self.version})"

  def __hash__(self):
    return hash(tuple(sorted(self.__dict__.items())))

  def __eq__(self, other):
    return repr(self) == repr(other)
