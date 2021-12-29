import logging

from .local_transport import LocalTransport
from .ssh_transport import SshTransport
from .package import Package


class PackageFetcher(object):
  LIST_PACKAGES_COMMAND = {
    "rpm": ["rpm", "-qa", "--queryformat", "%{NAME}-%{VERSION}\\t%{RELEASE}:%{ARCH}\\n"],
    "dpkg": ["dpkg-query", "--show", "--showformat", "${source:Package}\\t${Version}:${Architecture}\\n"]
  }

  def __init__(self):
    # self.transport = LocalTransport()
    self.transport = SshTransport()

  def get_packages(self, hostname, *, username=None, type="rpm"):
    logging.info(f"Fetching package from {username}@{hostname}...")

    out_lines, stderr_lines, exit_code = self.transport.exec_command(PackageFetcher.LIST_PACKAGES_COMMAND[type])
    return [Package.parse(line.strip(), type=type) for line in out_lines]
