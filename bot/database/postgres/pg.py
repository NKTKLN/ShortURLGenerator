import os
import random
import psycopg2
from dotenv import load_dotenv
# Connecting to the pg
load_dotenv('env/pg.env')
conn = psycopg2.connect(dbname=os.getenv('POSTGRES_DB'), user=os.getenv('POSTGRES_USER'), password=os.getenv('POSTGRES_PASSWORD'), host='pg')
cursor = conn.cursor()  
# Generating a url id and adding it to the db
def UrlIdGenerator(url, id):
    cursor.execute(f"SELECT id FROM shorturl WHERE url='{url}' AND userId={id};")
    urlId = cursor.fetchone()
    # Checking for a link in the db
    if not urlId:
        words = list('ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789')
        while True:
            # Generating a random id
            exitId = ''
            try:
                for _ in range(6):
                    exitId += random.choice(words)
                # Adding values to the db
                cursor.execute(f"INSERT INTO shorturl (id, url, userId, visits) VALUES ('{exitId}', '{url}', {id}, 0);")
                conn.commit()
                return exitId
            except Exception:
                continue
    return urlId[0]
