import aiopg
import psycopg2.extras

class Database:
    def __init__(self, dsn: str):
        self.dsn = dsn
        self.pool = None

    async def create_pool(self):
        self.pool = await aiopg.create_pool(self.dsn)

    async def is_registered(self, user_id: int) -> bool:
        async with self.pool.acquire() as conn:
            async with conn.cursor() as cur:
                await cur.execute(
                    "SELECT user_id FROM influence.users WHERE user_id = %s", 
                    (user_id,)
                )
                user = await cur.fetchone()
                return user is not None

    async def register_user(self, user_id: int, name: str):
        async with self.pool.acquire() as conn:
            async with conn.cursor() as cur:
                await cur.execute(
                    "INSERT INTO influence.users (user_id, name) VALUES (%s, %s) "
                    "ON CONFLICT (user_id) DO NOTHING",
                    (user_id, name)
                )