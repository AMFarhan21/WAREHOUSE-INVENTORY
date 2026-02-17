import { useState } from 'react'


interface BeliDetail {
    barang_id: number
    qty: number
}


const useCreatePembelian = () => {
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState("")
    const API_URL = process.env.NEXT_PUBLIC_API_URL


    const createPembelian = async (supplier: string, beliDetail: BeliDetail[]) => {
        try {
            setLoading(true)
            const token = localStorage.getItem("TOKEN")
            console.log("CEK BARANG", beliDetail)
            const res = await fetch(`${API_URL}/api/pembelian`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`
                },
                body: JSON.stringify({ supplier, beli_detail: beliDetail })
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

    return { createPembelian, loading, error, setError }
}

export default useCreatePembelian