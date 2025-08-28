from fastapi import HTTPException, status
from services.payin_services import PayinService
import os
from models.Payin import (
    PayinRead,
    PayinRequest
)


class PayinController:
    def __init__(self, service: PayinService | None = None):
        self.service = service or PayinService()


    async def initiate_payin(self, request):
        response = await self.service.initiate_payin(request)

        if not response:
            raise HTTPException(status_code=400, detail="Os dados tao mal formados, presta atencao mlk")




