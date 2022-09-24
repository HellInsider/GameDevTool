u_id=$1
input="SteamApiKey.txt"
output="UsersInfo/RecentGamesId_$u_id.txt"

key=`cat ${input}`
url="http://api.steampowered.com/IPlayerService/GetRecentlyPlayedGames/v0001/?key=${key}&steamid=${u_id}$format=json"

curl -L "${url}" -o "${output}"

#takes user id as cmd argument and returns file with recent games of user to UsersInfo
#example u_id = 76561198057071110