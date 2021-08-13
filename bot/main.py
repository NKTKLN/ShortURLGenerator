import os
import redis
import logging
import requests
from pg import *
from buttons import *
from dotenv import load_dotenv
from aiogram import Bot, Dispatcher, executor, types
# Connecting the bot to the api
load_dotenv('env/bot.env')
logging.basicConfig(level=logging.INFO)
bot = Bot(token=os.getenv('API_TOKEN'))
dp = Dispatcher(bot)
# /start command
@dp.message_handler(commands=['start'])
async def start(message: types.Message):
    cursor.execute(f'SELECT id FROM users WHERE telegramId={message.from_user.id};')
    # Check for a user in the database
    if cursor.fetchone():
        await message.answer("Hi I'm glad to see you again!", reply_markup=keyboardLogged)
    elif len(message.text.split()) > 1:
        # Get the user id from the key
        client = redis.StrictRedis(host='redis', port=6379, db=0)
        val = client.get(message.text.split()[1])
        if val:
            # Adding a user's telegram id to the database
            cursor.execute(f'UPDATE users SET telegramId={message.from_user.id} WHERE id={int(val)};')
            conn.commit()
            await message.answer("It's ok you logged successfully.", reply_markup=keyboardLogged)
        else:
            await message.answer(f"I think the time has expired and your key is not valid or you logged in from another account, try logging in again please.\n<a href='https://{os.getenv('VIRTUAL_HOST')}/telegramLogin'>Click</a>", parse_mode='HTML')
    else:
        load_dotenv('env/app.env')
        await message.answer(f"Hi!\nI see you decided to start me up for a more user-friendly service <a href='https://{os.getenv('VIRTUAL_HOST')}'>{os.getenv('VIRTUAL_HOST')}</a>.\nFor starters, do you want to use your account when creating links or be anonymous. Select the button at the bottom:", reply_markup=inlineBtnStart, parse_mode='HTML')
# User authentication
@dp.callback_query_handler(lambda message: message.data == 'anonym')
async def anonym(message: types.CallbackQuery):
    await bot.edit_message_text(chat_id=message.from_user.id, message_id=message.message.message_id, text='Yes, I can understand you. I, too, sometimes want to stay in the shadows and not leave a trail behind me.\nIf necessary, you will always be able to log into your account, at the bottom there is a button for login.')
    await bot.send_message(chat_id=message.from_user.id, text='All right, here we go. To shorten the link just send it to me.  If the link is working, I will send you a shortened link.', reply_markup=keyboardLogin)

@dp.callback_query_handler(lambda message: message.data == 'login')
async def login(message: types.CallbackQuery):
    load_dotenv('app.env')
    await bot.edit_message_text(chat_id=message.from_user.id, message_id=message.message.message_id, text=f"If you decide to sign into your account, it's a good thing because it gives you more control over your links.\n<a href='https://{os.getenv('VIRTUAL_HOST')}/telegramLogin'>Click</a>", parse_mode='HTML')

@dp.message_handler(lambda message: message.text == '🙎‍♂️ Login')
async def state(message: types.Message):
    load_dotenv('app.env')
    await message.answer(f"If you decide to sign into your account, it's a good thing because it gives you more control over your links.\n<a href='https://{os.getenv('VIRTUAL_HOST')}/telegramLogin'>Click</a>", parse_mode='HTML')
# Sign out
@dp.message_handler(lambda message: message.text == '👤 Sign out')
async def state(message: types.Message):
    await message.answer('Go to the main menu.', reply_markup=inlineBtnChoice)

@dp.callback_query_handler(lambda message: message.data == 'yes')
async def signOut(message: types.CallbackQuery):
    cursor.execute(f'UPDATE users SET telegramId=null WHERE telegramId={message.from_user.id};')
    conn.commit()
    await bot.edit_message_text(chat_id=message.from_user.id, message_id=message.message.message_id, text='For starters, do you want to use your account when creating links or be anonymous. Select the button at the bottom:', reply_markup=inlineBtnStart)

@dp.callback_query_handler(lambda message: message.data == 'no')
async def signOutNo(message: types.CallbackQuery):
    await bot.edit_message_text(chat_id=message.from_user.id, message_id=message.message.message_id, text='Ok')
# Archive
@dp.message_handler(lambda message: message.text == '🗄 Archive')
async def archive(message: types.Message):
    messageText, messageButtons = archiveMenu(message.from_user.id, 1)
    await message.answer(messageText, reply_markup=messageButtons)

@dp.callback_query_handler(lambda message: 'page' in message.data)
async def archive(message: types.CallbackQuery):
    messageText, messageButtons = archiveMenu(message.from_user.id, int(message.data.split()[1]))
    await bot.edit_message_text(chat_id=message.from_user.id, message_id=message.message.message_id, text=messageText, reply_markup=messageButtons)
# Delete url button
@dp.callback_query_handler(lambda message: 'del' in message.data)
async def linkDelete(message: types.CallbackQuery):
    cursor.execute(f'SELECT id FROM users WHERE telegramId={message.from_user.id};')
    userId = cursor.fetchone()
    # Authorization check
    if userId:
        # Url deletion
        urlId = message.data.split()[1]
        cursor.execute(f"DELETE FROM shorturl WHERE userId={userId[0]} AND id='{urlId}';")
        conn.commit()
        await bot.edit_message_text(chat_id=message.from_user.id, message_id=message.message.message_id, text=f'The link with the id "{urlId}" been deleted.')
    else:
        await message.answer('Ha-ha-ha. A cunning plan.')
# Generating the link along with adding it to the database
@dp.message_handler()
async def radius(message: types.Message):
    try:
        # Check the link to see if it works 
        requests.get(message.text)
        cursor.execute(f'SELECT id FROM users WHERE telegramId={message.from_user.id};')
        userId = cursor.fetchone()
        # Saving the link in the database
        load_dotenv('env/app.env')
        await message.answer(f"Here is your link.\n\n<code>https://{os.getenv('VIRTUAL_HOST')}/{UrlIdGenerator(message.text, userId[0] if userId else 0)}</code>", parse_mode='HTML')
    except Exception:
        await message.answer('Sorry, but your link is not correct.')
# Running the bot
if __name__ == "__main__":
    executor.start_polling(dp, skip_updates=False)
