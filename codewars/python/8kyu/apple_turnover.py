def apple(x):
  x = int(x) ** 2
  if x > 1000:
    return "It's hotter than the sun!!"
  return "Help yourself to a honeycomb Yorkie for the glovebox."


if __name__ == '__main__':
  print(apple('50'))
  print(apple(4))
  print(apple('12'))
  print(apple(60))
