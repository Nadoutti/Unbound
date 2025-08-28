from fastapi import APIRouter, HTTPException 
from controllers.payin_controller import PayinController
from models.Payin import (
    PayinRead,
    PayinRequest
)
from dotenv import load_dotenv
import httpx

load_dotenv()


router = APIRouter(prefix="/payin", tags=["noticias"])
controller = PayinController()


@router.post("/initiate", response_model=PayinRead)
async def initiate_payin(request: PayinRequest):
    return await PayinController.initiate_payin(request)

