import psycopg2
import matplotlib.pyplot as plt
import que
import genres
import reviews


if __name__ == '__main__':
    conn = que.init_dbtools()

    #genres.genres_pie(conn)

    #genres.genres_paper(conn)

    #genres.genres_addition(conn)

    genres.genres_detailed(conn)

    #reviews.reviews_distrib(conn)

    #reviews.reviews_genres(conn)

    #reviews.reviews_timeline(conn)

    cursor = conn.cursor()
    cursor.execute('SELECT * FROM "DB_schema"."Genre"')
    records = cursor.fetchall()

    #    records = cursor.fetchmany(300)

    plt.show()
    cursor.close()
    conn.close()

