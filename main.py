from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware


# Create FastAPI application instance
app = FastAPI()


# Add CORS middleware to allow cross-origin requests
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  
    allow_credentials=True,  
    allow_methods=["*"],  
    allow_headers=["*"],  
    expose_headers=["*"],  
)

