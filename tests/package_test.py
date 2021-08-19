import unittest
from pkgdiff.package import Package
from pkgdiff.differ import create_diff


class PackageTest(unittest.TestCase):
  def test_valid_name(self):
    p = Package.parse("a-v1-1.9")
    self.assertEqual(p.name, "a")
    self.assertEqual(p.version, "v1-1.9")

  def test_missing_version(self):
    with self.assertRaises(ValueError):
      Package.parse("a")
