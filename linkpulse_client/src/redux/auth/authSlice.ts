import { createSlice, Draft, PayloadAction } from '@reduxjs/toolkit'
import { AuthState, User } from './authTypes'
import { create } from 'domain'

const initialState: AuthState = {
    user: null,
    isLoggedIn: false,
    loading: false,
    error: null
}

const authSlice = createSlice({
    name: 'auth',
    initialState,
    reducers: {
        // registration start
        // registration success
        // registration failed
        // login start

        loginSuccess: (state: Draft<AuthState>, action: PayloadAction<{ user: User }>) => {
            state.user = action.payload.user;
            state.isLoggedIn = true;
        }
        //login failure

        //logout

        //validate session
        
    }
})