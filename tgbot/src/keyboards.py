from aiogram.types import ReplyKeyboardMarkup, KeyboardButton
from aiogram.utils.keyboard import ReplyKeyboardBuilder

def get_init_kb() -> ReplyKeyboardMarkup:
    builder = ReplyKeyboardBuilder()
    builder.row(KeyboardButton(text="Старт"))
    return builder.as_markup(resize_keyboard=True)

def get_unreg_kb() -> ReplyKeyboardMarkup:
    builder = ReplyKeyboardBuilder()
    builder.row(KeyboardButton(text="Старт"), KeyboardButton(text="Регистрация"))
    return builder.as_markup(resize_keyboard=True)

def get_reg_kb() -> ReplyKeyboardMarkup:
    builder = ReplyKeyboardBuilder()
    builder.row(KeyboardButton(text="Старт"))
    builder.row(KeyboardButton(text="Отправить"), KeyboardButton(text="Посмотреть посылки"))
    builder.row(KeyboardButton(text="Посмотреть Статистику"))
    return builder.as_markup(resize_keyboard=True)