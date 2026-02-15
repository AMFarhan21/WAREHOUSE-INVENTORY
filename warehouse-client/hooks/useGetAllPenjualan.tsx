'use client'
import { useRouter } from 'next/navigation'
import { useEffect, useState } from 'react'
import { Barang, User } from './useGetAllHistoryStok'
import { Meta } from './useGetAllBarang'

export interface Penjualan {
    id: number
    no_faktur: string
    customer: string
    total: number
    user_id: number
    status: string
    created_at: string
    user: User
    beli_detail: JualDetail
}


export interface JualDetail {
    id: number
    jual_header_id: number
    barang_id: number
    qty: number
    harga: number
    subtotal: number
    barang: Barang[]
}


const useGetAllPenjualan = () => {
    const [penjualans, setPenjualans] = useState<Penjualan[]>([])
    const [metaPenjualan, setMetaPenjualan] = useState<Meta | null>(null)
    const [loadingPenjualan, setLoading] = useState(false)
    const [limitNumPenjualan, setLimitNumPenjualan] = useState(5)
    const [pageNumPenjualan, setPageNumPenjualan] = useState(0)
    const [search, setSearch] = useState("")
    const [startDatePenjualan, setStartDatePenjualan] = useState("")
    const [endDatePenjualan, setEndDatePenjualan] = useState("")
    const [errorPenjualan, setError] = useState("")
    const router = useRouter()
    const API_URL = process.env.NEXT_PUBLIC_API_URL


    useEffect(() => {
        const token = localStorage.getItem("TOKEN")
        const getAllBarang = async () => {
            try {
                setLoading(true)
                const res = await fetch(`${API_URL}/api/penjualan?page=${pageNumPenjualan}&limit=${limitNumPenjualan}&start_date=${startDatePenjualan}&end_date=${endDatePenjualan}`, {
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


                setPenjualans(data.data)
                setMetaPenjualan(data.meta)
            } catch (error) {

                console.log(error)

            } finally {
                setLoading(false)
            }
        }

        getAllBarang()
    }, [API_URL, pageNumPenjualan, search, limitNumPenjualan, router, startDatePenjualan, endDatePenjualan])


    return { penjualans, setPenjualans, loadingPenjualan, errorPenjualan, setPageNumPenjualan, pageNumPenjualan, limitNumPenjualan, search, setSearch, metaPenjualan, setStartDatePenjualan, setEndDatePenjualan, startDatePenjualan, endDatePenjualan }
}

export default useGetAllPenjualan