id=$1
url="http://store.steampowered.com/api/appdetails?appids=${id}" 
output="GamesInfo/GameInfoId_$id.txt"
curl -L "${url}" -o "${output}"

#takes game id as cmd argument and returns file of game to GamesInfo