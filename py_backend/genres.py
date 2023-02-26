import matplotlib.pyplot as plt
import que

GENRES_NAMES = ['Adventure', 'RPG', 'Strategy', 'Racing', 'Indie']
RELEASE_PERIODS = ['1986-01-01', '1992-01-01', '1996-01-01', '2000-01-01',
                   '2004-01-01', '2008-01-01', '2012-01-01', '2016-01-01',
                   '2020-01-01', '2023-01-01']
RELEASE_PERIODS_LABELS = [RELEASE_PERIODS[i][0:4] + '-' + str(int(RELEASE_PERIODS[i + 1][0:4]) - 1)
                          for i in range(len(RELEASE_PERIODS) - 1)]

RELEASE_PERIODS_ALTERNATIVE = ['2003-01-01', '2012-01-01', '2013-01-01', '2014-01-01',
                               '2015-01-01', '2016-01-01', '2017-01-01', '2018-01-01',
                               '2019-01-01', '2020-01-01', '2021-01-01', '2022-01-01',
                               '2023-01-01']
RELEASE_PERIODS_ALT_LABELS = [RELEASE_PERIODS_ALTERNATIVE[i][0:4]
                              for i in range(len(RELEASE_PERIODS_ALTERNATIVE) - 1)]
RELEASE_PERIODS_ALT_LABELS[0] = RELEASE_PERIODS_ALTERNATIVE[0][0:4] + '-' + \
                              str(int(RELEASE_PERIODS_ALTERNATIVE[0 + 1][0:4]) - 1)


def genres_pie(conn):
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

    oth = 0
    i = 0
    while i < len(genres_count):
        if genres_count[i] < len(genres_in_games) * 0.03:
            oth += genres_count[i]
            genres_descs.remove(genres_descs[i])
            genres_count.remove(genres_count[i])
            i -= 1
        i += 1
    genres_count.append(oth)
    genres_descs.append('Other')

    fig1, ax1 = plt.subplots()
    ax1.pie(genres_count, labels=genres_descs, autopct='%1.1f%%', startangle=90)
    ax1.axis('equal')  # Equal aspect ratio ensures that pie is drawn as a circle.
    plt.savefig("../Pics/genres_piechart.png")
    cursor.close()



def genres_paper(conn):
    for gnr in GENRES_NAMES:
        plt.figure()
        num_games = [que.count_games(conn, gnr, RELEASE_PERIODS[i], RELEASE_PERIODS[i + 1])[0]
             for i in range(len(RELEASE_PERIODS) - 1)]
        plt.title(gnr)
        plt.bar(RELEASE_PERIODS_LABELS, num_games)
        plt.xticks(rotation=45)
        plt.subplots_adjust(bottom=0.16)
        plt.savefig("../Pics/paper_" + gnr + ".png")


def genres_addition(conn):
    plt.figure()
    num_games_total = [que.count_games_total(conn, RELEASE_PERIODS[i], RELEASE_PERIODS[i + 1])[0]
        for i in range(len(RELEASE_PERIODS) - 1)]
    plt.title('total games')
    plt.bar(RELEASE_PERIODS_LABELS, num_games_total)
    plt.xticks(rotation=45)
    plt.subplots_adjust(bottom=0.16)
    plt.savefig("../Pics/total_games.png")

    plt.figure()
    num_last = [0 for i in range(len(RELEASE_PERIODS) - 1)]

    for gnr in GENRES_NAMES:
        num_games = [que.count_games(conn, gnr, RELEASE_PERIODS[i], RELEASE_PERIODS[i + 1])[0] /
                     max(num_games_total[i], 1)
                     for i in range(len(RELEASE_PERIODS) - 1)]
        plt.title('genres relative')
        plt.bar(RELEASE_PERIODS_LABELS, num_games, bottom = num_last)
        num_last = [num_games[i] + num_last[i] for i in range(len(RELEASE_PERIODS) - 1)]
        plt.xticks(rotation=45)
        plt.subplots_adjust(bottom=0.16)
    plt.legend(GENRES_NAMES)
    plt.savefig("../Pics/relative_genres.png")

    num_last = [0 for i in range(len(RELEASE_PERIODS) - 1)]

    for gnr in GENRES_NAMES:
        plt.figure()
        num_games = [que.count_games(conn, gnr, RELEASE_PERIODS[i], RELEASE_PERIODS[i + 1])[0] /
                     max(num_games_total[i], 1)
                     for i in range(len(RELEASE_PERIODS) - 1)]
        plt.title(gnr + ' relative')
        plt.bar(RELEASE_PERIODS_LABELS, num_games)
        plt.xticks(rotation=45)
        plt.subplots_adjust(bottom=0.16)
        plt.savefig("../Pics/relative_" + gnr + ".png")


def genres_detailed(conn):
    plt.figure()
    num_games_total = [que.count_games_total(conn, RELEASE_PERIODS_ALTERNATIVE[i], RELEASE_PERIODS_ALTERNATIVE[i + 1])[0]
        for i in range(len(RELEASE_PERIODS_ALTERNATIVE) - 1)]
    plt.title('total games')
    plt.bar(RELEASE_PERIODS_ALT_LABELS, num_games_total)
    plt.xticks(rotation=45)
    plt.subplots_adjust(bottom=0.16)
    plt.savefig("../Pics/total2_games.png")

    for gnr in GENRES_NAMES:
        plt.figure()
        num_games = [que.count_games(conn, gnr, RELEASE_PERIODS_ALTERNATIVE[i], RELEASE_PERIODS_ALTERNATIVE[i + 1])[0]
                     for i in range(len(RELEASE_PERIODS_ALTERNATIVE) - 1)]
        plt.title(gnr)
        plt.bar(RELEASE_PERIODS_ALT_LABELS, num_games)
        plt.xticks(rotation=45)
        plt.subplots_adjust(bottom=0.16)
        plt.savefig("../Pics/detailed2_" + gnr + ".png")

        plt.figure()
        num_games_rel = [num_games[i] / max(num_games_total[i], 1)
                         for i in range(len(RELEASE_PERIODS_ALTERNATIVE) - 1)]
        plt.title(gnr + ' relative')
        plt.bar(RELEASE_PERIODS_ALT_LABELS, num_games_rel)
        plt.xticks(rotation=45)
        plt.subplots_adjust(bottom=0.16)
        plt.savefig("../Pics/reltive2_" + gnr + ".png")

    plt.figure()
    num_last = [0 for i in range(len(RELEASE_PERIODS_ALTERNATIVE) - 1)]

    for gnr in GENRES_NAMES:
        num_games = [que.count_games(conn, gnr, RELEASE_PERIODS_ALTERNATIVE[i], RELEASE_PERIODS_ALTERNATIVE[i + 1])[0] /
                     max(num_games_total[i], 1)
                     for i in range(len(RELEASE_PERIODS_ALTERNATIVE) - 1)]
        plt.title('genres relative')
        plt.bar(RELEASE_PERIODS_ALT_LABELS, num_games, bottom = num_last)
        num_last = [num_games[i] + num_last[i] for i in range(len(RELEASE_PERIODS_ALTERNATIVE) - 1)]
        plt.xticks(rotation=45)
        plt.subplots_adjust(bottom=0.16)
    plt.legend(GENRES_NAMES)
    plt.savefig("../Pics/relative2_genres.png")
