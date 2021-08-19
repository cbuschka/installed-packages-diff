def print_diff(serverA, serverB, pkgDiff, *, file=None):
  print(f"\n= {serverA} {serverB} =", file=file)
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

    print(f"{action} {packageName:40s} {result[0]:40s} {result[1]:40s}",
          file=file)
