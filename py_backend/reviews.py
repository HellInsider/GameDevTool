import matplotlib.pyplot as plt
import que

GENRES_NAMES = ['Adventure', 'RPG', 'Strategy', 'Racing', 'Indie']
RELEASE_PERIODS_ALTERNATIVE = ['2003-01-01', '2012-01-01', '2013-01-01', '2014-01-01',
                               '2015-01-01', '2016-01-01', '2017-01-01', '2018-01-01',
                               '2019-01-01', '2020-01-01', '2021-01-01', '2022-01-01',
                               '2023-01-01']
RELEASE_PERIODS_ALT_LABELS = [RELEASE_PERIODS_ALTERNATIVE[i][0:4] + '-' + str(int(RELEASE_PERIODS_ALTERNATIVE[i + 1][0:4]) - 1)
                              for i in range(len(RELEASE_PERIODS_ALTERNATIVE) - 1)]


def reviews_distrib(conn):
    cursor = conn.cursor()

    cursor.execute('SELECT review_score, COUNT(review_score) '
                   'FROM "DB_schema"."Reviews" '
                   'GROUP BY review_score '
                   'ORDER BY review_score')
    records = cursor.fetchall()
    x_labels = [records[i][0] for i in range(1, len(records))]
    vals = [records[i][1] for i in range(1, len(records))]
    plt.figure()
    plt.title('reviews')
    plt.bar(x_labels, vals)
    plt.savefig("../Pics/reviews.png")

    cursor.close()


def reviews_genres(conn):
    for gnr in GENRES_NAMES:
        records = que.genre_reviews(conn, gnr)
        x_labels = [records[i][0] for i in range(1, len(records))]
        vals = [records[i][1] for i in range(1, len(records))]
        plt.figure()
        plt.title(gnr + ' reviews')
        plt.bar(x_labels, vals)
        plt.savefig("../Pics/reviews_" + gnr + ".png")


def reviews_timeline(conn):
    fig = plt.figure(figsize=(1280 / 80, 720 / 80))

    plt.title('timeline median reviews')
    i = -2
    for gnr in GENRES_NAMES:
        records2 = que.median_genre_review_by_year(conn, gnr)
        rec = records2[2:]
        x = [r[0] for r in rec]
        y = [r[1] + 0.05 * i for r in rec]
        i = i + 1
        plt.plot(x, y, alpha= 0.7, label=gnr, linewidth=2)
        plt.scatter(x, y, alpha=0.5)
    plt.legend()
    plt.savefig("../Pics/timeline_median_reviews.png", dpi = 80)

    fig = plt.figure(figsize=(1280 / 80, 720 / 80))

    plt.title('timeline average reviews')
#    i = -2
    i = 0
    for gnr in GENRES_NAMES:
        records2 = que.avg_genre_review_by_year(conn, gnr)
        rec = records2[2:]
        x = [r[0] for r in rec]
        y = [r[1] + 0.05 * i for r in rec]
#        i = i + 1
        plt.plot(x, y, alpha= 0.7, label=gnr, linewidth=2)
        plt.scatter(x, y, alpha=0.5)
    plt.legend()
    plt.savefig("../Pics/timeline_avg_reviews.png", dpi = 80)
