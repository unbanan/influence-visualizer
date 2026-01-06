from aiogram.fsm.state import State, StatesGroup

class UserStates(StatesGroup):
    unregistered = State()
    waiting_for_name = State()
    registered = State()