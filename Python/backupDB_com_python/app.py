from decouple import config #pip install python-decouple - lidar com as configs
from pathlib import Path
from datetime import datetime
import subprocess
import gzip
import os
import schedule #pip install schedule - rodar rotina de jobs
from time import sleep

DB_HOST = config('DB_HOST')
DB_PORT = config('DB_PORT')
DB_USER = config('DB_USER')
DB_PASSWORD = config('DB_PASSWORD')
DB_NAME = config('DB_NAME')

def backup_database():
    backup_dir = Path(__file__).resolve().parent
    file_name = f"{DB_NAME}_{datetime.now().strftime('%d-%m-%Y_%H-%M-%S')}"
    backup_path = f"{backup_dir}/{file_name}.sql"
    
    database_dump_command = f"pg_dump -h {DB_HOST} -p {DB_PORT} -U {DB_USER} -d {DB_NAME} -f {backup_path}"

    print("LOG >> INIT DATABASE DUMP")
    subprocess.run(database_dump_command, shell=True, check=True, env={'PGPASSWORD': DB_PASSWORD})
    print("LOG >> FINISH DATABASE DUMP")
 
    compress_backup_path = f"{backup_dir}/{file_name}.sql.gz"

    with open(backup_path, 'rb') as original_file:
        with gzip.open(compress_backup_path, 'wb') as compress_file:
            print("LOG >> COMPRESS SQL FILE")
            compress_file.write(original_file.read())
            print("LOG >> FINISH COMPRESS")

    print("LOG >> REMOVE ORIGINAL FILE")
    os.remove(backup_path)

schedule.every().friday.at('20:00').do(backup_database)

while True:
    print("Verificando...")
    schedule.run_pending()
    sleep(5)
