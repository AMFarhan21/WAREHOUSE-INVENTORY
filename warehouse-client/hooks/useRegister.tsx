'use client'

import { useState } from "react";

const useRegister = () => {
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null)

    const register = async (username: string, password: string, email: string, fullname: string) => {
        try {
            setLoading(true)
            setError(null)

            const API_URL = process.env.NEXT_PUBLIC_API_URL

            const res = await fetch(`${API_URL}/user/register`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    username,
                    password,
                    email,
                    full_name: fullname,
                })
            })

            const data = await res.json()

            if (!res.ok) {
                setError(data.message || "Registered failed");
                return null
            }

            return data.data
        } catch (err) {
            if (err instanceof Error) {
                setError(err.message)
            }
            console.log(err)
            return null;
        } finally {
            setLoading(false)
        }
    }


    return { register, loading, error }
}

export default useRegister