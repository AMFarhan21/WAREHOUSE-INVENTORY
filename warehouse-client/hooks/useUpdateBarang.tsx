import React, { useState } from 'react'

const useUpdateBarang = () => {
    // const [createdBarang, setCreatedBarang] = useState<MasterBarang | null>(null)
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState("")
    const API_URL = process.env.NEXT_PUBLIC_API_URL


    const updateBarang = async (barangID: number, nama_barang: string, deskripsi: string, satuan: string, harga_beli: number, harga_jual: number) => {
        try {
            setLoading(true)
            const token = localStorage.getItem("TOKEN")
            const res = await fetch(`${API_URL}/api/barang/${barangID}`, {
                method: "PUT",
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

    return { updateBarang, loading, error, setError }
}

export default useUpdateBarang