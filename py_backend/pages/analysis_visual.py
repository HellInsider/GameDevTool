import que
import globals

import flask
from dash import Dash, html, dcc, Input, Output, State, register_page, callback
from datetime import date


register_page(__name__)
genres_list = globals.data_storage.genres_list #['rpg', 'action', 'shooter', 'strategy']
attribs_list = ['genre', 'date', 'rewiews', 'price']
prices_range = (0, 1000)

layout = html.Div([
    dcc.Checklist(genres_list, id='chklst', style= {'display': 'flex', 'flex-wrap': 'wrap', 'align-content': 'left', 'justify-content': 'left'}),
    html.Div(children='select date'),
    dcc.DatePickerRange(
        id='release-date',
        start_date=date(2017, 5, 3),
        end_date=date.today()
    ),
    dcc.RangeSlider(prices_range[0], prices_range[1], id='price'),
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


@callback(Output('tabscontent', 'children'),
          Input('tabs', 'value'))
def render_content(tab):
    if tab == 'tab-1':
        return html.Div([
            html.P('Rewiews depends on:'),
            dcc.Dropdown(attribs_list, id='tab1-attrib'),
            html.Div(id = 'content0'),
            dcc.Graph(
                figure={
                    'data': [{
                        'x': [1, 2, 3],
                        'y': [3, 1, 2],
                        'type': 'plot'
                    }]
                }
            )
        ])
    elif tab == 'tab-2':
        return html.Div(
            id='tab2div', children = [
            html.H3('Distribution by:'),
            dcc.Dropdown(attribs_list, id='tab2-attrib'),
            html.Div(id = 'content'),
            dcc.Graph(
                id='graph-2-tabs-dcc',
                figure={
                    'data': [{
                        'x': [1, 2, 3],
                        'y': [5, 10, 6],
                        'type': 'plot'
                    }]
                }
            )
        ])

@callback(Output("tab2div", 'children')
    , Input("button-submit", 'n_clicks')
    , State("chklst", 'value')
    , State("release-date", 'start_date')
    , State("release-date", 'end_date')
    , State("price", 'value')
    )
def submit_message(n, chekss, start, end, price):
    print(n)
    if (n == 0):
        return None
    # print(droplist)
    print(chekss)
    print(start)
    print(end)
    print(price)
    return html.P(f"{que.records}")
