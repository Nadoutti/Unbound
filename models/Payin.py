from pydantic import BaseModel


class PayinRequest(BaseModel):
    amount: float  # Amount in BRL
    currency_type: str  # Default to BRL
    payment_method: str
    request_amount: int
    token: str = "USDT"  # Target stablecoin


class PayinRead(BaseModel):
    status: str;
    payin_id: str;
    pix_code: str;
    payment_track: str;
    reciever_id: str;
