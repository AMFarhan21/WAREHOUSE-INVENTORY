'use client'
import { useRouter } from 'next/navigation'
import { useEffect, useState } from 'react'
import { Barang, User } from './useGetAllHistoryStok'
import { Meta } from './useGetAllBarang'

export interface Pembelian {
    id: number
    no_faktur: string
    supplier: string
    total: number
    user_id: number
    status: string
    created_at: string
    user: User
    beli_detail: BeliDetail
}


export interface BeliDetail {
    id: number
    jual_header_id: number
    barang_id: number
    qty: number
    harga: number
    subtotal: number
    barang: Barang[]
}


const useGetAllPembelian = () => {
    const [pembelians, setPembelians] = useState<Pembelian[]>([])
    const [meta, setMeta] = useState<Meta | null>(null)
    const [loading, setLoading] = useState(false)
    const [limitNum, setLimitNum] = useState(5)
    const [pageNum, setPageNum] = useState(0)
    const [search, setSearch] = useState("")
    const [startDate, setStartDate] = useState("")
    const [endDate, setEndDate] = useState("")
    const [error, setError] = useState("")
    const router = useRouter()
    const API_URL = process.env.NEXT_PUBLIC_API_URL


    useEffect(() => {
        const token = localStorage.getItem("TOKEN")
        const getAllBarang = async () => {
            try {
                setLoading(true)
                const res = await fetch(`${API_URL}/api/pembelian?page=${pageNum}&limit=${limitNum}&start_date=${startDate}&end_date=${endDate}`, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                })

                const data = await res.json()

                if (data.error_code == "UNAUTHORIZED") {
                    localStorage.removeItem("TOKEN")
                    router.replace("/auth/login")
                    return
                }


                if (!res.ok) {
                    setError(data.message)
                    return
                }


                setPembelians(data.data)
                setMeta(data.meta)
            } catch (error) {

                console.log(error)

            } finally {
                setLoading(false)
            }
        }

        getAllBarang()
    }, [API_URL, pageNum, search, limitNum, router, startDate, endDate])


    return { pembelians, setPembelians, loading, error, setPageNum, pageNum, limitNum, search, setSearch, meta, setStartDate, setEndDate, startDate, endDate }
}

export default useGetAllPembelian