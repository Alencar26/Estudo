#usr/bin/python
import time
import itertools
import string
import hashlib
import sys
import signal
import threading

done = False
def animate():
    for c in itertools.cycle(['|', '/', '-', '\\']):
        if done is True:
            break
        
        sys.stdout.write('\rloading ' + c)
        sys.stdout.flush()
        time.sleep(0.2)
        
    

def _attack(chrs, inputt):
    global done
    done = False
    print("[+] Start Time: ", time.strftime('%H:%M:%S'))
    start_time = time.time()
    t = threading.Thread(target=animate)
    t.start()
    total_pass_try=0
    for n in range(1, 31+1):
      characterstart_time = time.time()
      
      for xs in itertools.product(chrs, repeat=n):
          saved = ''.join(xs)

          stringg = saved
          m = hashlib.md5()
          m.update(bytes(saved, encoding='utf-8'))
          total_pass_try +=1
          if m.hexdigest() == inputt:
              done = True
              print("\r\r _ _ _ _ _ _ _ _ _ _\n\n", stringg+'\n _ _ _ _ _ _ _ _ _ _')
              print("\n[-] End Time: ", time.strftime('%H:%M:%S'),"(%s sec)" % round((time.time() - start_time),1))
              print("\n[-] Total Keyword attempted: ", total_pass_try)
            #   controlador = True
              break
      if done is True:
        print("\r[!]",n,"-character finished in %s sec\r" % round((time.time() - characterstart_time),1))
        break
        

def brute_attack(hash_list:list[str]):
    
    chrs = string.printable.replace(' \t\n\r\x0b\x0c', '')

    #comment line below for full (Slow) bruteforce
    chrs='ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklnmopqrstuvwxyz1234567890!@#$%¨&*()_+=-{[^~]};:.,<>' 

    #uncomment line below for Russian letters
    #chrs=chrs+'абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ' #uncomment for Russian letters

    print('\nCharacter list:\n'+chrs+'\n')
    for hash in hash_list:
        print(f"\nHASH DA VEZ ---> {hash}\n")
        _attack(chrs,hash.lower())
   
   
   