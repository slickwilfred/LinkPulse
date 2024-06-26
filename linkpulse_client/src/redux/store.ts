import { configureStore } from '@reduxjs/toolkit'
import authReducer from './auth/authSlice'@types/express

export const store = configureStore({
    reducer: {
        auth: authReducer
    },
});