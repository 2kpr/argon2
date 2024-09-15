import os
import sys
import argon2
import getpass

key = argon2.low_level.hash_secret_raw(
  getpass.getpass().encode(),
  getpass.getpass().encode(),
  time_cost=1,
  memory_cost=1024,
  parallelism=1,
  hash_len=512,
  type=argon2.low_level.Type.ID
)

with open(getpass.getpass(prompt='Key File: ').encode(), "wb") as f:
	f.write(key)
