export interface AuthState {
    user:       User | null;
    isLoggedIn: boolean;
    loading:    boolean;
    error:      string | null;
    //token: string | null;
}

export interface User {
    id:      string;
    name:    string;
    company: string;
    email:   string;
    tier:    number;
}

export interface RegisterUser {
    name:     string;
    email:    string;
    password: string;
    company:  string;
    dob:      string;
}

export interface LoginUser {
    email:    string;
    password: string;
}

export interface ResetPasswordUrl {
    url: string | null;
}

export interface ValidationResponse {
    status: number;
    user?:  User;
    error?: string;
}
