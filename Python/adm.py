import subprocess
import time
from ppadb.client import Client


subprocess.run(['adb', '-s', '875997070405', 'shell', 'input tap 200 200'])