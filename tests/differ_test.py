import unittest
from pkgdiff.package import Package
from pkgdiff.differ import create_diff


class DifferTest(unittest.TestCase):
  def test_all(self):
    a = Package("a", "v1-1.0")
    b = Package("b", "v1-1.0")
    c1 = Package("c", "v1-1.0")
    c2 = Package("c", "v1-1.1")
    d1 = Package("d", "v1-1.0")
    d2 = Package("d", "v2-1.0")

    diff = create_diff([a, c1, d1], [b, c2, d2])

    self.assertEqual({'a': ['v1-1.0', 'missing'],
                      'b': ['missing', 'v1-1.0'],
                      'c': ['v1-1.0', 'v1-1.1'],
                      'd': ['v1-1.0', 'v2-1.0']},
                     diff)
