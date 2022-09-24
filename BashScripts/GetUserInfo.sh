u_id=$1
input="SteamApiKey.txt"
output="UsersInfo/UserInfoId_$u_id.txt"

key=`cat ${input}`

url="https://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=${key}&include_appinfo=true&steamid=${u_id}&format=json"

curl -L "${url}" -o "${output}"
#takes user id as cmd argument and returns file of user to UsersInfo
#example u_id = 76561198057071110