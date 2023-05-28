import que
import globals

from dash import Dash, page_container, page_registry, html, dcc

app = Dash(__name__, use_pages=True)


app.layout = html.Div(id='global-scope', children=[
                      html.H1('Multi-page app with Dash Pages'),

                      html.Div(
                          [
                              html.Div(
                                  dcc.Link(
                                      f"{page['name']}", href=page["relative_path"]
                                  )
                              )
                              for page in page_registry.values()
                          ]
                      ),

                      page_container
                      ])

if __name__ == '__main__':
    global conn
    conn = que.que_instance.connection

    app.run_server(debug=False, port=80)
    conn.close()

#if __name__ == '__main__':
    #conn = que.init_dbtools()
    #genres.genres_pie(conn)
    #genres.genres_paper(conn)
    #genres.genres_addition(conn)
    #genres.genres_detailed(conn)
    #reviews.reviews_distrib(conn)
    #reviews.reviews_genres(conn)
    #reviews.reviews_timeline(conn)

    #cursor = conn.cursor()
    #cursor.execute('SELECT * '
    #               'FROM "DB_schema"."Games" '
    #               'WHERE appid = ANY(SELECT dlc FROM "DB_schema"."Games" WHERE appid=400) ')
    #records = cursor.fetchall()
    #records = cursor.fetchmany(300)

    #server.server()

    #plt.show()
    #cursor.close()
    #conn.close()

