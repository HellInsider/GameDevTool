import que
import globals

import flask
from dash import Dash, html, dcc, Input, Output, State, register_page, callback
from datetime import date


register_page(__name__)
genres_list = globals.data_storage.genres_list
review_attribs_list = globals.data_storage.review_attribs_list
distrib_attribs_list = globals.data_storage.distribution_attribs_list
prices_range = (0, 1000)

layout = html.Div([
    dcc.Checklist(genres_list, id='chklst', style= {'display': 'flex', 'flex-wrap': 'wrap', 'align-content': 'left', 'justify-content': 'left'}),
    html.Div(children='select date'),
    dcc.DatePickerRange(
        id='release-date',
        start_date=date(2017, 5, 3),
        end_date=date.today()
    ),
    #dcc.RangeSlider(prices_range[0], prices_range[1], id='price'),
    dcc.Tabs(id="tabs", value='tab-1', children=[
        dcc.Tab(label='plot', value='tab-1'),
        dcc.Tab(label='bar', value='tab-2'),
    ]),
    html.Div(id='tabscontent'),
    html.Br(),
    html.Div(id='div-button', children=[
        html.Button('Submit'
                   , id='button-submit'
                   , n_clicks=0)])

    ], style = {'max-width': '70%', 'position': 'relative', 'top': '40px', 'left': '40px'})


def bar_query(attrib):
    que.records.clear()

    if (attrib not in distrib_attribs_list):
        attrib = distrib_attribs_list[0]
    if (attrib == 'year'):
        que.count_games_by_year(que.que_instance.connection)
    elif (attrib == 'genre'):
        que.count_games_by_genre(que.que_instance.connection)
    elif (attrib == 'reviews'):
        que.count_games_by_reviews(que.que_instance.connection)

@callback(Output('tabscontent', 'children'),
          Input('tabs', 'value'))
def render_content(tab):
    if tab == 'tab-1':
        return html.Div(id='tab1div', children = [
            html.P('Rewiews depends on:'),
            dcc.Dropdown(review_attribs_list, id='tab1-attrib'),
            html.Div(id = 'tab1content'),
        ])
    elif tab == 'tab-2':
        return html.Div(
            id='tab2div', children = [
            html.H3('Distribution by:'),
            dcc.Dropdown(distrib_attribs_list, id='tab2-attrib'),
            html.Div(id = 'tab2content'),
        ])

@callback(Output("tab1content", 'children')
    , Input("button-submit", 'n_clicks')
    , State("chklst", 'value')
    , State("release-date", 'start_date')
    , State("release-date", 'end_date')
    #, State("price", 'value')
    )
def tab1_submit(n, chekss, start, end, price):
    print(n)
    if (n == 0):
        return None

    que.records = []
    que.avg_review_by_year(que.que_instance.connection)
    return dcc.Graph(
        figure={
            'data': [{
                'x': [r[0] for r in que.records],
                'y': [r[1] for r in que.records],
                'type': 'plot'
            }]
        }
    )

    #return html.P(f"{que.records}")

@callback(Output("tab2content", 'children')
    , Input("button-submit", 'n_clicks')
    , State("chklst", 'value')
    , State("release-date", 'start_date')
    , State("release-date", 'end_date')
    , State("price", 'value')
    , State("tab2-attrib", 'value')
    )
def tab2_submit(n, chekss, start, end, price, attrib):
    print(n)
    if (n == 0):
        return None

    bar_query(attrib)
    return dcc.Graph(
        figure={
            'data': [{
                'x': [r[0] for r in que.records],
                'y': [r[1] for r in que.records],
                'type': 'bar'
            }]
        }
    )
