import { useEffect, useState } from 'react'
import { Barang } from './useGetAllHistoryStok'

export interface Mstok {
    id: number
    barang_id: number
    stok_akhir: number
    updated_at: string
    barang: Barang
}

export interface Stok {
    id: number
    barang_id: number
    stok_akhir: number
    updated_at: string
}

export const useGetStokByBarangID = (barangID: number) => {
    const [stok, setStok] = useState<Mstok | null>(null)
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState("")
    const API_URL = process.env.NEXT_PUBLIC_API_URL


    useEffect(() => {
        const token = localStorage.getItem("TOKEN")
        const getAllBarang = async () => {
            try {
                setLoading(true)
                const res = await fetch(`${API_URL}/api/stok/${barangID}`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                })

                const data = await res.json()
                if (!res.ok) {
                    setError(data.message)
                    return
                }

                setStok(data.data)
            } catch (error) {

                console.log(error)

            } finally {
                setLoading(false)
            }
        }

        getAllBarang()
    }, [API_URL])


    return { stok, setStok, loading, error }
}

export default useGetStokByBarangID

