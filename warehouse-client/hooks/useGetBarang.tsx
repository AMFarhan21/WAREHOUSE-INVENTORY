import { useEffect, useState } from 'react'
import { MasterBarang } from './useGetAllBarang'

const useGetBarang = (barangID: number) => {
    const [barang, setBarang] = useState<MasterBarang | null>(null)
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState("")
    const API_URL = process.env.NEXT_PUBLIC_API_URL


    useEffect(() => {
        const token = localStorage.getItem("TOKEN")
        const getAllBarang = async () => {
            try {
                setLoading(true)
                const res = await fetch(`${API_URL}/api/barang/${barangID}`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                })

                const data = await res.json()
                if (!res.ok) {
                    setError(data.message)
                    return
                }

                setBarang(data.data)
            } catch (error) {

                console.log(error)

            } finally {
                setLoading(false)
            }
        }

        getAllBarang()
    }, [API_URL])


    return { barang, setBarang, loading, error }
}

export default useGetBarang
