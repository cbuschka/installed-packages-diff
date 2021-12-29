import logging

from .transport import Transport
import paramiko

class SshTransport(Transport):
  def __init__(self, hostname, *, username=None):
    self.hostname = hostname
    self.username = username

  def exec_command(self, command):
    logging.info(f"Fetching package from {self.username}@{self.hostname}...")
    with paramiko.SSHClient() as client:
      client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
      client.connect(self.hostname, username=self.username)
      session = client.get_transport().open_session()
      paramiko.agent.AgentRequestHandler(session)
      session.exec_command(command)
      stdin = session.makefile_stdin("wb", -1)
      stdin.close()
      stdout = session.makefile("r", -1)
      stderr = session.makefile_stderr("r", -1)
      out_lines = stdout.readlines()
      err_lines = stderr.readlines()
      exit_status = session.recv_exit_status()
      if exit_status != 0:
        print(err_lines)
        raise ValueError(
          f"Querying packages from {self.hostname} failed with exit status {exit_status}")
      return out_lines, err_lines, exit_status
