import io
import unittest

from installed_packages_diff.differ import create_diff
from installed_packages_diff.package import Package
from installed_packages_diff.printing import print_diff

input1 = """aaa_base-11-6.115.1
curl-7.37.0-70.66.1
gpg-pubkey-257a4686-60101325
gpg-pubkey-257a4686-60101325
gpg-pubkey-307e3d54-4be01a65
gpg-pubkey-39db7c82-510a966b
gpg-pubkey-3d25d3d9-36e12d04
gpg-pubkey-50a3dd1c-50f35137
gpg-pubkey-9c800aca-4be01999
gpg-pubkey-b37b98a9-4be01a1a
gpg-pubkey-de57bfbe-53a9be98
kernel-default-3.0.101-108.126.1
kernel-default-3.0.101-108.129.1
kernel-default-base-3.0.101-108.126.1
kernel-default-base-3.0.101-108.129.1
kernel-default-devel-3.0.101-108.126.1
kernel-default-devel-3.0.101-108.129.1
kernel-firmware-20110923-0.59.3.1
kernel-source-3.0.101-108.126.1
kernel-source-3.0.101-108.129.1
libstdc++33-32bit-3.3.3-11.9
libstdc++33-3.3.3-11.9
libstdc++43-4.6.9-0.14.1.9
libstdc++43-devel-4.3.4_20091019-37.9.1
libstdc++46-4.6.9-0.13.22
libstdc++6-32bit-5.3.1+r233831-14.1
libstdc++6-5.3.1+r233831-14.1
libstdc++-devel-4.3-62.200.2
"""

input2 = """aaa_base-11-6.115.1
curl-7.37.0-70.47.1
gpg-pubkey-257a4686-60101325
gpg-pubkey-307e3d54-4be01a65
gpg-pubkey-39db7c82-510a966b
gpg-pubkey-3d25d3d9-36e12d04
gpg-pubkey-50a3dd1c-50f35137
gpg-pubkey-9c800aca-4be01999
gpg-pubkey-b37b98a9-4be01a1a
gpg-pubkey-de57bfbe-53a9be98
kernel-default-3.0.101-107.1
kernel-default-3.0.101-108.126.1
kernel-default-base-3.0.101-107.1
kernel-default-base-3.0.101-108.126.1
kernel-default-devel-3.0.101-107.1
kernel-default-devel-3.0.101-108.126.1
kernel-firmware-20110923-0.59.3.1
kernel-source-3.0.101-107.1
kernel-source-3.0.101-108.126.1
libstdc++33-32bit-3.3.3-11.9
libstdc++33-3.3.3-11.9
libstdc++43-4.6.9-0.14.1.9
libstdc++43-devel-4.3.4_20091019-37.9.1
libstdc++46-4.6.9-0.13.22
libstdc++47-devel-32bit-4.7.2_20130108-0.19.3
libstdc++47-devel-4.7.2_20130108-0.19.3
libstdc++6-5.3.1+r233831-14.1
libstdc++-devel-4.3-62.200.2
"""


class IntegerationTest(unittest.TestCase):
  def test_all(self):
    aPackages = [Package.parse(line) for line in input1.splitlines()]
    bPackages = [Package.parse(line) for line in input2.splitlines()]
    installed_packages_diff = create_diff(aPackages, bPackages, includeEqual=False)
    with io.StringIO() as buf:
      print_diff("serverA", "serverB", installed_packages_diff, file=buf)
      self.assertEqual([l.strip() for l in """
= serverA serverB =
curl                                     7.37.0-70.66.1                           7.37.0-70.47.1
gpg-pubkey                               257a4686-60101325                        missing
kernel-default                           3.0.101-108.129.1                        missing
kernel-default                           missing                                  3.0.101-107.1
kernel-default-base                      3.0.101-108.129.1                        missing
kernel-default-base                      missing                                  3.0.101-107.1
kernel-default-devel                     3.0.101-108.129.1                        missing
kernel-default-devel                     missing                                  3.0.101-107.1
kernel-source                            3.0.101-108.129.1                        missing
kernel-source                            missing                                  3.0.101-107.1
libstdc++47-devel                        missing                                  4.7.2_20130108-0.19.3""".splitlines()],
                       [l.strip() for l in buf.getvalue().splitlines()])
