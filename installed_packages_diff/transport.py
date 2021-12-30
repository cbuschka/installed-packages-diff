class Transport(object):
  def __enter__(self):
    return self

  def exec_command(self, command):
    raise NotImplementedError()

  def close(self):
    pass

  def __exit__(self, type, value, traceback):
    self.close()
