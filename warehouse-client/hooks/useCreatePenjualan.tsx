import { useState } from 'react'


interface JualDetail {
    barang_id: number
    qty: number
}


const useCreatePenjualan = () => {
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState("")
    const API_URL = process.env.NEXT_PUBLIC_API_URL


    const createPenjualan = async (customer: string, jualDetail: JualDetail[]) => {
        try {
            setLoading(true)
            const token = localStorage.getItem("TOKEN")
            const res = await fetch(`${API_URL}/api/penjualan`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`
                },
                body: JSON.stringify({ customer, jual_detail: jualDetail })
            })

            const data = await res.json()

            if (!res.ok) {
                setError(data.message)
                return
            }

            return data.data
        } catch (error) {
            console.log(error)

        } finally {
            setLoading(false)
        }
    }

    return { createPenjualan, loading, error, setError }
}

export default useCreatePenjualan