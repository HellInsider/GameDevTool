import que

class data:
    genres_list = []
    def __init__(self):
        conn = que.que_instance.connection
        cursor = conn.cursor()
        cursor.execute('SELECT description FROM "DB_schema"."Genre"')
        que.records = cursor.fetchall()
        # records = cursor.fetchmany(300)
        self.genres_list = [r[0] for r in que.records]
        cursor.close()


data_storage = data()