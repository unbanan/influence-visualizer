import os
import asyncio
from dotenv import load_dotenv
from aiogram import Bot, Dispatcher, types, F
from aiogram.filters import Command
from aiogram.fsm.context import FSMContext
from aiogram.fsm.state import default_state, State

from states import UserStates
from keyboards import get_init_kb, get_reg_kb, get_unreg_kb
from database import Database

class MyBot:
    def __init__(self, token: str, db_url: str):
        self.bot = Bot(token=token)
        self.dp = Dispatcher()
        self.db = Database(db_url)
        self.setup_handlers()

    def setup_handlers(self):
        @self.dp.message(default_state)
        async def cmd_reboot_recovery(message: types.Message, state: FSMContext):
            if message.text == "Старт" or message.text == "/start":
                user_id = message.from_user.id
                is_reg = await self.db.is_registered(user_id)
                if is_reg:
                    await state.set_state(UserStates.registered)
                    await message.answer("Меню обновлено.", reply_markup=get_reg_kb())
                else:
                    await state.set_state(UserStates.unregistered)
                    await message.answer("Вы не зарегистрированы. Нажмите 'Регистрация'.", reply_markup=get_unreg_kb())
            else:
                await message.answer(
                    "Бот был обновлен. Нажмите кнопку 'Старт' для синхронизации.",
                    reply_markup=get_init_kb()
                )

        @self.dp.message(F.text == "Старт", UserStates.unregistered)
        async def reg_info(message: types.Message, state: FSMContext):
            await message.answer("Вам необходимо зарегистрироваться. Нажмите 'Регистрация'.")

        @self.dp.message(F.text == "Регистрация", UserStates.unregistered)
        async def start_reg(message: types.Message, state: FSMContext):
            await message.answer("Введите ваше имя:")
            await state.set_state(UserStates.waiting_for_name)

        @self.dp.message(UserStates.waiting_for_name)
        async def process_name(message: types.Message, state: FSMContext):
            await self.db.register_user(message.from_user.id, message.text)
            await state.set_state(UserStates.registered)
            await message.answer(f"Готово, {message.text}!", reply_markup=get_reg_kb())
        
        @self.dp.message(UserStates.registered)
        async def handle_command(message: types.Message, state: FSMContext):
            if not await self.db.is_registered(message.from_user.id):
                await state.set_state(UserStates.unregistered)
                await message.answer("Необходимо заново зарегистрироваться.", reply_markup=get_unreg_kb())
                return
            await message.answer(f"Команда: {message.text}")


    async def run(self):
        await self.db.create_pool()
        await self.dp.start_polling(self.bot)

if __name__ == "__main__":
    load_dotenv()
    token = os.getenv("TELEGRAM_TOKEN")
    db_url = os.getenv("DATABASE_URL")
    bot = MyBot(token, db_url)
    asyncio.run(bot.run())