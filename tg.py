from pyrogram import Client , filters
appid = "4"
apihash="014b35b6184100b085b0d0572f9b5103"
app = Client(
    "account",
    api_id=appid, api_hash=apihash,
)
with app:
    session_string = app.export_session_string() 
    #print(session_string)

with Client("account", session_string=session_string) as app:
    new_firstname = '🐤S.Gra-Gra'
    app.update_profile(first_name=new_firstname)
    #print(app.get_me())
    print(app.join_chat("grafunmeme"))
    msg_result =  app.send_message("GraFunBot", "/start  n_GKr4FkM49h5BFxpuGjgT3ryNzXNPeLG3V72PhEtBU=")
    print(msg_result)

@app.on_message(filters.text & filters.private)
async def echo(client, message):
    print(message.text)

app.run()
