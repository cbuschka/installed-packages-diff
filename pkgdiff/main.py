import sys

from .config import load_config
from .fetcher import PackageFetcher
from .differ import create_diff
from .printing import print_diff
import logging


def diff_server(serverA, serverB):
  pkgFetcher = PackageFetcher()
  devPkgs = pkgFetcher.get_packages(serverA.hostname, username=serverA.username)
  prodPkgs = pkgFetcher.get_packages(serverB.hostname,
                                     username=serverB.username)
  pkgDiff = create_diff(devPkgs, prodPkgs, aExcludes=serverA.excludes,
                        bExcludes=serverB.excludes, includeEqual=False)
  print_diff(serverA.hostname, serverB.hostname, pkgDiff)


def main():
  logging.basicConfig(stream=sys.stderr, encoding='utf-8',
                      format='%(asctime)s %(message)s',
                      level=logging.WARN)

  config = load_config(sys.argv[1] if len(sys.argv) > 1 else "config.yaml")
  for group in config.groups:
    serverA = group.servers[0]
    serverB = group.servers[1]

    diff_server(serverA, serverB)
