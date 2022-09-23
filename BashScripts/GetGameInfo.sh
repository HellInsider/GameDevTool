id=$1
url="http://store.steampowered.com/api/appdetails?appids=${id}" 
output="GamesInfo/GameInfoId_$id.txt"
curl -L "${url}" -o "${output}"
