import unittest
from pkgdiff.package import Package
from pkgdiff.differ import create_diff


class PackageTest(unittest.TestCase):
  def test_valid_name(self):
    p = Package.parse("a-v1-1.9")
    self.assertEqual(p.name, "a")
    self.assertEqual(p.version, "v1-1.9")

  def test_equal(self):
    p = Package.parse("a-v1-1")
    p2 = Package.parse("a-v1-1")
    self.assertEqual(p, p2)

  def test_missing_version(self):
    with self.assertRaises(ValueError):
      Package.parse("a")

  def test_valid_name2(self):
    p = Package.parse("libstdc++47-devel-32bit-4.7.2_20130108-0.19.3")
    self.assertEqual(p.name, "libstdc++47-devel-32bit")
    self.assertEqual(p.version, "4.7.2_20130108-0.19.3")
