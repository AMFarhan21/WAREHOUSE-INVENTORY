'use client'

import { useState } from "react";

const useLogin = () => {
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null)

    const login = async (email: string, password: string) => {
        try {
            setLoading(true)
            setError(null)

            const API_URL = process.env.NEXT_PUBLIC_API_URL

            const res = await fetch(`${API_URL}/user/login`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    email,
                    password
                })
            })

            const data = await res.json()

            if (!res.ok) {
                setError(data.message || "Login failed");
                return null
            }

            return data.data
        } catch (err) {
            setError("Email or password in invalid")
            console.log(err)
            return null;
        } finally {
            setLoading(false)
        }
    }


    return { login, loading, error }
}

export default useLogin