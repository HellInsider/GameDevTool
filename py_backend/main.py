import que
import globals

from dash import Dash, page_container, page_registry, html, dcc
import dash_bootstrap_components as dbc

app = Dash(__name__, use_pages=True)


app.layout = html.Div(id='global-scope', children=[
                      html.Div([
                          html.H1('Game Dev Tool', style={'color':'rgb(255, 255, 255)'}),
                          html.Div([
                              html.Div(
                                  html.A(
                                      html.Button(
                                          f"{page['name']}"
                                      ), href=page["relative_path"]
                                  )
                              )
                              for page in page_registry.values()
                          ], style={'display': 'flex', 'flex-wrap': 'wrap'})],
                          style={'background-color': 'rgba(64, 64, 64, 0.75)', 'max-width': '70%',
                                 'position': 'relative', 'top': '15%', 'left': '15%',
                                 'padding': '10px 30px 10px 30px', 'text-align': 'center'}
                      ),
                      page_container
                      ])

if __name__ == '__main__':
    global conn
    conn = que.que_instance.connection
    arr1 = [str(i) for i in range(1)]
    #print('WHERE g.description = \''
    #                           + '\' OR g.description = \''.join(arr1) + '\') as gi ')

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

