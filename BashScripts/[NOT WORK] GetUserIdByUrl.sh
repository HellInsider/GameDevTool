u_url=$1
input="SteamApiKey.txt"
output="UsersInfo/UserInfoId_$u_id.txt"

key=`cat ${input}`
url="http://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=${key}&vanityurl=${u_url}$format=json"

echo "curl -L ${url}"

sleep 3
#takes user id as cmd argument and returns file of user to UsersInfo
#example u_id = 76561198057071110

http://api.steampowered.com/IPlayerService/GetRecentlyPlayedGames/v0001/?key=33CF43A0E8B0A89B488CCE3063DED7FC&steamid=76561198057071110&format=json