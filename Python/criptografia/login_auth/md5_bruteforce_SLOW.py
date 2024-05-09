#usr/bin/python
import time
import itertools 
import string
import hashlib
import sys
import signal
import threading

done = False
def signal_handler(signal, frame):
    print('You pressed Ctrl+C!')
    global done
    done=True
    sys.exit(0)
def animate():
    for c in itertools.cycle(['|', '/', '-', '\\']):
        if done:
            break
            
        sys.stdout.write('\rloading ' + c)
        sys.stdout.flush()
        time.sleep(0.1)
    


def _attack(chrs, inputt):
    print("[+] Start Time: ", time.strftime('%H:%M:%S'))
    start_time = time.time()
    t = threading.Thread(target=animate)
    t.start()
    total_pass_try=0
    for n in range(1, 31+1):
      characterstart_time = time.time()
      print("\n[!] I'm at ", n , "-character")
      
      for xs in itertools.product(chrs, repeat=n):
          saved = ''.join(xs)
          stringg = saved
          m = hashlib.md5()
          m.update(bytes(saved, encoding='utf-8'))
          total_pass_try +=1
          if m.hexdigest() == inputt:
              time.sleep(10)
              global done
              done = True

              print("\n[!] found ", stringg)
              print("\n[-] End Time: ", time.strftime('%H:%M:%S'))
              print("\n[-] Total Keyword attempted: ", total_pass_try)
              print("\n---Md5 cracked at %s seconds ---" % (time.time() - start_time))
              sys.exit("Thank You !")
        
      print("\n[!]",n,"-character finished in %s seconds ---" % (time.time() - characterstart_time))

def brute_attack(hash:str):
    
    print(f"\nHASH DA VEZ ---> {hash}\n")
    
    inp_usr = hash
    chrs = string.printable.replace(' \t\n\r\x0b\x0c', '')
    print(chrs)
    signal.signal(signal.SIGINT, signal_handler)
    return threading.Thread(target=_attack( chrs,inp_usr )).start()
   