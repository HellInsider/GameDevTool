import psycopg2

def init_dbtools():
    psycopg2.lc_messages = 'en-US'
    psycopg2.lc_monetary = 'en-US'
    psycopg2.lc_numeric = 'en-US'
    psycopg2.lc_time = 'en-US'

    conn = psycopg2.connect(dbname='GameDevDB', user='postgres',
                        password='yellowbluevp1', host='127.0.0.1')
    cursor = conn.cursor()

    cursor.execute(
        'CREATE OR REPLACE FUNCTION to_date_or_null(ADate TEXT, AFormat TEXT) RETURNS CHAR AS\n'
        '$BODY$\n'
        'BEGIN\n'
        '   RETURN to_date(ADate,AFormat);\n'
        'EXCEPTION\n'
        'WHEN others THEN RETURN to_date(\'01 01 1000\',\'DD MM YYYY\');\n'
        'END;\n'
        '$BODY$ LANGUAGE plpgsql IMMUTABLE STRICT;\n')

    cursor.close()
    return conn


def count_games_total(connection, date1: str, date2: str):
    cursor = connection.cursor()
    cursor.execute('SELECT COUNT(*) AS conv_date FROM '
                       '(SELECT appid, release_date_id '
                       'FROM "DB_schema"."Games") AS gr '
                   'INNER JOIN "DB_schema"."Release_date" AS r '
                   'ON gr.appid = r.appid '
                   'WHERE TO_DATE(to_date_or_null(r.date, \'DD Mon, YYYY\'), \'YYYY-MM-DD\') >= \'' + date1 + '\' '
                   'AND TO_DATE(to_date_or_null(r.date, \'DD Mon, YYYY\'), \'YYYY-MM-DD\') < \'' + date2 + '\'')
    records = cursor.fetchall()
    cursor.close()
    return records[0]


def count_games(connection, genre_str: str, date1: str, date2: str):
    cursor = connection.cursor()
    cursor.execute('SELECT * FROM "DB_schema"."Genre" WHERE description = \'' + genre_str + '\'')
    records = cursor.fetchall()
    if len(records) != 1:
        cursor.close()
        return 0

    genre_id = records[0][0]

    cursor.execute('SELECT COUNT(*) AS conv_date FROM '
                   '(SELECT g.appid, g.release_date_id '
                       'FROM "DB_schema"."Games" AS g '
                       'INNER JOIN "DB_schema"."Games_Genre" AS gg '
                       'ON g."appid" = gg."Games_appid" '
                       'WHERE gg."Genre_genre_id" = ' + str(genre_id) + ') AS gr '
                   'INNER JOIN "DB_schema"."Release_date" AS r '
                   'ON gr.appid = r.appid '
                   'WHERE TO_DATE(to_date_or_null(r.date, \'DD Mon, YYYY\'), \'YYYY-MM-DD\') >= \'' + date1 + '\' '
                   'AND TO_DATE(to_date_or_null(r.date, \'DD Mon, YYYY\'), \'YYYY-MM-DD\') < \'' + date2 + '\'')

    records = cursor.fetchall()
    cursor.close()
    return records[0]


def genre_reviews(conn, genre_str: str):
    cursor = conn.cursor()
    cursor.execute('SELECT * FROM "DB_schema"."Genre" WHERE description = \'' + genre_str + '\'')
    records = cursor.fetchall()
    if len(records) != 1:
        cursor.close()
        return 0

    genre_id = records[0][0]
    cursor.execute('SELECT review_score, COUNT(review_score) '
                   'FROM "DB_schema"."Reviews" as r '
                   'INNER JOIN "DB_schema"."Games_Genre" AS gg '
                   'ON r."appid" = gg."Games_appid" '
                   'WHERE gg."Genre_genre_id" = ' + str(genre_id) + ' '
                   'GROUP BY review_score '
                   'ORDER BY review_score')
    records = cursor.fetchall()

    cursor.close()
    return records


def median_genre_review_by_year(conn, genre_str: str):
    cursor = conn.cursor()
    cursor.execute('SELECT ye, PERCENTILE_CONT(0.5) '
                   'WITHIN GROUP(ORDER BY dgri.review_score) '
                   'FROM '
                       '(SELECT *, '
                           'EXTRACT(YEAR FROM '
                               'TO_DATE(to_date_or_null(d.date, \'DD Mon, YYYY\'), \'YYYY-MM-DD\'))::INTEGER '
                           'AS ye '
                       'FROM '
                           '(SELECT * '
                           'FROM "DB_schema"."Reviews" as r '
                           'INNER JOIN '
                               '(SELECT gg."Games_appid" '
                               'FROM "DB_schema"."Games_Genre" AS gg '
                               'INNER JOIN "DB_schema"."Genre" AS g '
                               'ON gg."Genre_genre_id" = g.genre_id '
                               'WHERE g.description = \'' + genre_str + '\') as gi '
                           'ON r."appid" = gi."Games_appid" '
                           'WHERE r.review_score != 0) AS gri '
                       'INNER JOIN "DB_schema"."Release_date" AS d '
                       'ON gri.appid = d.appid) AS dgri '
                   'GROUP BY ye '
                   'ORDER BY ye'
                   )

    records = cursor.fetchall()
    cursor.close()
    return records


def avg_genre_review_by_year(conn, genre_str: str):
    cursor = conn.cursor()
    cursor.execute('SELECT ye, AVG(review_score)::FLOAT as avgrs '
                   'FROM '
                       '(SELECT *, '
                           'EXTRACT(YEAR FROM '
                               'TO_DATE(to_date_or_null(d.date, \'DD Mon, YYYY\'),'
                               ' \'YYYY-MM-DD\'))::INTEGER '
                           'AS ye '
                       'FROM '
                           '(SELECT * '
                           'FROM "DB_schema"."Reviews" as r '
                           'INNER JOIN '
                               '(SELECT gg."Games_appid" '
                               'FROM "DB_schema"."Games_Genre" AS gg '
                               'INNER JOIN "DB_schema"."Genre" AS g '
                               'ON gg."Genre_genre_id" = g.genre_id '
                               'WHERE g.description = \'' + genre_str + '\') as gi '
                           'ON r."appid" = gi."Games_appid" '
                           'WHERE r.review_score != 0) AS gri '
                       'INNER JOIN "DB_schema"."Release_date" AS d '
                       'ON gri.appid = d.appid) AS dgri '
                   'GROUP BY ye '
                   'ORDER BY ye')

    records = cursor.fetchall()
    cursor.close()
    return records
