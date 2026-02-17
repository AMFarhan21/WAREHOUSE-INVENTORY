import React, { useState } from 'react'
import useGetAllBarang, { MasterBarang } from './useGetAllBarang'

const useCreateBarang = () => {
    // const [createdBarang, setCreatedBarang] = useState<MasterBarang | null>(null)
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState("")
    const API_URL = process.env.NEXT_PUBLIC_API_URL


    const createBarang = async (nama_barang: string, deskripsi: string, satuan: string, harga_beli: number, harga_jual: number) => {
        try {
            setLoading(true)
            const token = localStorage.getItem("TOKEN")
            const res = await fetch(`${API_URL}/api/barang`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`
                },
                body: JSON.stringify({ nama_barang, deskripsi, satuan, harga_beli, harga_jual })
            })

            const data = await res.json()

            if (!res.ok) {
                setError(data.message)
                return { success: false, message: data.message }
            }

            return { success: true, data: data.data }
        } catch (error) {
            console.log(error)

        } finally {
            setLoading(false)
        }
    }

    return { createBarang, loading, error, setError }
}

export default useCreateBarang