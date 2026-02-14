'use client'

import { createContext, ReactNode, useContext, useEffect, useState } from "react"

interface AuthContextType {
    token: string | null,
    setToken: (token: string | null) => void;
    isAuthenticated: boolean;
    loading: boolean,
}

const AuthContext = createContext<AuthContextType>({
    token: null,
    setToken: () => { },
    isAuthenticated: false,
    loading: true,
});

export const AuthProvider = ({ children }: { children: ReactNode }) => {
    const [token, setToken] = useState<string | null>(null);
    const [loading, setLoading] = useState(true)

    useEffect(() => {
        const storedToken = localStorage.getItem("TOKEN")

        if (storedToken) {
            setTimeout(() => setToken(storedToken), 0)
        }

        setTimeout(() => setLoading(false), 0)
    }, [])

    const isAuthenticated = !!token

    return (
        <AuthContext.Provider value={{ token, setToken, isAuthenticated, loading }}>
            {children}
        </AuthContext.Provider>
    )
}

export const useAuth = () => useContext(AuthContext);