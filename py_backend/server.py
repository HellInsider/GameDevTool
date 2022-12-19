from flask import Flask
from flask import send_file
from flask import request
import os
import genres
import names_mapping
import que

app = Flask(__name__, static_folder='../static')
conn = None

@app.route("/")
def helo():
    return send_file("../HTML/index.html")

@app.route("/<page>")
def test(page):
    return send_file("../HTML/" + page)

@app.route("/img", methods=['GET'])
def img():
    fname = request.args.get('name')
    fpath = "../Pics/" + fname + ".png"
    if os.path.exists(fpath):
        return send_file(fpath, mimetype="image/png")
    elif names_mapping.PICTURES_FUNCTIONS_MAP.get(fname) is not None:
        (names_mapping.PICTURES_FUNCTIONS_MAP[fname])(conn)
        if os.path.exists(fpath):
            return send_file(fpath, mimetype="image/png")
    return None

# localhost:5000/genre_info?genre=Adventure
@app.route("/genre_info", methods=['GET'])
def aboutGenre():
    genre = request.args.get('genre')
    if genre not in genres.GENRES_NAMES:
        return "<html><head></head><body><h1>ACHTUNG!!! GENRE NOT FOUND</h1></body></html>"

    # I believe that it is an array of triples. First - steamID, second - Name, third - date
    # For example
    # ((570, "Dota 2", "09-06-2013"), (1687950, "Persona 5 Royal", "31-10-2019"), ...)
    games = ((570, "Dota 2", "09-06-2013"), (1687950, "Persona 5 Royal", "31-10-2019"),#que.getBestGamesForGenre(genre)
             (570, "Dota 2", "09-06-2013"), (1687950, "Persona 5 Royal", "31-10-2019"),
             (1687950, "Persona 5 Royal", "31-10-2019"))
    return "<html><head></head><body>" \
           "<h1>" + genre + "</h1>" \
           "<img src=\"genre_detailed_chart/" + genre + "\">" \
           "<img src=\"genre_reltive_chart/" + genre + "\"></br>" \
           "<h1>5 most popular of all time</h1>" \
           "<h2>1.<a href=\"https://store.steampowered.com/app/" + str(
        games[0][0]) + "\">" + games[0][1] + "</a> Release date:" + games[0][2] + "</h2>" \
           "<h2>2.<a href=\"https://store.steampowered.com/app/" + str(
        games[1][0]) + "\">" + games[1][1] + "</a> Release date:" + games[1][2] + "</h2>" \
           "<h2>3.<a href=\"https://store.steampowered.com/app/" + str(
        games[2][0]) + "\">" + games[2][1] + "</a> Release date:" + games[2][2] + "</h2>" \
           "<h2>4.<a href=\"https://store.steampowered.com/app/" + str(
        games[3][0]) + "\">" + games[3][1] + "</a> Release date:" + games[3][2] + "</h2>" \
           "<h2>5.<a href=\"https://store.steampowered.com/app/" + str(
        games[4][0]) + "\">" + games[4][1] + "</a> Release date:" + games[4][2] + "</h2>" \
           "</body></html>"

@app.route("/genre_detailed_chart/<param>", methods=['GET'])
def getDetailedChart(param):
    return send_file("../Pics/detailed2_" + param + ".png", mimetype="image/png")

@app.route("/genre_reltive_chart/<param>", methods=['GET'])
def getReltiveChart(param):
    return send_file("../Pics/reltive2_" + param + ".png", mimetype="image/png")

@app.route("/<param>", methods=['GET'])
def index(param):
    return "<html><head></head><body><p>" + param + " shit </p>" \
           "<img src=\"pic\" alt=\"alternatetext\"></body></html>"

@app.route("/pic")
def helou():
    return send_file("../Pics/genres_piechart.png", mimetype="image/png")

def server():
    names_mapping.generate_mapping()
    global conn
    conn = que.init_dbtools()

  #  app.config['TEMPLATES_AUTO_RELOAD'] = True
  #  app.config['SEND_FILE_MAX_AGE_DEFAULT'] = 0

    app.run()
    conn.close()