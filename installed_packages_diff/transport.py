import logging

import paramiko

from .package import Package


class Transport(object):
  def exec_command(self, command):
    raise NotImplementedError()
