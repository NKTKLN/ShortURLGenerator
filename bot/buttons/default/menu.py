from aiogram import types
from aiogram.types import *
from dotenv import load_dotenv
from database.postgres import *
# Buttons for the logged user
keyboardLogged = types.ReplyKeyboardMarkup(resize_keyboard=True)
keyboardLogged.add(types.KeyboardButton(text='🗄 Archive'))
keyboardLogged.add(types.KeyboardButton(text='👤 Sign out'))
# Authorization buttons
keyboardLogin = types.ReplyKeyboardMarkup(resize_keyboard=True)
keyboardLogin.add(types.KeyboardButton(text='🙎‍♂️ Login'))

inlineBtnStart = InlineKeyboardMarkup(row_width=1)
loginBtn = InlineKeyboardButton('🙎‍♂️ Login', callback_data='login')
anonymousBtn = InlineKeyboardButton('👤 Anonymous', callback_data='anonym')
inlineBtnStart.add(anonymousBtn, loginBtn)

inlineBtnChoice = InlineKeyboardMarkup(row_width=1)
yes = InlineKeyboardButton('Yes', callback_data='yes')
no = InlineKeyboardButton('No!', callback_data='no')
inlineBtnChoice.add(yes, no)
# Create a menu for the archive
def archiveMenu(telegramUserId, page):
    cursor.execute(f'SELECT id FROM users WHERE telegramId={telegramUserId};')
    userId = cursor.fetchone()
    # Authorization check
    if userId:
        cursor.execute(f'SELECT * FROM shorturl WHERE userId={userId[0]};')
        archive = cursor.fetchall()
        if archive:
            # Menu generation
            load_dotenv('env/app.env')
            inlineBtnArchive = InlineKeyboardMarkup(row_width=4)
            inlineBtnArchive.add(InlineKeyboardButton('ID', callback_data='#'), InlineKeyboardButton('URL', callback_data='#'), InlineKeyboardButton('VISITS', callback_data='#'), InlineKeyboardButton('DELETE', callback_data='#'))
            if len(archive) <= 10:
                for link in archive:
                    inlineBtnArchive.add(InlineKeyboardButton(link[0], callback_data='#'), InlineKeyboardButton('click', url=link[1]), InlineKeyboardButton(link[3], callback_data='#'), InlineKeyboardButton('🗑', callback_data=f'del {link[0]}'))
            else:
                pages = len(archive) // 10 if len(archive) % 10 == 0 else len(archive) // 10 + 1
                if pages >= page:
                    if len(archive) < page * 10:
                        for index in range(10 * (page - 1), len(archive)):
                            inlineBtnArchive.add(InlineKeyboardButton(archive[index][0], callback_data='#'), InlineKeyboardButton('click', url=archive[index][1]), InlineKeyboardButton(archive[index][3], callback_data='#'), InlineKeyboardButton('🗑', callback_data=f'del {archive[index][0]}'))
                    else:
                        for index in range(10 * (page - 1), 10 * page):
                            inlineBtnArchive.add(InlineKeyboardButton(archive[index][0], callback_data='#'), InlineKeyboardButton('click', url=archive[index][1]), InlineKeyboardButton(archive[index][3], callback_data='#'), InlineKeyboardButton('🗑', callback_data=f'del {archive[index][0]}'))
                else:
                    inlineBtnArchive.row(InlineKeyboardButton('There is no such page'), callback_data='#')
                inlineBtnArchive.row(InlineKeyboardButton('<<', callback_data='#' if page - 1 == 0 else f'page {page - 1}'), InlineKeyboardButton(f'{page}/{pages}', callback_data='#'), InlineKeyboardButton('>>', callback_data='#' if page + 1 > pages else f'page {page + 1}'))
            return 'Archive', inlineBtnArchive
        else:
            return "You don't have links yet.", None
    else:
        return 'Ha-ha-ha. A cunning plan.', None
