import sys

from .config import load_config
from .fetcher import PackageFetcher
from .differ import create_diff
import logging


def diff_server(serverA, serverB):
  pkgFetcher = PackageFetcher()
  devPkgs = pkgFetcher.get_packages(serverA.hostname, username=serverA.username)
  prodPkgs = pkgFetcher.get_packages(serverB.hostname,
                                     username=serverB.username)
  pkgDiff = create_diff(devPkgs, prodPkgs)
  print(f"\n= {serverA.hostname} {serverB.hostname} =")
  for packageName in pkgDiff:
    result = pkgDiff[packageName]
    if result[0] == "missing":
      action = "A"
    elif result[1] == "missing":
      action = "B"
    elif result[0] != result[1]:
      action = "U"
    else:
      action = " "

    print(f"{action} {packageName:40s} {result[0]:40s} {result[1]:40s}")


def main():
  logging.basicConfig(stream=sys.stderr, encoding='utf-8',
                      format='%(asctime)s %(message)s',
                      level=logging.WARN)

  config = load_config(sys.argv[1] if len(sys.argv) > 1 else "config.yaml")
  for group in config.groups:
    serverA = group.servers[0]
    serverB = group.servers[1]

    diff_server(serverA, serverB)
