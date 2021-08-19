from yaml import load
import logging

try:
  from yaml import CLoader as Loader, CDumper as Dumper
except ImportError:
  from yaml import Loader, Dumper


class Exclude(object):
  def __init__(self, name):
    self.name = name


class Server(object):
  def __init__(self, raw):
    self.hostname = raw["hostname"]
    self.username = raw.get("username", None)
    self.excludes = {Exclude(e) for e in raw.get("excludes", [])}


class Group(object):
  def __init__(self, name, raw):
    self.name = name
    self.servers = [Server(server) for server in raw["servers"]]


class Config(object):
  def __init__(self, raw):
    self.raw = raw
    groups_dict = self.raw.get("groups", {})
    self.groups = [Group(name, groups_dict[name]) for name in groups_dict]


def load_config(filename):
  logging.info(f"Opening config {filename}...")
  with open(filename, 'rb') as stream:
    data = load(stream, Loader=Loader)
    return Config(data)
