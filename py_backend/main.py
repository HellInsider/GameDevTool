import psycopg2
import matplotlib.pyplot as plt

if __name__ == '__main__':
    psycopg2.lc_messages = 'en-US'
    psycopg2.lc_monetary = 'en-US'
    psycopg2.lc_numeric = 'en-US'
    psycopg2.lc_time = 'en-US'

    conn = psycopg2.connect(dbname='GameDevDB', user='postgres',
                        password='yellowbluevp1', host='127.0.0.1')
    cursor = conn.cursor()
    cursor.execute('SELECT * FROM "DB_schema"."Games_Genre"')
    genres_in_games = cursor.fetchall()
    cursor.execute('SELECT * FROM "DB_schema"."Genre"')
    records = cursor.fetchall()
    genres_ids = [d[0] for d in records]
    genres_descs = [d[1] for d in records]
    genres_count = [0 for d in records]
    unknown_genre = [id[1] for id in genres_in_games if id[1] not in genres_ids]
    for record in genres_in_games:
         genres_count[genres_ids.index(record[1])] += 1

    fig1, ax1 = plt.subplots()
    ax1.pie(genres_count, labels=genres_descs, autopct='%1.1f%%', startangle=90)
    ax1.axis('equal')  # Equal aspect ratio ensures that pie is drawn as a circle.

    plt.show()
    cursor.close()
    conn.close()

# See PyCharm help at https://www.jetbrains.com/help/pycharm/
