import os
import logging
from dotenv import load_dotenv
from aiogram import Bot, Dispatcher
# Connecting the bot to the api
load_dotenv('env/bot.env')
logging.basicConfig(level=logging.INFO)
bot = Bot(token=os.getenv('API_TOKEN'))
dp = Dispatcher(bot)
