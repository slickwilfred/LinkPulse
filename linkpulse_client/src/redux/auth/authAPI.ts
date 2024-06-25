import { User, LoginUser, RegisterUser, ResetPasswordUrl, ValidationResponse } from './authTypes'
import dotenv from 'dotenv'
dotenv.config()

const BASE_URL = `${process.env.API_URL}/auth`;


export const register = async (credentials: RegisterUser): Promise<any> => {
    const response = await fetch(`${BASE_URL}/register`, {
        method: 'POST',
        mode: 'cors',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(credentials),
    })

    const data = await response.json()

    if (!response.ok) {
        const errorMessage = data.message || 'Registration failed';
        console.error('Registration failed:', errorMessage);
        throw new Error(errorMessage);
    }

    return data
}

export const login = async (credentials: LoginUser): Promise<{ user: User }> => {
    const response = await fetch(`${BASE_URL}/login`, {
        method: 'POST',
        mode: 'cors',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(credentials),
    })

    const data = await response.json()

    if (!response.ok) {
        const errorMessage = data.message || 'Login failed';
        console.error('Login failed:', errorMessage);
        throw new Error(errorMessage);
    }

    return data
}


export const logout = async (): Promise<any> => {
    const response = await fetch(`${BASE_URL}/logout`, {
        method: 'GET',
        mode: 'cors',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
        },
    })

    const data = await response.json()

    if (!response.ok) {
        const errorMessage = data.message || 'Logout failed'
        console.error('Logout failed:', errorMessage)
        throw new Error(errorMessage)
    }

    return data
}


export const resetPassword = async (email: string): Promise<{ resetPasswordUrl: ResetPasswordUrl }> => {
    const response = await fetch(`${BASE_URL}/reset-password`, {
        method: 'POST',
        mode: 'cors',
        headers: {
            'Content-Type': 'application/json', 
        },
        body: JSON.stringify(email)
    })

    const data = await response.json()

    if (!response.ok) {
        const errorMessage = data.message || 'Reset password request failed'
        console.error('Reset password failed:', errorMessage)
        throw new Error(errorMessage) 
    }

    return data
}

export const validateSession = async (): Promise<{ validationResponse: ValidationResponse }> => {
    const response = await fetch(`${BASE_URL}/validate`, {
        method: 'GET',
        mode: 'cors',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json'
        }
    })

    const data = await response.json()

    if (!response.ok) {
        const errorMessage = data.message
        console.error('Session validation failed:', errorMessage)
        throw new Error(errorMessage)
    }

    return { validationResponse:{
            status: response.status,
            user: data.user
        } 
    }
}

