'use client'
import { useState } from 'react'

const useDeleteBarang = () => {
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState("")
    const API_URL = process.env.NEXT_PUBLIC_API_URL
    const deleteBarang = async (barangID: number) => {
        try {
            setLoading(true)
            const token = localStorage.getItem("TOKEN")
            const res = await fetch(`${API_URL}/api/barang/${barangID}`, {
                method: "DELETE",
                headers: {
                    Authorization: `Bearer ${token}`
                }
            })

            const data = await res.json()
            if (!res.ok) {
                setError(data.message)
                return
            }

            return data.message
        } catch (error) {
            console.log(error)
        } finally {
            setLoading(false)
        }
    }

    return { deleteBarang, loading, error }
}

export default useDeleteBarang