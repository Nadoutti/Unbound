import os
from fastapi import HTTPException
import httpx


BLINDPAY_API_URL = "https://api.blindpay.com/v1"
INSTANCE_ID = os.getenv("INSTANCE_ID")
API_KEY = os.getenv("API_KEY")
WALLET_KEY = os.getenv("WALLET_KEY")



class PayinService:
    def __init__(self) -> None:
        
        pass



    async def blindpay_api_post(self, endpoint: str, data: dict) -> dict:
        headers = {
            "Authorization": f"Bearer {API_KEY}",
            "Content-Type": "application/json"
        }
        async with httpx.AsyncClient() as client:
            response = await client.post(f"{BLINDPAY_API_URL}{endpoint}", json=data, headers=headers)
            if response.status_code != 200:
                raise HTTPException(status_code=response.status_code, detail=response.json().get("error", "BlindPay API error"))
            return response.json()

    async def initiate_payin(self, request):

        # creating the quote

        quote_endpoint = f"/instances/{INSTANCE_ID}/payin-quotes"

        quote_data = {
            "blockchain_wallet_id": WALLET_KEY,
            "currency_type": request.currency_type,
            "payment_method": request.payment_method,
            "request_amount": request.request_amount,
            "token": request.token,

        }

        create_quote = await self.blindpay_api_post(quote_endpoint, quote_data)






