def create_diff(listA, listB, *, aExcludes=None, bExcludes=None):
  mapA = {p.name: p for p in listA}
  mapB = {p.name: p for p in listB}
  diffAB = {}
  for packageName in mapA:
    packageA = mapA[packageName]
    packageB = mapB.get(packageName, None)
    if packageB is None:
      if bExcludes is None or "missing" not in bExcludes:
        diffAB[packageName] = [packageA.version, "missing"]
    elif packageB != packageA:
      diffAB[packageName] = [packageA.version, packageB.version]
  for packageName in mapB:
    if not packageName in diffAB:
      packageB = mapB[packageName]
      if aExcludes is None or "missing" not in aExcludes:
        diffAB[packageName] = ["missing", packageB.version]

  return diffAB
