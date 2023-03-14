# GoLang

PLEX playlist compare for two users

Create folder named "data" and in it urls.txt
http://localIP:32400/playlists/PLAYLISTID1/items?X-Plex-Token=USERTOKEN1
http://localIP:32400/playlists/PLAYLISTID2/items?X-Plex-Token=USERTOKEN2

each url with token in its own row

if you execute the program in the terminal you will get the http get code for each url request (200 for OK or 404 if there was an error)
you can get the playlist id out of your url if you open the playlist in your browser. It will look like playlists%2F85144 - the ID starts after the F
use your own plextoken for both requests
