# friend-tech-utils
hacking away at friend tech

want to create some trackers for friend tech.

sniper bot

## deps
you'll need a db with block transactions in a very specific format. if this is on git look for `firstrepo`. get that running first. this should run in parallel to that


## getting started
1) youll need some env vars. look at db
2) run main.go.  should auto start printing out any users over 1k followers.


## future plans
- i quickly realized that sniping non crypto accounts is actually pretty dogwater. followers isnt the most valuable thing but rather how "in" someone is. 
- so i want to either (or maybe both):
    - add a "crypto quality" number. maybe some ml to decide if an account is quality
    - embedding lookup. create a vector db of good and bad crypto accounts. for each new account, create embedding and look up how similar to quality accounts. pfp may be valuable here
